package smoke_test

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	initcmd "github.com/portpowered/agent-factory/pkg/cli/init"
	"github.com/portpowered/agent-factory/pkg/interfaces"
	"github.com/portpowered/agent-factory/pkg/testutil"
)

var retiredInitFactoryContractFields = []string{`"work_types"`, `"work_type"`, `"on_failure"`}
var retiredInitWorkerContractFields = []string{"model_provider:", "provider:", "stop_token:", "skip_permissions:", "concurrency:", "sessionId:"}

func TestInitFactory_StructureIsValid(t *testing.T) {
	dir := t.TempDir()

	if err := initcmd.Init(initcmd.InitConfig{Dir: dir}); err != nil {
		t.Fatalf("Init failed: %v", err)
	}

	expectedFiles := []string{
		"factory.json",
		filepath.Join("workers", "README.md"),
		filepath.Join("workers", "processor", "AGENTS.md"),
		filepath.Join("workstations", "README.md"),
		filepath.Join("workstations", "process", "AGENTS.md"),
		filepath.Join("inputs", "README.md"),
	}
	for _, f := range expectedFiles {
		path := filepath.Join(dir, f)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("expected file %s to exist", f)
		}
	}

	expectedDirs := []string{
		filepath.Join("inputs", "tasks", "default"),
	}
	for _, d := range expectedDirs {
		path := filepath.Join(dir, d)
		info, err := os.Stat(path)
		if err != nil {
			t.Errorf("expected directory %s to exist: %v", d, err)
		} else if !info.IsDir() {
			t.Errorf("expected %s to be a directory", d)
		}
	}

	factoriesPath := filepath.Join(dir, "factories")
	if _, err := os.Stat(factoriesPath); err == nil {
		t.Error("expected 'factories/' directory to NOT be created by init")
	}

	tasksPath := filepath.Join(dir, "inputs", "task")
	if _, err := os.Stat(tasksPath); err == nil {
		t.Error("expected 'inputs/task/' (singular) to NOT be created; should be 'inputs/tasks/'")
	}
}

func TestInitFactory_EndToEnd(t *testing.T) {
	dir := t.TempDir()

	if err := initcmd.Init(initcmd.InitConfig{Dir: dir}); err != nil {
		t.Fatalf("Init failed: %v", err)
	}
	assertGeneratedInitScaffoldCanonical(t, dir, "gpt-5-codex", "codex")

	testutil.WriteSeedFile(t, dir, "tasks", []byte(`{"title": "init factory e2e test"}`))

	work := map[string][]testutil.WorkResponse{
		"processor": {
			{Content: "Task processed successfully."},
		},
	}
	provider := testutil.NewMockWorkerMapProviderWithDefault(work)

	h := testutil.NewServiceTestHarness(t, dir,
		testutil.WithProvider(provider),
		testutil.WithFullWorkerPoolAndScriptWrap(),
	)

	h.RunUntilComplete(t, 15*time.Second)

	h.Assert().
		HasTokenInPlace("tasks:complete").
		HasNoTokenInPlace("tasks:init").
		HasNoTokenInPlace("tasks:failed").
		TokenCount(1)

	if provider.CallCount("processor") != 1 {
		t.Errorf("expected provider called 1 time, got %d", provider.CallCount("processor"))
	}

	calls := provider.Calls("processor")
	if len(calls) == 0 {
		t.Fatal("expected at least 1 provider call")
	}
	if calls[0].UserMessage == "" {
		t.Error("expected non-empty user message from rendered workstations/process/AGENTS.md template")
	}
	assertInitProviderRequest(t, calls[0], "gpt-5-codex", "codex")
}

func TestInitFactory_ClaudeEndToEndUsesClaudeStarterWorker(t *testing.T) {
	dir := t.TempDir()

	if err := initcmd.Init(initcmd.InitConfig{Dir: dir, Executor: "claude"}); err != nil {
		t.Fatalf("Init failed: %v", err)
	}
	assertGeneratedInitScaffoldCanonical(t, dir, "claude-sonnet-4-20250514", "claude")

	testutil.WriteSeedFile(t, dir, "tasks", []byte(`{"title": "claude init factory e2e test"}`))

	work := map[string][]testutil.WorkResponse{
		"processor": {
			{Content: "Claude task processed successfully."},
		},
	}
	provider := testutil.NewMockWorkerMapProviderWithDefault(work)

	h := testutil.NewServiceTestHarness(t, dir,
		testutil.WithProvider(provider),
		testutil.WithFullWorkerPoolAndScriptWrap(),
	)

	h.RunUntilComplete(t, 15*time.Second)

	h.Assert().
		HasTokenInPlace("tasks:complete").
		HasNoTokenInPlace("tasks:init").
		HasNoTokenInPlace("tasks:failed").
		TokenCount(1)

	if provider.CallCount("processor") != 1 {
		t.Errorf("expected provider called 1 time, got %d", provider.CallCount("processor"))
	}

	calls := provider.Calls("processor")
	if len(calls) == 0 {
		t.Fatal("expected at least 1 provider call")
	}
	assertInitProviderRequest(t, calls[0], "claude-sonnet-4-20250514", "claude")
}

