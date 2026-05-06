# meta view

## world state

- as of `2026-05-05T22:05:18.8628642-07:00`, the local checkout is on
  `branch-check` at `2e7a3b5` (`Merge remote-tracking branch 'origin/main'
  into branch-check`)
- local `branch-check` now contains `origin/main` at `41ef8d9`
  (`Merge pull request #108 from portpowered/ralph/remove-deadcode-2026-may`)
- local branch-only commit `4d731a3` (`docs: refresh meta world state`) was
  already merged by PR `#107` (`Branch check`)
- the local worktree is not clean:
  - ignored workflow residue is visible under `factory/inputs/**`
  - untracked local residue includes
    `factory/logs/dedupe-functional-agent-config-and-arg-sequence-test-helpers.md`
    and `test-export.png`
- the canonical maintainer ask surface remains `factory/logs/meta/asks.md`

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
  contract and no longer accepts direct
  `factory/inputs/<work-type>/<file>` submissions as an implicit `default`
  channel fallback
- the visible canonical inboxes now contain the tracked `.gitkeep` sentinels
  plus two still-live ignored idea files:
  - `factory/inputs/idea/default/dedupe-functional-api-server-harnesses.md`
  - `factory/inputs/idea/default/consolidate-static-command-runner-test-helpers.md`
- the stale ignored task residue
  `factory/inputs/task/default/code-coverage-frontend.md` has been pruned
  locally because merged PR `#103` already closed that lane on `main`
- the ignored cleanup idea
  `factory/inputs/idea/default/retire-config-public-enum-forwarder-wrappers.md`
  has also been pruned locally because open PR `#108`
  (`remove-deadcode-2026-may`) already owns that exact enum-wrapper seam

## customer-ask truth

- the import/export P0 lane is now materially closed on `main`:
  - merged PRs `#67`, `#68`, `#69`, `#70`, `#71`, `#72`, `#93`, and `#109`
    now cover the prompt template, split-layout, non-success-route, dialog,
    and portable bundled-file portions of that ask
  - canonical export now includes inline content for the supported portable
    script-backed supporting-file set so exported payloads are self-contained
  - import now materializes those supported files back to disk at their
    declared target paths, overwrites differing existing files with an
    observable reportable signal, and persists a thin imported `factory.json`
    without the inline file bodies
  - package, integration, service, CLI, and functional coverage now protect
    the supported export/import round-trip behavior through payloads, written
    files, persisted authored layout, and runtime loading
- the broader selected-work current-selection ask is materially satisfied on
  `main` through merged PRs `#74`, `#77`, and `#99`
- one narrower current-selection cleanup lane is still in flight as open
  PR `#110` (`workstation-request-current-selection-cleanup`), which removes
  duplicate inference detail from workstation-request dispatch summaries while
  preserving the export dialog filename follow-up on the same branch
- the submit-work copy ask is satisfied on `main` through merged PR `#75`
- merged PR `#83` satisfied the header-verbosity copy reduction ask on `main`:
  visible `Factory state`, `Stream`, `Export PNG`, and `Current` toolbar text
  labels are gone
- merged PR `#84` satisfied the work-outcome chart ask on `main`:
  the chart no longer teaches detached axis labels and now carries increased
  axis spacing through rendered chart behavior
- merged PR `#85` satisfied the remaining dashboard branding and header
  iconography lane on `main`:
  - `ui/index.html`, `ui/fallback_dist/index.html`, and
    `ui/fallback_dist/assets/index.js` no longer ship the old
    `Agent Factory` / triangle shell contract
  - `ui/src/features/header/dashboard-header.tsx`,
    `ui/src/features/header/dashboard-status-panel.tsx`,
    `ui/src/features/bento/agent-bento.tsx`,
    `ui/src/features/workflow-activity/react-flow-current-activity-card-import.tsx`,
    `ui/src/features/header/dashboard-export-dialog.tsx`, and
    `ui/src/features/export/build-factory-export-filename.ts` now align with
    the Infinite You rename
  - the header stream/export/current controls now use the requested
    pulsating-dot, share-style, and play-style semantics
