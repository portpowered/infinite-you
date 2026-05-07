# meta view

## world state

- as of `2026-05-07T12:36:57.0000000-07:00`, this meta refresh branch
  `meta-refresh-world-state-20260507-110000` has been rebased onto live
  `origin/main`, which now includes merged PR `#164`
  (`localize-terminal-work-card-copy`) at `2026-05-07T19:10:30Z` and
  merged PR `#165` (`localize-workflow-activity-graph-import-copy`) at
  `2026-05-07T19:40:24Z`
- the canonical maintainer ask surface remains `factory/logs/meta/asks.md`
- the only tracked local dirtiness outside this meta refresh remains the
  user-maintained canonical ask file `factory/logs/meta/asks.md`
- canonical `factory/inputs/**` remains tracked-sentinel-only in git, while
  live maintainer submissions remain ignored operating state

## workflow truth

- `factory/factory.json` still defines five work types: `thoughts`, `idea`,
  `plan`, `task`, and `cron-triggers`
- the checked-in maintainer loop remains:
  `thoughts:init -> ideafy -> thoughts:complete`
  `idea:init -> plan -> idea:complete + plan:init`
  `plan:init -> setup-workspace -> plan:complete + task:init`
  `task:init -> process -> task:in-review -> review -> task:complete`
- topology details that still matter:
  - `process` and `review` run in `.claude/worktrees/{{name}}`
  - shared `executor-slot` capacity is `10`
  - loop breakers still guard repeated `process` and `review` retries

## input surface truth

- tracked `factory/inputs/**` content is still sentinel-only:
  - `factory/inputs/BATCH/default/.gitkeep`
  - `factory/inputs/idea/default/.gitkeep`
  - `factory/inputs/plan/default/.gitkeep`
  - `factory/inputs/task/default/.gitkeep`
  - `factory/inputs/thoughts/default/.gitkeep`
- `.gitignore` still keeps live workflow submissions under `factory/inputs/**`
  out of normal commits except for those sentinel paths
- the previously ignored idea files
  `audit-repository-against-2026-website-and-backend-checklists.md`,
  `cover-cli-docs-command-surface.md`,
  `localize-current-selection-workstation-detail-card-copy.md`,
  `localize-terminal-work-card-copy.md`, and
  `localize-workflow-activity-graph-import-copy.md` are now stale for
  active-queue purposes because the audit lane already advanced into open
  PR `#141` and merged PRs `#161`, `#160`, `#164`, and `#165` already landed
  the other queue slots on `main`
- after pruning that residue, the maintainer-owned ignored queue should carry
  two fresh non-overlapping replacement ideas:
  - `localize-work-outcome-trend-cards-copy.md`
  - `simplify-loaded-runtime-definition-lookups.md`
- combined with open PR `#141`, those two replacement ideas restore three
  productive non-overlapping maintainer lanes without counting stacked
  meta-refresh PRs as cleanup throughput

## customer-ask truth

- the canonical ask file still carries one broad active quality lane plus the
  autonomy notice through `2026-05-25`
- the active quality asks remain:
  - follow the external website and backend checklists and create alignment
    tasks
  - keep backend and website testing moving toward declared high-coverage goals
  - keep simplifying backend and website ownership where duplicate or stale
    logic remains
  - keep at least three non-overlapping tasks running at a time
  - establish a daily QA run lane focused on whether major features still work
- the external checklist links in `factory/logs/meta/asks.md` still point at
  the live `portpowered/checklists` repository, and the explicit linked docs
  remain the current `2026` checklist revisions:
  - `website-development-checklist.md`
  - `backend-development-checklist.md`
- the linked external `asks.md` source still has an evidence gap on
  `2026-05-07`: both the raw URL and the live repo listing still fail to
  surface a readable `asks.md`, so that reference remains unavailable evidence
  rather than assumed requirements
- there is still no merged checked-in repo-wide review record mapping this
  repository against those external checklist documents on live `main`; open
  PR `#141` currently owns that audit lane
- there is still no checked-in repo-owned daily QA schedule or canonical QA
  work item on live `main`; that remains an acknowledged ask gap rather than a
  completed capability
- the UI localization foothold now reaches terminal-work and the
  workflow-activity graph import flow, while the repo-owned CLI docs package
  lane is materially closed:
  - merged PR `#160` localized
    `ui/src/features/current-selection/workstation-detail-card.tsx`
  - merged PR `#162` localized
    `ui/src/features/workflow-activity/dashboard-flow-axis-legend.tsx`
  - merged PR `#164` localized
    `ui/src/features/terminal-work/terminal-work-card.tsx`
  - merged PR `#165` localized the workflow-activity graph-import overlay,
    import dialog shell, and import-preview touch points
  - merged PR `#161` raised `pkg/cli` coverage to `95.7%` on
    `2026-05-07`; the previously suspected `Execute()` wrapper and docs-topic
    writer-failure gaps are already covered
  - with `#164` and `#165` merged, the next concentrated non-overlapping
    English-only dashboard surface now sits in the work-outcome trend cards
    and shared throughput labels

## replay truth

- `factory/logs/agent-fails.json` and
  `factory/logs/agent-fails.replay.json` remain the checked-in replay sample
  pair described in `factory/README.md`
- the replay pair is still historical fixture coverage rather than an exact
  copy of the current workflow contract
- one replay rejection payload is still quoted oddly as `"\"<REJECTED>\"\n"`;
  treat that as fixture history rather than live workflow behavior

## recent repo movement

