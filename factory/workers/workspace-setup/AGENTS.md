---
type: SCRIPT_WORKER
command: python
args:
  - "factory/scripts/setup-workspace.py"
  - "{{ (index .Inputs 0).Name }}"
---
