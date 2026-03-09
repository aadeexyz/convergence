# World Attention Markets

Trade internet attention itself through onchain long/short markets settled by a live index.

World Attention Markets is a topic-based trading system for internet attention. Instead of buying a proxy asset and hoping it reflects a trend, users can trade a market tied directly to a live Attention Index for a keyword or narrative. The index is computed offchain from external data sources, smoothed for stability, and committed onchain through Chainlink CRE for deterministic settlement.

This repo contains:

- `contracts/`: Foundry smart contracts for market creation, trading, settlement, and oracle storage
- `frontend/`: Next.js app for creating markets, trading positions, and using the faucet
- `cre/`: Chainlink CRE workflow that computes the attention index from external data and submits reports onchain

## Overview

The system has two layers:

- CRE pipeline: fetches attention data from Google Trends, Twitter, and YouTube, combines it into a weighted score, smooths it with an EMA, and normalizes it into a bounded index.
- Onchain market layer: creates topic-specific markets, supports long and short exposure, and settles against the official index values written onchain by the CRE workflow.

Chainlink CRE is the orchestration layer between those two parts. It runs the workflow that reads external data, computes the latest Attention Index, and writes settlement data onchain.

### Files Using CRE

- [`cre/wam-workflow/main.go`](cre/wam-workflow/main.go)
- [`cre/wam-workflow/aggregator.go`](cre/wam-workflow/aggregator.go)
- [`cre/wam-workflow/fof/list.go`](cre/wam-workflow/fof/list.go)
- [`cre/wam-workflow/lens/read.go`](cre/wam-workflow/lens/read.go)
- [`cre/wam-workflow/mf/balance.go`](cre/wam-workflow/mf/balance.go)
- [`cre/wam-workflow/mf/read.go`](cre/wam-workflow/mf/read.go)
- [`cre/wam-workflow/mf/write.go`](cre/wam-workflow/mf/write.go)
- [`cre/wam-workflow/sourcer/google_trends.go`](cre/wam-workflow/sourcer/google_trends.go)
- [`cre/wam-workflow/sourcer/twitter.go`](cre/wam-workflow/sourcer/twitter.go)
- [`cre/wam-workflow/sourcer/youtube.go`](cre/wam-workflow/sourcer/youtube.go)

### Contracts Consuming CRE Workflow

- [`contracts/src/keystone/ReceiverTemplate.sol`](contracts/src/keystone/ReceiverTemplate.sol)
- [`contracts/src/MarketFactory.sol`](contracts/src/MarketFactory.sol)
- [`contracts/src/Oracle.sol`](contracts/src/Oracle.sol)

## Architecture

### CRE

- `cre/wam-workflow/main.go`: workflow entrypoint
- `cre/wam-workflow/aggregator.go`: weighted attention score aggregation
- `cre/wam-workflow/sourcer/google_trends.go`: Google Trends ingestion
- `cre/wam-workflow/sourcer/twitter.go`: Twitter ingestion
- `cre/wam-workflow/sourcer/youtube.go`: YouTube ingestion
- `cre/wam-workflow/mf/write.go`: report generation and onchain writes through CRE

### Onchain

- `FactoryOfMarketFactories`: deploys topic-specific `MarketFactory` clones and collects creation fees
- `MarketFactory`: owns one oracle, manages market rollover, receives CRE reports, settles the previous market, and seeds the next one
- `Market`: holds collateral, tracks long/short pools, and handles mint, burn, and redeem flows
- `Oracle`: stores index rounds and a rolling EMA window
- `Lens`: batches read operations for the frontend and workflow
- `Router`: shared app write entrypoint

## Market Lifecycle

1. A user creates a `MarketFactory` for a keyword.
2. The factory deploys its own `Oracle`.
3. The CRE workflow fetches external attention signals for that keyword.
4. The workflow computes the current Attention Index and EMA.
5. CRE writes the report onchain.
6. The `MarketFactory` settles the previous market and seeds the next one.
7. Users trade long and short position tokens against the active collateral pool.

## Repo Layout

```text
.
├── contracts/   Foundry project
├── frontend/    Next.js 16
├── cre/         Go-based Chainlink CRE workflow
├── deployment.txt
└── README.md
```

## Safety

This software is experimental and unaudited, and is provided on an 'as is' and 'as available' basis. We do not give any warranties and will not be liable for any loss incurred through any use of this codebase.