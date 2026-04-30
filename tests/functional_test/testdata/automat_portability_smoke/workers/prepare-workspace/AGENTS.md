---
type: SCRIPT_WORKER
command: powershell
args:
  - -File
  - scripts/prepare-automat-slice.ps1
  - -DependencyContract
  - portable-dependencies.json
  - -WorkflowGuide
  - docs/portable-workflow.md
timeout: 10m
---

Stage the bounded automat portability slice and surface the dispatch-ready
layout inputs.
