"use client";

import { MarketFactoryCard } from "@/components/market-factory-card";
import { useMarketFactories } from "@/hooks/use-market-factories";

export function MarketFactories() {
    const { factories, isLoading, isError } = useMarketFactories();

    if (isLoading) return <p className="text-muted-foreground">Loading markets...</p>;
    if (isError) return <p className="text-muted-foreground">Failed to load markets.</p>;

    return (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {factories.map((factory) => (
                <MarketFactoryCard
                    key={factory}
                    marketFactoryAddress={factory}
                />
            ))}
        </div>
    );
}
