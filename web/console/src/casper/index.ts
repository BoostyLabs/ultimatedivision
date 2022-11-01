// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import { CasperNetworkClient } from '@/api/casper';
import { CLValueBuilder, RuntimeArgs, CasperClient, Contracts, CLPublicKey, DeployUtil } from 'casper-js-sdk';

import { Buffer } from 'buffer';
import { toast } from 'react-toastify';

/** Desctibes parameters for transaction */
export class CasperTransactionIdentificators {
    /** Includes wallet address, and card id */
    constructor(
        public casperWallet: string,
        public cardId: string
    ) { }
}

const TTL = 1800000;
const PAYMENT_AMOUNT = 100000000;
const GAS_PRICE = 1;

/** CasperTransactionService describes casper transaction entity. */
class CasperTransactionService {
    private readonly paymentAmount = PAYMENT_AMOUNT;
    private readonly gasPrice = GAS_PRICE;
    private readonly ttl = TTL;
    private readonly client = new CasperNetworkClient();
    public walletAddress: string = '';

    /** default VelasTransactionService implementation */
    constructor(walletAddress: string) {
        this.walletAddress = walletAddress;
    }

    /** Gets transaction from api */
    async getTransaction(signature: CasperTransactionIdentificators): Promise<any> {
        return await this.client.getTransaction(signature);
    }

    async setContractHashToBytes(contractHash: string) {
        return await Uint8Array.from(Buffer.from(contractHash, 'hex'));
    }

    public async contractSign(
        entryPoint:any,
        runtimeArgs:any,
        paymentAmount:any
    ) {
        const contractHashToBytes = await this.setContractHashToBytes('011d2f5eed581e3750fa3d2fd15ef782aa66a55a679346c0a339c485c78fc9fe68');

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

            const signature = await window.casperlabsHelper.sign(json, this.walletAddress, '011d2f5eed581e3750fa3d2fd15ef782aa66a55a679346c0a339c485c78fc9fe68');

            return signature;
        }
        catch (e) {
            toast.error('Something went wrong', {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
        }
    }

    async mint(cardId: string) {
        try {
            const nftWailist = await this.getTransaction(new CasperTransactionIdentificators(this.walletAddress, cardId));

            window.casperlabsHelper.requestConnection();

            const runtimeArgs = RuntimeArgs.fromMap({
                signature: CLValueBuilder.string(nftWailist.password),
                token_id: CLValueBuilder.u64(nftWailist.tokenId),
            });

            const casperClient = new CasperClient('http://3.136.227.9:7777/rpc');

            const signature = await this.contractSign('claim', runtimeArgs, this.paymentAmount);

            const deploy = casperClient.putDeploy(DeployUtil.deployFromJson(signature).unwrap());
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
