# Functional Test Runtime Closeout

This closeout records the repository-owned command surface that enforces the
fast default functional lane and the current measured runtimes for the default,
extended, and remaining legacy-package paths.

## Inventory Date

- 2026-04-30

## Canonical Commands

| Lane | Command | Role |
| --- | --- | --- |
| Default fast lane | `make test-functional-default` | Runs only `tests/functional/default/...` |
| Default fast-lane budget gate | `make test-functional-default-budget` | Re-runs the canonical default lane and fails if it takes longer than `10s` |
| Ordinary developer verification path | `make test` | Runs the short non-stress package set, then enforces the default functional-lane budget gate |
| Opt-in slow lane | `make test-functional-extended` | Runs the full remaining legacy `tests/functional_test` package until those scenarios are physically moved into `tests/functional/extended/...` |
| Broad legacy-package baseline | `go test ./tests/functional_test -count=1` | Measures the remaining unsplit `tests/functional_test` package directly |

## Measured Runtime

Measured on 2026-04-30 from the repository root in this worktree:

| Command | Runtime | Result |
| --- | --- | --- |
| `make test-functional-default-budget` | `2.554s` | Pass |
| `make test-functional-extended` | `25.768s` | Pass |
| `go test ./tests/functional_test -count=1` | `75.157s` | Pass |

## Closeout Notes

- The enforced default lane stays well under the 10-second budget.
- `make test` no longer routes ordinary short-suite verification through the
  remaining `tests/functional_test` package.
- Slow legacy functional scenarios remain available intentionally through
  `make test-functional-extended` while the full physical migration into
  `tests/functional/extended/...` is still pending.
