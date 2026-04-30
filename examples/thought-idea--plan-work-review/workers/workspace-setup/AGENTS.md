---
args:
    - factory/scripts/setup-workspace.py
    - '{{ (index .Inputs 0).Name }}'
command: python
type: SCRIPT_WORKER
---
