# Dashboard Browser Support And Performance Evidence

## Why this should exist

The current dashboard has browser-backed replay and Storybook verification, but
the automated surface is Chromium-only and the existing memory or heap scripts
are not promoted into a documented performance budget or CI-owned release
criterion.

This gap will recur across many future dashboard lanes:

- browser-only regressions can slip through if support expectations for Firefox
  or WebKit remain implicit
- replay-heavy or long-lived dashboard performance work will keep producing
  one-off measurements until the repository defines one canonical evidence lane

## Desired outcome

Create one focused lane that makes browser and performance expectations
reviewable:

- document the intended supported browser set for the dashboard
- either extend automated checks beyond Chromium or explicitly constrain
  support
- define one repo-owned dashboard performance budget or smoke threshold using
  the existing replay, memory, or heap tooling

## Observable evidence that would close it

- maintained docs name the supported browser set and the proof surface for it
- CI or another repo-owned command runs the chosen browser-compatibility checks
- one documented dashboard performance threshold is enforced or regularly
  reported from repo-owned tooling
