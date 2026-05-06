# meta view

## world state

- as of `2026-05-06T02:05:34.3495634-07:00`, local `HEAD` on `main` points to
  `0a22988`
  (`Merge branch 'main' of https://github.com/portpowered/infinite-you`) and is
  ahead of published `origin/main` (`20f0504`) because this checkout merged the
  local `import-niceties` branch into `main` during the repo refresh
- authoritative published repo truth therefore comes from `origin/main`
  together with the current open-PR set, not from local `HEAD` alone
- the canonical maintainer ask surface remains `factory/logs/meta/asks.md`
- the local worktree is not clean:
  - tracked `factory/inputs/**` remains sentinel-only
  - one ignored idea is now staged locally under
    `factory/inputs/idea/default/retire-transition-topology-runtime-lookup-adapter.md`
  - unrelated untracked local residue remains at the repo root and under
    `factory/`

## workflow truth

- `factory/factory.json` still defines five work types: `thoughts`, `idea`,
  `plan`, `task`, and `cron-triggers`
- the checked-in maintainer loop remains:
  `thoughts:init -> ideafy -> thoughts:complete`
  `idea:init -> plan -> idea:to-complete + plan:init`
  `plan:init -> setup-workspace -> plan:complete + task:init`
  `task:init -> process -> task:in-review -> review -> task:to-complete`
  `consume` completes same-name `idea` + `task` pairs once both reach
  `to-complete`
- topology details that still matter:
  - `process` and `review` run in `.claude/worktrees/{{name}}`
  - shared `executor-slot` capacity is `10`; each staffed workstation requests
    `1`
  - hourly `cleaner` emits `cron-triggers:complete`
  - `executor-loop-breaker` fails `task:init` after `process` visit `50`
  - `review-loop-breaker` fails `task:in-review` after `review` visit `10`

## input surface truth

- tracked `factory/inputs/**` content is still sentinel-only:
  - `factory/inputs/BATCH/default/.gitkeep`
  - `factory/inputs/idea/default/.gitkeep`
  - `factory/inputs/plan/default/.gitkeep`
  - `factory/inputs/task/default/.gitkeep`
  - `factory/inputs/thoughts/default/.gitkeep`
- `.gitignore` still ignores live workflow submissions under `factory/inputs/**`
  except those sentinel paths
- the file watcher still enforces the documented three-segment watched-input
  contract and rejects implicit two-segment default-channel fallback paths
- the previously queued ignored idea
  `factory/inputs/idea/default/consolidate-functional-factory-event-tick-helpers.md`
  is stale because merged PR `#114` already landed that exact cleanup on
  `origin/main`
- the current ignored local idea is a fresh replacement operating-state file,
  not checked-in queue truth

## customer-ask truth

- the import/export P0 lane is materially advanced on published `main` through
  merged PRs `#67`, `#68`, `#69`, `#70`, `#71`, `#72`, `#93`, `#109`, and
  `#112`
- the remaining active import follow-up is now owned by open PR `#115`
  (`import-niceties`), which touches `pkg/config`, `pkg/service`,
  `pkg/workers/script.go`, and import/export regression tests
- the bad-interface-for-outputs ask is materially satisfied on published
  `main` through the non-success route-array work already merged earlier
- the selected-work current-selection ask is materially satisfied on published
  `main` through merged PRs `#74`, `#77`, and `#110`
- the submit-work copy ask is satisfied on published `main` through merged PR
  `#75`
- the header verbosity, chart layout, branding/iconography, and button-tone
  asks are materially satisfied on published `main` through merged PRs `#83`,
  `#84`, `#85`, `#86`, `#87`, and `#98`
- the remaining open asks in `factory/logs/meta/asks.md` are broader program
  work rather than one narrow unowned customer-visible regression:
  - standards-migration checklist tracking
  - website `90%` coverage target
  - docs audit
  - manual QA
  - systems-quality documentation

## replay truth

- `factory/logs/agent-fails.json` and
  `factory/logs/agent-fails.replay.json` remain the checked-in replay sample
  pair described in `factory/README.md`
- the replay pair is still historical fixture coverage rather than an exact
  copy of the current workflow contract; it predates `to-complete`, `consume`,
  and the current `executor-slot` capacity of `10`
- replay outcome counts remain unchanged in the sample:
  - `process`: `9 ACCEPTED <COMPLETE>`, `27 CONTINUE <CONTINUE>`
  - `review`: `5 ACCEPTED <COMPLETE>`, `4 REJECTED <REJECTED>`
- one replay rejection payload is still oddly quoted as `"\"<REJECTED>\"\n"`;
  treat that as a fixture quirk rather than current workflow truth

## recent repo movement

- recent merged PRs on published `main` now include:
  - `#114` `consolidate-functional-factory-event-tick-helpers`, merged on
    `2026-05-06T08:18:50Z`
  - `#112` `updated website export to support exporting bundled files`, merged
    on `2026-05-06T07:45:59Z`
  - `#111` `remove-init-default-models`, merged on `2026-05-06T07:09:23Z`
  - `#110` `workstation-request-current-selection-cleanup`, merged on
    `2026-05-06T03:28:00Z`
  - `#109` `inline-supporting-file-content-on-export-and-thin-factory-import`,
    merged on `2026-05-06T03:23:17Z`
- `gh pr list --state open` currently reports two open PRs:
  - `#115` `Import niceties`, opened on `2026-05-06T08:53:55Z`
  - `#113` `docs: refresh meta world state`, opened on `2026-05-06T08:08:40Z`
- PR `#115` owns the current import/export follow-up lane, so new cleanup
  dispatches should avoid `pkg/config`, `pkg/service`, `pkg/workers/script.go`,
  and the affected bootstrap/runtime import tests until that lane merges or
  closes

## next cleanup candidate

- there is no remaining narrow unowned customer-visible ask gap on published
  `main`
- the next non-overlapping cleanup seam is authored-config-only runtime lookup
  bridging in transition-topology normalization:
  - `pkg/config/config_mapper.go` still builds a private
    `factoryConfigWorkstationLookupAdapter`
  - that adapter only exists so `state.NormalizeTransitionTopology(...)` can
    ask whether a workstation is a repeater while mapping authored config
  - `pkg/factory/state/transition_topology.go` is the consumer of that question
  - `pkg/factory/workstationconfig/runtime_lookup.go` should remain the owner
    for true runtime lookup behavior outside this mapping seam
- the next dispatch should retire the private authored-config lookup adapter
  and preserve the existing topology behavior through package tests around
  repeater rejection arcs and default failure arcs

## theory of mind

- when local `main` diverges from published `origin/main`, treat local merge
  state as operating residue and rebuild the worldview from remote truth plus
  open PR ownership before dispatching work
- `factory/inputs/**` must always be reasoned about in two layers:
  checked-in contract versus ignored operating residue
- re-check queued ignored idea files after every merged PR; this loop often
  leaves stale local backlog residue behind even when the live code is already
  clean
- the codebase graph can lag live `main` after branch movement; if graph reads
  still reference deleted packages or handlers after a refresh, verify with
  direct file reads before dispatching a cleanup
- prefer shrinking authored-config-only adapters before deduping test plumbing
  when both are available; production-path simplification reduces more future
  reasoning surface than another test-only helper layer
