# meta view

## world state

- as of `2026-05-07T10:03:45.2950540-07:00`, this meta refresh rebases
  `meta-refresh-world-state-20260507-050344` onto live `origin/main` through
  merged PR `#159`, while open PR `#160` now owns the next workstation-detail
  localization lane
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
- the ignored idea file
  `factory/inputs/idea/default/cover-gocoveragecheck-malformed-percentage-parser-branches.md`
  is now stale on live `main` because merged PR `#159` landed on
  `2026-05-07T16:32:36Z`
- after pruning that stale residue, the maintainer-owned ignored queue still
  carries three narrow non-overlapping idea files so the autonomous lane
  matches the standing ask to keep at least three tasks running:
  - `audit-repository-against-2026-website-and-backend-checklists.md`
  - `localize-current-selection-workstation-detail-card-copy.md`
  - `cover-cli-docs-command-surface.md`

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
  - open PR `#160` now owns the remaining concentrated workstation-detail copy
    and locale-regression lane across
    `ui/src/features/current-selection/workstation-detail-card.tsx`,
    adjacent feature-local messages, and focused tests
  - live `main` therefore still needs that PR to land before the surface stops
    hardcoding the remaining workstation-detail copy
- the next narrow backend testing seam on live `main` is now `pkg/cli`:
  - `go test -cover ./pkg/cli` reports `93.5%` statement coverage on
    `2026-05-07`
  - `pkg/cli/root.go` leaves `newDocsTopicCommand` at `83.3%` and `Execute`
    at `0.0%` in focused coverage output
  - existing tests in `pkg/cli/root_test.go` already cover docs help and
    successful markdown rendering, but they still miss the docs-topic writer
    failure path and the thin `Execute()` wrapper path

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
  - `#159` `cover-gocoveragecheck-malformed-percentage-parser-branches`,
    merged on `2026-05-07T16:32:36Z`
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
  - `#160` `localize-current-selection-workstation-detail-card-copy`
  - `#152` `docs: refresh meta world state`
  - `#145` `docs: refresh meta world state`
  - `#143` `docs: refresh meta world state`
  - `#141` `audit-repository-against-2026-website-and-backend-checklists`
  - `#139` `docs: refresh meta world state`
  - `#123` `docs: refresh meta world state`
  - `#120` `docs: refresh meta world state`
- open PR `#141` already owns the current checklist-audit lane, and open
  PR `#160` already owns the next concentrated current-selection UI lane, so
  new backlog replacements must avoid both surfaces while still acting on live
  code gaps

## next cleanup candidates

- the repo-wide standards audit remains the highest-priority checklist ask on
  live `main`, but PR `#141` already owns that documentation lane
- the next concentrated current-selection UI seam is already owned by open
  PR `#160`, so the next replacement backlog slot should stay outside that
  feature until the PR lands
- the next narrow backend testing seam is `pkg/cli/root.go`:
  - `go test -cover ./pkg/cli` reports `93.5%` statement coverage
  - remaining gaps stay local to `newDocsTopicCommand` and `Execute`
  - the missing assertions are behavioral docs-command owner branches, not a
    broad runtime startup gap
- the next backup UI seam after that is the dashboard flow-axis legend in
  `ui/src/features/workflow-activity/dashboard-flow-axis-legend.tsx`:
  - the widget still hardcodes `Graph legend`, `Legend`, `Expand graph legend`,
    `Collapse graph legend`, `Active flow`, `Failure path`, and default icon
    labels
  - focused component, parent-card, and Storybook coverage already exist, but
    they still assert the English vocabulary directly instead of a locale-backed
    message source

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
- when that adjacent current-selection follow-up has already advanced into an
  open PR, refill the third backlog slot with a non-overlapping repo-owned
  backend package seam instead of stacking another sibling UI idea onto the
  same feature
- when a repo-owned command or package already covers happy-path docs rendering,
  inspect thin writer-failure and wrapper branches next; those can keep small
  owner coverage gaps alive without requiring broader runtime refactors
- when the linked external checklist repo omits one requested source such as
  `asks.md`, record that as an evidence gap and continue from the sources that
  are actually retrievable instead of inventing missing checklist content
