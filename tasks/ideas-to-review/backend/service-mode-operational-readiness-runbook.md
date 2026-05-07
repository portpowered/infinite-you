# Service-Mode Operational Readiness Runbook

## Why this should exist

The repository has release packaging and smoke coverage, but this audit did not
find one maintained operator note that explains how to start the long-running
service mode safely, which secrets or provider config are required, how to
verify readiness, or how to roll back a bad deployment.

This is a recurring repository gap rather than a one-off docs nicety:

- future backend work can keep landing without a single shared operational
  contract for startup, health validation, and rollback
- release automation proves artifacts can boot, but it does not replace an
  operator-facing runbook for real service-mode ownership

## Desired outcome

Create one focused lane that publishes the minimum service-mode operator
runbook:

- required secret-bearing config and startup prerequisites
- startup and readiness validation commands or log evidence
- safe shutdown, restart, and rollback expectations for service mode

## Observable evidence that would close it

- one checked-in backend runbook names required runtime secrets and config
- maintainers can follow documented startup and readiness checks without
  inferring behavior from code
- rollback and failure-recovery steps are explicit and versioned in the repo
