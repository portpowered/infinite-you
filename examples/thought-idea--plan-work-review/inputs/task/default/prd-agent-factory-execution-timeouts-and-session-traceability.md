# PRD: Agent Factory Execution Timeouts and Session Traceability

---
author: Codex
last modified: 2026, april, 8
status: draft
---

## Context

### Customer Ask

"We have the agent-factory, but sometimes the workers fail because of getting stuck running some script or failing during inference and doing the wrong thing. Can we write a plan to apply bounded execution timeouts, pass session IDs through to Codex and Claude dispatches so we can inspect rollout logs, and add those session IDs to the dashboards and UI so we can understand failures more easily?"

### Problem

AI agent operators and workflow engineers using Agent Factory cannot reliably diagnose failed or stalled work today. Script workers can hang for an unbounded time, tying up workstation capacity and delaying unrelated work. When Codex or Claude executions fail, the factory does not consistently preserve the provider session identifier through runtime state, dashboard views, and UI surfaces, so operators cannot quickly correlate a failed dispatch with provider-side rollout logs. This causes slow incident triage, repeated reruns without evidence, and poor confidence in whether a failure came from a script hang, provider capacity issue, or a bad agent turn.

### Solution

Introduce bounded execution timeouts that resolve from workstation `timeout`, worker `timeout`, then the default 2-hour subprocess fallback when no configured timeout applies. Treat timeout outcomes as intermittent execution failures that return work to its initial state under the existing retry and exhaustion logic. Extend Codex and Claude dispatch handling so the effective provider session ID is preserved in canonical runtime records: capture provider-generated session IDs from Codex CLI output, preserve caller-specified or provider-confirmed Claude session IDs when supported, and surface the resulting identifier through dashboard APIs, CLI/debug views, and website UI so operators can trace failed work back to the underlying provider logs.

## Project Acceptance Criteria

- [ ] Script-based worker executions are force-terminated after the effective timeout, using workstation `timeout`, worker `timeout`, then the bounded 2-hour subprocess fallback
- [ ] Script timeout outcomes are normalized into the existing intermittent failure path so the affected work item returns to its initial state instead of immediately routing to terminal failure
- [ ] Codex and Claude dispatches persist the effective provider session ID or equivalent rollout identifier in raw runtime records whenever the provider supplies or accepts one
- [ ] Dashboard and trace APIs expose the recorded session ID for active and completed work without requiring consumers to reconstruct it from free-form logs
- [ ] Operator-facing surfaces in the dashboard, website UI, and related CLI/debug output display the session ID alongside failed or selected work so the rollout can be inspected externally
- [ ] A multi-subsystem smoke test proves a dispatch can produce a session ID, surface it through the backend/dashboard path, and remain visible to the UI/debug consumer after a failure or retry event
- [ ] Quality checks pass (typecheck, lint, tests)

## Goals

- [ ] Prevent hung script workers from consuming execution capacity indefinitely
- [ ] Preserve enough provider execution identity to debug Codex and Claude failures from external rollout logs, even though the two providers expose session continuity differently
- [ ] Surface failure context consistently across runtime state, dashboards, and operator UI
- [ ] Reuse the existing intermittent failure and exhaustion model instead of inventing a separate retry path for timeouts
- [ ] Keep the feature configurable without requiring per-workflow code changes

## User Stories

### US-001: Add bounded script execution timeouts to worker runtime config (P1)

**Description:** As an AI workflow engineer, I want script workers to enforce a default execution timeout so a hung script cannot occupy a workstation indefinitely.

**Acceptance Criteria:**
- [ ] Script worker execution resolves an effective timeout from runtime-loaded config, using workstation `timeout`, worker `timeout`, then the bounded 2-hour subprocess fallback when neither scope configures a positive timeout
- [ ] When a script exceeds the effective timeout, the running process is terminated and the executor returns a normalized timeout/intermittent failure result instead of waiting indefinitely
- [ ] A functional test proves a long-running script is interrupted after a short test override and does not keep the worker process alive
- [ ] Typecheck, lint, and relevant automated tests pass

**Notes:**
Use the existing runtime-loaded worker/workstation config path rather than introducing ad hoc config loading in the executor.

### US-002: Requeue timed-out work through the existing intermittent failure path (P1)

**Description:** As an agent-factory operator, I want script timeout failures to behave like other intermittent execution failures so work retries safely under the current exhaustion and retry rules.

