// Copyright (C) 2023 Creditor Corp. Group.
// See LICENSE for copying information.

export const ACCOUNT_HASH_PREFIX = 'account-hash-';
export const CHAIN_NAME = 'casper-test';

export const TTL = 1800000;
export const GAS_PRICE = 1;

export const PAYMENT_AMOUNT = 50000000000;
export const TOKEN_PAYMENT_AMOUNT = 6000000000;
export const LOT_PAYMENT_AMOUNT = 40000000000;

/** Describes parameters for transaction */
export class CasperTransactionIdentificators {
    /** Includes wallet address, and card id */
    constructor(
        public casperWallet: string,
        public cardId: string
    ) { }
}

/** Describes parameters for casper token transaction */
export class CasperTokenContract {
    /** default CasperTokenContract implementation */
    constructor(
        public address: string = '0',
        public addressMethod: string = ''
    ) { }
}

/** Transaction describes transaction entity of match response. */
export class CasperSeasonRewardTransaction {
    /** Transaction contains of nonce, signature hash udtContract and value. */
    constructor(
        public ID: string,
        public userId: string,
        public seasonID: string,
        public walletAddress: string,
        public casperWalletAddress: string,
        public walleType: string,
        public status: number,
        public nonce: number,
        public signature: string,
        public value: string,
        public casperTokenContract: {
            address: string;
            addressMethod: string;
        },
        public rpcNodeAddress: string,
    ) { }
};
