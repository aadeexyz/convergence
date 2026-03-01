"use client";

import Link from "next/link";
import { RiAddLine } from "@remixicon/react";
import { MarketFactoryCard } from "@/components/market-factory-card";
import { Button } from "@/components/ui/button";
import { useMarketFactories } from "@/hooks/use-market-factories";

const heroImages = [
    { src: "/assets/1.avif", alt: "Featured 1" },
    { src: "/assets/2.avif", alt: "Featured 2" },
    { src: "/assets/3.jpg", alt: "Featured 3" },
    { src: "/assets/4.png", alt: "Featured 4" },
    { src: "/assets/5.webp", alt: "Featured 5" },
    { src: "/assets/6.webp", alt: "Featured 6" },
    { src: "/assets/7.jpg", alt: "Featured 7" },
];

export const MarketsClientPage = () => {
    const { factories, isLoading, isError, error } = useMarketFactories();

    console.log(factories);

    if (isLoading) return <main>loading...</main>;
    if (isError) return <main>error</main>;

    return (
        <main className="space-y-6">
            {/* Hero grid */}
            <div className="relative grid grid-cols-5 grid-rows-2 gap-0 h-[280px] rounded-xl overflow-hidden">
                {/* Col 1: featured image, 2 rows */}
                <div className="row-span-2 relative">
                    <picture>
                        <img
                            src={heroImages[0].src}
                            alt={heroImages[0].alt}
                            className="absolute inset-0 w-full h-full object-cover"
                        />
                    </picture>
                </div>

                {/* Row 1: img green img img */}
                <div className="relative">
                    <picture>
                        <img
                            src={heroImages[1].src}
                            alt={heroImages[1].alt}
                            className="absolute inset-0 w-full h-full object-cover"
                        />
                    </picture>
                </div>
                <div className="bg-emerald-500" />
                <div className="relative">
                    <picture>
                        <img
                            src={heroImages[2].src}
                            alt={heroImages[2].alt}
                            className="absolute inset-0 w-full h-full object-cover"
                        />
                    </picture>
                </div>
                <div className="relative">
                    <picture>
                        <img
                            src={heroImages[3].src}
                            alt={heroImages[3].alt}
                            className="absolute inset-0 w-full h-full object-cover"
                        />
                    </picture>
                </div>

                {/* Row 2: img img red img */}
                <div className="relative">
                    <picture>
                        <img
                            src={heroImages[4].src}
                            alt={heroImages[4].alt}
                            className="absolute inset-0 w-full h-full object-cover"
                        />
                    </picture>
                </div>
                <div className="relative">
                    <picture>
                        <img
                            src={heroImages[5].src}
                            alt={heroImages[5].alt}
                            className="absolute inset-0 w-full h-full object-cover"
                        />
                    </picture>
                </div>
                <div className="bg-red-500" />
                <div className="relative">
                    <picture>
                        <img
                            src={heroImages[6].src}
                            alt={heroImages[6].alt}
                            className="absolute inset-0 w-full h-full object-cover"
                        />
                    </picture>
                </div>

                {/* Text overlay */}
                <div className="absolute inset-0 pointer-events-none flex flex-col justify-between">
                    <h1 className="text-[clamp(3rem,10vw,9rem)] font-black leading-none tracking-tighter text-white drop-shadow-lg">
                        Buy the Trend
                    </h1>
                    <h1 className="text-[clamp(3rem,10vw,9rem)] font-black leading-none tracking-tighter text-white drop-shadow-lg text-right">
                        Not the Token
                    </h1>
                </div>
            </div>

            <div className="flex items-center justify-between">
                <h2 className="text-2xl font-bold">Markets</h2>

                <Button render={<Link href="/markets/create" />}>
                    <RiAddLine className="size-4 mr-1" />
                    Create Market
                </Button>
            </div>
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                {factories.map((factory) => (
                    <MarketFactoryCard
                        key={factory}
                        marketFactoryAddress={factory}
                    />
                ))}
            </div>
        </main>
    );
};