**Acceptance Criteria:**
- [ ] A timed-out script execution is classified through the same normalized failure handling path used for intermittent provider or capacity-style failures
- [ ] When the timeout classification is returned, the transitioner rebuilds the work token from the consumed dispatch snapshot and returns it to its pre-transition place instead of routing it directly to a terminal failure place
- [ ] History and metrics record that the attempt ended because of a timeout while still showing the work item as eligible for retry until the configured guarded loop-breaker threshold is hit
- [ ] A functional test proves a timed-out work item returns to its initial state and can be dispatched again on a later tick
- [ ] Typecheck, lint, and relevant automated tests pass

### US-003: Capture provider session IDs for Codex and Claude dispatches (P1)

**Description:** As an operator debugging agent output, I want Codex and Claude dispatches to record provider session IDs so I can correlate a factory work item with the provider rollout logs.

**Acceptance Criteria:**
- [ ] The Codex execution path captures the provider-generated session ID or equivalent rollout identifier from the CLI response stream or output when one is emitted
- [ ] The Claude execution path supports preserving a caller-specified session ID when the provider interface allows one and records the effective session ID used for the dispatch
- [ ] The captured session ID is stored on a shared raw dispatch/result record or other canonical runtime field rather than being embedded only in rendered log text
- [ ] When a provider does not expose a session ID, or does not honor a requested one, the dispatch still completes without error and the effective absence is represented explicitly rather than guessed downstream from rendered log text
- [ ] Automated tests cover both the presence and absence of session IDs for Codex and Claude provider flows
- [ ] Typecheck, lint, and relevant automated tests pass

### US-004: Expose session IDs through dashboard and trace read models (P1)

**Description:** As a dashboard consumer, I need session IDs available in backend read models and APIs so I can inspect them without parsing raw executor output.

**Acceptance Criteria:**
- [ ] Dashboard/session summary read models reconstruct and expose provider session IDs from the canonical raw runtime records
- [ ] Trace or work-detail API responses include the session ID for any dispatch attempt that recorded one
- [ ] Failed, retried, and completed work remain queryable with their associated session IDs after dispatch history retirement or snapshot reconstruction
- [ ] Automated tests verify session IDs survive the normal dashboard reconstruction path and are not sourced from renderer-local mirrors
- [ ] Typecheck, lint, and relevant automated tests pass

### US-005: Display session IDs in operator-facing dashboard, website, and CLI/debug surfaces (P2)

**Description:** As an operator investigating failures, I want session IDs shown in the dashboard, website UI, and related debug output so I can jump directly to the relevant rollout logs.

**Acceptance Criteria:**
- [ ] The primary operator-facing dashboard view shows the session ID for work items or dispatch attempts where one exists
- [ ] The website UI exposes the same session ID field in the relevant work detail or failure inspection surface without requiring raw JSON inspection
- [ ] Related CLI or debug output that already summarizes work attempts includes the session ID when present
- [ ] UI/browser verification confirms the session ID is visible in the intended website flow
- [ ] Typecheck, lint, and relevant automated tests pass
- [ ] Verify in browser using dev-browser skill

### US-006: Add an end-to-end smoke test for timeout and traceability diagnostics (P3)

**Description:** As a maintainer, I want an end-to-end smoke test across runtime, dashboard, and UI/debug surfaces so regressions in timeout handling or session traceability are caught before release.

**Acceptance Criteria:**
- [ ] The smoke test exercises a dispatch flow that records a provider session ID and verifies that the same ID is visible through runtime state plus at least one operator-facing read surface
- [ ] The smoke test or companion scenario covers a timed-out script attempt returning to its initial state under intermittent failure handling
- [ ] The test proves the end-to-end path across at least two independently buildable components, satisfying the cross-subsystem integration requirement for this PRD
- [ ] Typecheck, lint, and relevant automated tests pass

## Functional Requirements

