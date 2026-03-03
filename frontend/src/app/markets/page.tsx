import Link from "next/link";
import { RiAddLine } from "@remixicon/react";
import { Hero } from "@/components/hero";
import { MarketFactories } from "@/components/market-factories";
import { Button } from "@/components/ui/button";

export default function MarketsPage() {
    return (
        <main className="space-y-6">
            <Hero />

            <div className="flex items-center justify-between">
                <h2 className="text-2xl font-bold">Markets</h2>

                <Button render={<Link href="/markets/create" />}>
                    <RiAddLine className="size-4 mr-1" />
                    Create Market
                </Button>
            </div>

            <MarketFactories />
        </main>
    );
}
