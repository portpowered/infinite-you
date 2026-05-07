import {
  DETAIL_CARD_WIDE_CLASS,
  DETAIL_COPY_CLASS,
  EMPTY_STATE_CLASS,
  EMPTY_STATE_COMPACT_CLASS,
  WIDGET_SUBTITLE_CLASS,
} from "../../components/dashboard/widget-board";
import { DashboardWidgetFrame } from "../../components/ui";
import {
  DASHBOARD_BODY_TEXT_CLASS,
  DASHBOARD_SUPPORTING_LABEL_CLASS,
  DASHBOARD_SUPPORTING_LABELS_CLASS,
} from "../../components/ui/dashboard-typography";
import {
  formatDurationMillis,
  formatTraceOutcome,
} from "../../components/ui/formatters";
import { cx } from "../../lib/cx";
import {
  DASHBOARD_CHART_AXIS_CLASS,
  DASHBOARD_CHART_SURFACE_CLASS,
  getDashboardChartSemanticStyle,
} from "./chart-contract";
import { getWorkOutcomeTrendMessages } from "./messages/trend-messages";
import {
  type FailureTrendModel,
  getThroughputRangeOptions,
  isThroughputRangeID,
  type ReworkTrendModel,
  type ThroughputRangeID,
  type TimingTrendModel,
} from "./trends";

interface FailureTrendCardProps {
  className?: string;
  locale?: string;
  model: FailureTrendModel;
  onRangeChange: (rangeID: ThroughputRangeID) => void;
  rangeID: ThroughputRangeID;
  widgetId?: string;
}

interface ReworkTrendCardProps {
  className?: string;
  locale?: string;
  model: ReworkTrendModel;
  widgetId?: string;
}

interface TimingTrendCardProps {
  className?: string;
  locale?: string;
  model: TimingTrendModel;
  widgetId?: string;
}

const TREND_TOOLBAR_CLASS =
  "mb-4 flex items-start justify-between gap-3 max-[720px]:flex-col";
const TREND_RANGE_LABEL_CLASS =
  "grid shrink-0 basis-[8.5rem] gap-1 max-[720px]:w-full max-[720px]:basis-auto";
const TREND_RANGE_TEXT_CLASS = DASHBOARD_SUPPORTING_LABEL_CLASS;
const TREND_RANGE_SELECT_CLASS = cx(
  "rounded-lg border border-af-accent/35 bg-af-canvas/82 px-[0.55rem] py-[0.45rem] text-af-ink",
  DASHBOARD_BODY_TEXT_CLASS,
);
const TREND_SUMMARY_CLASS = cx(
  "mb-4 grid grid-cols-3 gap-3 max-[720px]:grid-cols-1 [&_dd]:m-0 [&_div]:rounded-lg [&_div]:border [&_div]:border-af-overlay/8 [&_div]:bg-af-overlay/4 [&_div]:p-[0.7rem] [&_dt]:mb-1",
  DASHBOARD_SUPPORTING_LABELS_CLASS,
);
const TREND_CHART_CLASS = cx(
  DASHBOARD_CHART_SURFACE_CLASS,
  "min-h-44 border border-af-overlay/8",
);
const TREND_CAUSE_LIST_CLASS = "mt-4 grid list-none gap-[0.55rem] p-0";
const TREND_CAUSE_ITEM_CLASS =
  "flex items-center justify-between gap-3 rounded-lg border border-af-overlay/7 bg-af-overlay/4 px-[0.7rem] py-[0.6rem]";
const TREND_CAUSE_LABEL_CLASS = cx(
  "min-w-0 text-af-ink/78 [overflow-wrap:anywhere]",
  DASHBOARD_BODY_TEXT_CLASS,
);
const TIMING_RANGE_SUMMARY_CLASS = cx(
  TREND_SUMMARY_CLASS,
  "mt-[0.85rem] grid-cols-2",
);
const TREND_SUMMARY_TERM_CLASS = DASHBOARD_SUPPORTING_LABEL_CLASS;
const TREND_SUMMARY_VALUE_CLASS = WIDGET_SUBTITLE_CLASS;
const FAILURE_TREND_CHART_STYLE =
  getDashboardChartSemanticStyle("failureTrend");
