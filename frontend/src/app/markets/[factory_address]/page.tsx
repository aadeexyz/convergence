import Link from "next/link";
import { RiArrowLeftSLine } from "@remixicon/react";
import { MarketHeader } from "@/components/market-header";
import { MarketContent } from "@/components/market-content";
import type { Address } from "viem";

type Props = {
    params: Promise<{ factory_address: string }>;
};

export default async function MarketDetailPage({ params }: Props) {
    const { factory_address } = await params;
    const address = factory_address as Address;

    return (
        <main className="space-y-4">
            <div>
                <Link
                    href="/markets"
                    className="text-sm text-muted-foreground hover:text-foreground transition-colors flex items-center gap-0.5"
                >
                    <RiArrowLeftSLine className="size-4" /> Back to markets
                </Link>
            </div>

            <MarketHeader factoryAddress={address} />
            <MarketContent factoryAddress={address} />
        </main>
    );
}
