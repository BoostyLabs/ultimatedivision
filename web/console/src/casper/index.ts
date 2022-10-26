// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import { CasperClient } from '@/api/casper';
import { CLPublicKey, DeployUtil } from 'casper-js-sdk';
import { CasperService } from './service';

const casperClient = new CasperClient();
const casperService = new CasperService(casperClient);

/** Desctibes parameters for transaction */
export class CasperTransactionIdentificators {
    /** Includes wallet address, and card id */
    constructor(
        public casperWallet: string,
        public cardId: string
    ) { }
}

/** CasperTransactionService describes velas transaction entity. */
class CasperTransactionService {
    private readonly to = 'a';
    private readonly amount = 10;
    private readonly transferId = 287821;
    private readonly paymentAmount = 100000000;
    private readonly gasPrice = 1;
    private readonly ttl = 1800000;
    private readonly service = casperService;
    public walletAddress: string = '';

    /** default VelasTransactionService implementation */
    constructor(walletAddress: string) {
        this.walletAddress = walletAddress;
    }

    /** Gets transaction from api */
    async getTransaction(signature: CasperTransactionIdentificators): Promise<any> {
        return await this.service.getTransaction(signature);
    }

    async sendTransaction(cardId:string) {
        const walletAddressConverted = CLPublicKey.fromHex(this.walletAddress);

        const address = await this.getTransaction(new CasperTransactionIdentificators(this.walletAddress, cardId));

        const deployParams = new DeployUtil.DeployParams(walletAddressConverted, 'casper-test', this.gasPrice, this.ttl);

        // We create a public key from account-address (it is the hex representation of the public-key with an added prefix).
        const toPublicKey = CLPublicKey.fromHex(address.nftCreateContract.address);

        const session = DeployUtil.ExecutableDeployItem.newTransfer(this.amount, toPublicKey, null, this.transferId);

        const payment = DeployUtil.standardPayment(this.paymentAmount);
        const deploy = DeployUtil.makeDeploy(deployParams, session, payment);

        // Turn your transaction data to format JSON
        const json = DeployUtil.deployToJson(deploy);


        // Sign transcation using casper-signer.
        const signature = await window.casperlabsHelper.sign(json, this.walletAddress, this.to);
        const deployObject = DeployUtil.deployFromJson(signature);
    }
}

export default CasperTransactionService;
