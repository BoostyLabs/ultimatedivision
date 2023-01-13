// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import { VelasData } from '@/app/types/velas';
import { APIClient } from '.';
import { SignedMessage, Transaction, TransactionIdentificators } from '@/ethers';

/** Ethers api client */
export class VelasClient extends APIClient {
    private readonly ROOT_PATH = '/api/v0';

    /** Gets transaction from api  */
    public async getTransaction(signature: TransactionIdentificators): Promise<Transaction> {
        const response = await this.http.post('/api/v0/nft-waitlist', JSON.stringify(signature));

        if (!response.ok) {
            await this.handleError(response);
        }
        const transaction = await response.json();

        return new Transaction(transaction.password, transaction.tokenId, transaction.nftCreateContract);
    }

    /** sends data to register user with velas wallet */
    public async register(walletAddress: string, accessToken: string, expiresAt: any, velasData: string): Promise<void> {
        const path = `${this.ROOT_PATH}/auth/velas/register`;
        const response = await this.http.post(path, JSON.stringify({ walletAddress, accessToken, expiresAt, velasData }));

        if (!response.ok) {
            await this.handleError(response);
        }
    }
    /** sends address to get nonce to login user */
    public async nonce(address: string): Promise<string> {
        const walletType = 'velas_wallet_address';
        const path = `${this.ROOT_PATH}/auth/velas/nonce?address=${address}&walletType=${walletType}`;

        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        }
        const status = response.json();

        return status;
    }
    /** sends data to login user with velas wallet */
    public async login(nonce: string, walletAddress: string, accessToken: string, expiresAt: any): Promise<void> {
        const path = `${this.ROOT_PATH}/auth/velas/login`;

        const response = await this.http.post(path, JSON.stringify({ walletAddress, accessToken, expiresAt, nonce }));

        if (!response.ok) {
            await this.handleError(response);
        }
    }

    /** gets token to login user with velas wallet */
    public async csrfToken(): Promise<string> {
        const path = 'https://api-sponsor.testnet.velas.com/csrf';
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        }
        const result = await response.json();

        return result.token;
    }

    /** gets creds to fill velas vaclient */
    public async vaclientCreds(): Promise<any> {
        const path = `${this.ROOT_PATH}/auth/velas/vaclient`;
        ;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        }
        const result = await response.json();

        return result;
    }

    /** gets creds to fill velas vaclient */
    public async vaclientData(userId: string): Promise<VelasData> {
        const path = `${this.ROOT_PATH}/auth/velas/register-data/${userId}`;

        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        }
        const result = await response.json();

        return result;
    }
}
