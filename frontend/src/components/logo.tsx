export function Logo({ size = 28 }: { size?: number }) {
    const width = Math.round(size * 2.8);
    const boxHeight = Math.round(size * 0.5);

    return (
        <div className="relative flex items-end" style={{ width, height: size }}>
            {/* Bottom box with glow effect */}
            <div
                className="absolute bottom-0 left-0 right-0 overflow-hidden rounded-sm"
                style={{ height: boxHeight }}
            >
                {/* Green glow — bottom left */}
                <div
                    className="absolute rounded-full bg-emerald-950 blur-md"
                    style={{
                        width: size * 1.4,
                        height: size * 1.4,
                        bottom: -size * 0.7,
                        left: -size * 0.3,
                    }}
                />
                {/* Red glow — bottom right */}
                <div
                    className="absolute rounded-full bg-rose-950 blur-md"
                    style={{
                        width: size * 1.4,
                        height: size * 1.4,
                        bottom: -size * 0.7,
                        right: -size * 0.3,
                    }}
                />
                {/* Fade — top of box fades from background to transparent */}
                <div
                    className="absolute inset-0"
                    style={{ background: "linear-gradient(to bottom, var(--background) 0%, transparent 100%)" }}
                />
            </div>
            {/* WAM text — sits on top, vertically centered on the box edge */}
            <span
                className="relative z-10 w-full text-center font-extrabold text-white tracking-widest leading-none"
                style={{ fontSize: size * 0.55, marginBottom: boxHeight * 0.15 }}
            >
                WAM
            </span>
        </div>
    );
}
