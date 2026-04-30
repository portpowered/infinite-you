# Repeated Identical Input Binding Keys

## Why this matters

Workstations can declare multiple inputs with the same `workType` and `state`.
Several runtime paths still infer binding identity from default arc names or
place IDs, which are not unique for repeated identical inputs.

The new `MATCHES_FIELDS` guard can be wired safely for mixed input types today,
but repeated identical inputs remain a structural edge case for:

- guard binding names
- dispatch `InputBindings`
- workstation input ordering derived from place IDs

## Observed risk

- Two authored inputs like `{ "workType": "asset", "state": "ready" }` produce
  the same default arc name.
- Downstream code that groups or reorders consumed tokens by place can no longer
  distinguish the first authored input from the second.
- Future grouped-input guards and worker logic may silently compare or dispatch
  the wrong tokens when identical input slots are used.

## Suggested follow-up

1. Introduce stable authored input-slot identities for transition input arcs.
2. Carry those identities through enablement bindings and dispatch payloads.
3. Update workstation input ordering helpers to consume each authored slot once
   instead of grouping only by place.
4. Add focused tests for repeated identical inputs at mapper, scheduler, and
   worker-dispatch boundaries.
