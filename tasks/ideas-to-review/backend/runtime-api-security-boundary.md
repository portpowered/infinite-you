# Runtime API Security Boundary

## Why this should exist

The current runtime API contract exposes write and read routes with
`security: []` in `api/openapi-main.yaml`, and this audit did not find a
checked-in transport or deployment boundary that explains how those endpoints
are protected in shared or remote environments.

This is an architecture-level gap rather than a one-off polish item:

- future backend stories will keep tripping the same audit finding until the
  repository either adds an explicit auth boundary or documents a hard
  local-only deployment assumption with enforceable checks
- the current repo already supports named-factory creation, work submission,
  and event streaming over HTTP, so the blast radius is larger than a read-only
  internal diagnostic port

## Desired outcome

Create one focused lane that does one of these explicitly:

- add an authenticated runtime API boundary for service-mode HTTP routes, or
- document and enforce a local-only deployment contract that makes the absence
  of auth an intentional, reviewable product constraint

## Observable evidence that would close it

- `api/openapi-main.yaml` declares the chosen security scheme or clearly scoped
  local-only assumptions
- `pkg/api/` enforces the boundary instead of relying on tribal knowledge
- tests prove protected and unprotected route behavior at the public HTTP layer
- docs explain how operators are expected to run the service safely
