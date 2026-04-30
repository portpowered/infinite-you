# PRD: Agent Factory Workflow Dashboard Redesign

---
author: Codex
last modified: 2026, april, 9
status: draft
---

## Context

### Customer Ask

"The current dashboard is fantastic, but the components need to be more usable. The current flow view is not handling even our small workflow easily. Expand the layout algorithm, allow panning, move selection into a floating shrinkable panel, remove useless text, make the dashboard full width, make traces larger, show timings in human units, break workstation and work metrics into separate widgets, and add Grafana-like graphs so we can understand throughput, failures, retries, rework, and what is blocking the system."

### Problem Statement

Agent Factory operators use the website dashboard to understand what work is running, what failed, why work took too long, and what is blocking throughput. The current dashboard makes that difficult because the work graph is cramped, non-pannable, overloaded with low-value text, and unclear about the difference between workstation definitions, work items, resources, constraints, and retries. Important drill-down information is placed in awkward side panels, timing values are hard to read, and the dashboard does not provide clear time-based views for throughput, failure causes, rework, or bottlenecks. As a result, operators spend extra time interpreting the screen and still cannot reliably answer the core questions "what is happening now?" and "what is slowing us down over time?"

### Solution

Redesign the Agent Factory dashboard in `website/` around a full-width, operator-first layout with a more spacious pannable work graph, clearer node semantics, a floating collapsible inspector, and a dedicated lower section for traces and selected-item detail. Add dashboard widgets and Grafana-like time-series panels for throughput, latency, failures, retries, and rework using existing website-accessible data where possible, while keeping the scope limited to website UX and presentation.

## Project Acceptance Criteria

- [ ] The dashboard uses a full-width layout and removes the currently identified low-value headings and panels, including the live workstation dashboard and terminal summary.
- [ ] The work graph supports a more expanded layout, viewport panning, and clearer visual differentiation between workstation definitions, work items, resources, constraints, and retry/rejected states.
- [ ] Selecting a workstation and selecting a work item produce distinct highlight states and distinct inspector content labels.
- [ ] Selected work or trace detail renders in a full-width section beneath the work graph, and clicking completed or failed items opens useful lower-page detail instead of relying on the current side panel layout.
- [ ] The dashboard exposes separate widgets or panels for workstation selection, work selection, timings, and aggregate operational metrics over a controllable time range.
- [ ] Timing values shown to operators default to human-readable units such as hours, minutes, and seconds instead of raw milliseconds.
- [ ] Automated coverage validates graph rendering and dashboard layout behavior across representative workflow sizes, node compositions, and viewport widths.
- [ ] Quality checks pass (typecheck, lint, tests).

## Goals

- Make the workflow topology readable and navigable for real operator workflows, not only small demos.
- Reduce low-value text so the dashboard emphasizes graph structure, state, and operational signals.
- Make selection and drill-down behavior predictable across workstation, work item, completed item, and failed item interactions.
- Expose the main causes of throughput loss through time-based metrics for failures, retries, latency, and rework.
- Add automated UI coverage that proves graph and dashboard behavior stays usable as workflow size and viewport size change.
- Preserve website-only scope by improving visualization and interaction without requiring a backend telemetry initiative in this phase.

## User Stories

### US-001: Simplify the dashboard shell and copy (P1)
**Description:** As an Agent Factory operator, I want the dashboard chrome and text reduced to only useful labels so that I can focus on the work graph and operational signals.

**Acceptance Criteria:**
- [ ] The page uses a full-width dashboard layout instead of a centered content column.
- [ ] The workflow topology area is labeled `Work Graph`.
- [ ] The `Live Workstation Dashboard` panel is removed from the page.
- [ ] The `Reconstruction-first workflow graph with live workstation overlays, work selection, and trace entry points` copy is removed.
- [ ] The terminal summary section is removed.
- [ ] Remaining inspector labels use `Workstation Info` when a workstation is selected and `Work Info` when a work item is selected.
- [ ] Verify in browser using dev-browser skill.
- [ ] Typecheck, lint, and tests pass.

### US-002: Expand and pan the work graph canvas (P1)
**Description:** As an Agent Factory operator, I want a more spacious pannable graph canvas so that I can inspect workflows without the graph collapsing into an unreadable cluster.

