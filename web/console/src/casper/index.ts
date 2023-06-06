// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import { Buffer } from 'buffer';
import { JsonTypes } from 'typedjson';
import { CLPublicKey, CLValueBuilder, DeployUtil, RuntimeArgs } from 'casper-js-sdk';

import { CasperNetworkClient } from '@/api/casper';
import { CasperMatchTransaction } from '@/matches';
import { ToastNotifications } from '@/notifications/service';
import { BidsMakeOfferTransaction, MarketCreateLotTransaction } from '@/marketplace';
import { ACCOUNT_HASH_PREFIX, CHAIN_NAME, CasperSeasonRewardTransaction, CasperTransactionIdentificators, GAS_PRICE, LOT_PAYMENT_AMOUNT, PAYMENT_AMOUNT, TOKEN_PAYMENT_AMOUNT, TTL } from '@/casper/types';

enum CasperRuntimeArgs {
    RECIPIENT = 'recipient',
    TOKEN_ID = 'token_id',
    NFT_CONTRACT_HASH = 'nft_contract_hash',
    MIN_BID_PRICE = 'min_bid_price',
    REDEMPRION_PRICE = 'redemption_price',
    AUCTION_DURATION = 'auction_duration',
    VALUE = 'value',
    NONCE = 'nonce',
    SIGNATURE = 'signature',
    SPENDER = 'spender',
    OFFER_PRICE = 'offer_price',
    ERC20_CONTRACT = 'erc20_contract',
    AMOUNT = 'amount'
}

/** CasperTransactionService describes casper transaction entity. */
class CasperTransactionService {
    private readonly paymentAmount: number = PAYMENT_AMOUNT;
    private readonly lotPaymentAmount: number = LOT_PAYMENT_AMOUNT;
    private readonly gasPrice: number = GAS_PRICE;
    private readonly ttl: number = TTL;
    private readonly client: CasperNetworkClient = new CasperNetworkClient();
    public walletAddress: string = '';

    /** default CasperTransactionService implementation */
    constructor(walletAddress: string) {
        this.walletAddress = walletAddress;
    }

    /** Gets minting signature with contract address from api */
    async getTransaction(signature: CasperTransactionIdentificators): Promise<any> {
        return await this.client.getTransaction(signature);
    }

    /** Converts contract hash to bytes  */
    public static async convertContractHashToBytes(contractHash: string): Promise<Uint8Array> {
        return await Uint8Array.from(Buffer.from(contractHash, 'hex'));
    }

    /** Converts account hash to string without prefix 'hash-'  */
    public convertAccountHash(walletAddress: string): string {
        const accountHash = CLPublicKey.fromHex(walletAddress).toAccountHashStr();
        const accountHashConverted = accountHash.replace(ACCOUNT_HASH_PREFIX, '');

        return accountHashConverted;
    }

    /** Signs a contract */
    public async contractSign(
        entryPoint: string,
        runtimeArgs: RuntimeArgs,
        paymentAmount: number,
        contractAddress: string
    ): Promise<JsonTypes> {
        const contractHashToBytes = await CasperTransactionService.convertContractHashToBytes(contractAddress);

        const walletAddressConverted = CLPublicKey.fromHex(this.walletAddress);

        const deployParams = new DeployUtil.DeployParams(walletAddressConverted, CHAIN_NAME, this.gasPrice, this.ttl);

        const deploy = DeployUtil.makeDeploy(
            deployParams,
            DeployUtil.ExecutableDeployItem.newStoredContractByHash(
                contractHashToBytes,
                entryPoint,
                runtimeArgs),
            DeployUtil.standardPayment(paymentAmount)
        );

        const deployJson = DeployUtil.deployToJson(deploy);

        const signature = await window.casperlabsHelper.sign(deployJson, this.walletAddress, contractAddress);

        return signature;
    }

    /** Mints a nft */
    async mint(cardId: string): Promise<void> {
        try {
            const accountHash = CLPublicKey.fromHex(this.walletAddress).toAccountHashStr();
            const accountHashConverted = accountHash.replace(ACCOUNT_HASH_PREFIX, '');

            const nftWaitlist = await this.getTransaction(new CasperTransactionIdentificators(accountHashConverted, cardId));

            const runtimeArgs = RuntimeArgs.fromMap({
                [CasperRuntimeArgs.TOKEN_ID]: CLValueBuilder.string(nftWaitlist.tokenId),
                [CasperRuntimeArgs.RECIPIENT]: CLValueBuilder.string(`account-hash-${accountHashConverted}`),
            });

            const isConnected = window.casperlabsHelper.isConnected();

            if (!isConnected) {
                await window.casperlabsHelper.requestConnection();
            }

            const signature = await this.contractSign('mint_one', runtimeArgs, this.paymentAmount, '05560ca94e73f35c5b9b8a0f8b66e56238169e60ae421fb7b71c7ac3c6c744e2');

            await this.client.claim(nftWaitlist.rpcNodeAddress, JSON.stringify(signature));
        }
        catch (error: any) {
            ToastNotifications.casperError(`${error.error}`);
        }
    }

