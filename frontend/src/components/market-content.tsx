"use client";

import { useMarketFactory } from "@/hooks/use-market-factory";
import { useOracle } from "@/hooks/use-oracle";
import { OracleHistoryChart } from "@/components/oracle-history-chart";
import { SettledMarkets } from "@/components/settled-markets";
import { TradePanel } from "@/components/trade-panel";
import { MarketStats } from "@/components/market-stats";
import type { Address } from "viem";

type MarketContentProps = {
    factoryAddress: Address;
};

export function MarketContent({ factoryAddress }: MarketContentProps) {
    const { factory } = useMarketFactory(factoryAddress);
    const { oracle } = useOracle(factoryAddress);

    if (!factory) return null;

    if (factory.state === 0) {
        return (
            <p className="animate-pulse text-muted-foreground">
                Market is being created...
            </p>
        );
    }

    return (
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-4">
            {/* Chart + settled markets — left 2/3 */}
            <div className="lg:col-span-2 space-y-4">
                <OracleHistoryChart
                    factoryAddress={factoryAddress}
                    oracleDecimals={oracle?.decimals ?? 8}
                />
                <SettledMarkets factoryAddress={factoryAddress} />
            </div>

            {/* Right column — trade + stats stacked */}
            <div className="flex flex-col gap-4">
                <TradePanel factoryAddress={factoryAddress} />
                <MarketStats factoryAddress={factoryAddress} />
            </div>
        </div>
    );
}
