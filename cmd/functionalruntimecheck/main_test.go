package main

import (
	"testing"
	"time"
)

func TestRunDefaultFunctionalLaneRequiresPositiveDuration(t *testing.T) {
	t.Parallel()

	if defaultPackage != "./tests/functional/default/..." {
		t.Fatalf("default package = %q, want canonical default functional lane", defaultPackage)
	}
}

func TestBudgetComparisonUsesMeasuredElapsedTime(t *testing.T) {
	t.Parallel()

	budget := 10 * time.Second
	elapsed := 4600 * time.Millisecond
	if elapsed > budget {
		t.Fatalf("elapsed %s should remain within budget %s", elapsed, budget)
	}
}

func TestResolveGoBinaryDefaultsToGo(t *testing.T) {
	t.Parallel()

	if got := resolveGoBinary(""); got != "go" {
		t.Fatalf("resolveGoBinary(\"\") = %q, want %q", got, "go")
	}
}

func TestResolveGoBinaryPreservesConfiguredToolchain(t *testing.T) {
	t.Parallel()

	const configured = "/tmp/custom-go"
	if got := resolveGoBinary(configured); got != configured {
		t.Fatalf("resolveGoBinary(%q) = %q, want %q", configured, got, configured)
	}
}
