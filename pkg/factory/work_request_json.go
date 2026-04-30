package factory

import (
	"encoding/json"
	"fmt"

	"github.com/portpowered/agent-factory/pkg/interfaces"
)

// ParseCanonicalWorkRequestJSON parses a public FACTORY_REQUEST_BATCH JSON
// payload and rejects retired aliases that must not be accepted on public
// submit boundaries.
func ParseCanonicalWorkRequestJSON(data []byte) (interfaces.WorkRequest, error) {
	var request interfaces.WorkRequest
	if err := json.Unmarshal(data, &request); err != nil {
		return interfaces.WorkRequest{}, err
	}
	if err := RejectRetiredWorkRequestFieldAliases(data); err != nil {
		return interfaces.WorkRequest{}, err
	}
	if err := RejectConflictingWorkRequestTraceFields(data); err != nil {
		return interfaces.WorkRequest{}, err
	}
	return request, nil
}

// RejectRetiredWorkRequestFieldAliases rejects retired public submit fields
// that should no longer be accepted on canonical work-request JSON inputs.
func RejectRetiredWorkRequestFieldAliases(data []byte) error {
	var raw struct {
		WorkTypeID json.RawMessage `json:"work_type_id"`
		Works      []struct {
			WorkTypeID  json.RawMessage `json:"work_type_id"`
			TargetState json.RawMessage `json:"target_state"`
		} `json:"works"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("parse work request retired fields: %w", err)
	}
	if raw.WorkTypeID != nil {
		return fmt.Errorf("work request batch uses retired work_type_id field; use work_type_name")
	}
	for i := range raw.Works {
		if raw.Works[i].WorkTypeID != nil {
			return fmt.Errorf("work request batch works[%d] uses retired work_type_id field; use work_type_name", i)
		}
		if raw.Works[i].TargetState != nil {
			return fmt.Errorf("work request batch works[%d] uses retired target_state field; use state", i)
		}
	}
	return nil
}

func RejectConflictingWorkRequestTraceFields(data []byte) error {
	var raw struct {
		Works []struct {
			CurrentChainingTraceID json.RawMessage `json:"current_chaining_trace_id"`
			TraceID                json.RawMessage `json:"trace_id"`
		} `json:"works"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("parse work request chaining traces: %w", err)
	}
	for i := range raw.Works {
		if raw.Works[i].CurrentChainingTraceID == nil || raw.Works[i].TraceID == nil {
			continue
		}
		var current string
		if err := json.Unmarshal(raw.Works[i].CurrentChainingTraceID, &current); err != nil {
			return fmt.Errorf("parse work request works[%d] current chaining trace: %w", i, err)
		}
		var legacy string
		if err := json.Unmarshal(raw.Works[i].TraceID, &legacy); err != nil {
			return fmt.Errorf("parse work request works[%d] trace_id: %w", i, err)
		}
		if current != "" && legacy != "" && current != legacy {
			return fmt.Errorf("work request batch works[%d] current_chaining_trace_id and trace_id must match when both are provided", i)
		}
	}
	return nil
}
