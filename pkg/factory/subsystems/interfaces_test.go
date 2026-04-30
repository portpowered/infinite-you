package subsystems

import "testing"

func TestActiveRuntimeTickGroupsAreOrdered(t *testing.T) {
	activeRuntimeGroups := []TickGroup{
		CircuitBreaker,
		Dispatcher,
		History,
		Transitioner,
		CascadingFailure,
		TerminationCheck,
	}

	for i := 1; i < len(activeRuntimeGroups); i++ {
		if activeRuntimeGroups[i] <= activeRuntimeGroups[i-1] {
			t.Fatalf("tick group %d = %d, want greater than previous %d", i, activeRuntimeGroups[i], activeRuntimeGroups[i-1])
		}
	}
}
