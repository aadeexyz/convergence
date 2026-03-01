import { useReadContract } from "wagmi";
import { lensAbi } from "@/lib/abis";
import { lensAddress } from "@/lib/metadata";
import type { Address } from "viem";
import type { MarketFactoryView } from "@/types";

export function useMarketFactory(marketFactoryAddress: Address) {
    const { data, isLoading, isError, error } = useReadContract({
        address: lensAddress,
        abi: lensAbi,
        functionName: "getFactory",
        args: [marketFactoryAddress],
    });

    if (!data)
        return {
            factory: null,
            isLoading,
            isError,
            error,
        };

    const factory = data ? (data as unknown as MarketFactoryView) : null;

    return {
        factory,
        isLoading,
        isError,
        error,
    };
}
