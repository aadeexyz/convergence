"use client";

import { ConnectKitButton } from "connectkit";
import { cn } from "@/lib/utils";
import { Spinner } from "@/components/ui/spinner";
import { Button } from "@/components/ui/button";

export function ConnectWalletButton({ className }: { className?: string }) {
    return (
        <ConnectKitButton.Custom>
            {({ isConnected, isConnecting, show, truncatedAddress }) => (
                <Button
                    onClick={show}
                    className={cn("hover:cursor-pointer", className)}
                    disabled={isConnecting}
                    variant={isConnected ? "secondary" : "outline"}
                >
                    {isConnected ? (
                        truncatedAddress
                    ) : isConnecting ? (
                        <Spinner />
                    ) : (
                        "Connect Wallet"
                    )}
                </Button>
            )}
        </ConnectKitButton.Custom>
    );
}
