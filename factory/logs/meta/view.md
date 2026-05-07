# meta view

## world state

- as of `2026-05-07T11:03:26.5424733-07:00`, this meta refresh runs from
  branch `meta-refresh-world-state-20260507-110000` created directly from live
  `origin/main`, which now includes merged PR `#160`
  (`localize-current-selection-workstation-detail-card-copy`) at
  `2026-05-07T17:53:59Z` and merged PR `#161`
  (`cover-cli-docs-command-surface`) at `2026-05-07T17:25:09Z`
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
  `cover-cli-docs-command-surface.md`, and
  `localize-current-selection-workstation-detail-card-copy.md` were stale for
  active-queue purposes because the audit lane already advanced into open
  PR `#141` and merged PRs `#161` and `#160` already landed the latter two
  lanes on `main`
- after pruning that residue, the maintainer-owned ignored queue now carries
  one fresh non-overlapping replacement idea:
  - `localize-terminal-work-card-copy.md`
- combined with open PR `#141` and open PR `#162`, the active maintainers'
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
  detail and the repo-owned CLI docs package lane is materially closed:
  - merged PR `#160` localized
    `ui/src/features/current-selection/workstation-detail-card.tsx`
  - merged PR `#161` raised `pkg/cli` coverage to `95.7%` on
    `2026-05-07`; the previously suspected `Execute()` wrapper and docs-topic
    writer-failure gaps are already covered
  - open PR `#162` now owns the workflow-activity legend localization lane
  - `ui/src/features/terminal-work/terminal-work-card.tsx` is now the next
    concentrated non-overlapping English-only dashboard surface

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
  - `#160` `localize-current-selection-workstation-detail-card-copy`, merged
    on `2026-05-07T17:53:59Z`
  - `#161` `cover-cli-docs-command-surface`, merged on
    `2026-05-07T17:25:09Z`
  - `#159` `cover-gocoveragecheck-malformed-percentage-parser-branches`,
    merged on `2026-05-07T16:18:02Z`
  - `#158` `localize-selected-work-dispatch-history-card-copy`, merged on
    `2026-05-07T15:22:57Z`
  - `#157` `cover-releasetagcheck-command-owner-parse-failure-branch`, merged
    on `2026-05-07T15:12:40Z`
- `gh pr list --state open` now reports:
  - `#162` `localize-dashboard-flow-axis-legend-copy`
  - `#152` `docs: refresh meta world state`
  - `#145` `docs: refresh meta world state`
  - `#143` `docs: refresh meta world state`
  - `#141` `audit-repository-against-2026-website-and-backend-checklists`
  - `#139` `docs: refresh meta world state`
  - `#123` `docs: refresh meta world state`
  - `#120` `docs: refresh meta world state`
- open PR `#141` owns the current checklist-audit lane and open PR `#162`
  owns the current workflow-activity legend lane, so fresh backlog work must
  stay outside both surfaces

## next cleanup candidates

- the repo-wide standards audit remains the highest-priority checklist ask on
  live `main`, but PR `#141` already owns that documentation lane
- the current workflow-activity legend seam is already owned by open PR `#162`
- the next replacement backlog slot on live `main` is now the terminal-work
  detail card:
  - `ui/src/features/terminal-work/terminal-work-card.tsx` still hardcodes
    user-facing and accessible copy such as `Completed and failed work`,
    `Terminal work outcomes`, `Completed`, `Failed`, `Expand`, `Collapse`,
    `Completed work`, `Failed work`, and the empty and fallback summary text
  - focused coverage already exists in
    `ui/src/features/terminal-work/terminal-work-card.test.tsx` and
    `ui/src/features/terminal-work/terminal-work-card.stories.tsx`
  - the lane is feature-local, implementation-ready, and does not overlap the
    open checklist-audit or workflow-activity legend branches
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
- when a repo-owned coverage package has just merged a focused branch-coverage
  PR, re-run live coverage and inspect exact uncovered lines before re-queueing
  the same seam; stale summaries can miss that the remaining gap moved to a
  different branch entirely
- when the linked external checklist repo omits one requested source such as
  `asks.md`, record that as an evidence gap and continue from the sources that
  are actually retrievable instead of inventing missing checklist content
