package functional_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"

	factoryapi "github.com/portpowered/agent-factory/pkg/api/generated"
	agentcli "github.com/portpowered/agent-factory/pkg/cli"
	"github.com/portpowered/agent-factory/pkg/factory"
	"github.com/portpowered/agent-factory/pkg/factory/projections"
	"github.com/portpowered/agent-factory/pkg/interfaces"
	"github.com/portpowered/agent-factory/pkg/testutil"
	functionalharness "github.com/portpowered/agent-factory/tests/functional/support/harness"
)

const cleanupSmokeProject = "acme-inventory"

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

func TestCleanupSmoke_RuntimeContextAndCanonicalHistoryStayProjectAgnostic(t *testing.T) {
	dir := testutil.CopyFixtureDir(t, fixtureDir(t, "tags_test"))
	setWorkingDirectory(t, dir)
	rewriteCleanupSmokeProject(t, dir, cleanupSmokeProject)

	provider := testutil.NewMockWorkerMapProvider(map[string][]interfaces.InferenceResponse{
		"checker": {{Content: "cleanup COMPLETE"}},
	})
	h := testutil.NewServiceTestHarness(t, dir,
		testutil.WithProvider(provider),
		testutil.WithFullWorkerPoolAndScriptWrap(),
	)

	h.SubmitWorkRequest(context.Background(), interfaces.WorkRequest{
		RequestID: "request-project-cleanup-smoke",
		Type:      interfaces.WorkRequestTypeFactoryRequestBatch,
		Works: []interfaces.Work{{
			Name:       "cleanup smoke work",
			WorkID:     "work-project-cleanup-smoke",
			WorkTypeID: "task",
			TraceID:    "trace-project-cleanup-smoke",
			Payload:    "cleanup smoke payload",
			Tags: map[string]string{
				"branch":  "feature/acme-cleanup",
				"project": cleanupSmokeProject,
			},
		}},
	})

	h.RunUntilComplete(t, 10*time.Second)
	h.Assert().PlaceTokenCount("task:complete", 1)

	calls := provider.Calls("checker")
	if len(calls) != 1 {
		t.Fatalf("checker provider calls = %d, want 1", len(calls))
	}
	assertCleanupSmokeRuntimeContext(t, dir, calls[0])
	assertInferenceRequestsDoNotContainPortOS(t, calls)

	events, err := h.GetFactoryEvents(context.Background())
	if err != nil {
		t.Fatalf("GetFactoryEvents: %v", err)
	}
	assertCleanupSmokeEvents(t, events)
	assertFactoryEventsDoNotContainPortOS(t, events)

	data, err := json.Marshal(events)
	if err != nil {
		t.Fatalf("marshal cleanup smoke events: %v", err)
	}
	assertTextOmitsRetiredEventNames(t, "cleanup smoke canonical events", string(data))
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

func rewriteCleanupSmokeProject(t *testing.T, dir, project string) {
	t.Helper()

	configPath := filepath.Join(dir, interfaces.FactoryConfigFile)
	data, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatalf("read cleanup smoke factory config: %v", err)
	}

	var cfg map[string]any
	if err := json.Unmarshal(data, &cfg); err != nil {
		t.Fatalf("unmarshal cleanup smoke factory config: %v", err)
	}
	cfg["project"] = project

	updated, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		t.Fatalf("marshal cleanup smoke factory config: %v", err)
	}
	if err := os.WriteFile(configPath, updated, 0o644); err != nil {
		t.Fatalf("write cleanup smoke factory config: %v", err)
	}
}

func assertCleanupSmokeRuntimeContext(t *testing.T, dir string, call interfaces.ProviderInferenceRequest) {
	t.Helper()

	if call.WorkingDirectory != resolvedRuntimePath(dir, "/workspaces/acme-inventory/feature/acme-cleanup") {
		t.Fatalf("working directory = %q, want acme project path", call.WorkingDirectory)
	}
	if call.EnvVars["PROJECT"] != cleanupSmokeProject {
		t.Fatalf("PROJECT env = %q, want %s", call.EnvVars["PROJECT"], cleanupSmokeProject)
	}
	if call.EnvVars["CONTEXT_PROJECT"] != cleanupSmokeProject {
		t.Fatalf("CONTEXT_PROJECT env = %q, want %s", call.EnvVars["CONTEXT_PROJECT"], cleanupSmokeProject)
	}
	if call.EnvVars["BRANCH"] != "feature/acme-cleanup" {
		t.Fatalf("BRANCH env = %q, want feature/acme-cleanup", call.EnvVars["BRANCH"])
	}
	if len(call.InputTokens) != 1 {
		t.Fatalf("provider input tokens = %d, want 1", len(call.InputTokens))
	}
	assertCleanupSmokeTags(t, "provider input token", functionalharness.FirstInputToken(call.InputTokens).Color.Tags)
}

