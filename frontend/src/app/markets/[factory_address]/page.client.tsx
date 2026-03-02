"use client";

import { useState, useEffect } from "react";
import Link from "next/link";
import { useWriteContract, useReadContract, useTransactionReceipt, useAccount } from "wagmi";
import { useQueryClient } from "@tanstack/react-query";
import { toast } from "sonner";
import { Card, CardContent } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import {
    InputGroup,
    InputGroupAddon,
    InputGroupInput,
    InputGroupText,
} from "@/components/ui/input-group";
import { Tabs, TabsList, TabsTrigger, TabsContent } from "@/components/ui/tabs";
import { Spinner } from "@/components/ui/spinner";
import { SonnerTxLink } from "@/components/sonner-tx-link";
import { useMarketFactory } from "@/hooks/use-market-factory";
import { useLatestMarket } from "@/hooks/use-latest-market";
import { useOracle } from "@/hooks/use-oracle";
import { useSettlementCountdown } from "@/hooks/use-settlement-countdown";
import { SettledMarkets } from "@/components/settled-markets";
import { OracleHistoryChart } from "@/components/oracle-history-chart";
import { erc20Abi, routerAbi } from "@/lib/abis";
import { routerAddress } from "@/lib/metadata";
import { cn } from "@/lib/utils";
import { MAX_ALLOWANCE } from "@/lib/constants";
import { formatUnits, maxUint256, parseUnits } from "viem";
import type { Address } from "viem";

const AMOUNT_PRESETS = ["1", "5", "10", "100"];

type Props = {
    factoryAddress: string;
};