- merged PR `#86` satisfied the remaining header button-convergence seam on
  `main`:
  - `ui/src/features/header/dashboard-header.tsx` and
    `ui/src/features/header/tick-slider-control.tsx` now route the export and
    return-to-current actions through
    `ui/src/features/header/dashboard-header-action-button.tsx`
  - `ui/src/features/header/dashboard-header.test.tsx` and
    `ui/src/App.stories.tsx` now protect the converged neutral header-action
    treatment through rendered behavior
- merged PR `#87` satisfied the remaining submit-work button-tone seam on
  `main`:
  - `ui/src/features/submit-work/submit-work-card.tsx` now renders the
    `Submit work` CTA through the shared neutral `outline` tone instead of the
    accent-filled default treatment
  - `ui/src/features/submit-work/submit-work-widget.test.tsx` and
    `ui/src/App.stories.tsx` now protect that neutral rendered treatment
- the tracked local ask diff now explicitly includes the array-valued
  non-success output-interface request, but that ask is already materially
  satisfied on `main` through merged PR `#69`
- there is no remaining unowned customer-visible dashboard ask on `main`; the
  next narrow cleanup candidate now comes from the broader backend quality
  lane:
- merged PR `#88` retired the duplicated legacy dispatch compat payload
  aliases from `ui/src/features/timeline/state/timeline/replayWorldStateTypes.ts`
  and merged PR `#89` centralized the remaining replay
  `completionToTraceDispatch` projection duplication without reopening the
  broader legacy compat lane
- merged PR `#90` closed the authored-throttle transition-ID fallback lane on
  `main`:
  - `pkg/petri/inference_throttle_guard.go` no longer carries
    `WatchedTransitionIDs`
  - `pkg/config/config_mapper.go` no longer lowers authored throttle guards
    through transition-ID watch sets
  - throttle tests now assert runtime-observable lane behavior instead of the
    retired fallback shape
- merged PR `#91` closed the last hidden non-success route compatibility seam
  on `main`:
  - `pkg/config/openapi_factory.go` no longer carries the singular
    `onContinue` / `onRejection` / `onFailure` coercion helper
  - `pkg/api/factory_config_smoke_test.go` now rejects singular non-success
    route objects at the generated boundary
  - checked-in factory fixtures across `factory/`, `examples/`, and
    `tests/functional_test/testdata/` now author canonical array-valued
    non-success routes
- merged PR `#92` closed the stale OpenAPI cron post-processing compatibility
  seam on `main`:
  - `pkg/config/factory_config_mapping.go` no longer calls
    `applyOpenAPICronCompatibility`
  - `pkg/config/openapi_factory.go` no longer reparses normalized authored JSON
    through `buildRawOpenAPIWorkstationCronIndex`
  - `pkg/config/openapi_factory_test.go` and
    `tests/functional/runtime_api/api_runtime_config_alignment_smoke_test.go`
    now cover canonical cron loading through the generated boundary and runtime
    config path
- merged PR `#93` closed the supported portable bundled-file portability seam
  on `main`:
  - `pkg/config/factory_config_mapping.go` now strips supported portable inline
    bundled-file content from canonical export while preserving the portable
    transport entries
  - `pkg/config/runtime_config.go`, `pkg/config/layout.go`, and
    `pkg/config/config_validator.go` now treat supported portable bundled files
    as disk-backed during runtime loading, expansion, and validation
  - `pkg/config/portable_bundled_files*.go`, `pkg/service/factory_test.go`,
    and functional portability smoke coverage now protect the supported
    export/import/runtime round-trip behavior
- merged PR `#94` closed the remaining replay-wrapper cleanup seam on `main`:
  - `ui/src/features/timeline/state/timeline/replayCompletion.ts` and
    `ui/src/features/timeline/state/timeline/replayWorldState.ts` no longer
    carry the cast-wrapper-only replay payload helpers
  - replay timeline tests remain the proof path for the live replay behavior
- merged PR `#95` closed the duplicated submit-request shaping drift on `main`:
  - `pkg/factory/work_request.go` now owns the canonical
    `[]interfaces.SubmitRequest -> interfaces.WorkRequest` shaping path
  - `pkg/internal/submission/work_request.go` is now only a pass-through alias
    onto `factory.WorkRequestFromSubmitRequests`
  - functional runtime smoke coverage now protects trace and request-shaping
    parity through the runtime-visible submission path
