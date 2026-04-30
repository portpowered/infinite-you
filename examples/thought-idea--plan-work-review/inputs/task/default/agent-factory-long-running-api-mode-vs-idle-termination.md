# Agent Factory long-running API mode vs idle termination

## Problem

The current factory runtime terminates immediately when it reaches an idle, fully-completed state. That is correct for one-shot batch execution, but it makes browser-facing and API-facing observability harder to exercise and reason about:

- a freshly started factory with no preseeded work can terminate before a live client submits work
- functional smoke tests for live dashboard behavior must preseed blocked work to keep the process alive
- runtime submission after startup is harder to verify because the lifecycle depends on the engine still being active

## Suggested direction

Split the current execution semantics into two explicit modes:

1. **batch mode**: current behavior, terminate on idle completion
2. **service mode**: remain alive until context cancellation and continue accepting new work submissions

Possible implementation shapes:

- add a `WithServiceMode()` or similar factory/service option that changes termination behavior
- keep the current default for CLI batch-style workflows, but use service mode when we set the --continuously
- make the status response report whether the runtime is intentionally idle vs terminally finished

## Why it matters

This would simplify operator-facing dashboard behavior, browser testing, and live API use without weakening the current deterministic batch flow that existing tests rely on.
