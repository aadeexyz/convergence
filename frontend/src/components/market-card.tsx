import type { Address } from "viem";

type MarketCardProps = {
    marketAddress: Address;
};

export function MarketCard({ marketAddress }: MarketCardProps) {
    return <div>MarketCard</div>;
}
