# Port OS Reference Inventory

This inventory defines the Port OS coupling audit boundary for
`libraries/agent-factory`. It separates the customer-facing release surface
from internal-only matches so later cleanup stories can use one scope,
classification rule set, and rerunnable collection process.

## Collection Process

### Customer-Facing Scan

Use this scan for the in-scope public surface: `README.md`, shipped docs,
examples, checked-in scaffold content, and emitted scaffold source. The
maintained enforcement command below uses the same explicit surface list and
the same internal-doc exclusion rule.

```powershell
cd libraries/agent-factory
go run ./cmd/publicsurfacecheck
```

Result on 2026-04-21 after US-005: 0 matching lines across 0 files.

Equivalent ad hoc collection scan:

```powershell
rg -n -i "portos|port os|port_os" `
  libraries/agent-factory/README.md `
  libraries/agent-factory/docs `
  libraries/agent-factory/examples `
  libraries/agent-factory/factory `
  libraries/agent-factory/pkg/cli/init/init.go `
  --glob '!libraries/agent-factory/docs/development/**'
```

### Whole-Tree Triage Scan

Use this broader scan only to prove the remaining matches are internal test
fixtures, standards metadata, historical artifacts, or audit references rather
than customer-facing defaults.

```powershell
rg -n -i "portos|port os|port_os" libraries/agent-factory --glob '!**/portos-reference-inventory.md'
```

Result on 2026-04-21 on the current review branch head: 205 matching lines
across 62 files.

## Classification Rules

- `remove` means the reference is only product coupling and can be deleted
  outright.
- `generalize` means the reference should be rewritten to neutral project or
  repository language while preserving the behavior.
- `retain` means the reference is an explicit internal-only compatibility
  check, historical artifact, standards tag, or audit artifact that is not the
  default customer path.

## Summary

- The enforced customer-facing scan now finds no Port OS matches in the
  customer-facing Agent Factory release surface.
- `libraries/agent-factory/README.md`, shipped docs under
  `libraries/agent-factory/docs/**`, `libraries/agent-factory/factory/**`, and
  `libraries/agent-factory/pkg/cli/init/init.go` currently contain no Port OS
  references.
- The whole-tree triage scan still finds internal-only package and test
  metadata, the `Makefile` target that runs the guard, UI review fixtures,
  historical contributor docs, historical replay artifacts, and development
  doc references back to this inventory and guard.
- The remaining whole-tree triage matches are all explicit internal-only,
  historical, or audit exceptions outside the customer-facing release surface.
- `libraries/agent-factory/Makefile` runs `go run ./cmd/publicsurfacecheck`
  during `make lint`, so review-time verification now fails deterministically
  when a new customer-facing match is introduced.

## In-Scope Customer-Facing Matches

None. The customer-facing scan above returns zero matches after US-005.

## In-Scope Surfaces With Zero Matches

- `libraries/agent-factory/README.md`
- `libraries/agent-factory/docs/**`
- `libraries/agent-factory/factory/**`
- `libraries/agent-factory/pkg/cli/init/init.go`
- `libraries/agent-factory/examples/basic/**`
- `libraries/agent-factory/examples/idea-plan-code-review/**`
- `libraries/agent-factory/examples/simple-tasks/**`
- `libraries/agent-factory/examples/write-code-review/**`

## Out-Of-Scope Matches From Whole-Tree Triage

| Bucket | Lines / files | Classification | Rationale |
| --- | --- | --- | --- |
| Internal package, command, guard, and test matches under `libraries/agent-factory/cmd/**`, `libraries/agent-factory/pkg/**`, `libraries/agent-factory/tests/**`, and `libraries/agent-factory/Makefile`, excluding the replay artifact files listed below | 156 lines / 53 files | retain | These matches are internal-only standards metadata such as `// portos:func-length-exception`, compatibility fixtures such as explicit `PORTOS_*` pass-through coverage, the `cmd/publicsurfacecheck` guard and its tests, guard assertions that prove product coupling does not leak into runtime behavior, and the `Makefile` target that runs the `TestInit_GeneratedCustomerFacingFilesDoNotContainPortOS` regression. They are not customer-facing release content. |
| Internal UI review fixtures in `libraries/agent-factory/ui/src/App.stories.tsx`, `libraries/agent-factory/ui/src/components/dashboard/fixtures/workstation-requests.ts`, and `libraries/agent-factory/ui/src/features/current-selection/*.test.tsx` | 16 lines / 4 files | retain | These are Storybook and Vitest review fixtures with example local paths. They are not part of the shipped docs, examples, or scaffold defaults. |
| Historical contributor docs in `libraries/agent-factory/docs/development/record-replay-design.md` and `libraries/agent-factory/docs/development/cleanup-analyzer-reports/2026-04-19-retire-raw-worker-emitted-batch-output.md` | 23 lines / 2 files | retain | These files capture internal design history and analyzer output, including absolute local paths. They are not customer-facing release guidance. |
| Historical replay artifacts in `libraries/agent-factory/tests/adhoc/factory-recording-04-11-02.json` and `libraries/agent-factory/tests/functional_test/testdata/adhoc-recording-batch-event-log.json` | 8 lines / 2 files | retain | These are recorded test artifacts with captured local paths and prompts. They are explicit historical fixtures, not defaults used by the public examples or scaffold flow. |
| Development guide references in `libraries/agent-factory/docs/development/development.md` | 2 lines / 1 file | retain | One line names the `go run ./cmd/publicsurfacecheck` guard in the release-surface smoke description and one line links to `portos-reference-inventory.md` from the development index. Both references are intentional because the guide documents the canonical audit command and links to this inventory artifact. |

## Verification

1. Run the customer-facing scan and confirm it returns zero matches.
2. Run the whole-tree triage scan and confirm all remaining matches still fit
   one of the out-of-scope buckets above.
3. When a later story changes the customer-facing surface, update the expected
   counts, the evidence tables in this inventory, and the
   `cmd/publicsurfacecheck` inclusion rules in the same change.
