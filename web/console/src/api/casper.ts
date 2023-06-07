// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import { Transaction } from '@/ethers';
import { CasperTransactionApprove, CasperTransactionIdentificators } from '@/casper/types';
import { APIClient } from '.';

/**
 * CasperNetworkClient is a http implementation of casper-wallet API.
 * Exposes all casper wallet related functionality.
 */
export class CasperNetworkClient extends APIClient {
    private readonly ROOT_PATH: string = '/api/v0';

    /** Sends signed message and registers user */
    public async register(walletAddress: string, accountHash: string): Promise<void> {
        const response = await this.http.post(`${this.ROOT_PATH}/auth/casper/register`, JSON.stringify({ casperWallet: walletAddress, casperWalletHash: accountHash }));

        if (!response.ok) {
            await this.handleError(response);
        }
    }

    /** Gets message from API for sign with casper */
    public async nonce(walletAddress: string): Promise<string> {
        const path = `${this.ROOT_PATH}/auth/casper/nonce?address=${walletAddress}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        }

        return await response.json();
    }

    /** Sends signed message, and logs-in */
    public async login(nonce: string, signature: string): Promise<void> {
        const response = await this.http.post(`${this.ROOT_PATH}/auth/casper/login`, JSON.stringify({ nonce, signature }));

        if (!response.ok) {
            await this.handleError(response);
        }
    }
    /** Gets minting signature with contract address from api */
    public async getTransaction(signature: CasperTransactionIdentificators): Promise<Transaction> {
        const response = await this.http.post(`${this.ROOT_PATH}/nft-waitlist`, JSON.stringify(signature));

        if (!response.ok) {
            await this.handleError(response);
        }
        const transaction = await response.json();

        return new Transaction(
            transaction.password,
            transaction.tokenId,
            transaction.nftCreateContract,
            transaction.nftCreateCasperContract,
            transaction.rpcNodeAddress
        );
    }
    /** Sends deploy data to api */
    public async sendTx(RPCNodeAddress: string, deploy: string, casperWallet?: string): Promise<void> {
        let response;

        if (casperWallet) {
            response = await this.http.post(`${this.ROOT_PATH}/casper/send-tx`, JSON.stringify({ RPCNodeAddress, deploy, casperWallet }));
        } else {
            response = await this.http.post(`${this.ROOT_PATH}/casper/send-tx`, JSON.stringify({ RPCNodeAddress, deploy }));
        }

        if (!response.ok) {
            throw await response.json();
        }
    }


    /** approve transaction casper */
    public async approve(cardId?: string): Promise<CasperTransactionApprove> {
        let path;

        if (cardId) {
            path = `/api/v0/casper-approve?card_id=${cardId}`;
        }
        else {
            path = '/api/v0/casper-approve?card_id=';
        }

        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        };

        const approveData = await response.json();


        return new CasperTransactionApprove(
            approveData.addressNodeServer,
            approveData.amount,
            approveData.nftContractAddress,
            approveData.tokenRewardContractAddress,
            approveData.tokenId,
            approveData.tokenContractPackageAddress,
            approveData.nftContractPackageAddress,
        );
    };
}