1. FR-1: Script worker execution must support an effective timeout value resolved from runtime-loaded configuration, using workstation `timeout`, worker `timeout`, then the bounded 2-hour subprocess fallback when no configured timeout applies.
2. FR-2: When a script worker exceeds its effective timeout, the factory must terminate the running process and classify the result as an intermittent execution failure.
3. FR-3: Timeout-classified script work must return to its initial or pre-transition state using the existing retry and exhaustion path rather than a new bespoke retry system.
4. FR-4: Codex dispatches must capture and persist a provider-generated session ID when one is emitted by the provider CLI or response stream.
5. FR-5: Claude dispatches must support passing through a configured or caller-provided session ID when supported, and must persist the effective session ID used for the dispatch.
6. FR-6: The canonical runtime dispatch/result record must store provider session IDs in a structured field that downstream read models can consume without re-parsing rendered logs.
7. FR-7: Dashboard, trace, and work-detail APIs must expose the structured session ID field when present.
8. FR-8: Operator-facing dashboard, website UI, and related CLI/debug output must render the session ID in the relevant failure or work-detail views.
9. FR-9: The system must continue to function when a provider returns no session ID, representing that absence without parsing free-form output text.
10. FR-10: Automated tests must cover timeout termination, intermittent requeue behavior, session ID capture, and cross-surface visibility.

## Non-Goals

- This work does not add durable persistence for factory runtime state across process restarts
- This work does not introduce a new generic provider-debugging console beyond surfacing session IDs in existing views
- This work does not guarantee that every provider returns a session ID; it only preserves and surfaces the identifier when available
- This work does not redesign the current retry or exhaustion policy beyond classifying script timeouts into the existing intermittent failure path
- This work does not add browser UI for editing timeout values; configuration remains code- or file-driven

## Design Considerations

- Prefer concise, copyable session ID presentation with clear labeling so operators know which external rollout system to inspect.
- If a view shows multiple attempts for one work item, the UI should make it clear which session ID belongs to which attempt.
- Timeout-related failure messaging should distinguish "process exceeded configured timeout" from generic script exit failure.

## Technical Considerations

- The Agent Factory intent and current development notes emphasize runtime-loaded config, reconstruction-first dashboard read models, and canonical raw dispatch/result records; this work should follow those existing patterns.
- Script execution already flows through workstation resolution, so timeout support should respect both workstation and worker configuration rather than hardcoding executor-local values.
- Session ID propagation likely spans executor/provider boundaries, runtime state snapshot fields, dashboard reconstruction models, backend APIs, and website rendering; the PRD intentionally leaves the exact field placement to implementation, but the canonical source should be a structured provider metadata field rather than free-form output text.
- Codex and Claude should not be forced into the same acquisition path: Codex may require extracting a provider-generated session ID from CLI output or event streams, while Claude may also support caller-provided session continuity that must be preserved distinctly from factory `DispatchID` or `TraceID`.
- Because timeout handling is being treated like intermittent capacity-style failure, implementation should reuse the same normalized error classification path and retry visibility semantics used elsewhere in the factory.
- Existing functional tests around provider subprocess handling, trace reconstruction, and dashboard snapshot rendering are likely the cheapest places to extend coverage.

## Success Metrics

| Metric | Target | Evaluation Date |
|--------|--------|-----------------|
| Script dispatches exceeding their timeout that remain stuck without a terminal or requeued outcome | 0 known occurrences in test and staging runs after release | 2026-05-08 |
| Failed Codex or Claude dispatches where an operator can retrieve a session ID from dashboard/API/UI surfaces when the provider returned one | 95% or better in sampled post-release failures | 2026-05-08 |
| Time from failure report to locating the corresponding provider rollout/session log | Reduced versus current manual triage baseline, validated qualitatively during first on-call cycle | 2026-05-08 |

## Open Questions

- [ ] Should the canonical provider metadata model expose one generic `session_id` field plus optional provider-specific metadata, or should Codex and Claude each surface distinct structured fields because their session semantics differ?
- [ ] Which exact website surfaces should render the session ID first: work list rows, detail panes, trace explorer, or only failed dispatch detail views?
- [ ] Should timeout overrides be allowed at both worker and workstation scope from day one, or should one scope simply inherit from the other until a later cleanup?
- [ ] Do we also want timeout and session ID data exported through any external telemetry sink, or is the in-product dashboard/UI surface sufficient for the first release?

## Changelog

| Date | Change | Author |
|------|--------|--------|
| 2026-04-08 | Initial version | Codex |
| 2026-04-08 | Clarified that Codex session IDs are provider-generated and captured from CLI output while Claude may support caller-provided session continuity; updated requirements to preserve the effective session ID in canonical runtime records | Codex |
