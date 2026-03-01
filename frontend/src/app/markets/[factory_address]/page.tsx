import { MarketDetailClientPage } from "./page.client";

type Props = {
    params: Promise<{ factory_address: string }>;
};

export default async function MarketDetailPage({ params }: Props) {
    const { factory_address } = await params;
    return <MarketDetailClientPage factoryAddress={factory_address} />;
}
