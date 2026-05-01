# Functional Test Suite Inventory

This inventory captures the current `tests/functional_test` surface during the
transition into explicit collections. It records what is still mixed together
in the legacy package, the dominant fixture/support seams, and the broad
package runtime that later cleanup stories should continue shrinking.

## Inventory Date

- 2026-04-30

## Current Surface Summary

- `tests/functional_test` currently contains 104 `_test.go` files.
- The package also contains 18 support files whose names explicitly mark them as
  helpers, fixtures, harnesses, or server scaffolding.
- The shared `tests/functional_test/testdata` tree currently contains 55
  top-level fixture directories.
- The package root also mixes in suite-local process artifacts:
  `CLAUDE.md`, `prd.json`, `prd.md`, and `progress.txt`.

## Support Files Mixed Into The Package

These are the most obvious non-scenario ownership seams currently sharing the
same directory as the behavioral tests:

- `agent_config_helpers_test.go`
- `agent_factory_export_import_fixture_test.go`
- `automat_portability_fixture_test.go`
- `current_factory_activation_fixture_test.go`
- `event_history_dispatch_helpers_test.go`
- `export_import_fixture_test.go`
- `export_import_harness_test.go`
- `functional_harness_guardrail_test.go`
- `functional_server_override_regression_test.go`
- `functional_server_test.go`
- `pipeline_helpers_test.go`
- `provider_error_corpus_helpers_test.go`
- `replay_regression_harness_test.go`
- `service_harness_test.go`
- `short_skip_test.go`
- `token_identity_helpers_test.go`
- `utilities_test.go`
- `work_request_helpers_test.go`

## Dominant Coverage Themes

The current package mixes several distinct coverage themes that would be easier
to own and tune as separate collections:

| Theme | Representative files | Representative fixtures | Notes |
| --- | --- | --- | --- |
| CLI, bootstrap, and customer-facing smoke coverage | `cleanup_smoke_test.go`, `cli_docs_smoke_test.go`, `init_factory_test.go`, `ralph_init_smoke_test.go`, `legacy_unary_retirement_smoke_test.go` | `happy_path`, `ralph_worktree` | Fast repository-surface checks currently live beside much heavier runtime scenarios. |
| Workflow, routing, and engine behavior | `dispatcher_workflow_test.go`, `workflow_modification_test.go`, `dependency_tracking_test.go`, `multi_input_guard_test.go`, `multi_output_test.go`, `repeater_parameterized_test.go` | `dispatcher_workflow`, `workflow_v1_dir`, `workflow_v2_dir`, `multi_input_guard_dir`, `multi_output_dir`, `repeater_workstation` | This is the broadest cluster and mixes ordinary routing regressions with timing-sensitive engine behavior. |
| API, dashboard, and service-mode coverage | `e2e_test.go`, `generated_api_smoke_test.go`, `named_factory_api_test.go`, `dashboard_engine_state_test.go`, `dashboard_snapshot_test.go`, `current_factory_watcher_switch_test.go` | `e2e`, `service_simple`, `submitted_parent_child_filewatcher` | These tests pay for live server or service harness setup but are still stored in the same bucket as pure fixture checks. |
| Replay, export/import, serialization, and portability | `export_import_e2e_smoke_test.go`, `record_replay_end_to_end_test.go`, `replay_regression_harness_test.go`, `event_replay_smoke_test.go`, `factory_only_serialization_smoke_test.go`, `automat_portability_fixture_test.go` | `automat_portability_smoke`, `factory_request_batch`, `service_parameterized_success` | The suite combines artifact-format checks, replay projections, and full round-trip runtime scenarios. |
| Provider, script, and mock-worker execution behavior | `provider_error_smoke_test.go`, `script_executor_test.go`, `mock_workers_smoke_test.go`, `integration_smoke_test.go`, `timeout_cleanup_smoke_test.go` | `script_executor_dir`, `retry_exhaustion`, `review_retry_exhaustion` | These tests are some of the most expensive because they exercise retries, throttling, subprocesses, or timeout cleanup. |
| Fixture, harness, and policy guardrails | `functional_harness_guardrail_test.go`, `service_harness_test.go`, `provider_harness_helpers_test.go`, `short_skip_test.go` | `tags_test`, `same_name_guard_dir` | Support and guardrail seams are already first-class concerns, but they currently share ownership with scenario files instead of sitting in explicit support collections. |

## Major Fixture Groups

The 67 top-level `testdata` directories also break down into clear clusters:

- Pipeline and workflow fixtures:
  `batch_ideation_pipeline`, `full_ideation_pipeline`,
  `serial_ideation_pipeline`, `dispatcher_workflow`, `workflow_v1_dir`,
  `workflow_v2_dir`, `workflow_v2_rejection_dir`.
