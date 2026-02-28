import { useReadContract } from "wagmi";
import { factoryOfMarketFactoriesAddress } from "@/lib/metadata";
import { factoryOfMarketFactoriesAbi } from "@/lib/abis";
import { Address } from "viem";

export function useMarketFactories() {
    const { data, isLoading, isError, error } = useReadContract({
        address: factoryOfMarketFactoriesAddress,
        abi: factoryOfMarketFactoriesAbi,
        functionName: "marketFactories",
    });

    if (!data)
        return {
            factories: [],
            isLoading,
            isError,
            error,
        };

    const factories: readonly Address[] = data;

    if (!factories)
        return {
            factories: [],
            isLoading,
            isError,
            error,
        };

    return {
        factories,
        isLoading,
        isError,
        error,
    };
}
