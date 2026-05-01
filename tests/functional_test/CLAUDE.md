The functional tests are intended to test the system's overall behavior.

The intent is that we mock at the very edges of the system.
i.e. we mock out the os.exec/ we mock out script execution, etc.

If there are no hooks available to mock out at the edges, you should add them.

## Test Mocking Strategy: MockProvider vs MockWorker vs Custom Executor

### 1. WithProvider + MockProvider (preferred for most tests)

Use when the test intent is to verify **business logic above the executor boundary** — token routing, stop-word evaluation, multi-output splitting, persistence, etc.

```go
provider := testutil.NewMockProvider(
    workers.InferenceResponse{Content: "Done. COMPLETE"},  // worker accepts (stop_token match)
)
h := testutil.NewServiceTestHarness(t, dir,
    testutil.WithProvider(provider),
    testutil.WithFullWorkerPoolAndScriptWrap(),
)
```

- Exercises the full pipeline: Provider → AgentExecutor → stop_token evaluation → WorkstationExecutor → TransitionerSubsystem
- For **accept**: content contains the worker's stop_token (usually "COMPLETE")
- For **reject/fail**: content does NOT contain the stop_token
- Provider is shared across all workers — queue enough responses for each call
- Use `provider.CallCount()`, `provider.Calls()`, `provider.LastCall()` for verification

### 1b. WithProviderCommandRunner + full ScriptWrapProvider (for inference-provider behavior)

Use when the test intent is to verify the provider CLI contract itself: selected dispatcher, computed arguments, env merging, stdin usage, stdout/stderr capture, or exit-code handling.

```go
runner := testutil.NewProviderCommandRunner(
    workers.CommandResult{Stdout: []byte("Done. COMPLETE")},
)
h := testutil.NewServiceTestHarness(t, dir,
    testutil.WithProviderCommandRunner(runner),
    testutil.WithFullWorkerPoolAndScriptWrap(),
)
```

- Exercises the factory-facing path through `WorkstationExecutor -> AgentExecutor -> ScriptWrapProvider -> provider command runner`
- Preferred for tests whose value comes from the real Claude/Codex command construction
- Do not migrate unrelated workflow tests just to remove `MockProvider`; custom-executor and mock-provider coverage outside inference-provider behavior is intentionally deferred

### 2. MockWorker (for executor-level routing tests)

Use when the test needs **specific WorkResult fields** that MockProvider cannot produce:

```go
h.MockWorker("worker-name",
    workers.WorkResult{Outcome: workers.OutcomeFailed, Error: "specific error message"},
)
```

Appropriate when tests use:
- `Error` field — specific error messages for failure routing (cascading_failure, conflict_resolution, failed_immutability)
- `SpawnedWork` field — parent-child token spawning (multi_input_guard, multichannel_guard)
- `Feedback` field — rejection feedback propagation (review_retry_exhaustion, rejection_no_arcs)
- `OutputTokens` field — custom token color manipulation
- Mock call inspection (`mock.LastCall()`, `mock.Calls()`) — dispatch field verification (executor_context)

### 3. SetCustomExecutor (for timing/synchronization tests)

Use when tests need **channel-based blocking, sleep-based timing, or mid-execution state inspection**:

```go
h.SetCustomExecutor("worker", &blockingExecutor{ch: ch})
```

Remaining custom executors and why MockProvider is insufficient:
- `sleepyExecutor` / `channelExecutor` (dispatch_timing) — requires controlled execution timing
- `blockingExecutor` (dashboard_inflight, runtime_state) — blocks mid-execution for state snapshot
- `barrierMockExecutor` (e2e concurrency) — barrier synchronization across concurrent dispatches
- `support/harness.FanoutParserExecutor` — dynamic SpawnedWork generation based on input
- `capturePayloadExecutor` (logical_move) — payload inspection across logical moves
- `spawningExecutor` (name_propagation) — dynamic child token creation with names
- `multiCapturingExecutor` (ralph_loop) — captures dispatches across iteration loops
- `capturingExecutor` (FACTORY_REQUEST_BATCH) — dispatch inspection for tag/relation verification
