"use client";

import { Line, LineChart, CartesianGrid, XAxis, YAxis } from "recharts";
import {
    ChartContainer,
    ChartTooltip,
    ChartTooltipContent,
    type ChartConfig,
} from "@/components/ui/chart";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { useOracleHistory } from "@/hooks/use-oracle-history";
import { useMarkPriceHistory } from "@/hooks/use-market-history";
import { formatUnits } from "viem";
import type { Address } from "viem";

const chartConfig = {
    mark: {
        label: "Mark Price",
        color: "var(--chart-1)",
    },
    index: {
        label: "Underlying Index",
        color: "var(--chart-2)",
    },
} satisfies ChartConfig;

type DataPoint = {
    time: number;
    label: string;
    index?: number;
    mark?: number;
};

type Props = {
    factoryAddress: Address;
    oracleDecimals: number;
    collateralDecimals?: number;
};

export function OracleHistoryChart({ factoryAddress, oracleDecimals, collateralDecimals = 6 }: Props) {
    const { rounds } = useOracleHistory(factoryAddress);
    const { snapshots } = useMarkPriceHistory(factoryAddress, 10);

    if (!snapshots || snapshots.length === 0) {
        // Fallback: show index-only if no mark data yet
        if (!rounds || rounds.length === 0) return null;

        const indexOnly = rounds.map((r) => ({
            time: Number(r.timestamp),
            label: new Date(Number(r.timestamp) * 1000).toLocaleString(),
            index: Number(formatUnits(r.index, oracleDecimals)),
            mark: undefined as number | undefined,
        }));

        return renderChart(indexOnly);
    }

    // Mark price snapshots define the time window
    const minTime = Number(snapshots[0].timestamp);
    const maxTime = Number(snapshots[snapshots.length - 1].timestamp);

    // Parse rounds sorted by timestamp for index interpolation
    const parsedRounds = (rounds ?? [])
        .map((r) => ({
            ts: Number(r.timestamp),
            index: Number(formatUnits(r.index, oracleDecimals)),
        }))
        .sort((a, b) => a.ts - b.ts);

    // Find the index value at a given timestamp (most recent round <= ts)
    function indexAt(ts: number): number | undefined {
        let val: number | undefined;
        for (const r of parsedRounds) {
            if (r.ts <= ts) val = r.index;
            else break;
        }
        return val;
    }

    const pointMap = new Map<number, DataPoint>();

    // Add mark price snapshots as primary data
    for (const snap of snapshots) {
        const ts = Number(snap.timestamp);
        pointMap.set(ts, {
            time: ts,
            label: new Date(ts * 1000).toLocaleString(),
            mark: Number(formatUnits(snap.longPrice, collateralDecimals)),
            index: indexAt(ts),
        });
    }

    // Add index rounds that fall within the mark time window
    for (const r of parsedRounds) {
        if (r.ts >= minTime && r.ts <= maxTime && !pointMap.has(r.ts)) {
            pointMap.set(r.ts, {
                time: r.ts,
                label: new Date(r.ts * 1000).toLocaleString(),
                index: r.index,
            });
        }
    }

    // Sort by time
    const chartData = Array.from(pointMap.values()).sort((a, b) => a.time - b.time);

    return renderChart(chartData);
}

function renderChart(data: DataPoint[]) {
    return (
        <Card className="ring-0 border-none bg-muted/40">
            <CardHeader className="flex flex-row items-center justify-between">
                <CardTitle className="text-sm text-muted-foreground">
                    Price Chart
                </CardTitle>
                <div className="flex items-center gap-4 text-xs">
                    <span className="flex items-center gap-1.5">
                        <span className="inline-block h-2 w-2 rounded-full bg-chart-1" />
                        Mark Price
                    </span>
                    <span className="flex items-center gap-1.5">
                        <span className="inline-block h-2 w-2 rounded-full bg-chart-2" />
                        Underlying Index
                    </span>
                </div>
            </CardHeader>
            <CardContent>
                <ChartContainer config={chartConfig} className="min-h-62.5 w-full">
                    <LineChart
                        accessibilityLayer
                        data={data}
                        margin={{ left: 12, right: 12, top: 4, bottom: 4 }}
                    >
                        <CartesianGrid vertical={false} />
                        <XAxis
                            dataKey="time"
                            tickLine={false}
                            axisLine={false}
                            tickMargin={8}
                            tickFormatter={(value) =>
                                new Date(value * 1000).toLocaleTimeString([], {
                                    hour: "2-digit",
                                    minute: "2-digit",
                                })
                            }
                        />
                        <YAxis
                            tickLine={false}
                            axisLine={false}
                            tickMargin={8}
                            domain={[0, 1]}
                            tickFormatter={(value) => value.toFixed(2)}
                        />
                        <ChartTooltip
                            content={
                                <ChartTooltipContent
                                    labelFormatter={(_, payload) => {
                                        const p = payload?.[0]?.payload;
                                        return p?.label ?? "";
                                    }}
                                />
                            }
                        />
                        <Line
                            type="monotone"
                            dataKey="mark"
                            name="Mark Price"
                            stroke="var(--color-mark)"
                            strokeWidth={2}
                            dot={{ r: 3 }}
                            activeDot={{ r: 5 }}
                            connectNulls
                        />
                        <Line
                            type="monotone"
                            dataKey="index"
                            name="Underlying Index"
                            stroke="var(--color-index)"
                            strokeWidth={1.5}
                            strokeDasharray="4 4"
                            dot={false}
                            activeDot={{ r: 4 }}
                        />
                    </LineChart>
                </ChartContainer>
            </CardContent>
        </Card>
    );
}
