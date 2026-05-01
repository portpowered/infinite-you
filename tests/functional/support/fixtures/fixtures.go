package fixtures

import (
	"path/filepath"
	"runtime"
	"testing"
)

// SharedDir returns the absolute path to a shared functional fixture.
func SharedDir(t testing.TB, name string) string {
	t.Helper()

	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("cannot determine shared fixture path")
	}

	return filepath.Join(filepath.Dir(thisFile), "testdata", name)
}
