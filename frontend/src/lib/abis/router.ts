import type { Abi } from "viem";

export const routerAbi = [
    {
        type: "function",
        name: "burn",
        inputs: [
            {
                name: "factory_",
                type: "address",
                internalType: "address",
            },
            { name: "isLong_", type: "bool", internalType: "bool" },
            {
                name: "positionTokenAmount_",
                type: "uint256",
                internalType: "uint256",
            },
        ],
        outputs: [],
        stateMutability: "nonpayable",
    },
    {
        type: "function",
        name: "mint",
        inputs: [
            {
                name: "factory_",
                type: "address",
                internalType: "address",
            },
            { name: "isLong_", type: "bool", internalType: "bool" },
            {
                name: "collateralAmount_",
                type: "uint256",
                internalType: "uint256",
            },
        ],
        outputs: [],
        stateMutability: "nonpayable",
    },
    {
        type: "function",
        name: "redeem",
        inputs: [
            {
                name: "factory_",
                type: "address",
                internalType: "address",
            },
            {
                name: "marketIndex_",
                type: "uint256",
                internalType: "uint256",
            },
            { name: "isLong_", type: "bool", internalType: "bool" },
        ],
        outputs: [],
        stateMutability: "nonpayable",
    },
    {
        type: "function",
        name: "redeemAll",
        inputs: [
            {
                name: "factory_",
                type: "address",
                internalType: "address",
            },
        ],
        outputs: [],
        stateMutability: "nonpayable",
    },
    {
        type: "event",
        name: "Burned",
        inputs: [
            {
                name: "factory",
                type: "address",
                indexed: true,
                internalType: "address",
            },
            {
                name: "market",
                type: "address",
                indexed: true,
                internalType: "address",
            },
            {
                name: "account",
                type: "address",
                indexed: true,
                internalType: "address",
            },
            {
                name: "isLong",
                type: "bool",
                indexed: false,
                internalType: "bool",
            },
            {
                name: "positionTokenAmount",
                type: "uint256",
                indexed: false,
                internalType: "uint256",
            },
        ],
        anonymous: false,
    },
    {
        type: "event",
        name: "Minted",
        inputs: [
            {
                name: "factory",
                type: "address",
                indexed: true,
                internalType: "address",
            },
            {
                name: "market",
                type: "address",
                indexed: true,
                internalType: "address",
            },
            {
                name: "account",
                type: "address",
                indexed: true,
                internalType: "address",
            },
            {
                name: "isLong",
                type: "bool",
                indexed: false,
                internalType: "bool",
            },
            {
                name: "collateralAmount",
                type: "uint256",
                indexed: false,
                internalType: "uint256",
            },
        ],
        anonymous: false,
    },
    {
        type: "event",
        name: "Redeemed",
        inputs: [
            {
                name: "factory",
                type: "address",
                indexed: true,
                internalType: "address",
            },
            {
                name: "market",
                type: "address",
                indexed: true,
                internalType: "address",
            },
            {
                name: "account",
                type: "address",
                indexed: true,
                internalType: "address",
            },
            {
                name: "isLong",
                type: "bool",
                indexed: false,
                internalType: "bool",
            },
            {
                name: "collateralAmount",
                type: "uint256",
                indexed: false,
                internalType: "uint256",
            },
        ],
        anonymous: false,
    },
    { type: "error", name: "MarketNotLive", inputs: [] },
    { type: "error", name: "NoSettledMarkets", inputs: [] },
] as const satisfies Abi;
