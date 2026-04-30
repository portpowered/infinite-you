package functional_test

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReferenceDocsSurface_PackageIndexLinksLandingPageAndEveryStableTopic(t *testing.T) {
	readme := readAgentFactoryDoc(t, "docs/README.md")
	referenceIndex := readAgentFactoryDoc(t, "docs/reference/README.md")

	for _, rel := range []string{
		"reference/README.md",
		"reference/config.md",
		"reference/workstations.md",
		"reference/workers.md",
		"reference/resources.md",
		"reference/batch-work.md",
		"reference/templates.md",
	} {
		if !strings.Contains(readme, "("+rel+")") {
			t.Fatalf("docs/README.md missing reference-surface link %q", rel)
		}
	}

	for _, rel := range []string{
		"config.md",
		"workstations.md",
		"workers.md",
		"resources.md",
		"batch-work.md",
		"templates.md",
	} {
		if !strings.Contains(referenceIndex, "("+rel+")") {
			t.Fatalf("docs/reference/README.md missing topic link %q", rel)
		}
	}
}

func TestReferenceDocsSurface_StableTopicFilesExistExactlyOnceInReferenceSurface(t *testing.T) {
	referenceDir := agentFactoryPath(t, "docs/reference")
	counts := map[string]int{
		"config.md":       0,
		"workstations.md": 0,
		"workers.md":      0,
		"resources.md":    0,
		"batch-work.md":   0,
		"templates.md":    0,
	}

	err := filepath.WalkDir(referenceDir, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			return nil
		}
		name := filepath.Base(path)
		if _, ok := counts[name]; ok {
			counts[name]++
		}
		return nil
	})
	if err != nil {
		t.Fatalf("walk reference docs: %v", err)
	}

	for name, count := range counts {
		if count != 1 {
			t.Fatalf("%s count = %d, want 1", name, count)
		}
	}
}

func TestReferenceDocsSurface_TopicPagesKeepDeeperPackageGuideLinks(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		wantLinks  []string
	}{
		{
			name: "config",
			path: "docs/reference/config.md",
			wantLinks: []string{
				"(README.md)",
				"(../README.md)",
				"(../work.md)",
				"(../authoring-agents-md.md)",
			},
		},
		{
			name: "workstations",
			path: "docs/reference/workstations.md",
			wantLinks: []string{
				"(README.md)",
				"(../README.md)",
				"(../workstations.md)",
				"(../guides/workstation-guards-and-guarded-loop-breakers.md)",
			},
		},
		{
			name: "workers",
			path: "docs/reference/workers.md",
			wantLinks: []string{
				"(README.md)",
				"(../README.md)",
				"(../authoring-agents-md.md)",
				"(../workstations.md)",
			},
		},
		{
			name: "resources",
			path: "docs/reference/resources.md",
			wantLinks: []string{
				"(README.md)",
				"(../README.md)",
				"(../work.md)",
				"(../guides/workstation-guards-and-guarded-loop-breakers.md)",
			},
		},
		{
			name: "batch-work",
			path: "docs/reference/batch-work.md",
			wantLinks: []string{
				"(README.md)",
				"(../README.md)",
				"(../work.md)",
				"(../guides/batch-inputs.md)",
			},
		},
		{
			name: "templates",
			path: "docs/reference/templates.md",
			wantLinks: []string{
				"(README.md)",
				"(../README.md)",
				"(../prompt-variables.md)",
				"(../authoring-agents-md.md)",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content := readAgentFactoryDoc(t, tt.path)
			if !strings.Contains(content, "## Related") {
				t.Fatalf("%s missing Related section", tt.path)
			}
			for _, link := range tt.wantLinks {
				if !strings.Contains(content, link) {
					t.Fatalf("%s missing related link %q", tt.path, link)
				}
			}
		})
	}
}

func readAgentFactoryDoc(t *testing.T, rel string) string {
	t.Helper()

	data, err := os.ReadFile(agentFactoryPath(t, rel))
	if err != nil {
		t.Fatalf("read %s: %v", rel, err)
	}
	return string(data)
}
