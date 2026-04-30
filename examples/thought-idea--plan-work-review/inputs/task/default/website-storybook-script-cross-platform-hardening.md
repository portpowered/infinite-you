# Idea: Harden website Storybook CI scripts for cross-platform local and CI execution

## Problem

The website's `test-storybook:ci` script was using `npx http-server` and waiting only for the root Storybook URL. On Windows in this worktree, that combination produced `ECONNRESET` / index fetch failures for `test-storybook` even though the built Storybook output itself was valid.

## Proposed improvement

Standardize Storybook static-runner scripts on package-local CLI binaries and wait for `index.json` explicitly before starting the test runner.

## Why it helps

- Reduces Windows-specific process spawning failures from nested `npx` wrappers
- Aligns the readiness check with the actual asset that `test-storybook` requires
- Makes local verification closer to CI behavior and reduces false-negative review feedback

## Suggested follow-up

- Audit other repo npm scripts that wrap local dev dependencies in `npx` inside long-running `concurrently` commands
- Prefer waiting on the specific asset a downstream tool needs instead of only an HTML entrypoint
