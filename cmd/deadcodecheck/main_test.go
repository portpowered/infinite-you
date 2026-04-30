package main

import "testing"

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
