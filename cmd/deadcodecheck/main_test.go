package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMainRoutesThroughCommandMain(t *testing.T) {
	originalCommandMain := commandMain
	originalExitFunc := exitFunc
	originalStdout := stdout
	originalStderr := stderr
	originalArgs := os.Args
	t.Cleanup(func() {
		commandMain = originalCommandMain
		exitFunc = originalExitFunc
		stdout = originalStdout
		stderr = originalStderr
		os.Args = originalArgs
	})

	var gotArgs []string
	var gotStdout io.Writer
	var gotStderr io.Writer
	var exitCode int
	out := &bytes.Buffer{}
	errOut := &bytes.Buffer{}
	commandMain = func(args []string, stdout io.Writer, stderr io.Writer) int {
		gotArgs = append([]string(nil), args...)
		gotStdout = stdout
		gotStderr = stderr
		return 17
	}
	exitFunc = func(code int) {
		exitCode = code
	}
	stdout = out
	stderr = errOut
	os.Args = []string{"deadcodecheck", "-example"}

	main()

	if exitCode != 17 {
		t.Fatalf("main() exit code = %d, want 17", exitCode)
	}
	if len(gotArgs) != 1 || gotArgs[0] != "-example" {
		t.Fatalf("main() args = %v, want [-example]", gotArgs)
	}
	if gotStdout != out {
		t.Fatal("main() stdout writer mismatch")
	}
	if gotStderr != errOut {
		t.Fatal("main() stderr writer mismatch")
	}
}

func TestEnsureGoTypesAliasEnabled(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "empty", in: "", want: "gotypesalias=1"},
		{name: "preserves other flags", in: "gocachehash=1", want: "gocachehash=1,gotypesalias=1"},
		{name: "replaces disabled flag", in: "gotypesalias=0", want: "gotypesalias=1"},
		{name: "preserves flag order", in: "gocachehash=1,gotypesalias=0,inittrace=1", want: "gocachehash=1,gotypesalias=1,inittrace=1"},
		{name: "leaves enabled flag", in: "gotypesalias=1", want: "gotypesalias=1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ensureGoTypesAliasEnabled(tt.in); got != tt.want {
				t.Fatalf("ensureGoTypesAliasEnabled(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
