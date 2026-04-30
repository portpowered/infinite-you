# Example: Adhoc Planning Workflow

This adhoc fixture models a generic thought-to-plan-to-work workflow for manual record/replay checks.

## Workflow

```
idea:init → [plan] → idea:complete + plan:init
plan:init → [setup-workspace] → task:init
task:init → [process] → task:complete
                 ↑
                 └── onRejection
```

1. Ideas become plans.
2. Plans set up task work.
3. Tasks repeat through processing until complete or failed.

## Directory Structure

```
tests/adhoc/factory/
├── factory.json                    # Workflow: idea → plan → task
├── scripts/
│   └── setup-workspace.py          # Workspace setup helper
├── workers/
│   ├── README.md
│   ├── processor/AGENTS.md         # MODEL_WORKER: handles planning and task work
│   └── workspace-setup/AGENTS.md   # MODEL_WORKER: prepares workspaces
├── workstations/
│   ├── README.md
│   ├── ideafy/AGENTS.md            # Prompt template for idea generation
│   ├── plan/AGENTS.md              # Prompt template for plan generation
│   ├── process/AGENTS.md           # Prompt template for implementation
│   └── review/AGENTS.md            # Prompt template for review
└── README.md                       # This file
```

## Running

```bash
agent-factory run -dir tests/adhoc/factory
```

## Record/Replay Smoke

The adhoc test package includes an opt-in record/replay smoke that copies this
fixture, records a deterministic task run, then replays the generated artifact
from embedded config:

```bash
AGENT_FACTORY_ADHOC_RECORD_REPLAY=1 go test -v ./tests/adhoc -run TestAdHocRecordReplaySmoke -count=1
```

Set `AGENT_FACTORY_ADHOC_ARTIFACT=/path/to/artifact.json` to keep the generated
artifact. The test output prints the artifact path and replay result. The smoke
uses a mock provider so the common path does not require live model credentials
or source edits.

For command-surface checks, run:

```bash
AGENT_FACTORY_ADHOC_RUN=1 go test -v ./tests/adhoc -run TestAdHocPrepare -count=1
```

Optional environment variables:

- `AGENT_FACTORY_ADHOC_DIR` points the command or smoke flow at another factory directory.
- `AGENT_FACTORY_ADHOC_RECORD` adds `--record <path>` to `TestAdHocPrepare`.
- `AGENT_FACTORY_ADHOC_REPLAY` adds `--replay <path>` to `TestAdHocPrepare`.

## Retained Replay Artifact

`../factory-recording-04-11-02.json` is a historical replay artifact that still contains captured local paths. It remains explicit replay input data and is not a factory default.
