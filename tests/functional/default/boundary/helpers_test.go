package boundary_test

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func fixtureDir(t *testing.T, name string) string {
	t.Helper()

	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("cannot determine test file path")
	}

	baseDir := filepath.Dir(thisFile)
	for _, root := range []string{
		filepath.Join(baseDir, "testdata"),
		filepath.Join(baseDir, "..", "..", "..", "functional_test", "testdata"),
	} {
		candidate := filepath.Join(root, name)
		if stat, err := os.Stat(candidate); err == nil && stat.IsDir() {
			return candidate
		}
	}
	return filepath.Join(baseDir, "testdata", name)
}
