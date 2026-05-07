# Service-Mode Observability Inventory

## Why this should exist

The backend already emits structured logs and carries trace-oriented context in
several runtime paths, but this audit did not find one maintained inventory
that tells operators which signals exist today and which production diagnostics
are still missing.

This is an architecture and operability gap rather than a one-off cleanup:

- future runtime work will keep rediscovering log fields and trace seams unless
  the current observability surface is documented in one place
- later health, metrics, or tracing lanes need a baseline inventory so they can
  extend an existing contract instead of inventing parallel operator guidance

## Desired outcome

Create one focused lane that documents the current backend observability
surface:

- structured log fields and correlation identifiers that operators can rely on
- current trace or request identifiers across service-mode flows
- explicit statement of which health, metrics, or tracing signals are absent

## Observable evidence that would close it

- one checked-in observability note names the current log and trace signals
- maintainers can tell which operator-visible diagnostics are stable today
- later health or metrics work has one canonical inventory to extend
