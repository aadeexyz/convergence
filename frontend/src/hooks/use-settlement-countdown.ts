import { useEffect, useState } from "react";

const CRON_INTERVAL_SEC = 30 * 60; // 30 minutes

/** Formats remaining seconds into HH:MM:SS */
function formatCountdown(secs: number): string {
    if (secs <= 0) return "any moment";
    const h = Math.floor(secs / 3600);
    const m = Math.floor((secs % 3600) / 60);
    const s = Math.floor(secs % 60);
    return `${h.toString().padStart(2, "0")}:${m.toString().padStart(2, "0")}:${s.toString().padStart(2, "0")}`;
}

/**
 * Returns a live countdown to the next settlement for this factory.
 * Next settlement ≈ lastSettledTimestamp + (totalFactories * 30 min).
 */
export function useSettlementCountdown(lastSettledTimestamp?: bigint, totalFactories?: number) {
    const computeLabel = () => {
        if (!lastSettledTimestamp || lastSettledTimestamp === BigInt(0)) return "—";
        const n = totalFactories && totalFactories > 0 ? totalFactories : 1;
        const nextSettlement = Number(lastSettledTimestamp) + n * CRON_INTERVAL_SEC;
        const remaining = nextSettlement - Math.floor(Date.now() / 1000);
        return formatCountdown(remaining);
    };

    const [label, setLabel] = useState(computeLabel);

    useEffect(() => {
        if (!lastSettledTimestamp || lastSettledTimestamp === BigInt(0)) return;

        const interval = setInterval(() => setLabel(computeLabel()), 1000);
        return () => clearInterval(interval);
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [lastSettledTimestamp, totalFactories]);

    return label;
}
