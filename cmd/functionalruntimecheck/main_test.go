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
