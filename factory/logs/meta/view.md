# meta view

## world state

- as of `2026-05-07T07:05:05.7996914-07:00`, local `HEAD` on
  `meta-refresh-world-state-20260507-050344` points to `74dca31`
  (`docs: refresh meta world state`) and includes live `origin/main` through
  merged PR `#154` and merged PR `#153`
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
  `factory/inputs/idea/default/localize-current-selection-shell-and-execution-details.md`
  and
  `factory/inputs/idea/default/close-gocoveragecheck-helper-runtime-and-parser-branches.md`
  are now stale on live `main` because merged PR `#154` landed on
  `2026-05-07T13:44:01Z` and merged PR `#153` landed on
  `2026-05-07T13:56:22Z`
- after pruning those stale residues, the maintainer-owned ignored queue now
  carries three fresh non-overlapping idea files so the autonomous lane still
  matches the standing ask to keep at least three tasks running:
  - `audit-repository-against-2026-website-and-backend-checklists.md`
  - `localize-current-selection-shell-history-controls.md`
  - `cover-deadcodecheck-write-read-and-stderr-branches.md`

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
  PR `#141` currently owns that audit lane and already stages repo-owned
  follow-up ideas under `tasks/ideas-to-review/`
- the UI localization foothold now reaches the current-selection detail-body
  surfaces in addition to the earlier import, export, and header work on live
  `main`:
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
  - `ui/src/features/current-selection/current-selection-detail-layout.tsx`
    still hardcodes `Current selection`, `Undo selection`, `Undo`,
    `Redo selection`, and `Redo`
  - `ui/src/features/terminal-work/terminal-work-card.tsx` still hardcodes its
    widget title, legend, empty states, status meta copy, and generic
    `Expand`/`Collapse` toggle names
- the next narrow backend testing seam on live `main` is now
  `cmd/deadcodecheck`:
  - `go test ./cmd/deadcodecheck` reports `83.7%` statement coverage on
    `2026-05-07`
  - remaining command-owner gaps stay local to `cmd/deadcodecheck/main.go`
    and `cmd/deadcodecheck/main_test.go`, including the current-report write
    branch, baseline-read branch, and successful stderr passthrough in
    `runDeadcode`

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
  - `#154` `localize-current-selection-shell-and-execution-details`, merged on
    `2026-05-07T13:44:01Z`
  - `#153` `close-gocoveragecheck-helper-runtime-and-parser-branches`, merged
    on `2026-05-07T13:56:22Z`
  - `#151` `localize-dashboard-header-timeline-and-stream-status`, merged on
    `2026-05-07T12:21:01Z`
  - `#150` `retire-template-fields-variadic-worktree-shim`, merged on
    `2026-05-07T12:16:03Z`
  - `#149` `cover-releaseprep-help-and-parse-failure-branches`, merged on
    `2026-05-07T12:12:27Z`
  - `#148` `docs: refresh meta world state`, merged on
    `2026-05-07T12:06:55Z`
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
- the next narrow UI checklist-alignment seam is the current-selection shell
  history-control wrapper:
  - `ui/src/features/current-selection/current-selection-detail-layout.tsx`
    still hardcodes the widget title plus undo and redo button labels even
    after the current-selection detail-body localization landed
  - `ui/src/features/current-selection/current-selection-locale.tsx` and
    `ui/src/features/current-selection/messages/current-selection-shell.ts`
    already exist on live `main`, so this is now a one-file feature-local
    follow-up rather than a new i18n foundation project
- the next narrow backend testing seam is `cmd/deadcodecheck`:
  - `go test ./cmd/deadcodecheck` reports `83.7%` statement coverage
  - remaining coverage is concentrated in package-local branches for writing
    `bin/deadcode-current.txt`, reading
    `docs/internal/development/deadcode-baseline.txt`, and preserving stderr
    passthrough from successful `runDeadcode` execution
- the next backup UI seam after that is the dashboard event-stream message
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
- when a merged localization PR claims a broad dashboard surface, re-open the
  exact touched file before declaring the lane closed; merged PR `#154` still
  left English-only shell chrome in
  `ui/src/features/current-selection/current-selection-detail-layout.tsx`
- when one repo-owned command package just absorbed a large test expansion,
  prefer moving sideways to the next small command-owner testing seam unless
  the remaining same-file follow-up is clearly smaller than changing packages
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
