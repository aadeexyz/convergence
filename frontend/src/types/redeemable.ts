import type { Address } from "viem";

export type RedeemablePosition = {
    market: Address;
    marketIndex: bigint;
    isLong: boolean;
    positionTokenBalance: bigint;
    redeemableCollateral: bigint;
};