**Acceptance Criteria:**
- [ ] The graph layout algorithm spaces nodes more aggressively than the current layout for the same workflow data.
- [ ] The graph canvas supports pointer or trackpad panning without requiring a page scroll workaround.
- [ ] The graph remains usable when the workflow contains more nodes than the current small-workflow example.
- [ ] The pannable canvas preserves node click and hover interactions.
- [ ] Automated tests cover representative small and large graph sizes, including a 1-node graph and an approximately 20-node graph.
- [ ] Verify in browser using dev-browser skill.
- [ ] Typecheck, lint, and tests pass.

### US-003: Separate workstation and work-item selection behavior (P1)
**Description:** As an Agent Factory operator, I want workstation selection and work-item selection to look and behave differently so that I always know whether I selected a definition or an execution.

**Acceptance Criteria:**
- [ ] Selecting a workstation definition applies a different highlight treatment than selecting an individual work item or sub-element.
- [ ] The selected state label and detail content clearly identify whether the active selection is workstation-level or work-level.
- [ ] Selection changes update the inspector content without ambiguous mixed terminology such as `workstation/work detail`.
- [ ] Keyboard or pointer focus changes do not leave stale highlight states on previously selected elements.
- [ ] Verify in browser using dev-browser skill.
- [ ] Typecheck, lint, and tests pass.

### US-004: Move selection detail into a floating collapsible inspector (P1)
**Description:** As an Agent Factory operator, I want selection detail in a floating shrinkable panel over the canvas so that I can inspect items without giving up graph space.

**Acceptance Criteria:**
- [ ] The selection inspector floats over the graph canvas rather than occupying a fixed layout column.
- [ ] The inspector can be collapsed to a compact state and expanded back to a readable state.
- [ ] The collapsed state preserves enough context for operators to recover the selected item.
- [ ] The inspector does not block essential graph interactions such as panning or selecting nearby nodes.
- [ ] Verify in browser using dev-browser skill.
- [ ] Typecheck, lint, and tests pass.

### US-005: Make graph semantics visible for resources, constraints, state positions, and retries (P1)
**Description:** As an Agent Factory operator, I want the graph to distinguish execution semantics visually so that I can understand what each element does without reading dense text.

**Acceptance Criteria:**
- [ ] Resources and constraints are rendered with distinct visual treatments instead of appearing as undifferentiated generic nodes.
- [ ] The graph exposes the positions or state nodes needed for operators to understand where work is in the workflow.
- [ ] Active executions stay visually pinned to the owning workstation definition by `DispatchID`.
- [ ] Completed and rejected repeater outcomes visibly distinguish normal completion from retry-related rejection or retry state.
- [ ] A user can tell from the graph alone which elements are workstation definitions, active executions, resources, constraints, and state positions.
- [ ] Storybook coverage includes representative graph compositions that exercise these distinct element types.
- [ ] Verify in browser using dev-browser skill.
- [ ] Typecheck, lint, and tests pass.

### US-006: Move trace and selected-work detail beneath the graph (P1)
**Description:** As an Agent Factory operator, I want detailed trace and selected-item information in a wide lower section so that I can inspect execution context without compressing the graph.

**Acceptance Criteria:**
- [ ] The dashboard renders a full-width section beneath the graph for trace detail and selected-item detail.
- [ ] Clicking completed or failed items opens useful lower-page detail tied to the selected item.
- [ ] The lower section provides more vertical space for trace content than the current dashboard.
- [ ] The lower section content updates based on selection without forcing a page navigation.
- [ ] Verify in browser using dev-browser skill.
- [ ] Typecheck, lint, and tests pass.

### US-007: Humanize execution timing presentation (P2)
**Description:** As an Agent Factory operator, I want timing values presented in approximate human units so that I can interpret long-running work quickly.

**Acceptance Criteria:**
- [ ] Execution timing values default to human-readable units such as `2h 4m`, `3m 12s`, or `450ms` as appropriate.
- [ ] The timing formatter handles sub-second, multi-second, multi-minute, and multi-hour values consistently.
- [ ] Timing values shown in trace and metric panels no longer default to raw millisecond-only output.
- [ ] Verify in browser using dev-browser skill.
- [ ] Typecheck, lint, and tests pass.

