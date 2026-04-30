package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScanPublicSurface_FindsCustomerFacingCoupling(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, publicReadmePath, "# Agent Factory\nPort OS should not appear here.\n")
	writeTestFile(t, root, publicDocsDir+"/guides/quickstart.md", "portable docs\n")
	writeTestFile(t, root, publicExamplesDir+"/basic/README.md", "example\n")
	writeTestFile(t, root, publicFactoryDir+"/README.md", "factory\n")
	writeTestFile(t, root, publicInitPath, "package init\n")

	findings, err := scanPublicSurface(root)
	if err != nil {
		t.Fatalf("scanPublicSurface returned error: %v", err)
	}
	if len(findings) != 1 {
		t.Fatalf("findings = %d, want 1", len(findings))
	}
	if findings[0].path != publicReadmePath {
		t.Fatalf("finding path = %q, want %q", findings[0].path, publicReadmePath)
	}
	if findings[0].line != 2 {
		t.Fatalf("finding line = %d, want 2", findings[0].line)
	}
	if findings[0].term != "port os" {
		t.Fatalf("finding term = %q, want %q", findings[0].term, "port os")
	}
}

func TestScanPublicSurface_IgnoresInternalDocs(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, publicReadmePath, "# Agent Factory\n")
	writeTestFile(t, root, publicDocsDir+"/guides/quickstart.md", "portable docs\n")
	writeTestFile(t, root, publicDocsDir+"/development/notes.md", "PortOS inventory note\n")
	writeTestFile(t, root, publicExamplesDir+"/basic/README.md", "example\n")
	writeTestFile(t, root, publicFactoryDir+"/README.md", "factory\n")
	writeTestFile(t, root, publicInitPath, "package init\n")

	findings, err := scanPublicSurface(root)
	if err != nil {
		t.Fatalf("scanPublicSurface returned error: %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("findings = %d, want 0", len(findings))
	}
}

func writeTestFile(t *testing.T, root string, relativePath string, content string) {
	t.Helper()

	path := filepath.Join(root, filepath.FromSlash(relativePath))
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", relativePath, err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write %s: %v", relativePath, err)
	}
}
