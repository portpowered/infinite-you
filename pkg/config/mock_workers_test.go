package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParseMockWorkersConfig_ValidConfigPreservesSelectorsAndRunTypeOptions(t *testing.T) {
	cfg, err := ParseMockWorkersConfig([]byte(`{
		"mockWorkers": [
			{
				"id": "accept-reviewer",
				"workerName": "reviewer",
				"workstationName": "review",
				"workInputs": [
					{
						"workId": "work-1",
						"workType": "story",
						"state": "in-review",
						"inputName": "story",
						"traceId": "trace-1",
						"channel": "default",
						"payloadHash": "sha256:test"
					}
				],
				"runType": "accept"
			},
			{
				"id": "script-executor",
				"workerName": "executor",
				"workstationName": "execute",
				"runType": "script",
				"scriptConfig": {
					"command": "go",
					"args": ["test", "./..."],
					"env": {"AGENT_FACTORY_MOCK": "1"},
					"workingDirectory": "/tmp/work",
					"stdin": "script input",
					"timeout": "30s"
				}
			},
			{
				"id": "reject-reviewer",
				"workerName": "reviewer",
				"workstationName": "review",
				"runType": "reject",
				"rejectConfig": {
					"stdout": "review output",
					"stderr": "needs changes",
					"exitCode": 7
				}
			}
		]
	}`))
	if err != nil {
		t.Fatalf("ParseMockWorkersConfig returned error: %v", err)
	}

	if len(cfg.MockWorkers) != 3 {
		t.Fatalf("mock worker count = %d, want 3", len(cfg.MockWorkers))
	}

	accept := cfg.MockWorkers[0]
	if accept.ID != "accept-reviewer" ||
		accept.WorkerName != "reviewer" ||
		accept.WorkstationName != "review" ||
		accept.RunType != MockWorkerRunTypeAccept {
		t.Fatalf("accept entry = %#v, want selectors and accept run type preserved", accept)
	}
	if len(accept.WorkInputs) != 1 {
		t.Fatalf("accept work input count = %d, want 1", len(accept.WorkInputs))
	}
	input := accept.WorkInputs[0]
	if input.WorkID != "work-1" ||
		input.WorkType != "story" ||
		input.State != "in-review" ||
		input.InputName != "story" ||
		input.TraceID != "trace-1" ||
		input.Channel != "default" ||
		input.PayloadHash != "sha256:test" {
		t.Fatalf("accept work input = %#v, want all selectors preserved", input)
	}

	script := cfg.MockWorkers[1]
	if script.RunType != MockWorkerRunTypeScript {
		t.Fatalf("script run type = %q, want %q", script.RunType, MockWorkerRunTypeScript)
	}
	if script.ScriptConfig == nil {
		t.Fatal("scriptConfig was not preserved")
	}
	if script.ScriptConfig.Command != "go" ||
		strings.Join(script.ScriptConfig.Args, " ") != "test ./..." ||
		script.ScriptConfig.Env["AGENT_FACTORY_MOCK"] != "1" ||
		script.ScriptConfig.WorkingDirectory != "/tmp/work" ||
		script.ScriptConfig.Stdin != "script input" ||
		script.ScriptConfig.Timeout != "30s" {
		t.Fatalf("script config = %#v, want command options preserved", script.ScriptConfig)
	}

	reject := cfg.MockWorkers[2]
	if reject.RunType != MockWorkerRunTypeReject {
		t.Fatalf("reject run type = %q, want %q", reject.RunType, MockWorkerRunTypeReject)
	}
	if reject.RejectConfig == nil {
		t.Fatal("rejectConfig was not preserved")
	}
	if reject.RejectConfig.Stdout != "review output" ||
		reject.RejectConfig.Stderr != "needs changes" ||
		reject.RejectConfig.ExitCode == nil ||
		*reject.RejectConfig.ExitCode != 7 {
		t.Fatalf("reject config = %#v, want stdout, stderr, and exit code preserved", reject.RejectConfig)
	}
}