func assertCleanupSmokeEvents(t *testing.T, events []factoryapi.FactoryEvent) {
	t.Helper()

	var sawRequest bool
	var sawWorkInput bool
	var sawDispatch bool
	var sawTerminalOutput bool
	for _, event := range events {
		switch event.Type {
		case factoryapi.FactoryEventTypeWorkRequest:
			payload, err := event.Payload.AsWorkRequestEventPayload()
			if err != nil || stringPointerValue(event.Context.RequestId) != "request-project-cleanup-smoke" || payload.Works == nil {
				continue
			}
			sawRequest = true
			if len(*payload.Works) != 1 {
				t.Fatalf("cleanup smoke request works = %d, want 1", len(*payload.Works))
			}
			sawWorkInput = true
			assertCleanupSmokeTags(t, "WORK_REQUEST item", generatedTags((*payload.Works)[0].Tags))
		case factoryapi.FactoryEventTypeDispatchRequest:
			payload, err := event.Payload.AsDispatchRequestEventPayload()
			if err != nil || payload.TransitionId != "process" {
				continue
			}
			for _, input := range dispatchInputWorksFromHistory(t, events, event, payload) {
				if stringPointerValue(input.WorkId) != "work-project-cleanup-smoke" {
					continue
				}
				sawDispatch = true
				assertCleanupSmokeTags(t, "DISPATCH_CREATED input", generatedTags(input.Tags))
			}
		case factoryapi.FactoryEventTypeDispatchResponse:
			payload, err := event.Payload.AsDispatchResponseEventPayload()
			if err != nil || payload.OutputWork == nil {
				continue
			}
			for _, output := range *payload.OutputWork {
				if stringPointerValue(output.WorkId) != "work-project-cleanup-smoke" {
					continue
				}
				sawTerminalOutput = true
				assertCleanupSmokeTags(t, "DISPATCH_COMPLETED output work", generatedTags(output.Tags))
			}
		}
	}
	if !sawRequest || !sawWorkInput || !sawDispatch || !sawTerminalOutput {
		t.Fatalf(
			"cleanup smoke missing event boundary: request=%v input=%v dispatch=%v terminal=%v",
			sawRequest,
			sawWorkInput,
			sawDispatch,
			sawTerminalOutput,
		)
	}
}

func generatedTags(tags *factoryapi.StringMap) map[string]string {
	if tags == nil {
		return nil
	}
	return map[string]string(*tags)
}

func assertCleanupSmokeTags(t *testing.T, label string, tags map[string]string) {
	t.Helper()

	if tags["project"] != cleanupSmokeProject {
		t.Fatalf("%s project tag = %q, want %s", label, tags["project"], cleanupSmokeProject)
	}
	if tags["branch"] != "feature/acme-cleanup" {
		t.Fatalf("%s branch tag = %q, want feature/acme-cleanup", label, tags["branch"])
	}
	assertMapDoesNotContainPortOS(t, label, tags)
}

func assertInferenceRequestsDoNotContainPortOS(t *testing.T, calls []interfaces.ProviderInferenceRequest) {
	t.Helper()

	if len(calls) == 0 {
		t.Fatal("expected at least one provider request")
	}
	for i, call := range calls {
		data, err := json.Marshal(call)
		if err != nil {
			t.Fatalf("marshal provider request %d: %v", i, err)
		}
		assertValueDoesNotContainPortOS(t, fmt.Sprintf("provider request %d", i), string(data))
	}
}

func assertFactoryEventsDoNotContainPortOS(t *testing.T, events []factoryapi.FactoryEvent) {
	t.Helper()

	if len(events) == 0 {
		t.Fatal("expected at least one factory event")
	}
	for i, event := range events {
		data, err := json.Marshal(event)
		if err != nil {
			t.Fatalf("marshal factory event %d: %v", i, err)
		}
		assertValueDoesNotContainPortOS(t, fmt.Sprintf("factory event %d (%s)", i, event.Type), string(data))
	}
}

func assertMapDoesNotContainPortOS(t *testing.T, label string, values map[string]string) {
	t.Helper()

	for key, value := range values {
		assertValueDoesNotContainPortOS(t, label+" key", key)
		assertValueDoesNotContainPortOS(t, label+" value", value)
	}
}

func assertValueDoesNotContainPortOS(t *testing.T, label string, value string) {
	t.Helper()

	normalized := strings.ToLower(value)
	if strings.Contains(normalized, "portos") ||
		strings.Contains(normalized, "port os") ||
		strings.Contains(normalized, "port_os") {
		t.Fatalf("%s contains Port OS coupling: %q", label, value)
	}
}

func assertTextOmitsRetiredEventNames(t *testing.T, label string, value string) {
	t.Helper()

	for _, retired := range []string{
		"RUN_STARTED",
		"INITIAL_STRUCTURE",
		"RELATIONSHIP_CHANGE",
		"DISPATCH_CREATED",
		"DISPATCH_COMPLETED",
		"FACTORY_STATE_CHANGE",
		"RUN_FINISHED",
	} {
		if strings.Contains(value, `"`+retired+`"`) {
			t.Fatalf("%s contains retired public event name %q", label, retired)
		}
	}
}
