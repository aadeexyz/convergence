import Link from "next/link";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Progress } from "@/components/ui/progress";
import { useMarketFactory } from "@/hooks/use-market-factory";
import { useLatestMarket } from "@/hooks/use-latest-market";
import { useOracle } from "@/hooks/use-oracle";
import { useSettlementCountdown } from "@/hooks/use-settlement-countdown";
import { cn } from "@/lib/utils";
import { formatUnits } from "viem";
import type { Address } from "viem";

type MarketFactoryCardProps = {
    marketFactoryAddress: Address;
    totalFactories: number;
};

export function MarketFactoryCard({
    marketFactoryAddress,
    totalFactories,
}: MarketFactoryCardProps) {
    const { factory } = useMarketFactory(marketFactoryAddress);
    const { market } = useLatestMarket(marketFactoryAddress);
    const { oracle } = useOracle(marketFactoryAddress);
    const countdown = useSettlementCountdown(oracle?.latestRound?.timestamp, totalFactories);

    const longPrice = market ? Number(formatUnits(market.longPrice, 6)) : 0;
    const longPct = Math.round(longPrice * 100);

    return (
        <Link href={`/markets/${marketFactoryAddress}`}>
            <Card
                className={cn(
                    "ring-0 border-none bg-muted/40",
                    factory?.state === 0
                        ? "opacity-60"
                        : "transition-colors hover:bg-muted/70 cursor-pointer",
                )}
            >
                <CardHeader>
                    <CardTitle className="flex justify-between">
                        <p>{factory?.name}</p>
                        <p>${factory?.symbol}</p>
                    </CardTitle>
                </CardHeader>
                <CardContent>
                    {factory?.state == 0 ? (
                        <p className="animate-pulse">Should be live any moment</p>
                    ) : (
                        <div className="space-y-3">
                            <Progress value={longPct} max={100} />
                            <div className="flex justify-between text-sm text-muted-foreground">
                                <span>Long: {market ? Number(formatUnits(market.longPrice, 6)).toFixed(2) : "—"}</span>
                                <span>Short: {market ? Number(formatUnits(market.shortPrice, 6)).toFixed(2) : "—"}</span>
                            </div>
                        </div>
                    )}
                    {factory?.state === 1 && (
                        <div className="flex justify-between text-sm text-muted-foreground pt-2">
                            <span>Vol: ${market ? Number(formatUnits(market.totalLiquidity, 6)).toLocaleString() : "—"}</span>
                            <span>Settles in {countdown}</span>
                        </div>
                    )}
                </CardContent>
            </Card>
        </Link>
    );
}
