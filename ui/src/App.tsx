import "./styles.css";

import { DashboardScreen } from "./features/dashboard";
import { resolveSupportedLocale } from "./i18n";

function resolveAppLocale(): string | undefined {
  if (typeof navigator === "undefined") {
    return undefined;
  }

  return resolveSupportedLocale(navigator.language);
}

export function App() {
  return <DashboardScreen locale={resolveAppLocale()} />;
}
