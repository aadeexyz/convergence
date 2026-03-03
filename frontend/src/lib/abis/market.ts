export const marketAbi = [
    {
        type: "function",
        name: "burn",
        inputs: [
            { name: "isLong_", type: "bool", internalType: "bool" },
            {
                name: "positionTokenAmount_",
                type: "uint256",
                internalType: "uint256",
            },
            {
                name: "recipient_",
                type: "address",
                internalType: "address",
            },
        ],
        outputs: [],
        stateMutability: "nonpayable",
    },
    {
        type: "function",
        name: "redeem",
        inputs: [
            { name: "isLong_", type: "bool", internalType: "bool" },
            {
                name: "recipient_",
                type: "address",
                internalType: "address",
            },
        ],
        outputs: [],
        stateMutability: "nonpayable",
    },
    {
        type: "function",
        name: "settled",
        inputs: [],
        outputs: [{ name: "", type: "bool", internalType: "bool" }],
        stateMutability: "view",
    },
] as const;
