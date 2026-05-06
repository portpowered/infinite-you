import { cx } from "../../components/ui";

interface DashboardBrandLockupProps {
  className?: string;
  wordmarkClassName?: string;
}

const BRAND_MARK_CLASS =
  "inline-flex h-8 items-center justify-center rounded-full border border-af-accent/28 bg-af-accent/12 px-3 text-[1.05rem] font-black leading-none text-af-accent";

export function DashboardBrandLockup({
  className = "",
  wordmarkClassName = "",
}: DashboardBrandLockupProps) {
  return (
    <span className={cx("inline-flex min-w-0 items-center gap-3", className)}>
      <span aria-hidden="true" className={BRAND_MARK_CLASS}>
        ∞
      </span>
      <span className={cx("sr-only", wordmarkClassName)}>
        Infinite You
      </span>
    </span>
  );
}
