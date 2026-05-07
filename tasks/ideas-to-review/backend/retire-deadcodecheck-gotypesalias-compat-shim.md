# Retire Deadcodecheck Gotypesalias Compatibility Shim

## Why this should exist

The maintainer ask explicitly calls for simplification where duplicate or stale
logic remains, and the current narrow backend candidate is the
`cmd/deadcodecheck` compatibility shim that forces `GODEBUG=gotypesalias=1`.

The live code still carries this policy in `runDeadcode()`, `deadcodeEnv()`,
and `ensureGoTypesAliasEnabled()`, while the May 2026 dead-code candidate
ledger already records the same seam as a narrow simplification opportunity.

## Desired outcome

Create one focused lane that proves whether this shim is still required:

- verify deadcode runs on the supported Go `1.24.x` toolchain with and without
  the override
- remove the compatibility branch if it no longer protects real supported
  behavior
- keep command-owner tests and maintainer docs aligned with the final policy

## Observable evidence that would close it

- the deadcode command either runs without the shim or has a maintained record
  explaining why the shim is still required
- `cmd/deadcodecheck` keeps one explicit policy owner instead of layered legacy
  handling
- tests and docs prove the chosen behavior at the repo-owned command surface
