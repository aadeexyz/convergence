export function Hero() {
    return (
        <div className="rounded-xl bg-[radial-gradient(ellipse_at_bottom_left,var(--color-emerald-950)_0%,transparent_50%),radial-gradient(ellipse_at_bottom_right,var(--color-rose-950)_0%,transparent_50%)] px-8 py-14 md:px-14 md:py-20 flex flex-col gap-3">
            <h1 className="text-4xl md:text-6xl font-black tracking-tight leading-none">
                Trade the Trend
            </h1>
            <h1 className="text-4xl md:text-6xl font-black tracking-tight leading-none text-muted-foreground">
                Not the Token
            </h1>
            <p className="mt-3 text-sm md:text-base text-muted-foreground max-w-md">
                Trade attention markets powered by real-time search data. Go long or short on what the world is paying attention to.
            </p>
        </div>
    );
}
