// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import { Buffer } from 'buffer';
import { toast } from 'react-toastify';

import { CLValueBuilder, RuntimeArgs, CLPublicKey, DeployUtil } from 'casper-js-sdk';

import { CasperNetworkClient } from '@/api/casper';

/** Desctibes parameters for transaction */
export class CasperTransactionIdentificators {
    /** Includes wallet address, and card id */
    constructor(
        public casperWallet: string,
        public cardId: string
    ) { }
}

const ACCOUNT_HASH_PREFIX = 'account-hash-';

const TTL = 1800000;
const PAYMENT_AMOUNT = 4500000000;
const GAS_PRICE = 1;

/** CasperTransactionService describes casper transaction entity. */
class CasperTransactionService {
    private readonly paymentAmount: number = PAYMENT_AMOUNT;
    private readonly gasPrice: number = GAS_PRICE;
    private readonly ttl: number = TTL;
    private readonly client: any = new CasperNetworkClient();
    public walletAddress: string = '';

    /** default VelasTransactionService implementation */
    constructor(walletAddress: string) {
        this.walletAddress = walletAddress;
    }

    /** Gets transaction from api */
    async getTransaction(signature: CasperTransactionIdentificators): Promise<any> {
        return await this.client.getTransaction(signature);
    }

    /** Converts contract hash to bytes  */
    public static async convertContractHashToBytes(contractHash: string): Promise<Uint8Array> {
        return await Uint8Array.from(Buffer.from(contractHash, 'hex'));
    }

    /** Signs a contract */
    public async contractSign(
        entryPoint:any,
        runtimeArgs:any,
        paymentAmount: any,
        contractAddress: string
    ): Promise<any> {
        const contractHashToBytes = await CasperTransactionService.convertContractHashToBytes(contractAddress);

        try {
            const walletAddressConverted = CLPublicKey.fromHex(this.walletAddress);

            const deployParams = new DeployUtil.DeployParams(walletAddressConverted, 'casper-test', this.gasPrice, this.ttl);

            const deploy = DeployUtil.makeDeploy(
                deployParams,
                DeployUtil.ExecutableDeployItem.newStoredContractByHash(
                    contractHashToBytes,
                    entryPoint,
                    runtimeArgs),
                DeployUtil.standardPayment(paymentAmount)
            );

            const json = DeployUtil.deployToJson(deploy);

            const signature = await window.casperlabsHelper.sign(json, this.walletAddress, contractAddress);

            return signature;
        }
        catch (e) {
            toast.error('Something went wrong', {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
        }

        return false;
    }

    /** Mints a nft */
    async mint(cardId: string): Promise<void> {
        try {
            const accountHash = CLPublicKey.fromHex(this.walletAddress).toAccountHashStr();
            const accountHashConverted = accountHash.replace(ACCOUNT_HASH_PREFIX, '');

            const nftWaitlist = await this.getTransaction(new CasperTransactionIdentificators(accountHashConverted, cardId));

            const isConected = window.casperlabsHelper.isConnected();

            if (!isConected) {
                window.casperlabsHelper.requestConnection();
            }

            const runtimeArgs = RuntimeArgs.fromMap({
                signature: CLValueBuilder.string(nftWaitlist.password),
                /* eslint-disable */
                token_id: CLValueBuilder.u64(nftWaitlist.tokenId),
            });

            const signature = await this.contractSign('claim', runtimeArgs, this.paymentAmount, nftWaitlist.nftCreateCasperContract.address);

            this.client.claim(nftWaitlist.rpcNodeAddress, JSON.stringify(signature));
        }
        catch (e) {
            toast.error('Something went wrong', {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
        }
    }
}

export default CasperTransactionService;