const REWORK_TREND_CHART_STYLE = getDashboardChartSemanticStyle("reworkTrend");
const TIMING_TREND_CHART_STYLE = getDashboardChartSemanticStyle("timingTrend");

export function FailureTrendCard({
  className = "",
  locale,
  model,
  onRangeChange,
  rangeID,
  widgetId = "failure-trend",
}: FailureTrendCardProps) {
  const messages = getWorkOutcomeTrendMessages(locale);
  const throughputRangeOptions = getThroughputRangeOptions(locale);
  const changeRange = (value: string) => {
    if (isThroughputRangeID(value)) {
      onRangeChange(value);
    }
  };

  return (
    <DashboardWidgetFrame
      className={cx(DETAIL_CARD_WIDE_CLASS, className)}
      title={messages.failureCard.title}
      widgetId={widgetId}
    >
      <div className={TREND_TOOLBAR_CLASS}>
        <p className={WIDGET_SUBTITLE_CLASS}>{messages.failureCard.subtitle}</p>
        <label className={TREND_RANGE_LABEL_CLASS}>
          <span className={TREND_RANGE_TEXT_CLASS}>
            {messages.failureCard.timeRangeLabel}
          </span>
          <select
            aria-label={messages.failureCard.timeRangeLabel}
            className={TREND_RANGE_SELECT_CLASS}
            value={rangeID}
            onChange={(event) => changeRange(event.target.value)}
          >
            {throughputRangeOptions.map((option) => (
              <option key={option.id} value={option.id}>
                {option.label}
              </option>
            ))}
          </select>
        </label>
      </div>

      <dl className={TREND_SUMMARY_CLASS}>
        <div>
          <dt className={TREND_SUMMARY_TERM_CLASS}>
            {messages.failureCard.failedInRangeSummaryLabel}
          </dt>
          <dd className={TREND_SUMMARY_VALUE_CLASS}>{model.failureDelta}</dd>
        </div>
        <div>
          <dt className={TREND_SUMMARY_TERM_CLASS}>
            {messages.failureCard.totalFailedSummaryLabel}
          </dt>
          <dd className={TREND_SUMMARY_VALUE_CLASS}>{model.currentFailed}</dd>
        </div>
        <div>
          <dt className={TREND_SUMMARY_TERM_CLASS}>
            {messages.failureCard.causeGroupsSummaryLabel}
          </dt>
          <dd className={TREND_SUMMARY_VALUE_CLASS}>{model.groups.length}</dd>
        </div>
      </dl>

      {model.points.length > 0 ? (
        <svg
          className={TREND_CHART_CLASS}
          role="img"
          aria-label={messages.failureCard.chartAriaLabel(model.rangeLabel)}
          viewBox="0 0 320 120"
        >
          <TrendAxes />
          {model.path ? (
            <path
              className={FAILURE_TREND_CHART_STYLE.lineClassName}
              d={model.path}
              stroke={FAILURE_TREND_CHART_STYLE.color}
            />
          ) : null}
          {model.points.map((point) => (
            <circle
              key={`${point.label}-${point.x}-${point.y}`}
              className={FAILURE_TREND_CHART_STYLE.pointClassName}
              cx={point.x}
              cy={point.y}
              r={FAILURE_TREND_CHART_STYLE.pointRadius}
            >
              <title>{point.label}</title>
            </circle>
          ))}
        </svg>
      ) : (
        <div className={cx(EMPTY_STATE_CLASS, EMPTY_STATE_COMPACT_CLASS)}>
          <h3>{messages.failureCard.emptyTitle}</h3>
          <p>{messages.failureCard.emptyBody}</p>
        </div>
      )}

      {model.groups.length > 0 ? (
        <ul
          className={TREND_CAUSE_LIST_CLASS}
          aria-label={messages.failureCard.causeGroupsAriaLabel}
        >
          {model.groups.map((group) => (
            <li className={TREND_CAUSE_ITEM_CLASS} key={group.label}>
              <span className={TREND_CAUSE_LABEL_CLASS}>{group.label}</span>
              <strong className="shrink-0 text-af-danger-bright">
                {group.count}
              </strong>
            </li>
          ))}
        </ul>
      ) : (
        <p className={DETAIL_COPY_CLASS}>{messages.failureCard.emptyGroups}</p>
      )}
    </DashboardWidgetFrame>
  );
}

