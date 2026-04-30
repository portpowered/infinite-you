# Portable Workflow Slice

This bounded smoke represents the pre-dispatch setup that a portable
`translate/automat`-style factory needs before real chapter processing can run.

The fixture expects these bundled files to survive flatten and expand:

- `scripts/prepare-automat-slice.ps1`
- `scripts/verify-external-tools.ps1`
- `docs/portable-workflow.md`
- `portable-dependencies.json`

After `config expand`, those files should be readable from the restored runtime
layout at the same factory-relative paths shown above.

The fixture expects these tools to remain external:

- `mangaka.exe`
- `magick`

Those tools are intentionally declared in `portable-dependencies.json` instead
of being checked into the portable bundle.

This smoke ends once the expanded layout is proven dispatch-ready for that
bounded setup path. It does not validate full chapter translation, OCR, or
image-processing outputs.
