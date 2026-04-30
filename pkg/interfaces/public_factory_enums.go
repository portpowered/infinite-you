package interfaces

import (
	"strings"

	factoryapi "github.com/portpowered/agent-factory/pkg/api/generated"
)

var publicFactoryWorkerTypeAliases = map[string]string{
	WorkerTypeModel:  WorkerTypeModel,
	WorkerTypeScript: WorkerTypeScript,
}

var publicFactoryWorkerModelProviderAliases = map[string]string{
	"ANTHROPIC": publicFactoryWorkerModelProviderClaude,
	"CLAUDE":    publicFactoryWorkerModelProviderClaude,
	"CODEX":     publicFactoryWorkerModelProviderCodex,
	"OPENAI":    publicFactoryWorkerModelProviderCodex,
	"anthropic": publicFactoryWorkerModelProviderClaude,
	"claude":    publicFactoryWorkerModelProviderClaude,
	"codex":     publicFactoryWorkerModelProviderCodex,
	"openai":    publicFactoryWorkerModelProviderCodex,
}

var publicFactoryWorkerProviderAliases = map[string]string{
	"ANTHROPIC":    publicFactoryWorkerProviderScriptWrap,
	"CLAUDE":       publicFactoryWorkerProviderScriptWrap,
	"CLAUDE_CLI":   publicFactoryWorkerProviderScriptWrap,
	"CODEX_CLI":    publicFactoryWorkerProviderScriptWrap,
	"LOCAL":        publicFactoryWorkerProviderScriptWrap,
	"LOCAL_CLAUDE": publicFactoryWorkerProviderScriptWrap,
	"SCRIPT":       publicFactoryWorkerProviderScriptWrap,
	"SCRIPTWRAP":   publicFactoryWorkerProviderScriptWrap,
	"SCRIPT_WRAP":  publicFactoryWorkerProviderScriptWrap,
	"anthropic":    publicFactoryWorkerProviderScriptWrap,
	"claude":       publicFactoryWorkerProviderScriptWrap,
	"claude_cli":   publicFactoryWorkerProviderScriptWrap,
	"claude-cli":   publicFactoryWorkerProviderScriptWrap,
	"codex_cli":    publicFactoryWorkerProviderScriptWrap,
	"codex-cli":    publicFactoryWorkerProviderScriptWrap,
	"local":        publicFactoryWorkerProviderScriptWrap,
	"local_claude": publicFactoryWorkerProviderScriptWrap,
	"local-claude": publicFactoryWorkerProviderScriptWrap,
	"script":       publicFactoryWorkerProviderScriptWrap,
	"scriptwrap":   publicFactoryWorkerProviderScriptWrap,
	"script_wrap":  publicFactoryWorkerProviderScriptWrap,
	"script-wrap":  publicFactoryWorkerProviderScriptWrap,
}

var publicFactoryWorkstationTypeAliases = map[string]string{
	WorkstationTypeLogical: WorkstationTypeLogical,
	WorkstationTypeModel:   WorkstationTypeModel,
}

const (
	publicFactoryWorkerModelProviderClaude = "claude"
	publicFactoryWorkerModelProviderCodex  = "codex"
	publicFactoryWorkerProviderScriptWrap  = "script_wrap"
)

func canonicalPublicFactoryEnumValue(value string, aliases map[string]string) string {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return ""
	}
	if canonical, ok := aliases[trimmed]; ok {
		return canonical
	}
	return trimmed
}

// CanonicalPublicFactoryWorkerType returns the canonical public worker type.
func CanonicalPublicFactoryWorkerType(value string) string {
	return canonicalPublicFactoryEnumValue(value, publicFactoryWorkerTypeAliases)
}

// CanonicalPublicFactoryWorkerModelProvider returns the canonical public worker model provider.
func CanonicalPublicFactoryWorkerModelProvider(value string) string {
	return canonicalPublicFactoryEnumValue(value, publicFactoryWorkerModelProviderAliases)
}

// CanonicalPublicFactoryWorkerProvider returns the canonical public worker provider.
func CanonicalPublicFactoryWorkerProvider(value string) string {
	return canonicalPublicFactoryEnumValue(value, publicFactoryWorkerProviderAliases)
}

// CanonicalPublicFactoryWorkstationType returns the canonical public workstation type.
func CanonicalPublicFactoryWorkstationType(value string) string {
	return canonicalPublicFactoryEnumValue(value, publicFactoryWorkstationTypeAliases)
}

// GeneratedPublicFactoryWorkerType returns the generated worker type enum.
func GeneratedPublicFactoryWorkerType(value string) factoryapi.WorkerType {
	return factoryapi.WorkerType(CanonicalPublicFactoryWorkerType(value))
}

// GeneratedPublicFactoryWorkerTypePtr returns the generated worker type enum when non-empty.
func GeneratedPublicFactoryWorkerTypePtr(value string) *factoryapi.WorkerType {
	if strings.TrimSpace(value) == "" {
		return nil
	}
	enumValue := GeneratedPublicFactoryWorkerType(value)
	return &enumValue
}

// GeneratedPublicFactoryWorkerModelProvider returns the generated worker model provider enum.
func GeneratedPublicFactoryWorkerModelProvider(value string) factoryapi.WorkerModelProvider {
	return factoryapi.WorkerModelProvider(CanonicalPublicFactoryWorkerModelProvider(value))
}

// GeneratedPublicFactoryWorkerModelProviderPtr returns the generated worker model provider enum when non-empty.
func GeneratedPublicFactoryWorkerModelProviderPtr(value string) *factoryapi.WorkerModelProvider {
	if strings.TrimSpace(value) == "" {
		return nil
	}
	enumValue := GeneratedPublicFactoryWorkerModelProvider(value)
	return &enumValue
}

// GeneratedPublicFactoryWorkerProvider returns the generated worker provider enum.
func GeneratedPublicFactoryWorkerProvider(value string) factoryapi.WorkerProvider {
	return factoryapi.WorkerProvider(CanonicalPublicFactoryWorkerProvider(value))
}

// GeneratedPublicFactoryWorkerProviderPtr returns the generated worker provider enum when non-empty.
func GeneratedPublicFactoryWorkerProviderPtr(value string) *factoryapi.WorkerProvider {
	if strings.TrimSpace(value) == "" {
		return nil
	}
	enumValue := GeneratedPublicFactoryWorkerProvider(value)
	return &enumValue
}

// GeneratedPublicFactoryWorkstationType returns the generated workstation type enum.
func GeneratedPublicFactoryWorkstationType(value string) factoryapi.WorkstationType {
	return factoryapi.WorkstationType(CanonicalPublicFactoryWorkstationType(value))
}

// GeneratedPublicFactoryWorkstationTypePtr returns the generated workstation type enum when non-empty.
func GeneratedPublicFactoryWorkstationTypePtr(value string) *factoryapi.WorkstationType {
	if strings.TrimSpace(value) == "" {
		return nil
	}
	enumValue := GeneratedPublicFactoryWorkstationType(value)
	return &enumValue
}
