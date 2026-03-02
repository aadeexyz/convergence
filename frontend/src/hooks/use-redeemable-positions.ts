import { useReadContract, useAccount } from "wagmi";
import { lensAbi } from "@/lib/abis";
import { lensAddress } from "@/lib/metadata";
import type { Address } from "viem";
import type { RedeemablePosition } from "@/types";

export function useRedeemablePositions(marketFactoryAddress: Address) {
    const { address: userAddress } = useAccount();

    const { data, isLoading, isError, error } = useReadContract({
        address: lensAddress,
        abi: lensAbi,
        functionName: "getRedeemable",
        args: userAddress ? [marketFactoryAddress, userAddress] : undefined,
        query: { enabled: !!userAddress },
    });

    const positions = data
        ? (data as unknown as RedeemablePosition[])
        : [];

    return {
        positions,
        isLoading,
        isError,
        error,
    };
}
