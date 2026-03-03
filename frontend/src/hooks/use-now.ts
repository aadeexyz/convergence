import { useState, useEffect } from "react";

export function useNow(active: boolean, intervalMs: number = 1000) {
    const [now, setNow] = useState(() => Date.now());

    useEffect(() => {
        if (!active) return;
        const id = setInterval(() => setNow(Date.now()), intervalMs);
        return () => clearInterval(id);
    }, [active, intervalMs]);

    return now;
}