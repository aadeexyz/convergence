import { useReadContract } from "wagmi";
import { lensAbi } from "@/lib/abis";
import { lensAddress } from "@/lib/metadata";
import type { Address } from "viem";
import type { OracleView } from "@/types";

export function useOracle(marketFactoryAddress: Address) {
    const { data, isLoading, isError, error } = useReadContract({
        address: lensAddress,
        abi: lensAbi,
        functionName: "getOracle",
        args: [marketFactoryAddress],
    });

    if (!data)
        return {
            oracle: null,
            isLoading,
            isError,
            error,
        };

    const oracle = data as unknown as OracleView;

    return {
        oracle,
        isLoading,
        isError,
        error,
    };
}
