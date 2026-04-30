package functional_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"testing"
)

const fullWorkerPoolGuardrailMessage = "Agent Factory functional tests should build ServiceTestHarness with testutil.WithFullWorkerPoolAndScriptWrap(); mock at provider, provider command-runner, command-runner, or mock-worker command boundaries when possible"

type harnessShortcutUse struct {
	TestName string
	File     string
	Line     int
}

type harnessShortcutException struct {
	Count  int
	Reason string
}

var fullWorkerPoolGuardrailExceptions = map[string]harnessShortcutException{
	"TestConcurrencyLimit_ResourceReleasedOnExecutorPanic_Inline":    {Count: 1, Reason: "Intentionally verifies inline WorkerExecutor panic recovery, which edge mocks cannot trigger."},
	"TestConcurrencyLimit_ResourceReleasedOnExecutorPanic_Async":     {Count: 1, Reason: "Intentionally verifies async WorkerExecutor panic recovery, which edge mocks cannot trigger."},
	"TestDashboard_InFlightDispatches":                               {Count: 1, Reason: "Blocks a custom executor mid-dispatch to inspect in-flight dashboard state."},
	"TestDashboard_SingleWorkItemSnapshot":                           {Count: 1, Reason: "Uses a snapshot executor to inspect runtime/dashboard state during controlled execution."},
	"TestDashboard_ParallelWorkItemsSnapshot":                        {Count: 1, Reason: "Uses a barrier executor to inspect parallel in-flight snapshot timing."},
	"TestDispatchTiming_HistoryRecordsDuration":                      {Count: 1, Reason: "Uses a sleeping custom executor to assert dispatch duration accounting."},
	"TestDispatchTiming_InFlightStartTime":                           {Count: 1, Reason: "Uses a channel-blocking custom executor to assert in-flight start-time snapshots."},
	"TestExecutorContext_RejectionFeedback":                          {Count: 1, Reason: "Inspects rejection feedback propagation through MockWorker WorkResult fields."},
	"TestLogicalMove_Success":                                        {Count: 1, Reason: "Uses custom executors to inspect logical move payload routing."},
	"TestLogicalMove_PreservesTokenColor":                            {Count: 1, Reason: "Uses custom executors to inspect logical move color preservation."},
	"TestMultiInputGuard_PartialChildren_BlocksGuard":                {Count: 1, Reason: "Uses custom spawned-work executors to isolate multi-input guard behavior."},
	"TestMultiInputGuard_AllChildrenComplete":                        {Count: 1, Reason: "Uses custom spawned-work executors to isolate all-children guard release."},
	"TestMultiInputGuard_IndependentChapters":                        {Count: 1, Reason: "Uses custom spawned-work executors to isolate independent chapter guards."},
	"TestMultiInputGuard_IndependentChapters_StaggeredCompletion":    {Count: 1, Reason: "Uses custom spawned-work executors to isolate staggered guard release."},
	"TestMultiOutput_NoStopWordsConfigured":                          {Count: 1, Reason: "Uses MockWorker output tokens to isolate multi-output behavior without stop-word semantics."},
	"TestMultiChannelGuard_FileDropToCompletion":                     {Count: 1, Reason: "Uses custom spawned-work executors to isolate multichannel guard completion."},
	"TestMultiChannelGuard_GuardBlocksUntilAllPagesComplete":         {Count: 1, Reason: "Uses custom spawned-work executors to isolate multichannel guard blocking."},
	"TestRepeater_RefiresOnRejectedStopsOnAccepted":                  {Count: 1, Reason: "Uses MockWorker ordered outcomes to isolate repeater retry routing."},
	"TestRepeater_GuardedLoopBreakerTerminatesRejectedRepeater":      {Count: 1, Reason: "Uses MockWorker ordered outcomes to isolate repeater guarded loop-breaker routing."},
	"TestRepeater_YieldsBetweenIterations":                           {Count: 1, Reason: "Uses MockWorker ordered outcomes to inspect repeater yield behavior."},
	"TestParameterizedFields_WorkingDirectoryResolvesFromTags":       {Count: 1, Reason: "Captures raw dispatch fields to assert parameterized working-directory resolution."},
	"TestRepeater_ResourceReleaseBetweenIterations":                  {Count: 1, Reason: "Uses MockWorker ordered outcomes to isolate resource release between repeater iterations."},
	"TestReviewRetryLoopBreaker_FeedbackPropagated":                  {Count: 1, Reason: "Uses MockWorker feedback fields to assert guarded retry-loop feedback propagation."},
	"TestRuntimeState_ThreeStagePipeline":                            {Count: 1, Reason: "Uses sleeping custom executors to inspect runtime state across staged dispatches."},
	"TestRuntimeState_MidExecutionConsistency":                       {Count: 1, Reason: "Uses a blocking custom executor to inspect mid-execution runtime consistency."},
	"TestServiceHarness_MockWorker":                                  {Count: 1, Reason: "Harness contract test for default MockWorker behavior."},
	"TestServiceHarness_MockWorker_Idempotent":                       {Count: 1, Reason: "Harness contract test for MockWorker idempotency."},
	"TestServiceHarness_SetCustomExecutor":                           {Count: 1, Reason: "Harness contract test for SetCustomExecutor behavior."},
	"TestServiceHarness_CustomExecutor_Precedence":                   {Count: 1, Reason: "Harness contract test for custom executor precedence over MockWorker."},
	"TestStatelessExecution_ThinDispatchCarriesLookupReferencesOnly": {Count: 1, Reason: "Captures raw dispatch fields to assert thin stateless execution contracts."},
	"TestWorkflowModificationRejectionLoop":                          {Count: 1, Reason: "Uses MockWorker rejection loop outcomes to isolate workflow modification routing."},
	"TestFactoryRequestBatch_TagsAccessibleInTokenPayload":           {Count: 1, Reason: "Captures raw dispatch payload to assert request-batch tag accessibility."},
}

