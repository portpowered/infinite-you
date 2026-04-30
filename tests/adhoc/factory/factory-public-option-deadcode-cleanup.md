# Remove Dead Factory Public Option Hooks

Date: 2026-04-18
Source: cleaner sweep of `libraries/agent-factory`
Channel: local code and deadcode-baseline review
Decision: build

## Raw Signal

- `docs/development/deadcode-baseline.txt` still lists these Agent Factory production symbols:
  - `pkg/factory/options.go:183:6: unreachable func: WithChannelBuffer`
  - `pkg/factory/options.go:207:6: unreachable func: WithEventSink`
  - `pkg/factory/options.go:214:6: unreachable func: WithTracer`
  - `pkg/factory/options.go:259:6: unreachable func: WithWorkRequestRecorder`
  - `pkg/factory/engine/firing.go:16:6: unreachable func: applyCardinality`
- `rg --glob '!docs/development/deadcode-baseline.txt' "WithChannelBuffer|WithEventSink|WithTracer|WithWorkRequestRecorder"` shows the factory-level option declarations are not used inside the module. The engine-level `engine.WithWorkRequestRecorder` is the active recorder path.
- `pkg/factory/options.go` defines `factory.EventSink` and `factory.Tracer` beside `pkg/factory/subsystems/events.go`'s real event sink abstraction and the canonical generated event history path.
- `pkg/factory/runtime/factory.go` still carries config fields for `EventSink`, `Tracer`, `ChannelBuffer`, and `WorkRequestRecorder`, but normal runtime wiring records canonical factory events through `FactoryEventHistory` and engine recorder callbacks.

## Interpretation

The factory package has a leftover public extension surface from older runtime wiring. It duplicates lower-level abstractions, increases the apparent API area, and keeps known production deadcode in the accepted baseline. Because the project currently has no customers, we can break these unused public options instead of preserving compatibility shims.

## Proposed Cleanup

Remove the dead factory-level option hooks and the config fields they only exist to set:

- Delete `factory.WithChannelBuffer` and replace configurable channel buffer usage with a private runtime constant.
- Delete `factory.EventSink`, `factory.WithEventSink`, and `FactoryConfig.EventSink`; keep event delivery on canonical factory events and subsystem-local event emitters only where they are still directly tested or wired.
- Delete `factory.Tracer`, `factory.WithTracer`, and `FactoryConfig.Tracer`; no runtime code consumes the tracer field.
- Delete `factory.WithWorkRequestRecorder`, `FactoryConfig.WorkRequestRecorder`, and the extra runtime callback branch. Keep `engine.WithWorkRequestRecorder` as the internal bridge that records canonical work requests into `FactoryEventHistory`.
- Delete `pkg/factory/engine/firing.go`'s unused `applyCardinality` wrapper if no tests or callers need it after the option cleanup.

## Impact

- Shrinks the public factory API and the deadcode baseline.
- Reduces overlapping event/recorder abstractions in favor of the event history stream.
- Makes runtime configuration easier to review because supported extension points are the ones actually wired.

## Effort

Small. Expected scope is `pkg/factory/options.go`, `pkg/factory/runtime/factory.go`, `pkg/factory/engine/firing.go`, focused tests that compile against the deleted fields, and `docs/development/deadcode-baseline.txt`.

## Risk

Low to medium. The main risk is deleting a public library option that an untracked external caller used. The project guidance allows breaking unused legacy paths for now, and the module itself has no internal callers.

## Dependencies

- Keep `engine.WithWorkRequestRecorder` and `FactoryEventHistory.RecordWorkRequest` intact.
- Do not change `GET /events`, replay artifacts, or generated `FactoryEvent` payloads.
- If implementation is done as a cleanup sweep, add a scoped report under `docs/development/cleanup-analyzer-reports/`.

## Acceptance Criteria

- The listed factory-level options and dead wrapper no longer appear in `pkg/factory` production code.
- Runtime construction still creates result buffers and worker-pool dispatch hooks with a private default buffer size.
- Canonical work request events are still recorded through `FactoryEventHistory`.
- `go test ./pkg/factory ./pkg/factory/engine ./pkg/factory/runtime -count=1` passes.
- `make lint` passes, or the deadcode baseline is updated only after confirming the remaining findings are intentional.

## Time Box

One focused engineering session. If the cleanup requires changing generated API, replay artifact shape, or dashboard event semantics, stop and split those concerns into a separate idea.
