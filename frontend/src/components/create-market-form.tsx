"use client";

import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { useWriteContract, useReadContract, useTransactionReceipt, useAccount } from "wagmi";
import { useQueryClient } from "@tanstack/react-query";
import { toast } from "sonner";
import { formatUnits, maxUint256, parseUnits } from "viem";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Spinner } from "@/components/ui/spinner";
import { SonnerTxLink } from "@/components/sonner-tx-link";
import { mockUSDCAbi, factoryOfMarketFactoriesAbi } from "@/lib/abis";
import { MAX_ALLOWANCE } from "@/lib/constants";
import { factoryOfMarketFactoriesAddress, usdcAddress } from "@/lib/metadata";

const CREATION_COST = "11,000";
const LIQUIDITY_COST = "10,000";
const PROTOCOL_COST = "1,000";
const CREATION_COST_PARSED = parseUnits("11000", 6);

export function CreateMarketForm() {
    const router = useRouter();
    const { address: userAddress, isConnected } = useAccount();
    const queryClient = useQueryClient();

    const [name, setName] = useState("");
    const [symbol, setSymbol] = useState("");

    // --- Read USDC balance ---
    const { data: usdcBalance } = useReadContract({
        address: usdcAddress,
        abi: mockUSDCAbi,
        functionName: "balanceOf",
        args: userAddress ? [userAddress] : undefined,
        query: { enabled: !!userAddress },
    });

    // --- Read allowance ---
    const { data: allowance } = useReadContract({
        address: usdcAddress,
        abi: mockUSDCAbi,
        functionName: "allowance",
        args: userAddress ? [userAddress, factoryOfMarketFactoriesAddress] : undefined,
        query: { enabled: !!userAddress },
    });

    const needsApproval =
        allowance === undefined || (allowance as bigint) < MAX_ALLOWANCE;

    const formattedBalance =
        usdcBalance !== undefined
            ? Number(formatUnits(usdcBalance as bigint, 6)).toLocaleString(undefined, {
                  minimumFractionDigits: 2,
                  maximumFractionDigits: 2,
              })
            : "—";

    // --- Approve USDC ---
    const approveUsdc = useWriteContract({
        mutation: {
            onError() {
                toast.error("Approval failed");
            },
            onSuccess(data) {
                toast.loading("Approving USDC…", {
                    description: <SonnerTxLink txHash={data} />,
                    id: `approve-${data}`,
                });
            },
        },
    });
    const approveReceipt = useTransactionReceipt({ hash: approveUsdc.data });

    useEffect(() => {
        if (approveReceipt.status === "success") {
            queryClient.invalidateQueries();
            toast.success("USDC approved", {
                id: `approve-${approveUsdc.data}`,
            });
        } else if (approveReceipt.status === "error") {
            toast.error("Approval failed", {
                id: `approve-${approveUsdc.data}`,
            });
        }
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [approveReceipt.status]);

    // --- Create market factory ---
    const createFactory = useWriteContract({
        mutation: {
            onError() {
                toast.error("Failed to create market");
            },
            onSuccess(data) {
                toast.loading("Creating market…", {
                    description: <SonnerTxLink txHash={data} />,
                    id: data,
                });
            },
        },
    });
    const createReceipt = useTransactionReceipt({ hash: createFactory.data });

    useEffect(() => {
        if (createReceipt.status === "success") {
            queryClient.invalidateQueries();
            toast.success("Market created!", {
                description: <SonnerTxLink txHash={createFactory.data} />,
                id: createFactory.data,
            });
            router.push("/markets");
        } else if (createReceipt.status === "error") {
            toast.error("Market creation failed", {
                description: <SonnerTxLink txHash={createFactory.data} />,
                id: createFactory.data,
            });
        }
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [createReceipt.status]);

    // --- Handlers ---
    function handleApprove() {
        approveUsdc.mutate({
            address: usdcAddress,
            abi: mockUSDCAbi,
            functionName: "approve",
            args: [factoryOfMarketFactoriesAddress, maxUint256],
        });
    }

    function handleCreate() {
        if (!name || !symbol) return;
        createFactory.mutate({
            address: factoryOfMarketFactoriesAddress,
            abi: factoryOfMarketFactoriesAbi,
            functionName: "createMarketFactory",
            args: [name, symbol],
        });
    }

    const isApproving =
        approveUsdc.isPending ||
        (approveReceipt.status === "pending" && !!approveUsdc.data);

    const isCreating =
        createFactory.isPending ||
        (createReceipt.status === "pending" && !!createFactory.data);

    return (
        <Card className="mx-auto max-w-md ring-0 border border-foreground/10">
            <CardHeader>
                <CardTitle>Create Market</CardTitle>
            </CardHeader>
            <CardContent className="space-y-5">
                <div className="space-y-3">
                    <div className="space-y-1.5">
                        <label className="text-sm text-muted-foreground">
                            Name
                        </label>
                        <Input
                            placeholder="e.g. donald trump"
                            value={name}
                            onChange={(e) =>
                                setName(e.target.value.toLowerCase())
                            }
                        />
                        <p className="text-xs text-muted-foreground">
                            Lowercase only. This is the search keyword for the attention index.
                        </p>
                    </div>

                    <div className="space-y-1.5">
                        <label className="text-sm text-muted-foreground">
                            Ticker
                        </label>
                        <Input
                            placeholder="e.g. TRUMP"
                            value={symbol}
                            onChange={(e) =>
                                setSymbol(e.target.value.toUpperCase())
                            }
                        />
                        <p className="text-xs text-muted-foreground">
                            Uppercase only.
                        </p>
                    </div>
                </div>

                <div className="rounded-lg bg-muted p-3 space-y-1.5 text-sm">
                    <div className="flex justify-between">
                        <span className="text-muted-foreground">Liquidity</span>
                        <span>${LIQUIDITY_COST}</span>
                    </div>
                    <div className="flex justify-between">
                        <span className="text-muted-foreground">Maintenance</span>
                        <span>${PROTOCOL_COST}</span>
                    </div>
                    <div className="flex justify-between font-medium border-t border-foreground/10 pt-1.5">
                        <span>Total</span>
                        <span>${CREATION_COST}</span>
                    </div>
                </div>

                <div className="flex justify-between text-sm">
                    <span className="text-muted-foreground">Your balance</span>
                    <span>${formattedBalance}</span>
                </div>

                {needsApproval ? (
                    <Button
                        className="w-full"
                        disabled={!isConnected || isApproving}
                        onClick={handleApprove}
                    >
                        {isApproving ? <Spinner /> : "Approve USDC"}
                    </Button>
                ) : (
                    <Button
                        className="w-full"
                        disabled={
                            !isConnected ||
                            !name ||
                            !symbol ||
                            isCreating ||
                            (usdcBalance !== undefined && (usdcBalance as bigint) < CREATION_COST_PARSED)
                        }
                        onClick={handleCreate}
                    >
                        {isCreating ? <Spinner /> : "Create Market"}
                    </Button>
                )}
            </CardContent>
        </Card>
    );
}
