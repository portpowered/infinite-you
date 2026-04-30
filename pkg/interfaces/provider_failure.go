package interfaces

// ProviderErrorFamily captures the runtime behavior category for a normalized
// provider failure.
type ProviderErrorFamily string

const (
	ProviderErrorFamilyTerminal  ProviderErrorFamily = "terminal"
	ProviderErrorFamilyRetryable ProviderErrorFamily = "retryable"
	ProviderErrorFamilyThrottle  ProviderErrorFamily = "throttle"
)

// ProviderErrorType is the stable customer-facing normalized error type for
// inference-backed worker failures.
type ProviderErrorType string

const (
	ProviderErrorTypeAuthFailure         ProviderErrorType = "auth_failure"
	ProviderErrorTypePermanentBadRequest ProviderErrorType = "permanent_bad_request"
	ProviderErrorTypeThrottled           ProviderErrorType = "throttled"
	ProviderErrorTypeInternalServerError ProviderErrorType = "internal_server_error"
	ProviderErrorTypeTimeout             ProviderErrorType = "timeout"
	ProviderErrorTypeUnknown             ProviderErrorType = "unknown"
	ProviderErrorTypeMisconfigured       ProviderErrorType = "misconfigured"
)

// ProviderFailureDecision is the normalized behavior contract consumed by
// downstream retry, termination, and throttle-pause logic.
type ProviderFailureDecision struct {
	Retryable             bool
	Terminal              bool
	TriggersThrottlePause bool
}

// ProviderFailureMetadata carries the normalized provider-failure contract
// across runtime boundaries after the original error has been rendered.
type ProviderFailureMetadata struct {
	Family ProviderErrorFamily `json:"family"`
	Type   ProviderErrorType   `json:"type"`
}
