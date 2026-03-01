import { Address } from "viem";

export type MarketView = {
    market: Address;
    longPositionToken: Address;
    shortPositionToken: Address;
    longPrice: bigint;
    shortPrice: bigint;
    totalLiquidity: bigint;
    decimals: number;
    settled: boolean;
    settlementRoundId: bigint;
};

export type PriceSnapshot = {
    timestamp: bigint;
    longPrice: bigint;
};
