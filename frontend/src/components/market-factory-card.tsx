import Link from "next/link";
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/ui/card";
import { Progress } from "@/components/ui/progress";
import { useMarketFactory } from "@/hooks/use-market-factory";
import { useLatestMarket } from "@/hooks/use-latest-market";
import { useSettlementCountdown } from "@/hooks/use-settlement-countdown";
import { cn } from "@/lib/utils";
import { formatUnits } from "viem";
import type { Address } from "viem";

type MarketFactoryCardProps = {
    marketFactoryAddress: Address;
};

export function MarketFactoryCard({
    marketFactoryAddress,
}: MarketFactoryCardProps) {
    const {
        factory,
        isLoading: isLoadingFactory,
        isError: isErrorFactory,
        error: errorFactory,
    } = useMarketFactory(marketFactoryAddress);

    const {
        market,
        isLoading: isLoadingMarket,
        isError: isErrorMarket,
        error: errorMarket,
    } = useLatestMarket(marketFactoryAddress);

    const countdown = useSettlementCountdown();

    return (
        <Link href={`/markets/${marketFactoryAddress}`}>
            <Card
                className={cn(
                    "ring-0 border border-foreground/10",
                    factory?.state === 0
                        ? "border-dashed"
                        : "border-solid transition-colors hover:border-foreground cursor-pointer",
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
                            <Progress value={market ? Number(formatUnits(market.longPrice, 6)) * 100 : 0} max={100} />
                            <div className="flex justify-between text-sm text-muted-foreground">
                                <span>Long: {market ? Number(formatUnits(market.longPrice, 6)).toFixed(2) : "—"}</span>
                                <span>Short: {market ? Number(formatUnits(market.shortPrice, 6)).toFixed(2) : "—"}</span>
                            </div>
                        </div>
                    )}
                </CardContent>
                {factory?.state === 1 && (
                    <CardFooter className="flex justify-between text-sm text-muted-foreground">
                        <span>${market ? Number(formatUnits(market.totalLiquidity, 6)).toLocaleString() : "—"}</span>
                        <span>Settles in {countdown}</span>
                    </CardFooter>
                )}
            </Card>
        </Link>
    );
}