func TestFunctionalTestsUseFullWorkerPoolHarnessOrDocumentException(t *testing.T) {
	t.Parallel()

	uses, err := findServiceHarnessShortcuts(".")
	if err != nil {
		t.Fatalf("scan functional tests: %v", err)
	}

	actualCounts := make(map[string]int)
	var unexpected []harnessShortcutUse
	for _, use := range uses {
		actualCounts[use.TestName]++
		if _, ok := fullWorkerPoolGuardrailExceptions[use.TestName]; !ok {
			unexpected = append(unexpected, use)
		}
	}

	for testName, exception := range fullWorkerPoolGuardrailExceptions {
		if strings.TrimSpace(exception.Reason) == "" {
			t.Errorf("guardrail exception for %s must include a reviewable reason", testName)
		}
		if got := actualCounts[testName]; got != exception.Count {
			t.Errorf("guardrail exception count drift for %s: expected %d shortcut harness call(s), got %d", testName, exception.Count, got)
		}
	}

	if len(unexpected) > 0 {
		sort.Slice(unexpected, func(i, j int) bool {
			if unexpected[i].File == unexpected[j].File {
				return unexpected[i].Line < unexpected[j].Line
			}
			return unexpected[i].File < unexpected[j].File
		})

		var b strings.Builder
		b.WriteString(fullWorkerPoolGuardrailMessage)
		b.WriteString("\nUnexpected shortcut harness call(s):")
		for _, use := range unexpected {
			b.WriteString("\n- ")
			b.WriteString(use.File)
			b.WriteString(":")
			b.WriteString(strconv.Itoa(use.Line))
			b.WriteString(" in ")
			b.WriteString(use.TestName)
		}
		b.WriteString("\nUse the full worker-pool option or add a narrow exception with a reason and exact shortcut count.")
		t.Fatal(b.String())
	}
}

func TestFindServiceHarnessShortcutsFlagsMissingFullWorkerPoolOption(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	source := `package functional_test

import (
	"testing"
	"github.com/portpowered/agent-factory/pkg/testutil"
)

func TestShortcutHarness(t *testing.T) {
	dir := t.TempDir()
	_ = testutil.NewServiceTestHarness(t, dir)
}

func TestFullWorkerPoolHarness(t *testing.T) {
	dir := t.TempDir()
	_ = testutil.NewServiceTestHarness(t, dir, testutil.WithFullWorkerPoolAndScriptWrap())
}
`
	if err := os.WriteFile(filepath.Join(dir, "sample_test.go"), []byte(source), 0o600); err != nil {
		t.Fatalf("write sample: %v", err)
	}

	uses, err := findServiceHarnessShortcuts(dir)
	if err != nil {
		t.Fatalf("scan sample: %v", err)
	}
	if len(uses) != 1 {
		t.Fatalf("expected one shortcut harness use, got %d: %#v", len(uses), uses)
	}
	if uses[0].TestName != "TestShortcutHarness" {
		t.Fatalf("expected shortcut in TestShortcutHarness, got %s", uses[0].TestName)
	}
}

func findServiceHarnessShortcuts(dir string) ([]harnessShortcutUse, error) {
	fset := token.NewFileSet()
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var uses []harnessShortcutUse
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, "_test.go") {
			continue
		}

		path := filepath.Join(dir, name)
		file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return nil, err
		}

		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Body == nil || !strings.HasPrefix(fn.Name.Name, "Test") {
				continue
			}

			ast.Inspect(fn.Body, func(node ast.Node) bool {
				call, ok := node.(*ast.CallExpr)
				if !ok || !isServiceHarnessConstructor(call) {
					return true
				}
				if callContainsFullWorkerPoolOption(call) {
					return true
				}

				pos := fset.Position(call.Pos())
				uses = append(uses, harnessShortcutUse{
					TestName: fn.Name.Name,
					File:     filepath.Base(pos.Filename),
					Line:     pos.Line,
				})
				return true
			})
		}
	}

	return uses, nil
}

func isServiceHarnessConstructor(call *ast.CallExpr) bool {
	selector, ok := call.Fun.(*ast.SelectorExpr)
	if !ok || selector.Sel.Name != "NewServiceTestHarness" {
		return false
	}
	ident, ok := selector.X.(*ast.Ident)
	return ok && ident.Name == "testutil"
}

func callContainsFullWorkerPoolOption(call *ast.CallExpr) bool {
	found := false
	for _, arg := range call.Args {
		ast.Inspect(arg, func(node ast.Node) bool {
			selector, ok := node.(*ast.SelectorExpr)
			if !ok || selector.Sel.Name != "WithFullWorkerPoolAndScriptWrap" {
				return true
			}
			ident, ok := selector.X.(*ast.Ident)
			if ok && ident.Name == "testutil" {
				found = true
				return false
			}
			return true
		})
	}
	return found
}