func TestParseMockWorkersConfig_RejectsUnknownRunTypeWithActionableError(t *testing.T) {
	_, err := ParseMockWorkersConfig([]byte(`{
		"mockWorkers": [
			{"id": "bad", "runType": "maybe"}
		]
	}`))
	if err == nil {
		t.Fatal("expected unknown runType to fail validation")
	}
	if !strings.Contains(err.Error(), `runType must be one of "accept", "script", or "reject"; got "maybe"`) {
		t.Fatalf("error = %q, want actionable runType message", err)
	}
}

func TestParseMockWorkersConfig_RejectsScriptEntryWithoutScriptConfig(t *testing.T) {
	_, err := ParseMockWorkersConfig([]byte(`{
		"mockWorkers": [
			{"id": "script", "runType": "script"}
		]
	}`))
	if err == nil {
		t.Fatal("expected script runType without scriptConfig to fail validation")
	}
	if !strings.Contains(err.Error(), "scriptConfig is required") {
		t.Fatalf("error = %q, want missing scriptConfig message", err)
	}
}

func TestParseMockWorkersConfig_RejectsScriptEntryWithoutCommand(t *testing.T) {
	_, err := ParseMockWorkersConfig([]byte(`{
		"mockWorkers": [
			{"id": "script", "runType": "script", "scriptConfig": {"args": ["ok"]}}
		]
	}`))
	if err == nil {
		t.Fatal("expected script runType without scriptConfig.command to fail validation")
	}
	if !strings.Contains(err.Error(), "scriptConfig.command is required") {
		t.Fatalf("error = %q, want missing scriptConfig.command message", err)
	}
}

func TestParseMockWorkersConfig_RejectsRejectEntryWithInvalidExitCode(t *testing.T) {
	for _, exitCode := range []string{"-1", "0", "256"} {
		t.Run(exitCode, func(t *testing.T) {
			_, err := ParseMockWorkersConfig([]byte(`{
				"mockWorkers": [
					{"id": "reject", "runType": "reject", "rejectConfig": {"exitCode": ` + exitCode + `}}
				]
			}`))
			if err == nil {
				t.Fatal("expected invalid reject exit code to fail validation")
			}
			if !strings.Contains(err.Error(), "rejectConfig.exitCode must be between 1 and 255") {
				t.Fatalf("error = %q, want invalid exit-code message", err)
			}
		})
	}
}

func TestLoadMockWorkersConfig_EmptyPathReturnsEmptyDefaultAcceptConfig(t *testing.T) {
	cfg, err := LoadMockWorkersConfig("")
	if err != nil {
		t.Fatalf("LoadMockWorkersConfig empty path returned error: %v", err)
	}
	if cfg == nil {
		t.Fatal("LoadMockWorkersConfig empty path returned nil config")
	}
	if len(cfg.MockWorkers) != 0 {
		t.Fatalf("mock worker count = %d, want empty default config", len(cfg.MockWorkers))
	}
}

func TestLoadMockWorkersConfig_LoadsConfigFromPath(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "mock-workers.json")
	if err := os.WriteFile(path, []byte(`{"mockWorkers":[{"id":"accepted","runType":"accept"}]}`), 0o644); err != nil {
		t.Fatalf("write mock config: %v", err)
	}

	cfg, err := LoadMockWorkersConfig(path)
	if err != nil {
		t.Fatalf("LoadMockWorkersConfig returned error: %v", err)
	}
	if len(cfg.MockWorkers) != 1 {
		t.Fatalf("mock worker count = %d, want 1", len(cfg.MockWorkers))
	}
	if cfg.MockWorkers[0].ID != "accepted" || cfg.MockWorkers[0].RunType != MockWorkerRunTypeAccept {
		t.Fatalf("mock worker = %#v, want loaded accept entry", cfg.MockWorkers[0])
	}
}