func TestInitFactory_FailureRouting(t *testing.T) {
	dir := t.TempDir()

	if err := initcmd.Init(initcmd.InitConfig{Dir: dir}); err != nil {
		t.Fatalf("Init failed: %v", err)
	}

	testutil.WriteSeedFile(t, dir, "tasks", []byte(`{"title": "failing task"}`))

	work := map[string][]testutil.WorkResponse{
		"processor": {
			{Content: "something went wrong", Error: errors.New("provider execution failed")},
		},
	}
	provider := testutil.NewMockWorkerMapProviderWithDefault(work)

	h := testutil.NewServiceTestHarness(t, dir,
		testutil.WithProvider(provider),
		testutil.WithFullWorkerPoolAndScriptWrap(),
	)

	h.RunUntilComplete(t, 15*time.Second)

	h.Assert().
		HasTokenInPlace("tasks:failed").
		HasNoTokenInPlace("tasks:init").
		HasNoTokenInPlace("tasks:complete").
		TokenCount(1)
}

func TestInitFactory_Idempotent(t *testing.T) {
	dir := t.TempDir()

	if err := initcmd.Init(initcmd.InitConfig{Dir: dir}); err != nil {
		t.Fatalf("first Init failed: %v", err)
	}

	customContent := []byte("custom worker content")
	customPath := filepath.Join(dir, "workers", "processor", "AGENTS.md")
	if err := os.WriteFile(customPath, customContent, 0o644); err != nil {
		t.Fatalf("write custom file: %v", err)
	}

	if err := initcmd.Init(initcmd.InitConfig{Dir: dir}); err != nil {
		t.Fatalf("second Init failed: %v", err)
	}

	data, err := os.ReadFile(customPath)
	if err != nil {
		t.Fatalf("read custom file: %v", err)
	}
	if string(data) != string(customContent) {
		t.Error("expected Init to not overwrite existing AGENTS.md files")
	}
}

func assertInitProviderRequest(t *testing.T, req interfaces.ProviderInferenceRequest, wantModel, wantProvider string) {
	t.Helper()

	if req.Model != wantModel {
		t.Fatalf("provider request model = %q, want %q", req.Model, wantModel)
	}
	if req.ModelProvider != wantProvider {
		t.Fatalf("provider request model provider = %q, want %q", req.ModelProvider, wantProvider)
	}
	if req.SystemPrompt == "" {
		t.Fatal("expected provider request system prompt to be populated from worker AGENTS.md")
	}
	if !strings.Contains(req.SystemPrompt, "You are the processor. Complete the task.") {
		t.Fatalf("provider request system prompt = %q, want default processor instructions", req.SystemPrompt)
	}
	if req.UserMessage == "" {
		t.Fatal("expected provider request user message to be populated from workstation AGENTS.md")
	}
}

func assertGeneratedInitScaffoldCanonical(t *testing.T, dir, wantModel, wantProvider string) {
	t.Helper()

	factoryJSONBytes, err := os.ReadFile(filepath.Join(dir, "factory.json"))
	if err != nil {
		t.Fatalf("read generated factory.json: %v", err)
	}
	factoryJSON := string(factoryJSONBytes)
	for _, expected := range []string{`"workType"`, `"onFailure"`} {
		if !strings.Contains(factoryJSON, expected) {
			t.Fatalf("generated factory.json should contain %q:\n%s", expected, factoryJSON)
		}
	}
	for _, retired := range retiredInitFactoryContractFields {
		if strings.Contains(factoryJSON, retired) {
			t.Fatalf("generated factory.json should not contain retired %q:\n%s", retired, factoryJSON)
		}
	}

	workerAgentsBytes, err := os.ReadFile(filepath.Join(dir, "workers", "processor", "AGENTS.md"))
	if err != nil {
		t.Fatalf("read generated worker AGENTS.md: %v", err)
	}
	workerAgents := string(workerAgentsBytes)
	for _, expected := range []string{
		"model: " + wantModel,
		"modelProvider: " + wantProvider,
		"executorProvider: script_wrap",
		"skipPermissions: true",
		"timeout: 1h",
		"You are the processor. Complete the task.",
	} {
		if !strings.Contains(workerAgents, expected) {
			t.Fatalf("generated worker AGENTS.md should contain %q:\n%s", expected, workerAgents)
		}
	}
	for _, retired := range retiredInitWorkerContractFields {
		if strings.Contains(workerAgents, retired) {
			t.Fatalf("generated worker AGENTS.md should not contain retired %q:\n%s", retired, workerAgents)
		}
	}
}
