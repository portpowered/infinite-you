export const SUPPORTED_LOCALES = ["en", "zh", "ko", "ja"] as const;

export type SupportedLocale = (typeof SUPPORTED_LOCALES)[number];

export const DEFAULT_LOCALE: SupportedLocale = "en";

function normalizeLocaleCandidate(
  locale: string | undefined | null,
): string | undefined {
  if (!locale) {
    return undefined;
  }

  return locale.trim().toLowerCase().split(/[-_]/, 1)[0];
}

export function isSupportedLocale(locale: string): locale is SupportedLocale {
  return SUPPORTED_LOCALES.includes(locale as SupportedLocale);
}

export function resolveSupportedLocale(
  locale: string | undefined | null,
): SupportedLocale {
  const normalizedLocale = normalizeLocaleCandidate(locale);

  if (!normalizedLocale) {
    return DEFAULT_LOCALE;
  }

  return isSupportedLocale(normalizedLocale)
    ? normalizedLocale
    : DEFAULT_LOCALE;
}