### US-008: Add dashboard widgets for throughput, latency, failures, retries, and rework over time (P2)
**Description:** As an Agent Factory operator, I want dashboard widgets and graphs for system behavior over time so that I can see how much work is getting done and what is preventing higher throughput.

**Acceptance Criteria:**
- [ ] The dashboard includes separate panels or widgets for throughput/completion volume, failure reasons, and retry or rework rate over time.
- [ ] At least one graph supports an explicit time-range control.
- [ ] The dashboard exposes timing trend views that help explain why a trace or class of work took a long time.
- [ ] Failure views aggregate causes into operator-meaningful categories when that information is available in existing website data.
- [ ] The widgets are arranged as dashboard components rather than being merged into one overloaded panel.
- [ ] Automated tests validate that dashboard graph layouts remain readable at representative desktop widths and supported narrower widths.
- [ ] Verify in browser using dev-browser skill.
- [ ] Typecheck, lint, and tests pass.

### US-009: Introduce a dashboard widget-board foundation for future observability views (P3)
**Description:** As a product team member, I want the redesigned dashboard to use a widget-board pattern compatible with richer chart panels so that future operational views can be added without another page rewrite.

**Acceptance Criteria:**
- [ ] The redesigned dashboard organizes metrics, traces, and graph panels as discrete widgets or panels with a coherent board layout.
- [ ] The implementation evaluates and, if suitable, adopts the existing `approach graph UX` package direction for widget boards and movable dashboard elements; if not adopted, the code documents the alternative approach.
- [ ] The board foundation does not regress the core graph and inspector usability delivered by higher-priority stories.
- [ ] Verify in browser using dev-browser skill.
- [ ] Typecheck, lint, and tests pass.

### US-010: Add automated visual and interaction coverage for graph and dashboard behavior (P1)
**Description:** As a website maintainer, I want Storybook and automated UI coverage for the workflow dashboard so that graph rendering, layout responsiveness, and graph interactions stay reliable as the dashboard evolves.

**Acceptance Criteria:**
- [ ] Storybook stories cover representative graph states, including a single-node graph, a medium graph, and an approximately 20-node graph.
- [ ] Storybook stories cover representative element mixes, including workstation definitions, active executions, resources, constraints, and retry or rejected states.
- [ ] Storybook viewport coverage exercises the relevant dashboard widths called out by the website process for responsive validation.
- [ ] Automated tests verify that the graph canvas remains usable across those representative sizes and widths.
- [ ] Automated tests verify expected graph interactions such as panning and node movement when node movement is supported by the implemented UX.
- [ ] Automated tests verify that dashboard graph widgets render without overlap or unusable compression at supported widths.
- [ ] Verify in browser using dev-browser skill.
- [ ] Typecheck, lint, and tests pass.

## Functional Requirements

1. FR-1: The website dashboard must render in a full-width layout optimized for dashboard use rather than a centered article-style container.
2. FR-2: The workflow topology section must be labeled `Work Graph`.
3. FR-3: The dashboard must remove the current live workstation dashboard, terminal summary, and low-value explanatory copy identified in the customer ask.
4. FR-4: The work graph must support a more expanded automatic layout and direct panning.
5. FR-5: The dashboard must render workstation definitions, work items, resources, constraints, and state positions with clearly different visual semantics.
6. FR-6: Active executions must remain visually associated with their owning workstation by `DispatchID`.
7. FR-7: The dashboard must distinguish workstation selection from work-item selection in both highlight behavior and detail labeling.
8. FR-8: The selection inspector must float above the graph and support collapsed and expanded states.
9. FR-9: The dashboard must render trace and selected-item detail in a full-width section beneath the graph.
10. FR-10: Clicking completed or failed items must open lower-page detail for the chosen item.
11. FR-11: Execution timing values must default to human-readable approximate durations.
12. FR-12: The dashboard must provide separate widgets or panels for workstation information, work information, traces, and aggregate metrics.
13. FR-13: The dashboard must provide time-series style visualizations for throughput, failures, retries, rework, and timing trends using data already available to the website where possible.
14. FR-14: At least one metric panel must allow operators to change the displayed time span.
15. FR-15: The website must include automated UI coverage for representative workflow graph sizes, element compositions, and responsive widths.
16. FR-16: Storybook coverage must exercise graph renderer behavior for both minimal and larger workflow topologies.
17. FR-17: Automated interaction coverage must verify graph panning and any supported node-movement behavior.
18. FR-18: Automated layout coverage must verify dashboard graph widgets remain usable at supported viewport widths.

