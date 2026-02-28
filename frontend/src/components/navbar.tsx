import Image from "next/image";
import Link from "next/link";
import { ConnectWalletButton } from "./connect-wallet-button";

export function Navbar() {
    return (
        <header className="sticky left-0 top-0 z-[100] flex w-full flex-col border-b border-border bg-background">
            <nav className="flex py-3 bg-background justify-center items-center px-5">
                <div className="container flex flex-col md:flex-row items-center justify-between w-full gap-2">
                    <div className="flex justify-between items-center w-full">
                        <Link href="/" className="flex items-center gap-2">
                            {/* <Image
                                src="/static/brand/logo.png"
                                alt="Enchanted"
                                width={28}
                                height={28}
                            /> */}
                            <span className="font-bold text-lg text-foreground tracking-wide">
                                wam
                            </span>
                        </Link>

                        <div className="flex items-center gap-5">
                            <ConnectWalletButton className="min-w-[150px]" />
                        </div>
                    </div>
                </div>
            </nav>
        </header>
    );
}
