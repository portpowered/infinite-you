package replay

import "fmt"

const (
	DivergenceCategoryMissingDispatch    = "missing_dispatch"
	DivergenceCategoryDispatchMismatch   = "dispatch_mismatch"
	DivergenceCategoryUnknownCompletion  = "unknown_completion"
	DivergenceCategorySideEffectMismatch = "side_effect_mismatch"
	DivergenceCategoryConfigMismatch     = "config_mismatch"
)

// DivergenceReport describes a material difference between a replay artifact
// and the behavior observed while replaying it.
type DivergenceReport struct {
	Category        string `json:"category"`
	Tick            int    `json:"tick,omitempty"`
	DispatchID      string `json:"dispatch_id,omitempty"`
	ExpectedEventID string `json:"expected_event_id,omitempty"`
	ObservedEventID string `json:"observed_event_id,omitempty"`
	Expected        string `json:"expected,omitempty"`
	Observed        string `json:"observed,omitempty"`
}

// DivergenceError stops replay instead of allowing execution to continue after
// it no longer represents the recorded run.
type DivergenceError struct {
	Report DivergenceReport
}

func (e *DivergenceError) Error() string {
	if e == nil {
		return "replay divergence"
	}
	report := e.Report
	message := fmt.Sprintf("replay divergence: category=%s", report.Category)
	if report.Tick != 0 {
		message += fmt.Sprintf(" tick=%d", report.Tick)
	}
	if report.DispatchID != "" {
		message += fmt.Sprintf(" dispatch_id=%s", report.DispatchID)
	}
	if report.ExpectedEventID != "" {
		message += fmt.Sprintf(" expected_event_id=%s", report.ExpectedEventID)
	}
	if report.ObservedEventID != "" {
		message += fmt.Sprintf(" observed_event_id=%s", report.ObservedEventID)
	}
	if report.Expected != "" {
		message += fmt.Sprintf(" expected=%q", report.Expected)
	}
	if report.Observed != "" {
		message += fmt.Sprintf(" observed=%q", report.Observed)
	}
	return message
}

type divergenceOption func(*DivergenceReport)

func withExpectedEventID(eventID string) divergenceOption {
	return func(report *DivergenceReport) {
		report.ExpectedEventID = eventID
	}
}

func newDivergenceError(category string, tick int, dispatchID, expected, observed string, opts ...divergenceOption) error {
	report := DivergenceReport{
		Category:   category,
		Tick:       tick,
		DispatchID: dispatchID,
		Expected:   expected,
		Observed:   observed,
	}
	for _, opt := range opts {
		opt(&report)
	}
	return &DivergenceError{Report: report}
}
