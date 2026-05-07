import {
  type LocalizedMessages,
  resolveLocalizedMessages,
} from "../../../i18n";

export interface WorkOutcomeTrendMessages {
  failureCard: {
    causeGroupsAriaLabel: string;
    causeGroupsSummaryLabel: string;
    chartAriaLabel: (rangeLabel: string) => string;
    emptyGroups: string;
    emptyTitle: string;
    emptyBody: string;
    failedInRangeSummaryLabel: string;
    subtitle: string;
    timeRangeLabel: string;
    title: string;
    totalFailedSummaryLabel: string;
  };
  reworkCard: {
    chartAriaLabel: (workLabel: string) => string;
    emptyBody: string;
    emptyTitle: string;
    latestOutcomeSummaryLabel: string;
    pointTitle: (dispatchLabel: string, count: number) => string;
    retryOrReworkSummaryLabel: string;
    subtitle: string;
    title: string;
    traceWorkSummaryLabel: string;
  };
  throughput: {
    rangeLabels: {
      "5m": string;
      "15m": string;
      session: string;
    };
    seriesLabels: {
      completed: string;
      failed: string;
      inFlight: string;
      queued: string;
    };
  };
  timingCard: {
    averageDurationSummaryLabel: string;
    chartAriaLabel: (workLabel: string) => string;
    emptyBody: string;
    emptyTitle: string;
    fastestSummaryLabel: string;
    latestSummaryLabel: string;
    pointTitle: (dispatchLabel: string, durationLabel: string) => string;
    slowestDispatchSummaryLabel: string;
    subtitle: string;
    timingRangeAriaLabel: string;
    title: string;
    traceWorkSummaryLabel: string;
  };
  workTypeFailureGroupLabel: (workType: string) => string;
}

