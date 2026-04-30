---
author: ralph agent
last modified: 2026, april, 21
doc-id: AGF-DOC-005
---

# Migrate Event Type Names

After reading this guide, you will be able to update an older Agent Factory
event consumer from the retired public event names to the current canonical
request/response vocabulary.

## When To Use This Guide

Use this guide if your code filters, switches on, stores, or replays Agent
Factory event `type` values from `/events` or replay artifacts.

## What Changed

Agent Factory now uses one public event vocabulary that distinguishes
request-style events from response-style events by name alone. Events ending in
`_REQUEST` describe accepted work or scheduled lifecycle steps. Events ending
in `_RESPONSE` describe the resulting outcome or state snapshot.

`WORK_REQUEST`, `INFERENCE_REQUEST`, and `INFERENCE_RESPONSE` were already
canonical and did not change.

## Old-To-New Mapping

| Retired name | Canonical name | Lifecycle style | Meaning |
| --- | --- | --- | --- |
| `RUN_STARTED` | `RUN_REQUEST` | request-style | The run started and the effective runtime configuration is available. |
| `INITIAL_STRUCTURE` | `INITIAL_STRUCTURE_REQUEST` | request-style | The runtime published the initial topology snapshot before work moved. |
| `RELATIONSHIP_CHANGE` | `RELATIONSHIP_CHANGE_REQUEST` | request-style | The runtime recorded batch relations such as `PARENT_CHILD` or `DEPENDS_ON`. |
| `DISPATCH_CREATED` | `DISPATCH_REQUEST` | request-style | A workstation accepted input and created a dispatch attempt. |
| `DISPATCH_COMPLETED` | `DISPATCH_RESPONSE` | response-style | A workstation finished with `ACCEPTED`, `REJECTED`, or `FAILED`. |
| `FACTORY_STATE_CHANGE` | `FACTORY_STATE_RESPONSE` | response-style | The runtime reported a factory lifecycle state such as `RUNNING` or `COMPLETED`. |
| `RUN_FINISHED` | `RUN_RESPONSE` | response-style | The run ended and final metadata is available. |

## Unchanged Names

| Event name | Why it stayed the same |
| --- | --- |
| `WORK_REQUEST` | It already describes accepted work entering the factory. |
| `INFERENCE_REQUEST` | It already describes a provider attempt request. |
| `INFERENCE_RESPONSE` | It already describes the matching provider attempt outcome. |

## Update Consumer Switch Logic

Replace retired names in filters and switches with the canonical names:

```ts
// Before
switch (event.type) {
  case "RUN_STARTED":
  case "INITIAL_STRUCTURE":
  case "DISPATCH_CREATED":
    handleLifecycleRequest(event)
    break
  case "DISPATCH_COMPLETED":
  case "FACTORY_STATE_CHANGE":
  case "RUN_FINISHED":
    handleLifecycleResponse(event)
    break
}
```

```ts
// After
switch (event.type) {
  case "RUN_REQUEST":
  case "INITIAL_STRUCTURE_REQUEST":
  case "DISPATCH_REQUEST":
    handleLifecycleRequest(event)
    break
  case "DISPATCH_RESPONSE":
  case "FACTORY_STATE_RESPONSE":
  case "RUN_RESPONSE":
    handleLifecycleResponse(event)
    break
}
```

Apply the same mapping to stored artifacts, dashboard reducers, replay tooling,
and any alerting or metrics pipeline that keys on event names.

## Compatibility And Transition Behavior

The advertised public contract now emits and documents only the canonical event
names. Current `/events` streams, generated clients, examples, and checked-in
fixtures use the canonical vocabulary only.

Current binaries do not advertise transition-only alias reading for retired
public names in `/events` or replay artifacts. Internal branch fixtures were
migrated in the repository. If you still have a historical artifact with
retired names, either:

1. Re-record it with a current binary.
2. Rewrite the retired `type` values to the canonical names before replaying
   or testing it.

## Next Steps

- [Understand a run timeline](run-timeline.md)
- [Record and replay a run](record-replay.md)
