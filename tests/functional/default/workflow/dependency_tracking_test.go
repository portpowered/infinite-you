package workflow_test

import (
	"testing"
	"time"

	"github.com/portpowered/agent-factory/pkg/interfaces"
	"github.com/portpowered/agent-factory/pkg/testutil"
	functionalharness "github.com/portpowered/agent-factory/tests/functional/support/harness"
)

// TestDependencyTracking_BlocksUntilSatisfied validates that a token with a
// DEPENDS_ON relation is blocked from transitioning until its dependency
// reaches the required state.
func TestDependencyTracking_BlocksUntilSatisfied(t *testing.T) {
	dir := testutil.CopyFixtureDir(t, fixtureDir(t, "dependency_tracking_dir"))

	workIDA := "task-A-work-id"
	testutil.WriteSeedRequest(t, dir, interfaces.SubmitRequest{
		WorkTypeID: "task",
		WorkID:     workIDA,
		Payload:    []byte("task A"),
	})
	testutil.WriteSeedRequest(t, dir, interfaces.SubmitRequest{
		WorkTypeID: "task",
		Payload:    []byte("task B"),
		Relations: []interfaces.Relation{
			{Type: interfaces.RelationDependsOn, TargetWorkID: workIDA, RequiredState: "complete"},
		},
	})

	provider := testutil.NewMockProvider(
		functionalharness.AcceptedProviderResponse(),
		functionalharness.AcceptedProviderResponse(),
		functionalharness.AcceptedProviderResponse(),
		functionalharness.AcceptedProviderResponse(),
	)
	h := testutil.NewServiceTestHarness(t, dir,
		testutil.WithProvider(provider),
		testutil.WithFullWorkerPoolAndScriptWrap(),
	)

	h.RunUntilComplete(t, 10*time.Second)

	h.Assert().
		HasNoTokenInPlace("task:init").
		HasNoTokenInPlace("task:processing").
		PlaceTokenCount("task:complete", 2)

	if got := len(functionalharness.ProviderCallsForWorker(provider, "starter")); got != 2 {
		t.Errorf("expected starter called 2 times, got %d", got)
	}
}

// TestDependencyTracking_NoDepsPassThrough verifies that tokens without
// DEPENDS_ON relations pass through the DependencyGuard without blocking.
func TestDependencyTracking_NoDepsPassThrough(t *testing.T) {
	dir := testutil.CopyFixtureDir(t, fixtureDir(t, "dependency_tracking_simple_dir"))
	testutil.WriteSeedFile(t, dir, "task", []byte("no deps"))

	provider := testutil.NewMockProvider(functionalharness.AcceptedProviderResponse())
	h := testutil.NewServiceTestHarness(t, dir,
		testutil.WithProvider(provider),
		testutil.WithFullWorkerPoolAndScriptWrap(),
	)

	h.RunUntilComplete(t, 5*time.Second)

	h.Assert().
		HasTokenInPlace("task:complete").
		HasNoTokenInPlace("task:init").
		TokenCount(1)
}