- merged PR `#96` closed the dead submission-forwarder lane on `main`:
  - `pkg/internal/submission` is gone
  - live callers now import `pkg/factory.WorkRequestFromSubmitRequests`
    directly
  - the old meta note naming that package as the next cleanup candidate is now
    stale
- merged PR `#97` closed the last stale iconography branch residue without
  reopening the broader dashboard ask on `main`
- merged PR `#98` closed the remaining work-outcome chart padding follow-up on
  `main`
- merged PR `#100` closed the replay event-stream file-wrapper seam on `main`:
  - `pkg/replay/event_stream_artifact.go` is gone from production ownership
  - replay conversion behavior now lives in the replay contract smoke lane
- merged PR `#101` closed the backend `80%` coverage lane on `main`:
  - `Makefile` and `cmd/gocoveragecheck/main.go` now enforce the backend gate
  - backend runtime, replay, worker, and projection tests now cover the raised
    floor through observable behavior
- merged PR `#104` closed the replay-only shared-helper lane on `main`:
  - replay-only helper ownership no longer inflates the shared exported surface
    of `pkg/testutil` and `internal/testpath`
  - replay artifact smoke and replay regression coverage still protect copied
    factory replay success plus expected divergence behavior through observable
    runtime outcomes
- merged PR `#105` closed the duplicate functional dispatch-history helper
  lane on `main`:
  - `tests/functional/runtime_api/api_cron_workstations_smoke_test.go` now
    routes dispatch-input reconstruction through
    `tests/functional/internal/support/events.go`
  - `tests/functional/internal/support/events.go` no longer carries the extra
    local helper breadth that only existed to support the cron smoke's private
    copy
  - guards-batch and cron runtime assertions still prove the same observable
    history behavior through live factory events
- there is still no remaining narrow unowned customer-visible ask gap on
  `main`; merged PR `#106` closed the previously recorded shared functional
  helper seam:
  - `tests/functional/guards_batch/helpers_test.go`,
    `tests/functional/runtime_api/runtime_support_test.go`, and
    `tests/functional/smoke/service_config_override_alignment_test.go` no
    longer own duplicate `writeAgentConfig` / `assertArgsContainSequence`
    helpers already centralized under
    `tests/functional/internal/support/fixtures.go`
  - the older worldview note that still named that seam as next-up was stale
- merged PR `#108` closed the deadcode batch on `main`, including the
  pass-through public enum wrapper seam that the prior worldview still treated
  as open:
  - `pkg/config/public_factory_enums.go` is now reduced to the surviving live
    runtime helper path instead of the older wider forwarding surface
  - dead dashboard, logger, and test-only residue named in the May batch are
    now recorded as closed in the new deadcode closeout docs under
    `docs/development/`
- the next non-overlapping cleanup candidates now come from the broader
  backend quality lane and remain queued as ignored local idea files:
  - dedupe the near-copy functional API server harness ownership across
    `tests/functional/bootstrap_portability/helpers_test.go`,
    `tests/functional/runtime_api/helpers_test.go`,
    `tests/functional/replay_contracts/replay_live_helpers_test.go`, and the
    exported runtime helper in
    `tests/functional/runtime_api/functional_server_test.go`
  - consolidate duplicated static command-runner test helpers now split across
    `tests/functional/providers/helpers_test.go`,
    `tests/functional/bootstrap_portability/portability_helpers_test.go`,
    `tests/functional/workflow/helpers_test.go`,
    `tests/functional/replay_contracts/replay_script_boundary_events_test.go`,
    and `tests/functional/runtime_api/api_inference_events_test.go`
  - that command-runner lane remains behaviorally valid, but its queue note now
    explicitly records that `tests/functional/internal/support/command_runner.go`
    only covers the success/stdout case today and that
    `tests/functional/runtime_api/api_inference_events_test.go` is also touched
    by open PR `#110`
- the remaining ask surface beyond that is broader program work:
  - the general standards-migration checklist ask is still open in
    `factory/logs/meta/asks.md`
  - the website `90%` coverage target is still open in
    `factory/logs/meta/asks.md`
  - the manual QA and systems-quality documentation asks are still open in
    `factory/logs/meta/asks.md`

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

