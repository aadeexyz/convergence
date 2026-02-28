"use client";

import { WagmiProvider, createConfig, http } from "wagmi";
import { base } from "viem/chains";
import { getDefaultConfig, ConnectKitProvider } from "connectkit";
// import { useTheme } from "next-themes";

const config = createConfig(
    getDefaultConfig({
        chains: [base],
        transports: {
            [base.id]: http(process.env.NEXT_PUBLIC_BASE_RPC_URL!),
        },
        walletConnectProjectId:
            process.env.NEXT_PUBLIC_WALLETCONNECT_PROJECT_ID!,
        appName: "Lora Finance",
        appDescription: "Rent Upside Every Second",
        appUrl: "https://testnet.lora.finance",
        appIcon: "/static/brand/logo.png",
        enableFamily: false,
    }),
);

export function Web3Provider({ children }: { children: React.ReactNode }) {
    // const { theme } = useTheme();

    return (
        <WagmiProvider config={config}>
            <ConnectKitProvider
                customTheme={{
                    "--ck-font-family": "var(--font-sans)",
                }}
                // mode={
                //     theme === "auto"
                //         ? "auto"
                //         : theme === "dark"
                //         ? "dark"
                //         : "light"
                // }
            >
                {children}
            </ConnectKitProvider>
        </WagmiProvider>
    );
}
