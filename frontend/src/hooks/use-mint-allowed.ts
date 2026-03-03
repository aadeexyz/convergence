"use client";

import { useReadContract } from "wagmi";
import { mockUSDCAbi } from "@/lib/abis";
import { useNow } from "@/hooks/use-now";
import type { Address } from "viem";

const ONE_DAY = 86400;

function getTimeLeft(currentTimestamp: number, nextTimestamp: number) {
    if (currentTimestamp < nextTimestamp) {
        const diff = nextTimestamp - currentTimestamp;
        const hours = Math.floor(diff / 3600);
        const minutes = Math.floor((diff % 3600) / 60);
        const seconds = Math.floor(diff % 60);

        return {
            hours: hours.toString().padStart(2, "0"),
            minutes: minutes.toString().padStart(2, "0"),
            seconds: seconds.toString().padStart(2, "0"),
        };
    }

    return undefined;
}

export function useMintAllowed(tokenAddress: Address, userAddress?: Address) {
    const enabled = !!userAddress;

    const { data, isLoading, isError, error } = useReadContract({
        address: tokenAddress,
        abi: mockUSDCAbi,
        functionName: "lastMintedTimestamp",
        args: userAddress ? [userAddress] : undefined,
        query: { enabled },
    });

    const now = useNow(enabled);

    if (!enabled) {
        return {
            mintAllowed: false,
            timeLeft: undefined,
            isLoading: false,
            isError: false,
            error: undefined,
        };
    }

    if (data === undefined) {
        return {
            mintAllowed: false,
            timeLeft: undefined,
            isLoading,
            isError,
            error,
        };
    }

    const lastMinted = Number(data);
    const nextMintAllowedAt = lastMinted === 0 ? 0 : lastMinted + ONE_DAY;
    const currentTimestamp = Math.floor(now / 1000);

    return {
        mintAllowed: currentTimestamp >= nextMintAllowedAt,
        timeLeft: getTimeLeft(currentTimestamp, nextMintAllowedAt),
        isLoading,
        isError,
        error,
    };
}
