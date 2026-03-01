import { useReadContract } from "wagmi";
import { lensAbi } from "@/lib/abis";
import { lensAddress } from "@/lib/metadata";
import type { Address } from "viem";
import type { OracleRound } from "@/types";

export function useOracleHistory(marketFactoryAddress: Address) {
    const { data, isLoading, isError, error } = useReadContract({
        address: lensAddress,
        abi: lensAbi,
        functionName: "getAllRounds",
        args: [marketFactoryAddress],
    });

    if (!data)
        return {
            rounds: null,
            isLoading,
            isError,
            error,
        };

    const rounds = data as unknown as OracleRound[];

    return {
        rounds,
        isLoading,
        isError,
        error,
    };
}
