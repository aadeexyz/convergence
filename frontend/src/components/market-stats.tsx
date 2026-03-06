"use client";

import { Card, CardContent } from "@/components/ui/card";
import { useLatestMarket } from "@/hooks/use-latest-market";
import { useMarketFactory } from "@/hooks/use-market-factory";
import { useMarketFactories } from "@/hooks/use-market-factories";
import { useOracle } from "@/hooks/use-oracle";
import { useSettlementCountdown } from "@/hooks/use-settlement-countdown";
import { formatUnits } from "viem";
import type { Address } from "viem";

type MarketStatsProps = {
    factoryAddress: Address;
};

export function MarketStats({ factoryAddress }: MarketStatsProps) {
    const { market } = useLatestMarket(factoryAddress);
    const { factory } = useMarketFactory(factoryAddress);
    const { oracle } = useOracle(factoryAddress);
    const { factories } = useMarketFactories();
    const countdown = useSettlementCountdown(oracle?.latestRound?.timestamp, factories.length);

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
                            Next Settlement
                        </dt>
                        <dd>{countdown}</dd>
                    </div>
                    <div className="flex justify-between">
                        <dt className="text-muted-foreground">
                            Market
                        </dt>
                        <dd>
                            #{factory ? Number(factory.totalMarkets) : "—"}
                            {market?.settled && (
                                <span className="ml-1.5 text-xs text-amber-500">settled</span>
                            )}
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
                            Oracle Rounds
                        </dt>
                        <dd>{oracle ? Number(oracle.currentRoundId) : "—"}</dd>
                    </div>
                </dl>
            </CardContent>
        </Card>
    );
}
