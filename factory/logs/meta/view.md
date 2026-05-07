# meta view

## world state

- as of `2026-05-07T09:02:56.1302415-07:00`, this meta refresh rebases
  `meta-refresh-world-state-20260507-050344` onto live `origin/main` through
  merged PR `#158` and merged PR `#157`
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
  `factory/inputs/idea/default/localize-selected-work-dispatch-history-card-copy.md`
  and
  `factory/inputs/idea/default/cover-releasetagcheck-command-owner-parse-failure-branch.md`
  are now stale on live `main` because merged PR `#158` landed on
  `2026-05-07T15:53:40Z` and merged PR `#157` landed on
  `2026-05-07T15:20:04Z`
- after pruning those stale residues, the maintainer-owned ignored queue now
  carries three fresh non-overlapping idea files so the autonomous lane still
  matches the standing ask to keep at least three tasks running:
  - `audit-repository-against-2026-website-and-backend-checklists.md`
  - `localize-current-selection-workstation-detail-card-copy.md`
  - `cover-gocoveragecheck-malformed-percentage-parser-branches.md`

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
  `2026-05-07`: both
  `https://raw.githubusercontent.com/portpowered/checklists/main/asks.md` and
  the live `portpowered/checklists` repository root listing remain missing a
  readable `asks.md`, so that ask reference still needs to be treated as
  unavailable evidence rather than assumed input
- there is still no merged checked-in repo-wide review record mapping this
  repository against those external checklist documents on live `main`; open
  PR `#141` currently owns that audit lane and already stages repo-owned
  follow-up ideas under `tasks/ideas-to-review/`
- the UI localization foothold now reaches the selected-work dispatch-history
  card in addition to the earlier import, export, header, and current-selection
  shell/detail-body work on live `main`:
  - `ui/src/i18n/index.ts`, `ui/src/i18n/locales.ts`, and
    `ui/src/i18n/messages.ts` exist on `main`
  - merged PR `#142` localized and accessibility-hardened the import preview
    dialog
  - merged PR `#147` localized the export dialog and export trigger
  - merged PR `#151` localized the dashboard header timeline and stream-status
    accessibility labels
  - merged PR `#154` localized `no-selection-detail-card.tsx`,
    `execution-details.tsx`, `terminal-work-summary-detail.tsx`, and the
    current-selection locale plumbing
  - merged PR `#156` localized
    `ui/src/features/current-selection/current-selection-detail-layout.tsx`
  - merged PR `#158` localized
    `ui/src/features/current-selection/selected-work-dispatch-history-card.tsx`
    and added focused locale regressions for that surface
  - `ui/src/features/current-selection/workstation-detail-card.tsx` still
    hardcodes concentrated user-facing copy such as `Workstation summary`,
    `Active work`, `Historical requests`, `Historical runs`, `Request history`,
    `Run history`, `Expand`, `Collapse`, `Open request`, `Open request details`,
    `Open work item`, `Work selected`, `Request selected`, and multiple empty
    or unavailable-state strings
  - the workstation-detail surface has focused component tests, but no
    locale-aware regression for default and non-default locales
- the next narrow backend testing seam on live `main` is now
  `cmd/gocoveragecheck`:
  - `go test -cover ./cmd/gocoveragecheck` reports `94.3%` statement coverage
    on `2026-05-07`
  - the remaining command-owner gap stays local to
    `cmd/gocoveragecheck/main.go` and `cmd/gocoveragecheck/main_test.go`
  - parser helper coverage remains concentrated in malformed numeric tokens
    that still reach `strconv.ParseFloat` error branches in
    `parseTotalCoverage()` and `parseZeroCoveragePackagesFromReport()`

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
  - `#158` `localize-selected-work-dispatch-history-card-copy`, merged on
    `2026-05-07T15:53:40Z`
  - `#157` `cover-releasetagcheck-command-owner-parse-failure-branch`, merged
    on `2026-05-07T15:20:04Z`
  - `#156` `localize-current-selection-shell-history-controls`, merged on
    `2026-05-07T14:24:31Z`
  - `#155` `cover-deadcodecheck-write-read-and-stderr-branches`, merged on
    `2026-05-07T14:15:16Z`
  - `#154` `localize-current-selection-shell-and-execution-details`, merged on
    `2026-05-07T13:27:26Z`
  - `#153` `close-gocoveragecheck-helper-runtime-and-parser-branches`, merged
    on `2026-05-07T13:26:42Z`
- `gh pr list --state open` now reports:
  - `#152` `docs: refresh meta world state`
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
- the next narrow UI checklist-alignment seam is the current-selection
  workstation detail card:
  - `ui/src/features/current-selection/workstation-detail-card.tsx` still
    hardcodes dense user-facing copy across the summary, active-work list,
    request or run history labels, toggle labels, CTA labels, and empty-state
    messaging
  - `ui/src/features/current-selection/current-selection-locale.tsx` already
    exists on live `main`, so this is now a feature-local follow-up rather
    than a new i18n foundation project
  - focused tests already exist in
    `ui/src/features/current-selection/workstation-detail-card.test.tsx`, but
    the surface still lacks a locale-aware regression
- the next narrow backend testing seam is `cmd/gocoveragecheck`:
  - `go test -cover ./cmd/gocoveragecheck` reports `94.3%` statement coverage
  - remaining coverage is concentrated in one malformed-percentage parser lane
    that stays inside `parseTotalCoverage()` and
    `parseZeroCoveragePackagesFromReport()`
- the next backup UI seam after that is the dashboard terminal-work message
  lane in `ui/src/features/terminal-work/terminal-work-card.tsx`:
  - the widget still hardcodes `Completed and failed work`,
    `Terminal work outcomes`, row empty states, session-summary fallback copy,
    and generic `Expand`/`Collapse` labels that do not expose row-specific
    accessible names

## theory of mind

- the authoritative world model comes from live `main`, the checked-in workflow
  contract, the canonical ask file, current PR state, and current external
  checklist docs together
- `factory/inputs/**` must still be reasoned about in two layers:
  checked-in contract versus ignored operating residue
- when a previously queued idea lands on `main`, prune the ignored residue and
  refresh the next seam from live code and PR state before queuing anything
  else; merged PRs `#157` and `#158` invalidated two of the three active
  backlog slots within the same cycle
- when the customer ask requires at least three tasks in flight, satisfy that
  ask with three narrow non-overlapping idea files rather than one broad batch
  unless dependency ordering is actually required
- when checklist conformance is still undocumented, queue one audit lane plus
  smaller implementation-ready follow-ups instead of claiming alignment from
  standards intent alone
- when a current-selection dispatch-history localization follow-up merges,
  re-read the adjacent workstation-detail card before jumping to broader
  dashboard widgets; the next concentrated English-only residue can stay
  inside the same feature
- when a repo-owned parser already covers happy-path and missing-line shapes,
  inspect malformed numeric tokens next; regex-matched bad percentages can
  leave narrow parse-error branches uncovered without widening the lane
- when the linked external checklist repo omits one requested source such as
  `asks.md`, record that as an evidence gap and continue from the sources that
  are actually retrievable instead of inventing missing checklist content
