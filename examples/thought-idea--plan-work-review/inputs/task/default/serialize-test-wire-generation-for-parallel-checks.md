# Serialize `test-wire` generation for parallel backend checks

## Problem

Multiple backend verification targets regenerate `backend/test/functional/testwire/wire_gen.go` as part of their normal flow. Running those targets in parallel can fail with a file-lock error on Windows even when the underlying code is healthy.

## Why It Matters

- Agents often parallelize independent checks to keep iteration fast.
- Local contributors can hit a false-negative failure that looks like a broken Wire graph but is really concurrent file generation.
- The current behavior makes the backend validation workflow less predictable in worktrees and other concurrent environments.

## Suggested Direction

- Add a single serialized `test-wire` prerequisite step to backend verification flows instead of regenerating within each target.
- Or introduce a lock/guard around `go generate -tags=wireinject ./test/functional/testwire/...` so concurrent targets wait instead of failing.
- Document which targets are safe to parallelize until the generation step is consolidated.
