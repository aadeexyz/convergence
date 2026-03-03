import Link from "next/link";
import { RiArrowLeftSLine } from "@remixicon/react";
import { CreateMarketForm } from "@/components/create-market-form";

export default function CreateMarketPage() {
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

            <CreateMarketForm />
        </main>
    );
}
