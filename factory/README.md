# Checked-In Review-Loop Starter

This repository ships `./factory/` as a richer customer-facing starter example.
It is not the default `agent-factory` or `agent-factory init` scaffold, which
creates the single-step `tasks` workflow described in the package README.

This checked-in starter uses a neutral story workflow with one execution step,
one review step, and a guarded loop breaker for repeated review rejections.

## Input Layout

Submit Markdown story files under:

`inputs/story/default/`

General layout:

- `inputs/<work-type>/default/` for manual submissions
- `inputs/<work-type>/<execution-id>/` for executor-generated work

The file watcher monitors this directory tree and automatically watches new
subdirectories.
