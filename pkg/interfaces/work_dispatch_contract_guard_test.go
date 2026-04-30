package interfaces

import (
	"reflect"
	"strings"
	"testing"
)

var approvedWorkDispatchFields = map[string]string{
	"DispatchID":               "dispatch_id",
	"TransitionID":             "transition_id",
	"WorkerType":               "worker_type",
	"WorkstationName":          "workstation_name",
	"ProjectID":                "project_id",
	"CurrentChainingTraceID":   "current_chaining_trace_id",
	"PreviousChainingTraceIDs": "previous_chaining_trace_ids",
	"Execution":                "execution",
	"InputTokens":              "input_tokens",
	"InputBindings":            "input_bindings",
}

var retiredWorkDispatchWorkerFields = []string{
	"WorkstationType",
	"SystemPrompt",
	"UserMessage",
	"OutputSchema",
	"EnvVars",
	"Worktree",
	"WorkingDirectory",
	"Model",
	"ModelProvider",
	"Provider",
	"SessionID",
	"Command",
	"Args",
	"Stdin",
	"Env",
	"WorkDir",
}

func TestWorkDispatchContractGuard_FieldInventoryStaysDispatchOwned(t *testing.T) {
	t.Parallel()

	workDispatchType := reflect.TypeOf(WorkDispatch{})
	seen := make(map[string]struct{}, workDispatchType.NumField())
	for i := 0; i < workDispatchType.NumField(); i++ {
		field := workDispatchType.Field(i)
		wantJSONTag, ok := approvedWorkDispatchFields[field.Name]
		if !ok {
			t.Fatalf("WorkDispatch field %s is not in the approved dispatch-owned inventory; update the split-contract artifact and this guard before expanding the canonical dispatch payload", field.Name)
		}
		gotJSONTag := strings.Split(field.Tag.Get("json"), ",")[0]
		if gotJSONTag != wantJSONTag {
			t.Fatalf("WorkDispatch field %s json tag = %q, want %q", field.Name, gotJSONTag, wantJSONTag)
		}
		seen[field.Name] = struct{}{}
	}

	for fieldName := range approvedWorkDispatchFields {
		if _, ok := seen[fieldName]; !ok {
			t.Fatalf("WorkDispatch is missing approved field %s", fieldName)
		}
	}
}

func TestWorkDispatchContractGuard_WorkerOwnedFieldsStayDeleted(t *testing.T) {
	t.Parallel()

	workDispatchType := reflect.TypeOf(WorkDispatch{})
	for _, fieldName := range retiredWorkDispatchWorkerFields {
		if _, ok := workDispatchType.FieldByName(fieldName); ok {
			t.Fatalf("WorkDispatch must not reintroduce worker-owned field %s", fieldName)
		}
	}
}
