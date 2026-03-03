import Link from "next/link";
import { Logo } from "./logo";
import { ConnectWalletButton } from "./connect-wallet-button";

export function Navbar() {
    return (
        <header className="sticky left-0 top-0 z-[100] flex w-full flex-col border-b border-border bg-background">
            <nav className="flex py-3 bg-background justify-center items-center px-5">
                <div className="container flex flex-col md:flex-row items-center justify-between w-full gap-2">
                    <div className="flex justify-between items-center w-full">
                        <Link href="/" className="flex items-center">
                            <Logo size={40} />
                        </Link>

                        <div className="flex items-center gap-3 md:gap-5">
                            <Link href="/faucet" className="text-sm text-muted-foreground hover:text-foreground transition-colors">
                                Faucet
                            </Link>
                            <ConnectWalletButton className="min-w-30 md:min-w-37.5" />
                        </div>
                    </div>
                </div>
            </nav>
        </header>
    );
}