const workOutcomeTrendMessagesByLocale = {
  en: {
    failureCard: {
      causeGroupsAriaLabel: "Failure cause groups",
      causeGroupsSummaryLabel: "Cause groups",
      chartAriaLabel: (rangeLabel) => `Failed work trend for ${rangeLabel}`,
      emptyGroups: "No failed work has been grouped yet.",
      emptyTitle: "No failure samples",
      emptyBody:
        "Failure trend data appears after the event stream receives work history.",
      failedInRangeSummaryLabel: "Failed in range",
      subtitle:
        "Failed work and cause groups from the selected factory timeline.",
      timeRangeLabel: "Time range",
      title: "Failure trend",
      totalFailedSummaryLabel: "Total failed",
    },
    reworkCard: {
      chartAriaLabel: (workLabel) => `Retry and rework trend for ${workLabel}`,
      emptyBody:
        "Select active work with retained trace history to see retry activity.",
      emptyTitle: "No selected trace",
      latestOutcomeSummaryLabel: "Latest outcome",
      pointTitle: (dispatchLabel, count) =>
        `${dispatchLabel}: ${count} retry or rework events`,
      retryOrReworkSummaryLabel: "Retry or rework",
      subtitle:
        "Reject, retry, or rework activity from the selected work trace.",
      title: "Retry and rework trend",
      traceWorkSummaryLabel: "Trace work",
    },
    throughput: {
      rangeLabels: {
        "5m": "5m",
        "15m": "15m",
        session: "Session",
      },
      seriesLabels: {
        completed: "Completed",
        failed: "Failed/retried",
        inFlight: "In-flight",
        queued: "Queued",
      },
    },
    timingCard: {
      averageDurationSummaryLabel: "Average duration",
      chartAriaLabel: (workLabel) => `Timing trend for ${workLabel}`,
      emptyBody:
        "Select active work with retained trace history to compare dispatch timing.",
      emptyTitle: "No selected trace",
      fastestSummaryLabel: "Fastest",
      latestSummaryLabel: "Latest",
      pointTitle: (dispatchLabel, durationLabel) =>
        `${dispatchLabel}: ${durationLabel}`,
      slowestDispatchSummaryLabel: "Slowest dispatch",
      subtitle: "Dispatch duration trend from the selected work trace.",
      timingRangeAriaLabel: "Timing range",
      title: "Timing trend",
      traceWorkSummaryLabel: "Trace work",
    },
    workTypeFailureGroupLabel: (workType) => `Work type: ${workType}`,
  },
  ja: {
    failureCard: {
      causeGroupsAriaLabel: "失敗原因グループ",
      causeGroupsSummaryLabel: "原因グループ",
      chartAriaLabel: (rangeLabel) => `${rangeLabel} の失敗ワーク推移`,
      emptyGroups: "失敗したワークはまだグループ化されていません。",
      emptyTitle: "失敗サンプルはありません",
      emptyBody:
        "イベントストリームがワーク履歴を受信すると、失敗推移データが表示されます。",
      failedInRangeSummaryLabel: "範囲内の失敗数",
      subtitle:
        "選択中のファクトリー時系列における失敗ワークと原因グループです。",
      timeRangeLabel: "時間範囲",
      title: "失敗推移",
      totalFailedSummaryLabel: "累計失敗数",
    },
    reworkCard: {
      chartAriaLabel: (workLabel) => `${workLabel} の再試行と手戻りの推移`,
      emptyBody:
        "再試行アクティビティを確認するには、保持されたトレース履歴を持つアクティブなワークを選択してください。",
      emptyTitle: "選択中のトレースはありません",
      latestOutcomeSummaryLabel: "最新の結果",
      pointTitle: (dispatchLabel, count) =>
        `${dispatchLabel}: 再試行または手戻りイベント ${count} 件`,
      retryOrReworkSummaryLabel: "再試行または手戻り",
      subtitle:
        "選択したワークトレースにおける reject、retry、rework のアクティビティです。",
      title: "再試行と手戻りの推移",
      traceWorkSummaryLabel: "トレース対象ワーク",
    },
    throughput: {
      rangeLabels: {
        "5m": "5分",
        "15m": "15分",
        session: "セッション",
      },
      seriesLabels: {
        completed: "完了",
        failed: "失敗/再試行",
        inFlight: "進行中",
        queued: "待機中",
      },
    },
    timingCard: {
      averageDurationSummaryLabel: "平均時間",
      chartAriaLabel: (workLabel) => `${workLabel} の時間推移`,
      emptyBody:
        "ディスパッチ時間を比較するには、保持されたトレース履歴を持つアクティブなワークを選択してください。",
      emptyTitle: "選択中のトレースはありません",
      fastestSummaryLabel: "最速",
      latestSummaryLabel: "最新",
      pointTitle: (dispatchLabel, durationLabel) =>
        `${dispatchLabel}: ${durationLabel}`,
      slowestDispatchSummaryLabel: "最も遅いディスパッチ",
      subtitle: "選択したワークトレースにおけるディスパッチ時間の推移です。",
      timingRangeAriaLabel: "時間範囲",
      title: "時間推移",
      traceWorkSummaryLabel: "トレース対象ワーク",
    },
    workTypeFailureGroupLabel: (workType) => `ワークタイプ: ${workType}`,
  },
  ko: {
    failureCard: {
      causeGroupsAriaLabel: "실패 원인 그룹",
      causeGroupsSummaryLabel: "원인 그룹",
      chartAriaLabel: (rangeLabel) => `${rangeLabel} 실패 작업 추이`,
      emptyGroups: "실패한 작업이 아직 그룹화되지 않았습니다.",
      emptyTitle: "실패 샘플이 없습니다",
      emptyBody:
        "이벤트 스트림이 작업 기록을 수신하면 실패 추이 데이터가 표시됩니다.",
      failedInRangeSummaryLabel: "범위 내 실패",
      subtitle: "선택한 팩토리 타임라인의 실패 작업과 원인 그룹입니다.",
      timeRangeLabel: "시간 범위",
      title: "실패 추이",
      totalFailedSummaryLabel: "총 실패",
    },
    reworkCard: {
      chartAriaLabel: (workLabel) => `${workLabel} 재시도 및 재작업 추이`,
      emptyBody:
        "재시도 활동을 보려면 보존된 추적 기록이 있는 활성 작업을 선택하세요.",
      emptyTitle: "선택된 추적이 없습니다",
      latestOutcomeSummaryLabel: "최신 결과",
      pointTitle: (dispatchLabel, count) =>
        `${dispatchLabel}: 재시도 또는 재작업 이벤트 ${count}건`,
      retryOrReworkSummaryLabel: "재시도 또는 재작업",
      subtitle: "선택한 작업 추적의 reject, retry, rework 활동입니다.",
      title: "재시도 및 재작업 추이",
      traceWorkSummaryLabel: "추적 작업",
    },
    throughput: {
      rangeLabels: {
        "5m": "5분",
        "15m": "15분",
        session: "세션",
      },
      seriesLabels: {
        completed: "완료됨",
        failed: "실패/재시도",
        inFlight: "진행 중",
        queued: "대기 중",
      },
    },
    timingCard: {
      averageDurationSummaryLabel: "평균 시간",
      chartAriaLabel: (workLabel) => `${workLabel} 시간 추이`,
      emptyBody:
        "디스패치 시간을 비교하려면 보존된 추적 기록이 있는 활성 작업을 선택하세요.",
      emptyTitle: "선택된 추적이 없습니다",
      fastestSummaryLabel: "최단",
      latestSummaryLabel: "최신",
      pointTitle: (dispatchLabel, durationLabel) =>
        `${dispatchLabel}: ${durationLabel}`,
      slowestDispatchSummaryLabel: "가장 느린 디스패치",
      subtitle: "선택한 작업 추적의 디스패치 시간 추이입니다.",
      timingRangeAriaLabel: "시간 범위",
      title: "시간 추이",
      traceWorkSummaryLabel: "추적 작업",
    },
    workTypeFailureGroupLabel: (workType) => `작업 유형: ${workType}`,
  },
  zh: {
    failureCard: {
      causeGroupsAriaLabel: "失败原因分组",
      causeGroupsSummaryLabel: "原因分组",
      chartAriaLabel: (rangeLabel) => `${rangeLabel} 失败工作趋势`,
      emptyGroups: "失败工作尚未完成分组。",
      emptyTitle: "没有失败样本",
      emptyBody: "事件流接收到工作历史后，这里才会显示失败趋势数据。",
      failedInRangeSummaryLabel: "范围内失败数",
      subtitle: "来自所选工厂时间线的失败工作及其原因分组。",
      timeRangeLabel: "时间范围",
      title: "失败趋势",
      totalFailedSummaryLabel: "累计失败数",
    },
    reworkCard: {
      chartAriaLabel: (workLabel) => `${workLabel} 的重试与返工趋势`,
      emptyBody: "选择带有保留追踪历史的活动工作以查看重试活动。",
      emptyTitle: "未选择追踪",
      latestOutcomeSummaryLabel: "最新结果",
      pointTitle: (dispatchLabel, count) =>
        `${dispatchLabel}：${count} 次重试或返工事件`,
      retryOrReworkSummaryLabel: "重试或返工",
      subtitle: "来自所选工作追踪的 reject、retry 或 rework 活动。",
      title: "重试与返工趋势",
      traceWorkSummaryLabel: "追踪工作",
    },
    throughput: {
      rangeLabels: {
        "5m": "5 分钟",
        "15m": "15 分钟",
        session: "会话",
      },
      seriesLabels: {
        completed: "已完成",
        failed: "失败/重试",
        inFlight: "进行中",
        queued: "排队中",
      },
    },
    timingCard: {
      averageDurationSummaryLabel: "平均时长",
      chartAriaLabel: (workLabel) => `${workLabel} 的时长趋势`,
      emptyBody: "选择带有保留追踪历史的活动工作以比较分发时长。",
      emptyTitle: "未选择追踪",
      fastestSummaryLabel: "最快",
      latestSummaryLabel: "最新",
      pointTitle: (dispatchLabel, durationLabel) =>
        `${dispatchLabel}：${durationLabel}`,
      slowestDispatchSummaryLabel: "最慢分发",
      subtitle: "来自所选工作追踪的分发时长趋势。",
      timingRangeAriaLabel: "时长范围",
      title: "时长趋势",
      traceWorkSummaryLabel: "追踪工作",
    },
    workTypeFailureGroupLabel: (workType) => `工作类型：${workType}`,
  },
} satisfies LocalizedMessages<WorkOutcomeTrendMessages>;

export function getWorkOutcomeTrendMessages(
  locale?: string | null,
): WorkOutcomeTrendMessages {
  return resolveLocalizedMessages(workOutcomeTrendMessagesByLocale, locale);
}

export { workOutcomeTrendMessagesByLocale };