- recent merged PRs on `main` now include:
  - `#165` `localize-workflow-activity-graph-import-copy`, merged on
    `2026-05-07T19:40:24Z`
  - `#164` `localize-terminal-work-card-copy`, merged on
    `2026-05-07T19:10:30Z`
  - `#162` `localize-dashboard-flow-axis-legend-copy`, merged on
    `2026-05-07T17:24:30Z`
  - `#161` `cover-cli-docs-command-surface`, merged on
    `2026-05-07T17:12:20Z`
  - `#160` `localize-current-selection-workstation-detail-card-copy`, merged
    on `2026-05-07T16:32:50Z`
  - `#159` `cover-gocoveragecheck-malformed-percentage-parser-branches`,
    merged on `2026-05-07T16:18:02Z`
  - `#158` `localize-selected-work-dispatch-history-card-copy`, merged on
    `2026-05-07T15:22:57Z`
- `gh pr list --state open` now reports:
  - `#163` `docs: refresh meta world state`
  - `#152` `docs: refresh meta world state`
  - `#145` `docs: refresh meta world state`
  - `#143` `docs: refresh meta world state`
  - `#141` `audit-repository-against-2026-website-and-backend-checklists`
  - `#139` `docs: refresh meta world state`
  - `#123` `docs: refresh meta world state`
  - `#120` `docs: refresh meta world state`
- open PR `#141` is now the only productive non-meta feature lane still open,
  so the queue needs two fresh non-overlapping replacements to satisfy the
  standing three-task ask

## next cleanup candidates

- the repo-wide standards audit remains the highest-priority checklist ask on
  live `main`, but PR `#141` already owns that documentation lane
- the next UI replacement backlog slot on live `main` is now the work-outcome
  trend-card copy surface:
  - `ui/src/features/work-outcome/trend-cards.tsx` still hardcodes titles,
    subtitles, summary labels, empty states, and chart/list accessibility copy
    for the failure, rework, and timing cards
  - `ui/src/features/work-outcome/trends.ts` still hardcodes shared
    user-facing labels such as `Queued`, `In-flight`, `Completed`,
    `Failed/retried`, and `Session`
  - focused coverage already exists in
    `ui/src/features/work-outcome/trend-cards.test.tsx`
  - the lane is feature-local, implementation-ready, and does not overlap the
    open checklist-audit branch or the freshly merged terminal-work and
    workflow-activity lanes
- the next backend replacement backlog slot on live `main` is now the loaded
  runtime-definition lookup simplification seam:
  - `pkg/config/runtime_config.go` still carries one redundant internal lookup
    layer through `lookup`, `runtimeDefinitionLookup()`,
    `runtimeDefinitionLookupMaps`, the `runtimeDefinitionConfig` alias, and
    `newRuntimeDefinitionConfig()`
  - the public `interfaces.RuntimeDefinitionLookup` and
    `interfaces.RuntimeConfigLookup` behavior can stay unchanged while that
    package-local indirection is collapsed
  - focused behavior coverage already exists in
    `pkg/config/runtime_config_test.go`,
    `pkg/factory/projections/topology_projection_test.go`,
    `pkg/factory/event_history_test.go`, and
    `pkg/service/factory_test.go`
  - the lane is backend-local, implementation-ready, and does not overlap the
    open checklist-audit branch

## theory of mind

- the authoritative world model comes from live `main`, the checked-in workflow
  contract, the canonical ask file, current PR state, and current external
  checklist docs together
- `factory/inputs/**` must still be reasoned about in two layers:
  checked-in contract versus ignored operating residue
- when a queued idea has advanced into an open PR, prune the ignored local idea
  residue as well as merged residues; otherwise the canonical inbox keeps
  duplicating already-owned work
- when the customer ask requires at least three tasks in flight, open PRs can
  satisfy part of that count, so refill only the truly vacant backlog slots
  instead of keeping three ignored idea files regardless of PR state
- when only one productive non-meta feature PR remains open, refill the queue
  with two fresh non-overlapping ideas rather than counting stacked
  meta-refresh branches as satisfying the three-task ask
- when a current-selection localization seam merges and the workflow-activity
  legend is already owned by an open PR, move the next localization follow-up
  to another dashboard-adjacent feature such as terminal-work rather than
  stacking siblings onto the same owned surfaces
- when a queued ignored idea advances into an open PR, replace that queue slot
  from the next adjacent but non-identical feature seam instead of leaving the
  same idea file in place or hopping back into already-localized header or
  current-selection surfaces
- after the workflow-activity legend merges and terminal-work advances into an
  open PR, the least-overlap localization follow-up is the graph-import overlay
  and mutation-dialog shell inside `workflow-activity`, not another header or
  current-selection lane
- after a queued localization lane merges, prefer the next dashboard-adjacent
  feature that still lacks a `messages/` owner and already has focused tests;
  the work-outcome trend cards now fit that pattern better than revisiting
  recently localized workflow-activity files
- when a repo-owned coverage package has just merged a focused branch-coverage
  PR, re-run live coverage and inspect exact uncovered lines before re-queueing
  the same seam; stale summaries can miss that the remaining gap moved to a
  different branch entirely
- when a runtime config owner exposes the same lookup behavior through an alias
  type, an internal map holder, and a private bounce method, treat that
  package-local indirection as a valid simplification seam once behavior tests
  already pin the public contract
- when the linked external checklist repo omits one requested source such as
  `asks.md`, record that as an evidence gap and continue from the sources that
  are actually retrievable instead of inventing missing checklist content