## recent repo movement

- recent merged PRs on `main` now include:
  - `#108` `remove-deadcode-2026-may`, merged on `2026-05-06T04:08:33Z`
  - `#109` `inline-supporting-file-content-on-export-and-thin-factory-import`,
    merged on `2026-05-06T03:48:26Z`
  - `#107` `Branch check`, merged on `2026-05-06T02:56:30Z`
  - `#105` `dedupe-functional-dispatch-history-test-helpers`, merged on
    `2026-05-05T23:16:03Z`
  - `#104` `trim-replay-only-testutil-helper-surface`, merged on
    `2026-05-05T22:23:43Z`
  - `#102` `work-chaininig-trace-ids`, merged on `2026-05-05T21:39:26Z`
  - `#99` `workstation-current-selection-cleanup`, merged on
    `2026-05-05T21:20:11Z`
  - `#101` `code-coverage-backend`, merged on `2026-05-05T21:29:09Z`
  - `#100` `retire-replay-event-stream-file-wrapper-cluster`, merged on
    `2026-05-05T21:22:34Z`
  - `#98` `work-outcome-chart-padding`, merged on `2026-05-05T21:02:23Z`
  - `#97` `website-icon-removal`, merged on `2026-05-05T20:59:43Z`
  - `#96` `remove-dead-submission-work-request-forwarder-package`, merged on
    `2026-05-04T18:22:54Z`
  - `#95` `dedupe-submit-request-shaping-between-production-and-functional-runtime-tests`,
    merged on `2026-05-04T17:34:28Z`
  - `#94` `retire-replay-dispatch-payload-cast-wrappers`, merged on
    `2026-05-04T17:21:31Z`
  - `#93` `canonicalize-portable-bundled-files-without-inline-content`,
    merged on `2026-05-04T15:53:47Z`
  - `#92` `retire-openapi-cron-compatibility-patch`, merged on
    `2026-05-04T14:33:28Z`
  - `#91` `retire-singular-workstation-route-compatibility`, merged on
    `2026-05-04T13:28:02Z`
  - `#90` `retire-authored-throttle-transition-id-fallback`, merged on
    `2026-05-04T12:32:50Z`
  - `#89` `centralize-replay-trace-dispatch-projection`, merged on
    `2026-05-04T11:25:59Z`
  - `#88` `retire-replay-timeline-legacy-compat-duplication`, merged on
    `2026-05-04T10:20:32Z`
  - `#87` `normalize-submit-work-button-treatment`, merged on
    `2026-05-04T09:17:47Z`
  - `#86` `normalize-dashboard-header-button-treatment`, merged on
    `2026-05-04T08:36:11Z`
  - `#85` `align-dashboard-branding-and-header-iconography`, merged on
    `2026-05-04T08:00:00Z`
  - `#69` `workstation-non-success-route-arrays`, merged into `main` before the
    current refresh and now represented by `HEAD`
  - `#84` `align-work-outcome-chart-axis-labels-and-margins`, merged on
    `2026-05-04T06:31:54Z`
  - `#83` `simplify-dashboard-header-toolbar-verbosity`, merged on
    `2026-05-04T05:53:10Z`
  - `#82` `trim-ralph-starter-input-readme-contract`, merged on
    `2026-05-04T04:20:52Z`
  - `#81` `trim-starter-input-readme-contract`, merged on
    `2026-05-04T03:18:53Z`
  - `#80` `align-default-starter-task-input-contract`, merged on
    `2026-05-04T02:36:26Z`
  - `#79` `retire-filewatcher-default-channel-fallback`, merged on
    `2026-05-04T01:27:11Z`
  - `#78` `remove-list-work-legacy-pagination-shim`, merged on
    `2026-05-04T00:28:40Z`
- `gh pr list --state open` currently reports one open PR:
  - `#110` `workstation-request-current-selection-cleanup`, opened on
    `2026-05-06T03:28:00Z`

## theory of mind

- the authoritative world model still comes from live git state plus the
  checked-in workflow contract, not from replay fixtures alone