## Non-Goals

- Building a backend telemetry pipeline, new observability service, or cross-system monitoring platform in this phase.
- Replacing Grafana or creating a general-purpose dashboard product for the whole broader platform.
- Redesigning unrelated website areas outside the Agent Factory workflow dashboard.
- Defining new backend APIs solely for this PRD; if required data is unavailable, that work should be captured in a separate follow-up PRD.
- Delivering alerting, paging, anomaly detection, or SLO management.

## Design Considerations

- Prefer minimal text and high information density, but preserve the labels required to distinguish workstation, work item, resource, constraint, retry, and trace states.
- Keep the graph as the visual anchor of the page, with supporting drill-down detail below it rather than beside it.
- Use distinct color, shape, icon, or border semantics so that operators can identify resources and constraints without reading legends repeatedly.
- Treat widget separation as an information architecture decision, not just styling. Workstation, work, trace, and aggregate metrics should not compete for the same panel space.
- Preserve usability on common desktop widths first. Responsive behavior for smaller screens should degrade cleanly, but this dashboard is primarily an operator surface.

## Technical Considerations

- Scope is limited to `website/` changes. The PRD assumes the redesign should use existing website-accessible data sources where possible.
- The website intent document says the website is not the system of record for execution or observability. The redesign should therefore improve operator visibility without turning the website into a backend analytics platform.
- Existing graph rendering, selection, and panel composition patterns in the dashboard should be reused where they help, but they should not constrain the redesign if they are causing the current usability issues.
- The `approach graph UX` package direction is preferred when it materially accelerates widget-board, metrics, or trace panel work, but it is not a hard requirement if it conflicts with existing website patterns or timeline constraints.
- If implementation discovers missing data for failure categorization, retry semantics, or time-series views, the team should capture those gaps as explicit follow-up backend or telemetry work instead of improvising hidden API changes.
- The website feature process already calls for Storybook viewport coverage and browser-backed responsive checks where overflow or layout behavior matters. This dashboard redesign should treat those patterns as required validation, not optional polish.
- Graph renderer test fixtures should intentionally cover topology extremes and semantic mixes, including 1-node, medium-size, and approximately 20-node cases, so layout regressions are caught before manual QA.

## Success Metrics

- Operators can identify the current selected workstation or work item correctly in usability review without asking for clarification.
- Operators can answer "what failed?" and "what is slowing us down?" from the dashboard faster than with the current version.
- Operators can inspect completed and failed work details without losing graph context.
- The team gains visible trend views for throughput, failures, and rework over a chosen time window.
- The redesign reduces reliance on dense descriptive copy to explain the dashboard.

## Open Questions

- [ ] Which existing website data sources are available today for aggregated failure-cause reporting and rework-rate charts?
-- we already have the result error messages, we should use those, and for rework rates we should measure the number of reject operations based on the trace ids for say a non repeater workstation. 
- [ ] Which graph interactions should be supported beyond panning in this phase, such as zooming, minimaps, or fit-to-screen controls?
- zoom/pan is sufficient
- [ ] How much widget rearrangement should be enabled in the first release versus left as a future enhancement?
- it should be enabled completely by default given we plan to put all the current graph featuers as react-grid-layout should support this behavior
- [ ] What is the best operator-facing taxonomy for failure causes when raw provider or process errors are noisy?
- unclear, we should iiterate doy our best effort try
- [ ] Should completed and failed drill-downs share one unified lower panel component or render as mode-specific views?
- separate properly. 

We should note that drill downs should be closeable or hideable as well so that they don't appear forever. 


## Changelog

| Date | Change | Author |
|------|--------|--------|
| 2026-04-09 | Initial version | Codex |
| 2026-04-09 | Added automated Storybook and interaction test requirements for graph rendering, responsive layout, and dashboard usability | Codex |
