package functional_test

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"testing"
	"time"

	factoryapi "github.com/portpowered/agent-factory/pkg/api/generated"
	agentcli "github.com/portpowered/agent-factory/pkg/cli"
	"github.com/portpowered/agent-factory/pkg/factory"
	"github.com/portpowered/agent-factory/pkg/factory/projections"
	"github.com/portpowered/agent-factory/pkg/interfaces"
)

func TestCleanupSmoke_BackendDashboardAndCLIExposeOnlyCleanedFactorySurfaces(t *testing.T) {
	dir := scaffoldFactory(t, simplePipelineConfig())
	server := StartFunctionalServer(t, dir, true, factory.WithServiceMode())

	traceID := server.SubmitWork(t, "task", []byte(`{"title":"cleanup smoke"}`))
	work := waitForGeneratedWorkComplete(t, server.URL(), traceID, 10*time.Second)
	if len(work.Results) != 1 {
		t.Fatalf("GET /work result count = %d, want 1", len(work.Results))
	}
	completed := work.Results[0]
	if completed.TraceId != traceID {
		t.Fatalf("GET /work trace_id = %q, want %q", completed.TraceId, traceID)
	}
	if completed.PlaceId != "task:complete" {
		t.Fatalf("GET /work place_id = %q, want task:complete", completed.PlaceId)
	}

	statusRead := getGeneratedJSON[factoryapi.StatusResponse](t, server.URL()+"/status")
	if statusRead.TotalTokens != 1 {
		t.Fatalf("GET /status total_tokens = %d, want 1", statusRead.TotalTokens)
	}
	if statusRead.Categories.Terminal != 1 {
		t.Fatalf("GET /status terminal count = %d, want 1", statusRead.Categories.Terminal)
	}
	assertCleanupSmokeCanonicalFactoryEvents(t, server, completed.WorkId)
	assertGeneratedEventsStreamHasCanonicalHistory(t, server.URL())
	assertCleanupSmokeRemovedHTTPRoutes(t, server.URL(), traceID, completed.Id)
	assertCleanupSmokeDashboardShell(t, server.URL())
	assertCleanupSmokeCLI(t)
}

func assertCleanupSmokeCanonicalFactoryEvents(t *testing.T, server *FunctionalServer, workID string) {
	t.Helper()

	events, err := server.service.GetFactoryEvents(context.Background())
	if err != nil {
		t.Fatalf("GetFactoryEvents: %v", err)
	}
	assertCleanupSmokeHasEventType(t, events, factoryapi.FactoryEventTypeWorkRequest)
	assertCleanupSmokeHasEventType(t, events, factoryapi.FactoryEventTypeDispatchRequest)
	assertCleanupSmokeHasEventType(t, events, factoryapi.FactoryEventTypeDispatchResponse)

	worldState, err := projections.ReconstructFactoryWorldState(events, cleanupSmokeMaxTick(events))
	if err != nil {
		t.Fatalf("ReconstructFactoryWorldState: %v", err)
	}
	worldView := projections.BuildFactoryWorldView(worldState)
	if worldView.Runtime.Session.CompletedCount != 1 {
		t.Fatalf("canonical world view completed count = %d, want 1", worldView.Runtime.Session.CompletedCount)
	}
	if got := worldView.Runtime.PlaceTokenCounts["task:complete"]; got != 1 {
		t.Fatalf("canonical world view task:complete count = %d, want 1", got)
	}
	if !cleanupSmokePlaceContainsWork(worldView.Runtime.PlaceOccupancyWorkItemsByPlaceID["task:complete"], workID) {
		t.Fatalf("canonical world view task:complete occupancy = %#v, want work %q", worldView.Runtime.PlaceOccupancyWorkItemsByPlaceID["task:complete"], workID)
	}
}

func assertCleanupSmokeHasEventType(t *testing.T, events []factoryapi.FactoryEvent, eventType factoryapi.FactoryEventType) {
	t.Helper()

	for _, event := range events {
		if event.Type == eventType {
			return
		}
	}
	t.Fatalf("GetFactoryEvents missing %s in canonical history", eventType)
}

func cleanupSmokeMaxTick(events []factoryapi.FactoryEvent) int {
	maxTick := 0
	for _, event := range events {
		if event.Context.Tick > maxTick {
			maxTick = event.Context.Tick
		}
	}
	return maxTick
}

func cleanupSmokePlaceContainsWork(items []interfaces.FactoryWorldWorkItemRef, workID string) bool {
	for _, item := range items {
		if item.WorkID == workID {
			return true
		}
	}
	return false
}

func assertCleanupSmokeRemovedHTTPRoutes(t *testing.T, baseURL, traceID, tokenID string) {
	t.Helper()

	for _, path := range []string{
		"/dashboard",
		"/dashboard/stream",
		"/state",
		"/traces/" + url.PathEscape(traceID),
		"/work/" + url.PathEscape(tokenID) + "/trace",
		"/workflows",
		"/workflows/default",
	} {
		resp, err := http.Get(baseURL + path)
		if err != nil {
			t.Fatalf("GET %s: %v", path, err)
		}
		_, _ = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()
		if resp.StatusCode != http.StatusNotFound && resp.StatusCode != http.StatusMethodNotAllowed {
			t.Fatalf("GET %s status = %d, want route removed", path, resp.StatusCode)
		}
	}
}

