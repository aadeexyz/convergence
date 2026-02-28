import type { NextConfig } from "next";

const nextConfig: NextConfig = {
    redirects: async () => [
        {
            source: "/",
            destination: "/markets",
            permanent: true,
        },
    ],
};

export default nextConfig;
