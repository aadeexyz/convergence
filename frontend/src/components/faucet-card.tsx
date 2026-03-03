"use client";

import { useEffect } from "react";
import { useWriteContract, useTransactionReceipt, useAccount } from "wagmi";
import { useQueryClient } from "@tanstack/react-query";
import { toast } from "sonner";
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Spinner } from "@/components/ui/spinner";
import { SonnerTxLink } from "@/components/sonner-tx-link";
import { mockUSDCAbi } from "@/lib/abis";
import { usdcAddress } from "@/lib/metadata";
import { useMintAllowed } from "@/hooks/use-mint-allowed";

export function FaucetCard() {
    const { address, isConnected } = useAccount();
    const queryClient = useQueryClient();
    const { mintAllowed, timeLeft, isLoading: isMintAllowedLoading } = useMintAllowed(usdcAddress, address);

    const mint = useWriteContract({
        mutation: {
            onError() {
                toast.error("Failed to submit transaction");
            },
            onSuccess(data) {
                toast.loading("Minting USDC…", {
                    description: <SonnerTxLink txHash={data} />,
                    id: data,
                });
            },
        },
    });
    const mintReceipt = useTransactionReceipt({ hash: mint.data });

    useEffect(() => {
        if (mintReceipt.status === "success") {
            queryClient.invalidateQueries();
            toast.success("USDC minted!", {
                description: <SonnerTxLink txHash={mint.data} />,
                id: mint.data,
            });
        } else if (mintReceipt.status === "error") {
            toast.error("Mint transaction failed", {
                description: <SonnerTxLink txHash={mint.data} />,
                id: mint.data,
            });
        }
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [mintReceipt.status]);

    function handleMint() {
        mint.mutate({
            address: usdcAddress,
            abi: mockUSDCAbi,
            functionName: "mint",
        });
    }

    const isPending =
        mint.isPending || (mintReceipt.status === "pending" && !!mint.data);

    return (
        <Card className="mx-auto max-w-sm ring-0 border-none bg-muted/40">
            <CardHeader>
                <CardTitle>USDC Faucet</CardTitle>
            </CardHeader>
            <CardContent>
                <p className="text-sm text-muted-foreground">
                    100 USDC · once every 24h
                </p>
            </CardContent>
            <CardFooter className="bg-transparent border-t-0">
                <Button
                    className="w-full"
                    disabled={!isConnected || isPending || isMintAllowedLoading || !mintAllowed}
                    onClick={handleMint}
                >
                    {isPending || isMintAllowedLoading ? (
                        <Spinner />
                    ) : !mintAllowed && timeLeft ? (
                        `${timeLeft.hours}h ${timeLeft.minutes}m ${timeLeft.seconds}s`
                    ) : (
                        "Mint USDC"
                    )}
                </Button>
            </CardFooter>
        </Card>
    );
}
