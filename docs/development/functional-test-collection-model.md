# Functional Test Collection Model

This document defines the target structure for splitting the current
`tests/functional_test` package into explicit functional-test collections with
clear ownership and execution lanes.

It is the design contract for the cleanup stories that follow the inventory in
[Functional Test Suite Inventory](functional-test-suite-inventory.md). The
structure described here is the target state; later stories perform the
physical moves and command updates.

## Goals

- Replace the single `tests/functional_test` bucket with explicit collections.
- Keep the default functional lane focused on fast regression coverage.
- Preserve slower high-value scenarios behind opt-in commands.
- Keep helpers and `testdata` owned by the collection that uses them.
- Make lane membership obvious from repository paths and `Makefile` targets.

## Boundary Rules

- The primary enforcement boundary is the Go package directory, not
  `testing.Short()` checks inside one broad package.
- Each collection becomes its own directory under `tests/functional/`.
- Each directory is a standalone Go test package with its own local helpers and
  fixtures.
- `Makefile` owns the canonical commands that define which collections run by
  default and which run only on opt-in paths.
- New scenarios should join an explicit collection directory; they should not
  land in a revived mixed package.

## Target Directory Layout

The functional suite should move toward this structure:

```text
tests/
  functional/
    default/
      smoke/
      workflow/
      boundary/
    extended/
      service/
      replay/
      provider/
    support/
      harness/
      fixtures/
      assertions/
```

## Collection Definitions

### Default Lane Collections

These collections are intended to stay in the everyday regression path and
converge on the sub-10-second budget together.

| Collection | Scope | Expected traits | Current examples to migrate |
| --- | --- | --- | --- |
| `tests/functional/default/smoke` | CLI/bootstrap/docs/public-surface smokes | Small checked-in fixtures, little or no long-lived service orchestration | `cleanup_smoke_test.go`, `cli_docs_smoke_test.go`, `init_factory_test.go`, `ralph_init_smoke_test.go` |
| `tests/functional/default/workflow` | Core routing, token flow, guard, and ordinary workflow behavior | Fixture-driven engine assertions, provider edge mocks, no watcher or timeout orchestration | `dependency_tracking_test.go`, `dispatcher_workflow_test.go`, `multi_output_test.go`, `workflow_modification_test.go` |
| `tests/functional/default/boundary` | Fast serialization, schema, and contract regressions that still belong in functional coverage | Read-model or contract validation with bounded fixture cost | `generated_schema_deserialization_smoke_test.go`, `archive_terminal_test.go`, `workstation_stopwords_test.go` |

### Opt-In Extended Collections

These collections keep meaningful coverage, but they should not define the
default functional runtime budget.

| Collection | Scope | Why opt-in | Current examples to migrate |
| --- | --- | --- | --- |
| `tests/functional/extended/service` | Live server, dashboard, file-watcher, timeout, current-factory, and subprocess-heavy scenarios | Real service startup, watcher ownership, or timeout cleanup dominate runtime | `e2e_test.go`, `integration_smoke_test.go`, `current_factory_watcher_switch_test.go`, `timeout_cleanup_smoke_test.go` |
| `tests/functional/extended/replay` | Replay, export/import, portability, and end-to-end artifact roundtrip coverage | Artifact roundtrips and replay hydration are valuable but not routine fast-lane checks | `record_replay_end_to_end_test.go`, `export_import_e2e_smoke_test.go`, `event_replay_smoke_test.go`, `factory_only_serialization_smoke_test.go` |
| `tests/functional/extended/provider` | Retry, throttle, provider normalization, and other high-cost worker/runtime scenarios | Repeated attempts, throttling, and observability assertions are among the slowest surfaces in the inventory | `provider_error_smoke_test.go`, `review_retry_exhaustion_test.go`, `script_executor_test.go` |

### Shared Support Packages

Support code should not re-form another mixed grab-bag package. Shared helpers
move only when they are genuinely cross-collection seams.

| Support package | Ownership rule |
| --- | --- |
| `tests/functional/support/harness` | Harness builders and service-start helpers used by multiple collections |
| `tests/functional/support/fixtures` | Cross-collection fixture loading helpers only when the same helper is reused in multiple collections |
| `tests/functional/support/assertions` | Shared assertion helpers with stable domain meaning across collections |

Keep the shared seam narrow. The current extracted support surface uses
`tests/functional/support/harness` for provider-response helpers plus shared
token and executor helpers that already span multiple scenario files.
Collection-specific helpers should continue to stay beside the scenarios that
own them.

Anything used by only one collection stays local to that collection directory,
even if it is a helper file.

## Helper And Fixture Ownership Rules

- Collection-local helpers stay beside the tests that use them and keep the
  existing `*_helpers_test.go`, `*_fixture_test.go`, or `*_harness_test.go`
  naming pattern.
- Shared helpers move into `tests/functional/support/...` only after at least
  two collections need the same seam.
- `testdata` should mirror collection ownership. The target shape is:

```text
tests/
  functional/
    default/
      smoke/testdata/
      workflow/testdata/
      boundary/testdata/
    extended/
      service/testdata/
      replay/testdata/
      provider/testdata/
    support/
      fixtures/testdata/
```

- Large fixture trees that are reused across lanes should move once into the
  narrowest shared seam that still reflects real ownership.
- Suite-local process artifacts such as `CLAUDE.md`, `prd.json`, `prd.md`, and
  `progress.txt` should not remain mixed into collection directories after the
  split; they belong with worktree process state, not functional coverage
  ownership.

## Naming And Command Strategy

- Directory names define collection identity. Avoid generic names like
  `misc`, `helpers`, or `other`.
- Scenario files keep behavior-oriented names such as
  `current_factory_watcher_switch_test.go` rather than lane-oriented names.
- The canonical lane commands should be `Makefile` targets that expand to
  package globs, not ad hoc `go test` invocations copied from progress logs.
- `testing.Short()` may still skip unusually expensive cases inside an explicit
  slow collection, but it is not the mechanism that decides fast-lane
  membership.

## Planned Verification Surface

Later stories should converge on these command responsibilities:

| Target | Responsibility |
| --- | --- |
| `make test-functional-default` | Run only the default functional collections under `tests/functional/default/...` |
| `make test-functional-extended` | Run the opt-in extended collections under `tests/functional/extended/...` |
| `make test` | Keep the repository's ordinary developer path wired to the default functional lane rather than the full historical package |
| Focused smoke targets | Continue to exist for repeated or high-signal scenarios when they provide extra stability evidence |

## Migration Order

1. Create the destination collection directories and move the clearest fast-lane
   smoke and workflow tests first.
2. Extract or relocate helper files only after the target collection ownership
   is clear.
3. Move heavy service, replay, and provider scenarios into explicit extended
   collections.
4. Replace broad `tests/functional_test` package commands with lane-specific
   `Makefile` targets once migration coverage is in place.
5. Remove the legacy mixed package after every test, helper, and fixture has an
   explicit owner.

## Review Expectations

When reviewing later migration stories, confirm:

- The moved test lands in the correct collection for both behavior and runtime.
- The helper or fixture moved with the collection unless it is genuinely shared.
- The change improves or preserves lane clarity rather than moving files
  mechanically.
- The `Makefile` command surface remains the single source of truth for lane
  membership.
