package handwrittensourceguard

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestTargetedBroadContractGuardsUseSharedHandwrittenSourceOwner(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		path      string
		snippets  []string
		forbidden []string
	}{
		{
			name: "api guard uses shared repo-root policy",
			path: filepath.Join("..", "..", "pkg", "api", "legacy_model_guard_test.go"),
			snippets: []string{
				`handwrittensourceguard.ShouldSkipDir("pkg/api/legacy_model_guard_test.go", moduleRoot, path)`,
			},
		},
		{
			name: "petri guard uses shared repo-root policy",
			path: filepath.Join("..", "..", "pkg", "petri", "transition_contract_guard_test.go"),
			snippets: []string{
				`handwrittensourceguard.ShouldSkipDir("pkg/petri/transition_contract_guard_test.go", moduleRoot, path)`,
			},
		},
		{
			name: "config guard uses shared pkg-root policy",
			path: filepath.Join("..", "..", "pkg", "config", "exhaustion_rule_contract_guard_test.go"),
			snippets: []string{
				`handwrittensourceguard.ShouldSkipDir("pkg/config/exhaustion_rule_contract_guard_test.go", pkgRoot, path)`,
			},
			forbidden: []string{
				"contractguard.ShouldSkipDir(",
			},
		},
		{
			name: "interfaces boundary scan uses shared boundary policy",
			path: filepath.Join("..", "..", "pkg", "interfaces", "world_view_contract_guard_test.go"),
			snippets: []string{
				`handwrittensourceguard.ShouldSkipDir("pkg/interfaces/world_view_contract_guard_test.go#boundary", ".", path)`,
				`handwrittensourceguard.ShouldSkipDir("pkg/interfaces/world_view_contract_guard_test.go#canonical", root, path)`,
			},
			forbidden: []string{
				"contractguard.ShouldSkipDir(",
			},
		},
		{
			name: "interfaces runtime lookup uses shared pkg-root policy",
			path: filepath.Join("..", "..", "pkg", "interfaces", "runtime_lookup_contract_guard_test.go"),
			snippets: []string{
				`handwrittensourceguard.ShouldSkipDir("pkg/interfaces/runtime_lookup_contract_guard_test.go", root, path)`,
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			data, err := os.ReadFile(tc.path)
			if err != nil {
				t.Fatalf("read %s: %v", tc.path, err)
			}
			source := string(data)

			for _, snippet := range tc.snippets {
				if !strings.Contains(source, snippet) {
					t.Fatalf("%s no longer contains required shared-owner snippet %q", tc.path, snippet)
				}
			}
			for _, snippet := range tc.forbidden {
				if strings.Contains(source, snippet) {
					t.Fatalf("%s still contains stale non-canonical owner snippet %q", tc.path, snippet)
				}
			}
		})
	}
}
