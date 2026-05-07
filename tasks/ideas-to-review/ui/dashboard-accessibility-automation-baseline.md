# Dashboard Accessibility Automation Baseline

## Why this should exist

The dashboard already has strong semantic and keyboard-oriented component and
Storybook coverage, but this audit did not find a dedicated accessibility
automation lane such as axe or pa11y in CI.

This is a recurring repository gap rather than a one-off UI polish item:

- future dashboard changes can keep passing rendered-behavior tests while still
  regressing labels, roles, focus handling, or contrast-related semantics
- the repository already has stable browser-backed verification seams in
  `ui/scripts/run-storybook-ci.mjs` and
  `ui/integration/event-stream-replay.integration.test.mjs`, so adding
  accessibility assertions has a clear home instead of needing a new testing
  architecture

## Desired outcome

Create one focused lane that adds automated accessibility verification for the
highest-value dashboard flows:

- dashboard header and timeline controls
- submit-work surface
- export dialog
- import preview dialog

## Observable evidence that would close it

- one repo-owned accessibility command is wired into `ui/package.json` and CI
- browser-backed tests fail on missing labels, invalid dialog semantics, or
  similar high-signal accessibility regressions
- maintained docs explain which flows are covered automatically and which still
  require manual spot checks
