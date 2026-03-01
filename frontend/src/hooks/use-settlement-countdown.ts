import { useEffect, useState } from "react";

export function useSettlementCountdown() {
    const [timeLeft, setTimeLeft] = useState("");

    useEffect(() => {
        function update() {
            const now = new Date();
            const midnight = new Date(now);
            midnight.setUTCDate(midnight.getUTCDate() + 1);
            midnight.setUTCHours(0, 0, 0, 0);
            const diff = midnight.getTime() - now.getTime();
            const h = Math.floor(diff / 3600000);
            const m = Math.floor((diff % 3600000) / 60000);
            const s = Math.floor((diff % 60000) / 1000);
            setTimeLeft(`${h.toString().padStart(2, "0")}:${m.toString().padStart(2, "0")}:${s.toString().padStart(2, "0")}`);
        }
        update();
        const interval = setInterval(update, 1000);
        return () => clearInterval(interval);
    }, []);

    return timeLeft;
}
