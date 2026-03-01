import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"
import type { Hex } from "viem"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export function truncateHex(hex: Hex | undefined) {
    if (!hex) return "";

    return `${hex.slice(0, 6)}•••${hex.slice(-4)}`;
}