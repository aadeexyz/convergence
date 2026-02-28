"use client";

import { useMarketFactories } from "@/hooks/use-market-factories";

export const MarketsClientPage = () => {
    const { factories, isLoading, isError, error } = useMarketFactories();

    console.log(factories);

    if (isLoading) return <main>loading...</main>;
    if (isError) return <main>error</main>;

    return <main>hello</main>;
};
