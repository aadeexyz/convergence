import type { Abi } from "viem";

export const factoryOfMarketFactoriesAbi = [
    {
        type: "constructor",
        inputs: [
            {
                name: "collateralToken_",
                type: "address",
                internalType: "address",
            },
            {
                name: "liquidityFee_",
                type: "uint256",
                internalType: "uint256",
            },
            {
                name: "protocolFee_",
                type: "uint256",
                internalType: "uint256",
            },
            {
                name: "forwarderAddress_",
                type: "address",
                internalType: "address",
            },
            {
                name: "oracleDecimals_",
                type: "uint8",
                internalType: "uint8",
            },
            {
                name: "owner_",
                type: "address",
                internalType: "address",
            },
        ],
        stateMutability: "nonpayable",
    },
    {
        type: "function",
        name: "cancelOwnershipHandover",
        inputs: [],
        outputs: [],
        stateMutability: "payable",
    },
    {
        type: "function",
        name: "collateralToken",
        inputs: [],
        outputs: [
            {
                name: "",
                type: "address",
                internalType: "address",
            },
        ],
        stateMutability: "view",
    },
    {
        type: "function",
        name: "completeOwnershipHandover",
        inputs: [
            {
                name: "pendingOwner",
                type: "address",
                internalType: "address",
            },
        ],
        outputs: [],
        stateMutability: "payable",
    },
    {
        type: "function",
        name: "createMarketFactory",
        inputs: [
            {
                name: "name_",
                type: "string",
                internalType: "string",
            },
            {
                name: "symbol_",
                type: "string",
                internalType: "string",
            },
        ],
        outputs: [
            {
                name: "",
                type: "address",
                internalType: "address",
            },
        ],
        stateMutability: "nonpayable",
    },
    {
        type: "function",
        name: "creationFee",
        inputs: [],
        outputs: [
            {
                name: "",
                type: "uint256",
                internalType: "uint256",
            },
        ],
        stateMutability: "view",
    },
    {
        type: "function",
        name: "forwarderAddress",
        inputs: [],
        outputs: [
            {
                name: "",
                type: "address",
                internalType: "address",
            },
        ],
        stateMutability: "view",
    },
    {
        type: "function",
        name: "liquidityFee",
        inputs: [],
        outputs: [
            {
                name: "",
                type: "uint256",
                internalType: "uint256",
            },
        ],
        stateMutability: "view",
    },
    {
        type: "function",
        name: "marketFactories",
        inputs: [],
        outputs: [
            {
                name: "",
                type: "address[]",
                internalType: "address[]",
            },
        ],
        stateMutability: "view",
    },
    {
        type: "function",
        name: "marketFactory",
        inputs: [
            {
                name: "index_",
                type: "uint256",
                internalType: "uint256",
            },
        ],
        outputs: [
            {
                name: "",
                type: "address",
                internalType: "address",
            },
        ],
        stateMutability: "view",
    },
    {
        type: "function",
        name: "oracleDecimals",
        inputs: [],
        outputs: [
            {
                name: "",
                type: "uint8",
                internalType: "uint8",
            },
        ],
        stateMutability: "view",
    },
    {
        type: "function",
        name: "owner",
        inputs: [],
        outputs: [
            {
                name: "result",
                type: "address",
                internalType: "address",
            },
        ],
        stateMutability: "view",
    },
    {
        type: "function",
        name: "ownershipHandoverExpiresAt",
        inputs: [
            {
                name: "pendingOwner",
                type: "address",
                internalType: "address",
            },
        ],
        outputs: [
            {
                name: "result",
                type: "uint256",
                internalType: "uint256",
            },
        ],
        stateMutability: "view",
    },
    {
        type: "function",
        name: "protocolFee",
        inputs: [],
        outputs: [
            {
                name: "",
                type: "uint256",
                internalType: "uint256",
            },
        ],
        stateMutability: "view",
    },
    {
        type: "function",
        name: "renounceOwnership",
        inputs: [],
        outputs: [],
        stateMutability: "payable",
    },
    {
        type: "function",
        name: "requestOwnershipHandover",
        inputs: [],
        outputs: [],
        stateMutability: "payable",
    },
    {
        type: "function",
        name: "setLiquidityFee",
        inputs: [
            {
                name: "liquidityFee_",
                type: "uint256",
                internalType: "uint256",
            },
        ],
        outputs: [],
        stateMutability: "nonpayable",
    },
    {
        type: "function",
        name: "setProtocolFee",
        inputs: [
            {
                name: "protocolFee_",
                type: "uint256",
                internalType: "uint256",
            },
        ],
        outputs: [],
        stateMutability: "nonpayable",
    },
    {
        type: "function",
        name: "totalMarketFactories",
        inputs: [],
        outputs: [
            {
                name: "",
                type: "uint256",
                internalType: "uint256",
            },
        ],
        stateMutability: "view",
    },
    {
        type: "function",
        name: "transferOwnership",
        inputs: [
            {
                name: "newOwner",
                type: "address",
                internalType: "address",
            },
        ],
        outputs: [],
        stateMutability: "payable",
    },
    {
        type: "event",
        name: "LiquidityFeeUpdated",
        inputs: [
            {
                name: "oldFee",
                type: "uint256",
                indexed: false,
                internalType: "uint256",
            },
            {
                name: "newFee",
                type: "uint256",
                indexed: false,
                internalType: "uint256",
            },
        ],
        anonymous: false,
    },
    {
        type: "event",
        name: "MarketFactoryCreated",
        inputs: [
            {
                name: "marketFactory",
                type: "address",
                indexed: true,
                internalType: "address",
            },
            {
                name: "name",
                type: "string",
                indexed: false,
                internalType: "string",
            },
            {
                name: "symbol",
                type: "string",
                indexed: false,
                internalType: "string",
            },
        ],
        anonymous: false,
    },
    {
        type: "event",
        name: "OwnershipHandoverCanceled",
        inputs: [
            {
                name: "pendingOwner",
                type: "address",
                indexed: true,
                internalType: "address",
            },
        ],
        anonymous: false,
    },
    {
        type: "event",
        name: "OwnershipHandoverRequested",
        inputs: [
            {
                name: "pendingOwner",
                type: "address",
                indexed: true,
                internalType: "address",
            },
        ],
        anonymous: false,
    },
    {
        type: "event",
        name: "OwnershipTransferred",
        inputs: [
            {
                name: "oldOwner",
                type: "address",
                indexed: true,
                internalType: "address",
            },
            {
                name: "newOwner",
                type: "address",
                indexed: true,
                internalType: "address",
            },
        ],
        anonymous: false,
    },
    {
        type: "event",
        name: "ProtocolFeeUpdated",
        inputs: [
            {
                name: "oldFee",
                type: "uint256",
                indexed: false,
                internalType: "uint256",
            },
            {
                name: "newFee",
                type: "uint256",
                indexed: false,
                internalType: "uint256",
            },
        ],
        anonymous: false,
    },
    {
        type: "error",
        name: "AlreadyInitialized",
        inputs: [],
    },
    {
        type: "error",
        name: "InvalidName",
        inputs: [],
    },
    {
        type: "error",
        name: "InvalidSymbol",
        inputs: [],
    },
    {
        type: "error",
        name: "MarketFactoryAlreadyExists",
        inputs: [
            {
                name: "marketFactory",
                type: "address",
                internalType: "address",
            },
        ],
    },
    {
        type: "error",
        name: "NewOwnerIsZeroAddress",
        inputs: [],
    },
    {
        type: "error",
        name: "NoHandoverRequest",
        inputs: [],
    },
    {
        type: "error",
        name: "Unauthorized",
        inputs: [],
    },
] as const satisfies Abi;