export function ReworkTrendCard({
  className = "",
  locale,
  model,
  widgetId = "rework-trend",
}: ReworkTrendCardProps) {
  const messages = getWorkOutcomeTrendMessages(locale);

  return (
    <DashboardWidgetFrame
      className={cx(DETAIL_CARD_WIDE_CLASS, className)}
      title={messages.reworkCard.title}
      widgetId={widgetId}
    >
      <p className={WIDGET_SUBTITLE_CLASS}>{messages.reworkCard.subtitle}</p>

      <dl className={TREND_SUMMARY_CLASS}>
        <div>
          <dt className={TREND_SUMMARY_TERM_CLASS}>
            {messages.reworkCard.traceWorkSummaryLabel}
          </dt>
          <dd className={TREND_SUMMARY_VALUE_CLASS}>
            {model.currentWorkLabel}
          </dd>
        </div>
        <div>
          <dt className={TREND_SUMMARY_TERM_CLASS}>
            {messages.reworkCard.retryOrReworkSummaryLabel}
          </dt>
          <dd className={TREND_SUMMARY_VALUE_CLASS}>
            {model.retryOrReworkCount}
          </dd>
        </div>
        <div>
          <dt className={TREND_SUMMARY_TERM_CLASS}>
            {messages.reworkCard.latestOutcomeSummaryLabel}
          </dt>
          <dd className={TREND_SUMMARY_VALUE_CLASS}>
            {formatTraceOutcome(model.terminalOutcome)}
          </dd>
        </div>
      </dl>

      {model.points.length > 0 ? (
        <svg
          className={TREND_CHART_CLASS}
          role="img"
          aria-label={messages.reworkCard.chartAriaLabel(
            model.currentWorkLabel,
          )}
          viewBox="0 0 320 120"
        >
          <TrendAxes />
          {model.path ? (
            <path
              className={REWORK_TREND_CHART_STYLE.lineClassName}
              d={model.path}
              stroke={REWORK_TREND_CHART_STYLE.color}
            />
          ) : null}
          {model.points.map((point) => (
            <circle
              key={`${point.dispatchLabel}-${point.x}-${point.y}`}
              className={REWORK_TREND_CHART_STYLE.pointClassName}
              cx={point.x}
              cy={point.y}
              r={REWORK_TREND_CHART_STYLE.pointRadius}
            >
              <title>
                {messages.reworkCard.pointTitle(
                  point.dispatchLabel,
                  point.reworkCount,
                )}
              </title>
            </circle>
          ))}
        </svg>
      ) : (
        <div className={cx(EMPTY_STATE_CLASS, EMPTY_STATE_COMPACT_CLASS)}>
          <h3>{messages.reworkCard.emptyTitle}</h3>
          <p>{messages.reworkCard.emptyBody}</p>
        </div>
      )}
    </DashboardWidgetFrame>
  );
}