    /** Mints a token */
    async mintUDT(transaction: CasperMatchTransaction | CasperSeasonRewardTransaction, rpcNodeAddress: string): Promise<void> {
        try {
            const runtimeArgs = RuntimeArgs.fromMap({
                [CasperRuntimeArgs.VALUE]: CLValueBuilder.u256(transaction.value),
                [CasperRuntimeArgs.NONCE]: CLValueBuilder.u64(7),
                [CasperRuntimeArgs.SIGNATURE]: CLValueBuilder.string(transaction.signature),
            });

            const isConnected = window.casperlabsHelper.isConnected();

            if (!isConnected) {
                await window.casperlabsHelper.requestConnection();
            }

            const signature = await this.contractSign('claim', runtimeArgs, TOKEN_PAYMENT_AMOUNT, transaction.casperTokenContract.address);

            await this.client.claim(rpcNodeAddress, JSON.stringify(signature), this.walletAddress);
        }
        catch (error: any) {
            ToastNotifications.casperError(`${error.error}`);
        }
    }

    /** Mints a token */
    async approveNftMinting(): Promise<void> {
        try {
            const runtimeArgs = RuntimeArgs.fromMap({
                [CasperRuntimeArgs.TOKEN_ID]: CLValueBuilder.string('7d27e29d-3f30-425d-8abc-5e9195243132'),
                [CasperRuntimeArgs.SPENDER]: CLValueBuilder.string('contract-package-wasm701ed1a382367a6016f3b389f75177030fd583c5b8838b4c04e92da6b4a11928'),
            });

            const isConnected = window.casperlabsHelper.isConnected();

            if (!isConnected) {
                await window.casperlabsHelper.requestConnection();
            }

            const a = 2500000000;
            const signature = await this.contractSign('approve', runtimeArgs, a, '05560ca94e73f35c5b9b8a0f8b66e56238169e60ae421fb7b71c7ac3c6c744e2');

            await this.client.claim('http://116.202.169.210:7777/rpc', JSON.stringify(signature));
        }
        catch (error: any) {
            ToastNotifications.casperError(`${error.error}`);
        }
    }

    /** Mints a token */
    async approveTokenRevard(transaction: any): Promise<void> {
        try {
            const accountHash = CLPublicKey.fromHex(this.walletAddress);

            const runtimeArgs = RuntimeArgs.fromMap({
                [CasperRuntimeArgs.SPENDER]: CLValueBuilder.key(accountHash),
                [CasperRuntimeArgs.AMOUNT]: CLValueBuilder.u256(transaction.account),
            });

            const isConnected = window.casperlabsHelper.isConnected();

            if (!isConnected) {
                await window.casperlabsHelper.requestConnection();
            }

            const signature = await this.contractSign('approve', runtimeArgs, TOKEN_PAYMENT_AMOUNT, transaction.contractAddress);

            await this.client.claim(transaction.rpcNodeAddress, JSON.stringify(signature), this.walletAddress);
        }
        catch (error: any) {
            ToastNotifications.casperError(`${error.error}`);
        }
    }


    /** Creates a lot */
    async createLot(transaction: MarketCreateLotTransaction): Promise<void> {
        try {
            const runtimeArgs = RuntimeArgs.fromMap({
                [CasperRuntimeArgs.NFT_CONTRACT_HASH]: CLValueBuilder.string('contract-05560ca94e73f35c5b9b8a0f8b66e56238169e60ae421fb7b71c7ac3c6c744e2'),
                [CasperRuntimeArgs.TOKEN_ID]: CLValueBuilder.string('7d27e29d-3f30-425d-8abc-5e9195243132'),
                [CasperRuntimeArgs.MIN_BID_PRICE]: CLValueBuilder.u256(3000),
                [CasperRuntimeArgs.REDEMPRION_PRICE]: CLValueBuilder.u256(30000),
                [CasperRuntimeArgs.AUCTION_DURATION]: CLValueBuilder.u256(300000),
            });

            const isConnected = window.casperlabsHelper.isConnected();

            if (!isConnected) {
                await window.casperlabsHelper.requestConnection();
            }

            const signature = await this.contractSign('create_listing', runtimeArgs, this.lotPaymentAmount, 'feed638f60f5a2840656d86e0e51dc62c092e79d980ba8dc281387dbb8f80c42');

            await this.client.claim('http://116.202.169.210:7777/rpc', JSON.stringify(signature));
        }
        catch (error: any) {
            ToastNotifications.casperError(`${error.error}`);
        }
    }

