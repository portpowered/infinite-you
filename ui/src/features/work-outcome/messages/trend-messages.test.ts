import { SUPPORTED_LOCALES } from "../../../i18n";
import {
  getWorkOutcomeTrendMessages,
  workOutcomeTrendMessagesByLocale,
} from "./trend-messages";

describe("getWorkOutcomeTrendMessages", () => {
  it("supports the expected work-outcome locales", () => {
    expect(Object.keys(workOutcomeTrendMessagesByLocale).sort()).toEqual(
      [...SUPPORTED_LOCALES].sort(),
    );
  });

  it.each([
    ["en", "Failure trend", "Queued"],
    ["zh", "失败趋势", "排队中"],
    ["ko", "실패 추이", "대기 중"],
    ["ja", "失敗推移", "待機中"],
  ] as const)("resolves %s trend-card and throughput labels", (locale, expectedTitle, expectedQueuedLabel) => {
    const messages = getWorkOutcomeTrendMessages(locale);

    expect(messages.failureCard.title).toBe(expectedTitle);
    expect(messages.throughput.seriesLabels.queued).toBe(expectedQueuedLabel);
  });

  it("falls back to the default locale when locale is missing or unsupported", () => {
    const defaultMessages = getWorkOutcomeTrendMessages("en");

    expect(getWorkOutcomeTrendMessages(undefined).failureCard.title).toBe(
      defaultMessages.failureCard.title,
    );
    expect(
      getWorkOutcomeTrendMessages("fr").throughput.rangeLabels.session,
    ).toBe(defaultMessages.throughput.rangeLabels.session);
  });

  it.each([
    "ja",
    "ko",
    "zh",
  ] as const)("keeps interpolation helpers available for %s", (locale) => {
    const messages = getWorkOutcomeTrendMessages(locale);

    expect(messages.failureCard.chartAriaLabel("15m")).toContain("15");
    expect(messages.reworkCard.pointTitle("Review", 3)).toContain("Review");
    expect(messages.timingCard.pointTitle("Dispatch A", "450ms")).toContain(
      "450ms",
    );
    expect(messages.workTypeFailureGroupLabel("story")).toContain("story");
  });
});