func assertCleanupSmokeDashboardShell(t *testing.T, baseURL string) {
	t.Helper()

	resp, err := http.Get(baseURL + "/dashboard/ui")
	if err != nil {
		t.Fatalf("GET /dashboard/ui: %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read /dashboard/ui: %v", err)
	}
	shell := string(body)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("GET /dashboard/ui status = %d, want 200: %s", resp.StatusCode, shell)
	}
	for _, want := range []string{
		"<title>Agent Factory Dashboard</title>",
		"<div id=\"root\"></div>",
		"/dashboard/ui/assets/",
	} {
		if !strings.Contains(shell, want) {
			t.Fatalf("dashboard shell missing %q", want)
		}
	}

	routeResp, err := http.Get(baseURL + "/dashboard/ui/work/" + url.PathEscape("work-from-cleanup-smoke"))
	if err != nil {
		t.Fatalf("GET /dashboard/ui/work/...: %v", err)
	}
	defer routeResp.Body.Close()
	routeBody, err := io.ReadAll(routeResp.Body)
	if err != nil {
		t.Fatalf("read dashboard client route: %v", err)
	}
	if routeResp.StatusCode != http.StatusOK {
		t.Fatalf("dashboard client route status = %d, want 200", routeResp.StatusCode)
	}
	if string(routeBody) != shell {
		t.Fatal("dashboard client route should fall back to the embedded app shell")
	}

	assertCleanupSmokeDashboardBundleUsesCleanedAPI(t, baseURL, shell)
}

func assertCleanupSmokeDashboardBundleUsesCleanedAPI(t *testing.T, baseURL, shell string) {
	t.Helper()

	matches := regexp.MustCompile(`(?:src|href)="(/dashboard/ui/assets/[^"]+)"`).FindAllStringSubmatch(shell, -1)
	if len(matches) == 0 {
		t.Fatalf("dashboard shell did not reference embedded assets: %s", shell)
	}

	foundEventsEndpoint := false
	workTracePattern := regexp.MustCompile(`["'` + "`" + `]/work/[^"'` + "`" + `]+/trace(?:["'` + "`" + `/?]|$)`)
	removedEndpointPatterns := map[string]*regexp.Regexp{
		"/state":            regexp.MustCompile(`["'` + "`" + `]/state(?:["'` + "`" + `/?]|$)`),
		"/traces/":          regexp.MustCompile(`["'` + "`" + `]/traces/`),
		"/workflows":        regexp.MustCompile(`["'` + "`" + `]/workflows(?:["'` + "`" + `/?]|$)`),
		"/dashboard/stream": regexp.MustCompile(`["'` + "`" + `]/dashboard/stream(?:["'` + "`" + `/?]|$)`),
	}
	for _, match := range matches {
		assetPath := match[1]
		if !strings.HasSuffix(assetPath, ".js") {
			continue
		}

		resp, err := http.Get(baseURL + assetPath)
		if err != nil {
			t.Fatalf("GET %s: %v", assetPath, err)
		}
		body, readErr := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if readErr != nil {
			t.Fatalf("read %s: %v", assetPath, readErr)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("GET %s status = %d, want 200", assetPath, resp.StatusCode)
		}

		bundle := string(body)
		foundEventsEndpoint = foundEventsEndpoint || strings.Contains(bundle, "/events")
		for removed, pattern := range removedEndpointPatterns {
			if pattern.MatchString(bundle) {
				t.Fatalf("dashboard bundle %s still references removed endpoint %q", assetPath, removed)
			}
		}
		if workTracePattern.MatchString(bundle) {
			t.Fatalf("dashboard bundle %s still references removed work trace endpoint", assetPath)
		}
	}
	if !foundEventsEndpoint {
		t.Fatal("dashboard bundle did not reference canonical /events stream endpoint")
	}
}

func assertCleanupSmokeCLI(t *testing.T) {
	t.Helper()

	root := agentcli.NewRootCommand()
	expectedCommands := map[string]bool{
		"config": false,
		"init":   false,
		"run":    false,
		"submit": false,
	}
	for _, subcommand := range root.Commands() {
		if _, ok := expectedCommands[subcommand.Name()]; ok {
			expectedCommands[subcommand.Name()] = true
		}
		if subcommand.Name() == "audit" {
			t.Fatal("removed audit command should not be registered")
		}
		if subcommand.Name() == "status" {
			t.Fatal("removed status command should not be registered")
		}
	}
	for command, found := range expectedCommands {
		if !found {
			t.Fatalf("expected CLI command %q to be registered", command)
		}
	}

	for _, args := range [][]string{
		{"audit", "state-surfaces"},
		{"formattraceexplorer"},
		{"status"},
		{"trace"},
	} {
		cmd := agentcli.NewRootCommand()
		cmd.SetOut(io.Discard)
		cmd.SetErr(io.Discard)
		cmd.SetArgs(args)
		if err := cmd.Execute(); err == nil {
			t.Fatalf("expected removed CLI command %q to fail", strings.Join(args, " "))
		}
	}
}
