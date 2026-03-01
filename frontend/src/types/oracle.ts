export type OracleRound = {
    id: bigint;
    timestamp: bigint;
    index: bigint;
};

export type OracleView = {
    currentRoundId: bigint;
    latestRound: OracleRound;
    decimals: number;
    keyword: string;
    rollingEMAWindow: bigint[];
};
