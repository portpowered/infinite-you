# Backend Functional Coverage Evidence Lane

## Why this should exist

The canonical ask sets a stronger backend testing target than the current audit
can prove: functional tests should cover at least `90%` of non-generated
`pkg/` code.

The repository already has real backend test layers and a repo-owned coverage
entrypoint, but the current checked-in evidence still leaves repeated review
work:

- future audits will keep rediscovering backend test breadth without one stable
  proof surface for the declared target
- local command runs alone do not create a durable artifact that later
  reviewers can inspect without repeating the same coverage derivation

## Desired outcome

Create one focused lane that makes the backend coverage target reviewable:

- define the canonical repo-owned proof surface for the `90%` functional
  coverage target
- document or automate how non-generated `pkg/` scope is measured
- keep the result inspectable through maintained docs, command output, or CI

## Observable evidence that would close it

- one repo-owned command or maintained report proves the current backend
  functional coverage result against the declared target
- the scope for non-generated `pkg/` code is explicit instead of inferred
- future reviewers can verify the target without re-deriving the full coverage
  story from scratch
