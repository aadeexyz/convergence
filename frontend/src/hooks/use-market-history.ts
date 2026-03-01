import { useReadContract } from "wagmi";
import { lensAbi } from "@/lib/abis";
import { lensAddress } from "@/lib/metadata";
import type { Address } from "viem";
import type { PriceSnapshot } from "@/types";

export function useMarkPriceHistory(marketFactoryAddress: Address, count: number = 10) {
    const { data, isLoading, isError, error } = useReadContract({
        address: lensAddress,
        abi: lensAbi,
        functionName: "getMarkPriceHistory",
        args: [marketFactoryAddress, BigInt(count)],
    });

    if (!data)
        return {
            snapshots: null,
            isLoading,
            isError,
            error,
        };

    const snapshots = data as unknown as PriceSnapshot[];

    return {
        snapshots,
        isLoading,
        isError,
        error,
    };
}
