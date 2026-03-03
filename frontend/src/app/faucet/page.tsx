import Link from "next/link";
import { RiArrowLeftSLine } from "@remixicon/react";
import { FaucetCard } from "@/components/faucet-card";

export default function FaucetPage() {
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

            <FaucetCard />
        </main>
    );
}