    /** Accepts offer */
    async acceptOffer(transaction: any): Promise<void> {
        try {
            const runtimeArgs = RuntimeArgs.fromMap({
                [CasperRuntimeArgs.NFT_CONTRACT_HASH]: CLValueBuilder.string('hash-4ff1e5e37b8720e8049bfff88676d8e27c1037c02e1172a1006c6d2a535607da'),
                [CasperRuntimeArgs.TOKEN_ID]: CLValueBuilder.string('746c85ba-583c-4c45-9af7-dce858c4e121'),
            });

            const isConnected = window.casperlabsHelper.isConnected();

            if (!isConnected) {
                await window.casperlabsHelper.requestConnection();
            }

            const signature = await this.contractSign('accept_offer', runtimeArgs, this.lotPaymentAmount, transaction.address);

            await this.client.claim(transaction.rpcNodeAddress, JSON.stringify(signature), this.walletAddress);
        }
        catch (error: any) {
            ToastNotifications.casperError(`${error.error}`);
        }
    }

    /** Makes offer */
    async makeOffer(transaction: BidsMakeOfferTransaction): Promise<void> {
        try {
            const runtimeArgs = RuntimeArgs.fromMap({
                [CasperRuntimeArgs.NFT_CONTRACT_HASH]: CLValueBuilder.string('hash-05560ca94e73f35c5b9b8a0f8b66e56238169e60ae421fb7b71c7ac3c6c744e2'),
                [CasperRuntimeArgs.TOKEN_ID]: CLValueBuilder.string('7d27e29d-3f30-425d-8abc-5e9195243132'),
                [CasperRuntimeArgs.OFFER_PRICE]: CLValueBuilder.u256(3500),
            });

            const isConnected = window.casperlabsHelper.isConnected();

            if (!isConnected) {
                await window.casperlabsHelper.requestConnection();
            }

            const signature = await this.contractSign('make_offer', runtimeArgs, this.lotPaymentAmount, 'feed638f60f5a2840656d86e0e51dc62c092e79d980ba8dc281387dbb8f80c42');

            await this.client.claim('http://116.202.169.210:7777/rpc', JSON.stringify(signature), this.walletAddress);
        }
        catch (error: any) {
            ToastNotifications.casperError(`${error.error}`);
        }
    }

    /** Buys listing */
    async buyListing(transaction: any): Promise<void> {
        try {
            const runtimeArgs = RuntimeArgs.fromMap({
                [CasperRuntimeArgs.NFT_CONTRACT_HASH]: CLValueBuilder.string(transaction.contractHash),
                [CasperRuntimeArgs.TOKEN_ID]: CLValueBuilder.string(transaction.tokenId),
            });

            const isConnected = window.casperlabsHelper.isConnected();

            if (!isConnected) {
                await window.casperlabsHelper.requestConnection();
            }

            const signature = await this.contractSign('buy_listing', runtimeArgs, this.lotPaymentAmount, transaction.address);

            await this.client.claim(transaction.rpcNodeAddress, JSON.stringify(signature), this.walletAddress);
        }
        catch (error: any) {
            ToastNotifications.casperError(`${error.error}`);
        }
    }

    /** finalListing */
    async finalListing(transaction: any): Promise<void> {
        try {
            const runtimeArgs = RuntimeArgs.fromMap({
                [CasperRuntimeArgs.NFT_CONTRACT_HASH]: CLValueBuilder.string(transaction.contractHash),
                [CasperRuntimeArgs.TOKEN_ID]: CLValueBuilder.string(transaction.tokenId),
            });

            const isConnected = window.casperlabsHelper.isConnected();

            if (!isConnected) {
                await window.casperlabsHelper.requestConnection();
            }

            const signature = await this.contractSign('final_listing', runtimeArgs, this.lotPaymentAmount, transaction.address);

            await this.client.claim(transaction.rpcNodeAddress, JSON.stringify(signature), this.walletAddress);
        }
        catch (error: any) {
            ToastNotifications.casperError(`${error.error}`);
        }
    }
}

export default CasperTransactionService;
