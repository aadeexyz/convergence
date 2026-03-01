"use client";

import { RiExternalLinkLine } from "@remixicon/react";
import { useConfig, useChainId } from "wagmi";
import { truncateHex } from "@/lib/utils";

type SonnerTxLinkProps = {
    txHash: `0x${string}` | undefined;
};

export function SonnerTxLink({ txHash }: SonnerTxLinkProps) {
    const { chains } = useConfig();
    const chainId = useChainId();

    const chain = chains.find((c) => c.id === chainId);
    const base = chain?.blockExplorers?.default?.url;

    return (
        <a
            href={`${base}/tx/${txHash}`}
            target="_blank"
            rel="noreferrer"
            className="underline text-muted-foreground hover:text-primary transition-colors flex items-center gap-1 text-sm [&_svg:not([class*='size-'])]:size-4"
        >
            <span>{truncateHex(txHash)}</span>
            <RiExternalLinkLine />
        </a>
    );
}
