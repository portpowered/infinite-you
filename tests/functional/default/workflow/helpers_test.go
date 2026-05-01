package workflow_test

import (
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
	return filepath.Join(filepath.Dir(thisFile), "testdata", name)
}

func skipSlowFunctionalSmokeInShort(t *testing.T, reason string) {
	t.Helper()
	if testing.Short() {
		t.Skip(reason)
	}
}
