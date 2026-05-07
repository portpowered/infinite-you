# meta view

## world state

- as of `2026-05-07T12:06:28.9599999-07:00`, this meta refresh branch
  `meta-refresh-world-state-20260507-110000` has been rebased onto live
  `origin/main`, which now includes merged PR `#162`
  (`localize-dashboard-flow-axis-legend-copy`) at `2026-05-07T17:24:30Z`
  and keeps open PR `#164`
  (`localize-terminal-work-card-copy`) active at `2026-05-07T18:23:49Z`
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
  `localize-current-selection-workstation-detail-card-copy.md`, and
  `localize-terminal-work-card-copy.md` were stale for active-queue purposes
  because the audit lane already advanced into open PR `#141`, merged
  PRs `#161` and `#160` already landed the earlier queue slots on `main`, and
  the terminal-work lane has now advanced into open PR `#164`
- after pruning that residue, the maintainer-owned ignored queue now carries
  one fresh non-overlapping replacement idea:
  - `localize-workflow-activity-graph-import-copy.md`
- combined with open PR `#141` and open PR `#164`, the active maintainers'
  work now again spans three non-overlapping lanes

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
- the UI localization foothold now reaches both current-selection workstation
  detail and workflow-activity legend, while the repo-owned CLI docs package
  lane is materially closed:
  - merged PR `#160` localized
    `ui/src/features/current-selection/workstation-detail-card.tsx`
  - merged PR `#162` localized
    `ui/src/features/workflow-activity/dashboard-flow-axis-legend.tsx`
  - merged PR `#161` raised `pkg/cli` coverage to `95.7%` on
    `2026-05-07`; the previously suspected `Execute()` wrapper and docs-topic
    writer-failure gaps are already covered
  - open PR `#164` now owns the terminal-work localization lane
  - the next concentrated non-overlapping English-only dashboard surface now
    sits in the workflow-activity graph import overlay and dialog shell

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
  - `#164` `localize-terminal-work-card-copy`
  - `#163` `docs: refresh meta world state`
  - `#152` `docs: refresh meta world state`
  - `#145` `docs: refresh meta world state`
  - `#143` `docs: refresh meta world state`
  - `#141` `audit-repository-against-2026-website-and-backend-checklists`
  - `#139` `docs: refresh meta world state`
  - `#123` `docs: refresh meta world state`
  - `#120` `docs: refresh meta world state`
- open PR `#141` owns the current checklist-audit lane and open PR `#164`
  owns the current terminal-work lane, so fresh backlog work must stay
  outside both surfaces

## next cleanup candidates

- the repo-wide standards audit remains the highest-priority checklist ask on
  live `main`, but PR `#141` already owns that documentation lane
- the terminal-work seam is already owned by open PR `#164`
- the next replacement backlog slot on live `main` is now the workflow-
  activity graph-import copy surface:
  - `ui/src/features/workflow-activity/react-flow-current-activity-card-import.tsx`
    still hardcodes user-facing and accessible copy such as
    `Drop an Infinite You PNG onto this graph to start import.`,
    `Import factory PNG`, `Validating factory PNG`, `Factory import failed`,
    `Dismiss`, and the PNG import error-copy mapping
  - `ui/src/features/workflow-activity/mutation-dialog.tsx` still hardcodes
    the shared dialog-shell copy `Close dialog` and `Mutation flow`
  - focused coverage already exists in
    `ui/src/features/workflow-activity/react-flow-current-activity-card-import.test.tsx`
    and `ui/src/features/workflow-activity/mutation-dialog.test.tsx`
  - the lane is feature-local, implementation-ready, and does not overlap the
    open checklist-audit or terminal-work branches
- the `pkg/cli` docs lane is no longer the next candidate:
  - `go test -cover ./pkg/cli` now reports `95.7%` on `2026-05-07`
  - the remaining uncovered `newDocsTopicCommand` branch is the
    `docscli.Markdown(topic)` error path, not the previously suspected writer
    failure or `Execute()` wrapper path

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
- when a repo-owned coverage package has just merged a focused branch-coverage
  PR, re-run live coverage and inspect exact uncovered lines before re-queueing
  the same seam; stale summaries can miss that the remaining gap moved to a
  different branch entirely
- when the linked external checklist repo omits one requested source such as
  `asks.md`, record that as an evidence gap and continue from the sources that
  are actually retrievable instead of inventing missing checklist content
