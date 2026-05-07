import { DASHBOARD_WIDGET_CLASS } from "../../components/dashboard/widget-board";
import { AgentBentoCard } from "../../components/ui";
import { cx } from "../../lib/cx";
import { getDashboardWorkChartSeriesDefinitions } from "./chart-contract";
import type { WorkChartModel } from "./trends";
import type { WorkChartSeriesDefinition, WorkChartState } from "./work-chart";
import { WorkChart } from "./work-chart";

export interface WorkChartCardProps {
  chartState?: WorkChartState;
  className?: string;
  model: WorkChartModel;
  title?: string;
  widgetId?: string;
}
export type D3CompletionInformationCardProps = WorkChartCardProps;

const WORK_CHART_BODY_CLASS = "!flex !gap-0 !overflow-hidden !p-0";
const WORK_CHART_REGION_CLASS = "min-h-0 flex-1 px-4 sm:px-5";

export function WorkChartCard({
  chartState,
  className = "",
  model,
  title = "Work outcome chart",
  widgetId,
}: WorkChartCardProps) {
  const chartRegionID = widgetId
    ? `${widgetId}-chart-region`
    : "work-outcome-chart-region";
  const cardClassName = cx(DASHBOARD_WIDGET_CLASS, className);
  const chartSeries: readonly WorkChartSeriesDefinition[] =
    getDashboardWorkChartSeriesDefinitions(
      model.series.map((series) => ({
        key: series.key,
        label: series.label,
      })),
    );

  return (
    <AgentBentoCard
      bodyClassName={WORK_CHART_BODY_CLASS}
      className={cardClassName}
      title={title}
    >
      <section
        aria-label="Work outcome chart region"
        className={WORK_CHART_REGION_CLASS}
        id={chartRegionID}
      >
        <WorkChart
          ariaLabel={`Work outcome chart for ${model.rangeLabel}`}
          className="h-full"
          model={model}
          series={chartSeries}
          state={chartState}
        />
      </section>
    </AgentBentoCard>
  );
}

export const D3CompletionInformationCard = WorkChartCard;
