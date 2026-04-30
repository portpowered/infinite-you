---
type: SCRIPT_WORKER
command: powershell
args:
  - -File
  - scripts/verify-external-tools.ps1
  - -DependencyContract
  - portable-dependencies.json
timeout: 5m
---

Read the dependency contract and surface the required external tools without
bundling them.
