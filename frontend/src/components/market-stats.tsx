"use client";

import { useState, useEffect } from "react";
import { Card, CardContent } from "@/components/ui/card";
import { useLatestMarket } from "@/hooks/use-latest-market";
import { useOracle } from "@/hooks/use-oracle";
import { useSettlementCountdown } from "@/hooks/use-settlement-countdown";
import { formatUnits } from "viem";
import type { Address } from "viem";

type MarketStatsProps = {
    factoryAddress: Address;
};

export function MarketStats({ factoryAddress }: MarketStatsProps) {
    const { market } = useLatestMarket(factoryAddress);
    const { oracle } = useOracle(factoryAddress);
    const countdown = useSettlementCountdown();
    const [now, setNow] = useState(() => Math.floor(Date.now() / 1000));

    useEffect(() => {
        const interval = setInterval(() => setNow(Math.floor(Date.now() / 1000)), 1000);
        return () => clearInterval(interval);
    }, []);

    const tvl = market ? Number(formatUnits(market.totalLiquidity, 6)) : 0;
    const oracleIndex = oracle?.latestRound
        ? Number(formatUnits(oracle.latestRound.index, oracle.decimals))
        : null;

    return (
        <Card className="ring-0 border-none bg-muted/40">
            <CardContent>
                <dl className="space-y-2 text-sm">
                    <div className="flex justify-between">
                        <dt className="text-muted-foreground">
                            Index
                        </dt>
                        <dd>
                            {oracleIndex !== null
                                ? oracleIndex.toFixed(2)
                                : "—"}
                        </dd>
                    </div>
                    <div className="flex justify-between">
                        <dt className="text-muted-foreground">
                            Index Updated
                        </dt>
                        <dd>
                            {oracle?.latestRound?.timestamp
                                ? (() => {
                                      const secs = now - Number(oracle.latestRound.timestamp);
                                      if (secs < 60) return `${secs}s ago`;
                                      if (secs < 3600) return `${Math.floor(secs / 60)}m ago`;
                                      if (secs < 86400) return `${Math.floor(secs / 3600)}h ago`;
                                      return `${Math.floor(secs / 86400)}d ago`;
                                  })()
                                : "—"}
                        </dd>
                    </div>
                    <div className="flex justify-between">
                        <dt className="text-muted-foreground">
                            TVL
                        </dt>
                        <dd>${tvl.toLocaleString()}</dd>
                    </div>
                    <div className="flex justify-between">
                        <dt className="text-muted-foreground">
                            Settlement
                        </dt>
                        <dd>{countdown}</dd>
                    </div>
                </dl>
            </CardContent>
        </Card>
    );
}
