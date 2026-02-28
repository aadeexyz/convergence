import type { Abi } from "viem";

export const marketFactoryAbi: Abi = [
    {
        type: "constructor",
        inputs: [
            {
                name: "collateralToken_",
                type: "address",
                internalType: "address",
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
            { name: "name_", type: "string", internalType: "string" },
            {
                name: "symbol_",
                type: "string",
                internalType: "string",
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
        outputs: [{ name: "", type: "address", internalType: "address" }],
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
        name: "decimals",
        inputs: [],
        outputs: [{ name: "", type: "uint8", internalType: "uint8" }],
        stateMutability: "view",
    },
    {
        type: "function",
        name: "getExpectedAuthor",
        inputs: [],
        outputs: [{ name: "", type: "address", internalType: "address" }],
        stateMutability: "view",
    },
    {
        type: "function",
        name: "getExpectedWorkflowId",
        inputs: [],
        outputs: [{ name: "", type: "bytes32", internalType: "bytes32" }],
        stateMutability: "view",
    },
    {
        type: "function",
        name: "getExpectedWorkflowName",
        inputs: [],
        outputs: [{ name: "", type: "bytes10", internalType: "bytes10" }],
        stateMutability: "view",
    },
    {
        type: "function",
        name: "getForwarderAddress",
        inputs: [],
        outputs: [{ name: "", type: "address", internalType: "address" }],
        stateMutability: "view",
    },
    {
        type: "function",
        name: "latestMarket",
        inputs: [],
        outputs: [{ name: "", type: "address", internalType: "address" }],
        stateMutability: "view",
    },
    {
        type: "function",
        name: "markets",
        inputs: [{ name: "", type: "uint256", internalType: "uint256" }],
        outputs: [{ name: "", type: "address", internalType: "address" }],
        stateMutability: "view",
    },
    {
        type: "function",
        name: "name",
        inputs: [],
        outputs: [{ name: "", type: "string", internalType: "string" }],
        stateMutability: "view",
    },
    {
        type: "function",
        name: "onReport",
        inputs: [
            {
                name: "metadata",
                type: "bytes",
                internalType: "bytes",
            },
            { name: "report", type: "bytes", internalType: "bytes" },
        ],
        outputs: [],
        stateMutability: "nonpayable",
    },
    {
        type: "function",
        name: "oracle",
        inputs: [],
        outputs: [{ name: "", type: "address", internalType: "address" }],
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
        name: "setExpectedAuthor",
        inputs: [
            {
                name: "_author",
                type: "address",
                internalType: "address",
            },
        ],
        outputs: [],
        stateMutability: "nonpayable",
    },
    {
        type: "function",
        name: "setExpectedWorkflowId",
        inputs: [{ name: "_id", type: "bytes32", internalType: "bytes32" }],
        outputs: [],
        stateMutability: "nonpayable",
    },
    {
        type: "function",
        name: "setExpectedWorkflowName",
        inputs: [{ name: "_name", type: "string", internalType: "string" }],
        outputs: [],
        stateMutability: "nonpayable",
    },
    {
        type: "function",
        name: "setForwarderAddress",
        inputs: [
            {
                name: "_forwarder",
                type: "address",
                internalType: "address",
            },
        ],
        outputs: [],
        stateMutability: "nonpayable",
    },
    {
        type: "function",
        name: "state",
        inputs: [],
        outputs: [
            {
                name: "",
                type: "uint8",
                internalType: "enum IMarketFactory.State",
            },
        ],
        stateMutability: "view",
    },
    {
        type: "function",
        name: "supportsInterface",
        inputs: [
            {
                name: "interfaceId",
                type: "bytes4",
                internalType: "bytes4",
            },
        ],
        outputs: [{ name: "", type: "bool", internalType: "bool" }],
        stateMutability: "view",
    },
    {
        type: "function",
        name: "symbol",
        inputs: [],
        outputs: [{ name: "", type: "string", internalType: "string" }],
        stateMutability: "view",
    },
    {
        type: "function",
        name: "totalMarkets",
        inputs: [],
        outputs: [{ name: "", type: "uint256", internalType: "uint256" }],
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
        name: "ExpectedAuthorUpdated",
        inputs: [
            {
                name: "previousAuthor",
                type: "address",
                indexed: true,
                internalType: "address",
            },
            {
                name: "newAuthor",
                type: "address",
                indexed: true,
                internalType: "address",
            },
        ],
        anonymous: false,
    },
    {
        type: "event",
        name: "ExpectedWorkflowIdUpdated",
        inputs: [
            {
                name: "previousId",
                type: "bytes32",
                indexed: true,
                internalType: "bytes32",
            },
            {
                name: "newId",
                type: "bytes32",
                indexed: true,
                internalType: "bytes32",
            },
        ],
        anonymous: false,
    },
    {
        type: "event",
        name: "ExpectedWorkflowNameUpdated",
        inputs: [
            {
                name: "previousName",
                type: "bytes10",
                indexed: true,
                internalType: "bytes10",
            },
            {
                name: "newName",
                type: "bytes10",
                indexed: true,
                internalType: "bytes10",
            },
        ],
        anonymous: false,
    },
    {
        type: "event",
        name: "ForwarderAddressUpdated",
        inputs: [
            {
                name: "previousForwarder",
                type: "address",
                indexed: true,
                internalType: "address",
            },
            {
                name: "newForwarder",
                type: "address",
                indexed: true,
                internalType: "address",
            },
        ],
        anonymous: false,
    },
    {
        type: "event",
        name: "MarketCreated",
        inputs: [
            {
                name: "market",
                type: "address",
                indexed: true,
                internalType: "address",
            },
        ],
        anonymous: false,
    },
    {
        type: "event",
        name: "MarketSeeded",
        inputs: [
            {
                name: "market",
                type: "address",
                indexed: true,
                internalType: "address",
            },
            {
                name: "longCollateral",
                type: "uint256",
                indexed: false,
                internalType: "uint256",
            },
            {
                name: "shortCollateral",
                type: "uint256",
                indexed: false,
                internalType: "uint256",
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
        name: "SecurityWarning",
        inputs: [
            {
                name: "message",
                type: "string",
                indexed: false,
                internalType: "string",
            },
        ],
        anonymous: false,
    },
    {
        type: "event",
        name: "StateUpdated",
        inputs: [
            {
                name: "oldState",
                type: "uint8",
                indexed: false,
                internalType: "enum IMarketFactory.State",
            },
            {
                name: "newState",
                type: "uint8",
                indexed: false,
                internalType: "enum IMarketFactory.State",
            },
        ],
        anonymous: false,
    },
    { type: "error", name: "AlreadyInitialized", inputs: [] },
    {
        type: "error",
        name: "InvalidAuthor",
        inputs: [
            {
                name: "received",
                type: "address",
                internalType: "address",
            },
            {
                name: "expected",
                type: "address",
                internalType: "address",
            },
        ],
    },
    { type: "error", name: "InvalidForwarderAddress", inputs: [] },
    {
        type: "error",
        name: "InvalidSender",
        inputs: [
            {
                name: "sender",
                type: "address",
                internalType: "address",
            },
            {
                name: "expected",
                type: "address",
                internalType: "address",
            },
        ],
    },
    {
        type: "error",
        name: "InvalidWorkflowId",
        inputs: [
            {
                name: "received",
                type: "bytes32",
                internalType: "bytes32",
            },
            {
                name: "expected",
                type: "bytes32",
                internalType: "bytes32",
            },
        ],
    },
    {
        type: "error",
        name: "InvalidWorkflowName",
        inputs: [
            {
                name: "received",
                type: "bytes10",
                internalType: "bytes10",
            },
            {
                name: "expected",
                type: "bytes10",
                internalType: "bytes10",
            },
        ],
    },
    { type: "error", name: "NewOwnerIsZeroAddress", inputs: [] },
    { type: "error", name: "NoHandoverRequest", inputs: [] },
    { type: "error", name: "Unauthorized", inputs: [] },
    {
        type: "error",
        name: "WorkflowNameRequiresAuthorValidation",
        inputs: [],
    },
] as const satisfies Abi;
