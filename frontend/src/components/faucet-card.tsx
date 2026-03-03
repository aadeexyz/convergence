"use client";

import { useEffect } from "react";
import { useWriteContract, useTransactionReceipt, useAccount } from "wagmi";
import { useQueryClient } from "@tanstack/react-query";
import { toast } from "sonner";
import { parseUnits } from "viem";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Spinner } from "@/components/ui/spinner";
import { SonnerTxLink } from "@/components/sonner-tx-link";
import { usdcAddress } from "@/lib/metadata";

const MINT_AMOUNT = parseUnits("100", 6); // 100 USDC

const mockUsdcMintAbi = [
    {
        type: "function",
        name: "mint",
        inputs: [
            { name: "to_", type: "address", internalType: "address" },
            { name: "amount_", type: "uint256", internalType: "uint256" },
        ],
        outputs: [],
        stateMutability: "nonpayable",
    },
] as const;

export function FaucetCard() {
    const { address, isConnected } = useAccount();
    const queryClient = useQueryClient();

    const mint = useWriteContract({
        mutation: {
            onError() {
                toast.error("Failed to submit transaction");
            },
            onSuccess(data) {
                toast.loading("Minting 100 USDC…", {
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
            toast.success("100 USDC minted!", {
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
        if (!address) return;
        mint.mutate({
            address: usdcAddress,
            abi: mockUsdcMintAbi,
            functionName: "mint",
            args: [address, MINT_AMOUNT],
        });
    }

    const isPending =
        mint.isPending || (mintReceipt.status === "pending" && !!mint.data);

    return (
        <Card className="mx-auto max-w-sm ring-0 border border-foreground/10">
            <CardHeader>
                <CardTitle className="text-center">USDC Faucet</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
                <p className="text-sm text-center text-muted-foreground">
                    Mint 100 USDC to your wallet for testing.
                </p>
                <Button
                    className="w-full"
                    disabled={!isConnected || isPending}
                    onClick={handleMint}
                >
                    {isPending ? <Spinner /> : "Mint 100 USDC"}
                </Button>
            </CardContent>
        </Card>
    );
}
