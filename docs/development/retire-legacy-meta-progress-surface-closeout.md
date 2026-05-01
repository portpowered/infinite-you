# Retire Legacy Meta Progress Surface Closeout

Date: 2026-04-30
Scope: final verification for `prd.json` `US-004` on branch `ralph/retire-legacy-meta-progress-surface`

## Summary

This closeout proves the cleanup stayed limited to meta progress-surface
ownership and the minimum supporting control-plane enforcement needed to keep
that ownership stable.

Files changed in the reviewable diff against `main`:

- `docs/development/root-factory-artifact-contract-inventory.md`
- `docs/processes/factory-workstation-relevant-files.md`
- `factory/logs/meta/progress.tsx`
- `factory/logs/meta/view.md`
- `factory/workstations/cleaner/AGENTS.md`
- `internal/testpath/artifact_contract.go`
- `pkg/testutil/artifact_contract_test.go`

The branch does not widen into customer-facing features or unrelated workflow
cleanup. The diff is limited to:

- declaring and teaching `factory/logs/meta/progress.txt` as the canonical
  checked-in meta progress surface,
- deleting the competing checked-in `factory/logs/meta/progress.tsx` path,
- updating the checked-in meta view so it reflects the single remaining
  canonical surface, and
- adding the minimum artifact-contract inventory, guard, and regression-test
  updates needed to make future drift obvious.

## Validation

Commands run from the repository root:

```powershell
make artifact-contract-closeout
make lint
make test
```

Results on 2026-04-30:

- `make artifact-contract-closeout` passed.
- `make lint` passed.
- `make test` passed.

## What This Proves

- The reviewable diff stays narrow: it changes only meta progress-surface
  ownership, active maintainer guidance, and the guard/test surfaces that
  enforce the same path contract.
- `factory/logs/meta/progress.txt` is the only canonical checked-in meta
  progress surface left in the active maintainer control plane.
- Future divergence is harder to reintroduce silently because the legacy
  `factory/logs/meta/progress.tsx` path is removed and the artifact-contract
  checks now classify it as obsolete.