- `factory/inputs/**` must always be reasoned about in two layers:
  checked-in contract versus ignored operating residue
- ignored queue residue can become stale within a single merge cycle, so the
  meta loop has to reconcile local inbox files against the newest merged PRs
  before dispatching anything new
- merged PR state can flip between the start of a refresh and the end of it;
  open-PR assumptions need to be revalidated after `git pull` and before queue
  writes
- once a customer-visible ask lands, the right follow-up is usually not another
  broad lane but the smallest remaining seam that preserves the ask's public
  intent without reopening merged work
- once a portability lane looks closed, re-check export mapping, runtime load,
  validation, and service/functional tests together before retiring the idea;
  PR `#93` only became safe to close after all four layers moved to the same
  supported disk-backed ownership model
- once a narrow cleanup merges, the next best seam can still hide in boundary
  normalizers and fixtures rather than in the obvious runtime path; after
  PR `#90`, the stale throttle note was gone on `main` but the array-route ask
  still had a hidden compatibility layer in the OpenAPI factory boundary
- when a public JSON shape moves from singular objects to arrays, re-check
  import/load normalization helpers and checked-in factory fixtures; type and
  schema updates alone do not prove the old authored shape is actually rejected
- after a boundary migration lands, re-check for follow-on post-processing
  patches that still reparse the authored JSON out-of-band; once the generated
  model and mapper both carry the canonical field directly, those patches are
  often the next dead compatibility seam
- for import/export asks, verify both canonical `factory.json` flattening and
  expanded authored layout writes before declaring the lane closed; stripping a
  field during `WriteExpandedFactoryLayout` does not mean `FlattenFactoryConfig`
  and runtime loading have stopped rehydrating it back into the exported JSON
- for portable supporting files, export and import can intentionally be
  asymmetric: self-contained export may need inline bytes even when the correct
  imported steady state is a thin disk-backed `factory.json`
- after broad replay legacy-compat cleanups merge, scan for any remaining
  one-line cast wrappers or alias-only helpers on the live replay path before
  inventing a larger follow-up; those seams are often the next lowest-risk
  simplifications
- when production request-shaping and functional-test request-shaping both
  exist, compare chaining-trace propagation and batch metadata field-by-field
  before assuming the duplicate helper is harmless; this repo already allowed a
  test-only copy to drift away from the live submit path
- once a dedupe PR lands, re-scan the old owner package for alias-only
  wrappers; this repo now shows that a merged centralization can leave a whole
  package behind as dead forwarding surface even after the behavioral drift is
  fixed
- once a cleanup seam is recorded in the checked-in worldview, re-validate it
  against live `main` before dispatching; merged PR `#96` proved the meta view
  can go stale within a single refresh cycle
- when a broad cleanup batch PR is open, compare queued narrow ideas against
  the actual changed-file ownership before treating them as separate work; PR
  `#108` already absorbed the public-enum forwarder cleanup even though that
  seam was still queued locally
- after a broad cleanup batch merges, re-validate the exact remaining helper
  owners before keeping a queued idea unchanged; this pass showed both
  surviving queued ideas were still real, but each needed more precise
  ownership notes after `#108` landed
- when an open PR touches a surface already marked closed in the worldview,
  compare the PR diff against the canonical ask text before dispatching any
  sibling work; PR `#97` is a good example of overlap hiding behind a new
  branch name
- deadcode-baseline entries that are only reachable from one long functional
  smoke are strong candidates for the next narrow cleanup idea, especially when
  the live runtime, API, and CLI paths have no direct callers
- when a deadcode-baseline entry still has live callers, compare it against
  suite-local helper copies before queueing deletion; duplicated functional
  helper ownership can be the real simplification seam
- when a helper file is split by build tags, validate usage in both the short
  and `functionallong` lanes before treating it as dead; replay-contract helper
  files can look unused in one lane while still serving the other
- graph and explorer results can lag behind a fast-forwarded `main`; verify any
  suggested cleanup seam against live `git log`, `rg`, and direct file reads
  before writing a new queue item
- when reusing a shared test helper as the proposed replacement for duplicated
  suite-local helpers, verify the shared helper actually covers the same
  stderr, exit-code, capture, and retry semantics before narrowing the queue
  note around it
