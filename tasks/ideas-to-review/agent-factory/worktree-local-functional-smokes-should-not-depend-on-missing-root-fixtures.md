# Worktree-Local Functional Smokes Should Not Depend on Missing Root Fixtures

## Why this matters

Several broad functional and release-surface tests assume a sibling or
monorepo-mounted `factory/` tree outside the current worktree. In isolated
checkouts like autonomous iteration branches, those paths are often absent, so
otherwise unrelated stories cannot prove broader package health.

## Observed risk

- `go test ./tests/functional_test -count=1` fails in isolated worktrees before
  reaching story-specific regressions because fixture setup reads
  `../../../../factory` or `.claude/factory`.
- Contributors fall back to narrowly scoped test runs even when they want a
  broader package signal.
- Repeated missing-path failures obscure whether a story actually regressed
  runtime behavior or just inherited checkout topology assumptions.

## Suggested follow-up

1. Move root-factory smoke inputs behind checked-in package-local fixtures or a
   test helper that synthesizes the required tree under `t.TempDir()`.
2. Make release-surface and replay-artifact smokes fail with explicit skip or
   setup guidance when the external fixture root is intentionally unavailable.
3. Add one package-local guard test that proves `tests/functional_test` broad
   smoke paths do not require undeclared filesystem mounts.
