export const marketAbi = [
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
