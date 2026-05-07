# meta view

## world state

- as of `2026-05-07T06:05:41.3101565-07:00`, local `HEAD` on
  `meta-refresh-world-state-20260507-050344` points to `27bea8c`
  (`localize-dashboard-header-timeline-and-stream-status (#151)`) and matches
  live `origin/main` through merged PR `#151`
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
- the ignored idea files
  `factory/inputs/idea/default/localize-dashboard-header-timeline-and-stream-status.md`
  and
  `factory/inputs/idea/default/retire-template-fields-variadic-worktree-shim.md`
  were stale on live `main` because merged PR `#151` landed on
  `2026-05-07T12:21:01Z` and merged PR `#150` landed on
  `2026-05-07T12:16:03Z`
- after pruning those stale residues, the maintainer-owned ignored queue now
  carries three fresh non-overlapping idea files so the autonomous lane still
  matches the standing ask to keep at least three tasks running:
  - `audit-repository-against-2026-website-and-backend-checklists.md`
  - `localize-current-selection-shell-and-execution-details.md`
  - `close-gocoveragecheck-helper-runtime-and-parser-branches.md`

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
- the external checklist links in `factory/logs/meta/asks.md` still point at
  the live `portpowered/checklists` repository, and the explicit linked docs
  remain the current `2026` checklist revisions:
  - `website-development-checklist.md`
  - `backend-development-checklist.md`
- the linked external `asks.md` source still has an evidence gap on
  `2026-05-07`: `https://raw.githubusercontent.com/portpowered/checklists/main/asks.md`
  does not resolve to a readable checklist document, so that ask reference
  still needs to be treated as unavailable evidence rather than assumed input
- there is still no merged checked-in repo-wide review record mapping this
  repository against those external checklist documents on live `main`; open
  PR `#141` currently owns that audit lane
- the UI localization foothold now reaches the dashboard header in addition to
  import and export surfaces on live `main`:
  - `ui/src/i18n/index.ts`, `ui/src/i18n/locales.ts`, and
    `ui/src/i18n/messages.ts` exist on `main`
  - merged PR `#142` localized and accessibility-hardened the import preview
    dialog
  - merged PR `#147` localized the export dialog and export trigger
  - merged PR `#151` localized the dashboard header timeline and stream-status
    accessibility labels
  - `ui/src/features/current-selection/current-selection-detail-layout.tsx`,
    `no-selection-detail-card.tsx`, `execution-details.tsx`, and
    `terminal-work-summary-detail.tsx` still keep the next concentrated
    hardcoded English dashboard copy block on the live path
- the next narrow backend testing seam on live `main` is now
  `cmd/gocoveragecheck`:
  - `go test -cover ./cmd/...` reports `75.0%` for `cmd/gocoveragecheck`
  - peer repo-owned command packages now sit materially higher after merged
    PR `#149`, leaving the remaining helper/runtime/parser branches as the next
    small command-owner closeout

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
  - `#151` `localize-dashboard-header-timeline-and-stream-status`, merged on
    `2026-05-07T12:21:01Z`
  - `#150` `retire-template-fields-variadic-worktree-shim`, merged on
    `2026-05-07T12:16:03Z`
  - `#149` `cover-releaseprep-help-and-parse-failure-branches`, merged on
    `2026-05-07T12:12:27Z`
  - `#148` `docs: refresh meta world state`, merged on
    `2026-05-07T12:06:55Z`
  - `#147` `localize-export-dialog-and-trigger`, merged on
    `2026-05-07T11:25:43Z`
- `gh pr list --state open` now reports:
  - `#145` `docs: refresh meta world state`
  - `#143` `docs: refresh meta world state`
  - `#141` `audit-repository-against-2026-website-and-backend-checklists`
  - `#139` `docs: refresh meta world state`
  - `#123` `docs: refresh meta world state`
  - `#120` `docs: refresh meta world state`
- open PR `#141` already owns the current checklist-audit lane, so new backlog
  replacements must avoid its doc surface while still acting on live code gaps

## next cleanup candidates

- the repo-wide standards audit remains the highest-priority checklist ask on
  live `main`, but PR `#141` already owns that documentation lane
- the next narrow UI checklist-alignment seam is the current-selection shared
  shell and execution-details surface:
  - `ui/src/features/current-selection/current-selection-detail-layout.tsx`
    still hardcodes `Current selection`, `Undo`, and `Redo`
  - `ui/src/features/current-selection/no-selection-detail-card.tsx`,
    `execution-details.tsx`, and `terminal-work-summary-detail.tsx` still
    hardcode the empty-state guidance, section headings, trace guidance,
    terminal status text, and failure fallback copy
  - the import/export/header feature-local i18n pattern already exists on live
    `main`, so this is now a small feature-local follow-up rather than a new
    i18n foundation project
- the next narrow backend testing seam is `cmd/gocoveragecheck`:
  - `go test -cover ./cmd/...` reports `75.0%` for
    `cmd/gocoveragecheck`
  - remaining coverage is concentrated in package-local helper/runtime/parser
    branches such as combined failure aggregation, repo-root discovery, package
    listing, profile parsing, and path normalization
- the next backup UI seam after that is the dashboard event-stream message
  pipeline:
  - `ui/src/api/events/api.ts` and
    `ui/src/features/dashboard/state/dashboardStreamStore.ts` still hardcode
    English stream lifecycle and fallback messages outside the feature-local
    message catalogs

## theory of mind

- the authoritative world model comes from live `main`, the checked-in workflow
  contract, the canonical ask file, current PR state, and current external
  checklist docs together
- `factory/inputs/**` must still be reasoned about in two layers:
  checked-in contract versus ignored operating residue
- when a previously queued idea lands on `main`, prune the ignored residue and
  refresh the next seam from live code and PR state before queuing anything
  else; merged PRs `#146` and `#147` invalidated two of the three active
  backlog slots within the same cycle
- when the customer ask requires at least three tasks in flight, satisfy that
  ask with three narrow non-overlapping idea files rather than one broad batch
  unless dependency ordering is actually required
- when checklist conformance is still undocumented, queue one audit lane plus
  smaller implementation-ready follow-ups instead of claiming alignment from
  standards intent alone
- when a localization follow-up merges on one dashboard dialog, re-read the
  adjacent header controls and accessibility labels next; the concentrated
  English-only residue often lives there rather than in a second dialog
- when a dashboard header-localization follow-up merges, re-read the adjacent
  current-selection shell and execution-details surfaces next; shared detail
  widgets can still keep the next concentrated block of English-only copy
- when a helper still advertises an explicitly backwards-compatible variadic or
  optional parameter but the live production caller already passes the
  canonical explicit shape, retire the shim before widening into larger runtime
  refactors
- when neighboring repo-owned command packages have already closed their thin
  owner seams, queue the next-lowest command package and keep the follow-up
  local to that command file and tests before widening into application
  packages
- when the linked external checklist repo omits one requested source such as
  `asks.md`, record that as an evidence gap and continue from the sources that
  are actually retrievable instead of inventing missing checklist content