export function TimingTrendCard({
  className = "",
  locale,
  model,
  widgetId = "timing-trend",
}: TimingTrendCardProps) {
  const messages = getWorkOutcomeTrendMessages(locale);

  return (
    <DashboardWidgetFrame
      className={cx(DETAIL_CARD_WIDE_CLASS, className)}
      title={messages.timingCard.title}
      widgetId={widgetId}
    >
      <p className={WIDGET_SUBTITLE_CLASS}>{messages.timingCard.subtitle}</p>

      <dl className={TREND_SUMMARY_CLASS}>
        <div>
          <dt className={TREND_SUMMARY_TERM_CLASS}>
            {messages.timingCard.traceWorkSummaryLabel}
          </dt>
          <dd className={TREND_SUMMARY_VALUE_CLASS}>
            {model.currentWorkLabel}
          </dd>
        </div>
        <div>
          <dt className={TREND_SUMMARY_TERM_CLASS}>
            {messages.timingCard.slowestDispatchSummaryLabel}
          </dt>
          <dd className={TREND_SUMMARY_VALUE_CLASS}>
            {formatDurationMillis(model.slowestDurationMillis)}
          </dd>
        </div>
        <div>
          <dt className={TREND_SUMMARY_TERM_CLASS}>
            {messages.timingCard.averageDurationSummaryLabel}
          </dt>
          <dd className={TREND_SUMMARY_VALUE_CLASS}>
            {formatDurationMillis(model.averageDurationMillis)}
          </dd>
        </div>
      </dl>

      {model.points.length > 0 ? (
        <>
          <svg
            className={TREND_CHART_CLASS}
            role="img"
            aria-label={messages.timingCard.chartAriaLabel(
              model.currentWorkLabel,
            )}
            viewBox="0 0 320 120"
          >
            <TrendAxes />
            {model.path ? (
              <path
                className={TIMING_TREND_CHART_STYLE.lineClassName}
                d={model.path}
                stroke={TIMING_TREND_CHART_STYLE.color}
              />
            ) : null}
            {model.points.map((point) => (
              <circle
                key={`${point.dispatchLabel}-${point.x}-${point.y}`}
                className={TIMING_TREND_CHART_STYLE.pointClassName}
                cx={point.x}
                cy={point.y}
                r={TIMING_TREND_CHART_STYLE.pointRadius}
              >
                <title>
                  {messages.timingCard.pointTitle(
                    point.dispatchLabel,
                    formatDurationMillis(point.durationMillis),
                  )}
                </title>
              </circle>
            ))}
          </svg>
          <dl
            className={TIMING_RANGE_SUMMARY_CLASS}
            aria-label={messages.timingCard.timingRangeAriaLabel}
          >
            <div>
              <dt className={TREND_SUMMARY_TERM_CLASS}>
                {messages.timingCard.fastestSummaryLabel}
              </dt>
              <dd className={TREND_SUMMARY_VALUE_CLASS}>
                {formatDurationMillis(model.fastestDurationMillis)}
              </dd>
            </div>
            <div>
              <dt className={TREND_SUMMARY_TERM_CLASS}>
                {messages.timingCard.latestSummaryLabel}
              </dt>
              <dd className={TREND_SUMMARY_VALUE_CLASS}>
                {formatDurationMillis(model.latestDurationMillis)}
              </dd>
            </div>
          </dl>
        </>
      ) : (
        <div className={cx(EMPTY_STATE_CLASS, EMPTY_STATE_COMPACT_CLASS)}>
          <h3>{messages.timingCard.emptyTitle}</h3>
          <p>{messages.timingCard.emptyBody}</p>
        </div>
      )}
    </DashboardWidgetFrame>
  );
}

function TrendAxes() {
  return (
    <>
      <line
        className={DASHBOARD_CHART_AXIS_CLASS}
        x1="14"
        x2="306"
        y1="106"
        y2="106"
      />
      <line
        className={DASHBOARD_CHART_AXIS_CLASS}
        x1="14"
        x2="14"
        y1="14"
        y2="106"
      />
    </>
  );
}
