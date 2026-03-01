import type { Address } from "viem";

export type MarketFactoryView = {
    name: string;
    symbol: string;
    oracle: Address;
    collateralToken: Address;
    decimals: number;
    state: 0 | 1;
    totalMarkets: bigint;
    latestMarket: Address;
};