- Routing and guard fixtures:
  `dependency_tracking_dir`, `dependency_tracking_simple_dir`,
  `multi_input_guard_dir`, `matches_fields_single_input_dir`,
  `matches_fields_pair_guard_dir`, `matches_fields_triple_guard_dir`,
  `same_name_guard_dir`, `tags_test`.
- API, service, and runtime fixtures:
  `e2e`, `service_simple`, `service_parameterized_success`,
  `submitted_parent_child_filewatcher`, `worktree_passthrough`.
- Replay, portability, and export/import fixtures:
  `automat_portability_smoke`, `factory_request_batch`,
  `ralph_worktree`, `script_executor_dir`.
- Small boundary or regression fixtures:
  `invalid_worker_reference`, `logical_move_dir`,
  `logical_move_pipeline_dir`, `noop_pipeline`, `happy_path`.

## Runtime Baseline

The broad legacy-package baseline was captured with:

```powershell
go test ./tests/functional_test -count=1
```

Result on 2026-04-30:

- Package runtime: `75.157s`

The repository default verification lane now uses:

```powershell
make test
```

Result on 2026-04-30:

- `make test-functional-default-budget` completed in `2.554s`
- `make test` now keeps the ordinary developer path out of
  `tests/functional_test` and enforces the `10s` default-lane budget through
  the repository-owned runtime checker
- `make test-functional-extended` remains the canonical opt-in slow lane and
  completed in `25.768s`

## Largest Runtime Contributors

The heaviest top-level tests from the baseline run were:

| Test | Runtime |
| --- | --- |
| `TestFixtureDirectories_Load` | `7.87s` |
| `TestProviderErrorSmoke_ScriptWrapScenariosStayNormalizedAcrossProviders` | `3.48s` |
| `TestProviderErrorSmoke_ThrottlePauseObservabilityFlowsThroughRuntimeSnapshotAndDashboard` | `2.57s` |
| `TestRuntimeConfigAlignmentSmoke_CanonicalOnlyBoundaryStaysAlignedAcrossExecutionAndRejectsRetiredAliases` | `1.96s` |
| `TestIntegrationSmoke_TimeoutCancelsProcessTreeAndClearsActiveExecution` | `1.70s` |
| `TestIntegrationSmoke_TimeoutRequeuesWorkAndSucceedsOnLaterAttempt` | `1.58s` |
| `TestLegacyUnaryRetirementSmoke_CanonicalSubmitPathsStayBatchOnly` | `1.32s` |
| `TestCurrentFactoryActivationFixture_WatchedFileExecutionFollowsActivatedFactory` | `1.17s` |
| `TestCurrentFactoryWatcherSwitchSmoke_ActivatedFactoryOwnsWatchedInputWithoutDuplicateConsumption` | `1.10s` |

Top-level timing distribution from the same run:

- `13` tests took `>=1.0s`
- `29` tests took `0.5s-0.99s`
- `95` tests took `0.2s-0.49s`
- `175` tests took `<0.2s`

## Likely Runtime Drivers

The baseline points to a few repeatable cost drivers:

- Full fixture-directory scans and runtime-config loading make
  `TestFixtureDirectories_Load` the single largest contributor before any live
  workflow assertions run.
- Provider error and retry scenarios are expensive because they combine service
  startup, multi-attempt worker execution, throttling logic, and dashboard or
  observability assertions in the same tests.
- Timeout and current-factory watcher scenarios pay for real subprocess or
  watched-file orchestration, which makes them poor candidates for the default
  fast lane.
- Replay, export/import, and live API/server coverage are valuable but not
  cheap; they are strong opt-in lane candidates once explicit collections
  exist.

## Plausible Fast-Lane Candidates

The current package already contains many scenarios that are likely compatible
with a sub-10-second default lane if they are grouped away from the heavier
surfaces above:

- Fixture and boundary validation tests that stay under roughly `0.2s` to
  `0.4s`, such as archive-terminal, portability layout, generated schema,
  stop-word, and many token-routing regressions.
- Pure CLI/docs/public-surface smoke tests that do not require long-lived
  service orchestration.
- Focused workflow routing regressions that use checked-in fixtures and edge
  mocks without retry loops, watcher orchestration, or timeout cleanup.

## Verification Commands

Use these commands to refresh the inventory after structural changes:

```powershell
go test ./tests/functional_test -count=1
make test-functional-default-budget
make test-functional-extended
make test
Get-ChildItem tests/functional_test -Filter *_test.go | Measure-Object
Get-ChildItem tests/functional_test/testdata -Directory | Measure-Object
```
