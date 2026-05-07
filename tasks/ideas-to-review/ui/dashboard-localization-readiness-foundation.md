# Dashboard Localization Readiness Foundation

## Why this should exist

The dashboard currently renders user-visible copy inline across reusable
components and feature dialogs, and this audit did not find a dedicated
`ui/src/i18n/` setup or feature-local message catalogs.

This is an architecture-level deficiency rather than a narrow copy cleanup:

- future frontend work will keep deepening hardcoded-string ownership unless a
  localization boundary exists first
- the repository already formats some values with locale-aware browser APIs, so
  the missing piece is message ownership and fallback policy rather than a
  totally greenfield internationalization effort

## Desired outcome

Create one focused lane that introduces the minimum dashboard localization
foundation without rewriting the whole UI:

- add centralized `ui/src/i18n/` setup and fallback-locale policy
- move one bounded dashboard slice onto feature-owned message catalogs
- prove one non-default locale path for formatting-sensitive UI

## Observable evidence that would close it

- `ui/src/i18n/` exists with explicit locale registration and fallback rules
- at least one feature-owned message package is checked in near the owning UI
- tests prove localized formatting or copy lookup for a non-default locale
- the standards audit can cite concrete localization boundaries instead of
  inline copy examples
