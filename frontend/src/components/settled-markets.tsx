"use client";

import { useEffect } from "react";
import { useWriteContract, useTransactionReceipt, useAccount } from "wagmi";
import { useQueryClient } from "@tanstack/react-query";
import { toast } from "sonner";
import { Card, CardContent } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { SonnerTxLink } from "@/components/sonner-tx-link";
import { useRedeemablePositions } from "@/hooks/use-redeemable-positions";
import { useMarketFactory } from "@/hooks/use-market-factory";
import { marketAbi } from "@/lib/abis";
import { cn } from "@/lib/utils";
import { formatUnits } from "viem";
import type { Address } from "viem";

type Props = {
    factoryAddress: Address;
};

export function SettledMarkets({ factoryAddress }: Props) {
    const { isConnected, address: userAddress } = useAccount();
    const queryClient = useQueryClient();
    const { factory } = useMarketFactory(factoryAddress);
    const { positions } = useRedeemablePositions(factoryAddress);

    const redeemPosition = useWriteContract({
        mutation: {
            onError() {
                toast.error("Redeem failed");
            },
            onSuccess(data) {
                toast.loading("Redeeming position…", {
                    description: <SonnerTxLink txHash={data} />,
                    id: data,
                });
            },
        },
    });
    const redeemReceipt = useTransactionReceipt({ hash: redeemPosition.data });

    useEffect(() => {
        if (redeemReceipt.status === "success") {
            queryClient.invalidateQueries();
            toast.success("Position redeemed!", {
                description: <SonnerTxLink txHash={redeemPosition.data} />,
                id: redeemPosition.data,
            });
        } else if (redeemReceipt.status === "error") {
            toast.error("Redeem transaction failed", {
                description: <SonnerTxLink txHash={redeemPosition.data} />,
                id: redeemPosition.data,
            });
        }
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [redeemReceipt.status]);

    const isRedeeming =
        redeemPosition.isPending ||
        (redeemReceipt.status === "pending" && !!redeemPosition.data);

    function handleRedeem(marketAddress: Address, isLong: boolean) {
        redeemPosition.mutate({
            address: marketAddress,
            abi: marketAbi,
            functionName: "redeem",
            args: [isLong, userAddress!],
        });
    }

    if (positions.length === 0) return null;

    return (
        <Card>
            <CardContent className="pt-5">
                <div className="mb-4">
                    <h2 className="text-lg font-semibold">Settled Markets</h2>
                </div>
                <div className="space-y-2">
                    {positions.map((pos) => {
                        const collateral = factory
                            ? Number(formatUnits(pos.redeemableCollateral, factory.decimals))
                            : 0;
                        const tokens = factory
                            ? Number(formatUnits(pos.positionTokenBalance, factory.decimals))
                            : 0;
                        return (
                            <div
                                key={`${pos.market}-${pos.isLong ? "long" : "short"}`}
                                className="flex items-center justify-between rounded-lg bg-muted/50 px-4 py-3"
                            >
                                <div className="flex items-center gap-3">
                                    <span
                                        className={cn(
                                            "text-xs font-semibold rounded px-2 py-0.5",
                                            pos.isLong
                                                ? "bg-emerald-600/20 text-emerald-500"
                                                : "bg-rose-600/20 text-rose-500",
                                        )}
                                    >
                                        {pos.isLong ? "LONG" : "SHORT"}
                                    </span>
                                    <div className="text-sm">
                                        <span className="text-muted-foreground">Market #{Number(pos.marketIndex) + 1}</span>
                                        <span className="mx-2 text-muted-foreground">·</span>
                                        <span>{tokens.toLocaleString(undefined, { maximumFractionDigits: 2 })} shares</span>
                                    </div>
                                </div>
                                <div className="flex items-center gap-3">
                                    <span className="text-sm font-medium">
                                        ${collateral.toLocaleString(undefined, { minimumFractionDigits: 2, maximumFractionDigits: 2 })}
                                    </span>
                                    <Button
                                        size="sm"
                                        variant="outline"
                                        disabled={!isConnected || isRedeeming}
                                        onClick={() => handleRedeem(pos.market, pos.isLong)}
                                    >
                                        Redeem
                                    </Button>
                                </div>
                            </div>
                        );
                    })}
                </div>
            </CardContent>
        </Card>
    );
}
