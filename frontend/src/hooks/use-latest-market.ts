import { useReadContract } from "wagmi";
import { lensAbi } from "@/lib/abis";
import { lensAddress } from "@/lib/metadata";
import type { Address } from "viem";
import type { MarketView } from "@/types";

export function useLatestMarket(marketFactoryAddress: Address) {
    const { data, isLoading, isError, error } = useReadContract({
        address: lensAddress,
        abi: lensAbi,
        functionName: "getLatestMarket",
        args: [marketFactoryAddress],
    });

    if (!data)
        return {
            market: null,
            isLoading,
            isError,
            error,
        };

    const market = data ? (data as unknown as MarketView) : null;

    return {
        market,
        isLoading,
        isError,
        error,
    };
}
