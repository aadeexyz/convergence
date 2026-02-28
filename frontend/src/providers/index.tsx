// import { Toaster } from "@/components/ui/sonner";
import { ThemeProvider } from "./theme";
import { ReactQueryProvider } from "./react-query";
import { Web3Provider } from "./web3";

export function Providers({ children }: { children: React.ReactNode }) {
    return (
        <ThemeProvider attribute="class" defaultTheme="dark" forcedTheme="dark">
            <ReactQueryProvider>
                <Web3Provider>{children}</Web3Provider>
            </ReactQueryProvider>
        </ThemeProvider>
    );
}
