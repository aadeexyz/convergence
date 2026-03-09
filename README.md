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

The system has an initial creation flow and a recurring update flow.

### Initial creation

1. A user calls `FactoryOfMarketFactories.createMarketFactory(name, symbol)`.
2. The factory-of-factories deploys a new `MarketFactory` clone for that keyword and transfers the creation liquidity into it.
3. The new `MarketFactory` initializes its own `Oracle` clone and enters a pre-live state, waiting for its first workflow report.
4. The CRE workflow detects the new factory, fetches external attention signals for the keyword, computes the Attention Index and EMA, and submits that report onchain.
5. The `MarketFactory` receives the report, stores the new oracle round, creates the first `Market`, and seeds its long and short collateral pools using the factory’s deposited liquidity.
6. The market is now live and users can mint or burn long and short position tokens against the active pool.

### Recurring updates

1. On each scheduled workflow run, CRE reads the latest onchain state for a factory and fetches fresh offchain attention data for the same keyword.
2. The workflow recomputes the latest Attention Index and EMA, then generates and submits a new signed report onchain.
3. When `MarketFactory` processes that report, it first updates the `Oracle` with the new round.
4. If a previous market is active, the factory settles it against the new oracle round and redeems the remaining long and short collateral back into the factory.
5. The factory then creates a fresh market and reseeds it using the recycled collateral balance, split between long and short according to the new index value.
6. Users continue trading on the new active market until the next workflow update rolls the system forward again.

## Repo Layout

```text
├── contracts/   Foundry project
├── frontend/    Next.js 16
├── cre/         Go-based Chainlink CRE workflow
├── deployment.txt
├── LICENSE
└── README.md
```

## Safety

This software is experimental and unaudited, and is provided on an 'as is' and 'as available' basis. We do not give any warranties and will not be liable for any loss incurred through any use of this codebase.