export function MarketDetailClientPage({ factoryAddress }: Props) {
    const address = factoryAddress as Address;
    const { factory } = useMarketFactory(address);
    const { market } = useLatestMarket(address);
    const { oracle } = useOracle(address);
    const countdown = useSettlementCountdown();
    const { address: userAddress, isConnected } = useAccount();
    const queryClient = useQueryClient();

    const [outcome, setOutcome] = useState<"long" | "short">("long");
    const [amount, setAmount] = useState("");
    const [shares, setShares] = useState("");

    // --- Read allowances ---
    const { data: collateralAllowance } = useReadContract({
        address: factory?.collateralToken,
        abi: erc20Abi,
        functionName: "allowance",
        args: userAddress && factory ? [userAddress, routerAddress] : undefined,
        query: { enabled: !!userAddress && !!factory },
    });

    const positionTokenAddress = market
        ? outcome === "long"
            ? market.longPositionToken
            : market.shortPositionToken
        : undefined;

    const { data: positionAllowance } = useReadContract({
        address: positionTokenAddress,
        abi: erc20Abi,
        functionName: "allowance",
        args: userAddress && positionTokenAddress ? [userAddress, routerAddress] : undefined,
        query: { enabled: !!userAddress && !!positionTokenAddress },
    });

    // --- Read balances ---
    const { data: collateralBalance } = useReadContract({
        address: factory?.collateralToken,
        abi: erc20Abi,
        functionName: "balanceOf",
        args: userAddress ? [userAddress] : undefined,
        query: { enabled: !!userAddress && !!factory },
    });

    const { data: positionBalance } = useReadContract({
        address: positionTokenAddress,
        abi: erc20Abi,
        functionName: "balanceOf",
        args: userAddress ? [userAddress] : undefined,
        query: { enabled: !!userAddress && !!positionTokenAddress },
    });

    // --- Buy flow: approve collateral (if needed) then mint ---
    const approveCollateral = useWriteContract({
        mutation: {
            onError() {
                toast.error("Approval failed");
            },
            onSuccess(data) {
                toast.loading("Approving collateral…", {
                    description: <SonnerTxLink txHash={data} />,
                    id: `approve-${data}`,
                });
            },
        },
    });
    const approveCollateralReceipt = useTransactionReceipt({
        hash: approveCollateral.data,
    });

    const mintPosition = useWriteContract({
        mutation: {
            onError() {
                toast.error("Mint failed");
            },
            onSuccess(data) {
                toast.loading("Minting position…", {
                    description: <SonnerTxLink txHash={data} />,
                    id: data,
                });
            },
        },
    });
    const mintReceipt = useTransactionReceipt({ hash: mintPosition.data });

    // When collateral approval succeeds, fire the mint
    useEffect(() => {
        if (approveCollateralReceipt.status !== "success") return;
        if (!factory) return;

        toast.success("Collateral approved", {
            id: `approve-${approveCollateral.data}`,
        });

        const decimals = factory.decimals;
        const parsedAmount = parseUnits(amount, decimals);
        const isLong = outcome === "long";

        mintPosition.mutate({
            address: routerAddress,
            abi: routerAbi,
            functionName: "mint",
            args: [address, isLong, parsedAmount],
        });
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [approveCollateralReceipt.status]);

    // When mint succeeds, invalidate queries
    useEffect(() => {
        if (mintReceipt.status === "success") {
            queryClient.invalidateQueries();
            toast.success("Position minted!", {
                description: <SonnerTxLink txHash={mintPosition.data} />,
                id: mintPosition.data,
            });
            setAmount("");
        } else if (mintReceipt.status === "error") {
            toast.error("Mint transaction failed", {
                description: <SonnerTxLink txHash={mintPosition.data} />,
                id: mintPosition.data,
            });
        }
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [mintReceipt.status]);

    // --- Sell flow: approve position token (if needed) then burn ---
    const approvePosition = useWriteContract({
        mutation: {
            onError() {
                toast.error("Approval failed");
            },
            onSuccess(data) {
                toast.loading("Approving position token…", {
                    description: <SonnerTxLink txHash={data} />,
                    id: `approve-${data}`,
                });
            },
        },
    });
    const approvePositionReceipt = useTransactionReceipt({
        hash: approvePosition.data,
    });

    const burnPosition = useWriteContract({
        mutation: {
            onError() {
                toast.error("Burn failed");
            },
            onSuccess(data) {
                toast.loading("Burning position…", {
                    description: <SonnerTxLink txHash={data} />,
                    id: data,
                });
            },
        },
    });
    const burnReceipt = useTransactionReceipt({ hash: burnPosition.data });

    // When position token approval succeeds, fire the burn
    useEffect(() => {
        if (approvePositionReceipt.status !== "success") return;
        if (!market) return;

        toast.success("Position token approved", {
            id: `approve-${approvePosition.data}`,
        });

        const decimals = market.decimals;
        const parsedShares = parseUnits(shares, decimals);
        const isLong = outcome === "long";

        burnPosition.mutate({
            address: routerAddress,
            abi: routerAbi,
            functionName: "burn",
            args: [address, isLong, parsedShares],
        });
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [approvePositionReceipt.status]);

    // When burn succeeds, invalidate queries
    useEffect(() => {
        if (burnReceipt.status === "success") {
            queryClient.invalidateQueries();
            toast.success("Position burned!", {
                description: <SonnerTxLink txHash={burnPosition.data} />,
                id: burnPosition.data,
            });
            setShares("");
        } else if (burnReceipt.status === "error") {
            toast.error("Burn transaction failed", {
                description: <SonnerTxLink txHash={burnPosition.data} />,
                id: burnPosition.data,
            });
        }
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [burnReceipt.status]);

    // --- Approval state ---
    const needsCollateralApproval =
        collateralAllowance === undefined || (collateralAllowance as bigint) < MAX_ALLOWANCE;

    const needsPositionApproval =
        positionAllowance === undefined || (positionAllowance as bigint) < MAX_ALLOWANCE;

    // --- Handlers ---
    function handleApproveCollateral() {
        if (!factory) return;
        approveCollateral.mutate({
            address: factory.collateralToken,
            abi: erc20Abi,
            functionName: "approve",
            args: [routerAddress, maxUint256],
        });
    }

    function handleBuy() {
        if (!factory || !amount) return;
        const decimals = factory.decimals;
        const parsedAmount = parseUnits(amount, decimals);
        const isLong = outcome === "long";

        mintPosition.mutate({
            address: routerAddress,
            abi: routerAbi,
            functionName: "mint",
            args: [address, isLong, parsedAmount],
        });
    }

    function handleApprovePosition() {
        if (!positionTokenAddress) return;
        approvePosition.mutate({
            address: positionTokenAddress,
            abi: erc20Abi,
            functionName: "approve",
            args: [routerAddress, maxUint256],
        });
    }

    function handleSell() {
        if (!market || !shares) return;
        const decimals = market.decimals;
        const parsedShares = parseUnits(shares, decimals);
        const isLong = outcome === "long";

        burnPosition.mutate({
            address: routerAddress,
            abi: routerAbi,
            functionName: "burn",
            args: [address, isLong, parsedShares],
        });
    }

    const isApproving =
        approveCollateral.isPending ||
        (approveCollateralReceipt.status === "pending" && !!approveCollateral.data) ||
        approvePosition.isPending ||
        (approvePositionReceipt.status === "pending" && !!approvePosition.data);

    const isBuying =
        mintPosition.isPending ||
        (mintReceipt.status === "pending" && !!mintPosition.data);

    const isSelling =
        burnPosition.isPending ||
        (burnReceipt.status === "pending" && !!burnPosition.data);

    if (!factory)
        return <main className="text-muted-foreground">Loading...</main>;

    const longPrice = market ? Number(formatUnits(market.longPrice, 6)) : 0;
    const shortPrice = market ? Number(formatUnits(market.shortPrice, 6)) : 0;
    const tvl = market ? Number(formatUnits(market.totalLiquidity, 6)) : 0;
    const oracleIndex = oracle?.latestRound
        ? Number(formatUnits(oracle.latestRound.index, oracle.decimals))
        : null;

    const formattedCollateralBalance =
        collateralBalance !== undefined
            ? Number(formatUnits(collateralBalance as bigint, factory.decimals)).toLocaleString(undefined, { minimumFractionDigits: 2, maximumFractionDigits: 2 })
            : "—";

    const formattedPositionBalance =
        positionBalance !== undefined && market
            ? Number(formatUnits(positionBalance as bigint, market.decimals)).toLocaleString(undefined, { minimumFractionDigits: 2, maximumFractionDigits: 2 })
            : "—";

    return (
        <main className="space-y-4">
            <div>
                <Link
                    href="/markets"
                    className="text-sm text-muted-foreground hover:text-foreground transition-colors"
                >
                    ← Back to markets
                </Link>
            </div>

            {/* Header */}
            <div className="flex items-center justify-between">
                <div className="flex items-baseline gap-3">
                    <h1 className="text-2xl font-bold">{factory.name}</h1>
                    <span className="text-muted-foreground">
                        ${factory.symbol}
                    </span>
                </div>
            </div>

            {factory.state === 0 ? (
                <p className="animate-pulse text-muted-foreground">
                    Market is being created...
                </p>
            ) : (
                <div className="grid grid-cols-1 lg:grid-cols-3 gap-4">
                    {/* Chart + settled markets — left 2/3 */}
                    <div className="lg:col-span-2 space-y-4">
                        <OracleHistoryChart
                            factoryAddress={address}
                            oracleDecimals={oracle?.decimals ?? 8}
                        />
                        <SettledMarkets factoryAddress={address} />
                    </div>

                    {/* Right column — trade + stats stacked */}
                    <div className="flex flex-col gap-4">
                        {/* Trade Panel */}
                        <Card className="ring-0 border border-foreground/10">
                            <CardContent className="pt-5 space-y-5">
                                <Tabs defaultValue="buy">
                                    <TabsList className="w-full">
                                        <TabsTrigger
                                            value="buy"
                                            className="flex-1"
                                        >
                                            Buy
                                        </TabsTrigger>
                                        <TabsTrigger
                                            value="sell"
                                            className="flex-1"
                                        >
                                            Sell
                                        </TabsTrigger>
                                    </TabsList>

                                    <TabsContent
                                        value="buy"
                                        className="space-y-4 pt-4"
                                    >
                                        {/* Outcome selector */}
                                        <div className="grid grid-cols-2 gap-2">
                                            <button
                                                onClick={() =>
                                                    setOutcome("long")
                                                }
                                                className={cn(
                                                    "rounded-lg py-2 text-sm font-medium transition-colors",
                                                    outcome === "long"
                                                        ? "bg-emerald-600 text-white"
                                                        : "bg-muted text-muted-foreground hover:text-foreground",
                                                )}
                                            >
                                                Long{" "}
                                                {(longPrice * 100).toFixed(0)}¢
                                            </button>
                                            <button
                                                onClick={() =>
                                                    setOutcome("short")
                                                }
                                                className={cn(
                                                    "rounded-lg py-2 text-sm font-medium transition-colors",
                                                    outcome === "short"
                                                        ? "bg-rose-600 text-white"
                                                        : "bg-muted text-muted-foreground hover:text-foreground",
                                                )}
                                            >
                                                Short{" "}
                                                {(shortPrice * 100).toFixed(0)}¢
                                            </button>
                                        </div>

                                        {/* Amount */}
                                        <div className="space-y-2">
                                            <div className="flex justify-between items-center">
                                                <span className="text-sm text-muted-foreground">
                                                    Amount
                                                </span>
                                                <span className="text-xs text-muted-foreground">
                                                    Balance: ${formattedCollateralBalance}
                                                </span>
                                            </div>
                                            <InputGroup>
                                                <InputGroupInput
                                                    type="number"
                                                    placeholder="0"
                                                    value={amount}
                                                    onChange={(e) =>
                                                        setAmount(
                                                            e.target.value,
                                                        )
                                                    }
                                                    className="text-right text-lg"
                                                />
                                                <InputGroupAddon align="inline-start">
                                                    <InputGroupText>
                                                        $
                                                    </InputGroupText>
                                                </InputGroupAddon>
                                            </InputGroup>
                                            <div className="flex gap-1.5">
                                                {AMOUNT_PRESETS.map(
                                                    (preset) => (
                                                        <button
                                                            key={preset}
                                                            onClick={() =>
                                                                setAmount(
                                                                    String(Number(amount || "0") + Number(preset)),
                                                                )
                                                            }
                                                            className="flex-1 rounded-md bg-muted px-2 py-1 text-xs font-medium text-muted-foreground hover:text-foreground transition-colors"
                                                        >
                                                            +${preset}
                                                        </button>
                                                    ),
                                                )}
                                            </div>
                                        </div>

                                        {/* Trade button */}
                                        {needsCollateralApproval ? (
                                            <Button
                                                className="w-full"
                                                disabled={!isConnected || isApproving}
                                                onClick={handleApproveCollateral}
                                            >
                                                {isApproving ? <Spinner /> : "Approve USDC"}
                                            </Button>
                                        ) : (
                                            <Button
                                                className="w-full"
                                                disabled={!isConnected || !amount || isBuying || (collateralBalance !== undefined && factory && (collateralBalance as bigint) < parseUnits(amount || "0", factory.decimals))}
                                                onClick={handleBuy}
                                            >
                                                {isBuying ? <Spinner /> : "Trade"}
                                            </Button>
                                        )}
                                    </TabsContent>

                                    <TabsContent
                                        value="sell"
                                        className="space-y-4 pt-4"
                                    >
                                        {/* Outcome selector */}
                                        <div className="grid grid-cols-2 gap-2">
                                            <button
                                                onClick={() =>
                                                    setOutcome("long")
                                                }
                                                className={cn(
                                                    "rounded-lg py-2 text-sm font-medium transition-colors",
                                                    outcome === "long"
                                                        ? "bg-emerald-600 text-white"
                                                        : "bg-muted text-muted-foreground hover:text-foreground",
                                                )}
                                            >
                                                Long{" "}
                                                {(longPrice * 100).toFixed(0)}¢
                                            </button>
                                            <button
                                                onClick={() =>
                                                    setOutcome("short")
                                                }
                                                className={cn(
                                                    "rounded-lg py-2 text-sm font-medium transition-colors",
                                                    outcome === "short"
                                                        ? "bg-rose-600 text-white"
                                                        : "bg-muted text-muted-foreground hover:text-foreground",
                                                )}
                                            >
                                                Short{" "}
                                                {(shortPrice * 100).toFixed(0)}¢
                                            </button>
                                        </div>

                                        {/* Shares */}
                                        <div className="space-y-2">
                                            <div className="flex justify-between items-center">
                                                <span className="text-sm text-muted-foreground">
                                                    Shares
                                                </span>
                                                <span className="text-xs text-muted-foreground">
                                                    Balance: {formattedPositionBalance}
                                                </span>
                                            </div>
                                            <InputGroup>
                                                <InputGroupInput
                                                    type="number"
                                                    placeholder="0"
                                                    value={shares}
                                                    onChange={(e) =>
                                                        setShares(
                                                            e.target.value,
                                                        )
                                                    }
                                                    className="text-right text-lg"
                                                />
                                            </InputGroup>
                                            <div className="flex gap-1.5">
                                                {[25, 50, 75].map((pct) => (
                                                    <button
                                                        key={pct}
                                                        onClick={() => {
                                                            if (positionBalance !== undefined && market) {
                                                                const portion = (positionBalance as bigint) * BigInt(pct) / BigInt(100);
                                                                setShares(formatUnits(portion, market.decimals));
                                                            }
                                                        }}
                                                        className="flex-1 rounded-md bg-muted px-2 py-1 text-xs font-medium text-muted-foreground hover:text-foreground transition-colors"
                                                    >
                                                        {pct}%
                                                    </button>
                                                ))}
                                                <button
                                                    onClick={() => {
                                                        if (positionBalance !== undefined && market) {
                                                            setShares(formatUnits(positionBalance as bigint, market.decimals));
                                                        }
                                                    }}
                                                    className="flex-1 rounded-md bg-muted px-2 py-1 text-xs font-medium text-muted-foreground hover:text-foreground transition-colors"
                                                >
                                                    Max
                                                </button>
                                            </div>
                                        </div>

                                        {/* Trade button */}
                                        {needsPositionApproval ? (
                                            <Button
                                                className="w-full"
                                                disabled={!isConnected || isApproving}
                                                onClick={handleApprovePosition}
                                            >
                                                {isApproving ? <Spinner /> : "Approve Position Token"}
                                            </Button>
                                        ) : (
                                            <Button
                                                className="w-full"
                                                disabled={!isConnected || !shares || isSelling || (positionBalance !== undefined && !!market && (positionBalance as bigint) < parseUnits(shares || "0", market.decimals))}
                                                onClick={handleSell}
                                            >
                                                {isSelling ? <Spinner /> : "Trade"}
                                            </Button>
                                        )}
                                    </TabsContent>
                                </Tabs>
                            </CardContent>
                        </Card>

                        {/* Stats */}
                        <Card>
                            <CardContent>
                                <dl className="space-y-2 text-sm">
                                    <div className="flex justify-between">
                                        <dt className="text-muted-foreground">
                                            Index
                                        </dt>
                                        <dd>
                                            {oracleIndex !== null
                                                ? oracleIndex.toFixed(2)
                                                : "—"}
                                        </dd>
                                    </div>
                                    <div className="flex justify-between">
                                        <dt className="text-muted-foreground">
                                            Index Updated
                                        </dt>
                                        <dd>
                                            {oracle?.latestRound?.timestamp
                                                ? (() => {
                                                      const secs = Math.floor(Date.now() / 1000) - Number(oracle.latestRound.timestamp);
                                                      if (secs < 60) return `${secs}s ago`;
                                                      if (secs < 3600) return `${Math.floor(secs / 60)}m ago`;
                                                      if (secs < 86400) return `${Math.floor(secs / 3600)}h ago`;
                                                      return `${Math.floor(secs / 86400)}d ago`;
                                                  })()
                                                : "—"}
                                        </dd>
                                    </div>
                                    <div className="flex justify-between">
                                        <dt className="text-muted-foreground">
                                            TVL
                                        </dt>
                                        <dd>${tvl.toLocaleString()}</dd>
                                    </div>
                                    <div className="flex justify-between">
                                        <dt className="text-muted-foreground">
                                            Settlement
                                        </dt>
                                        <dd>{countdown}</dd>
                                    </div>
                                </dl>
                            </CardContent>
                        </Card>
                    </div>
                </div>
            )}

        </main>
    );
}
