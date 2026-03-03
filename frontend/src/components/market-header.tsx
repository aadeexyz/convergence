"use client";

import { useMarketFactory } from "@/hooks/use-market-factory";
import type { Address } from "viem";

type MarketHeaderProps = {
    factoryAddress: Address;
};

export function MarketHeader({ factoryAddress }: MarketHeaderProps) {
    const { factory } = useMarketFactory(factoryAddress);

    if (!factory) return null;

    return (
        <div className="flex items-center justify-between">
            <div className="flex items-baseline gap-3">
                <h1 className="text-2xl font-bold">{factory.name}</h1>
                <span className="text-muted-foreground">
                    ${factory.symbol}
                </span>
            </div>
        </div>
    );
}
