// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { VelasClient } from '../../api/velas';
import { VelasData } from '@/app/types/velas';
import { TransactionIdentificators } from '../ethers';

/**
 * Exposes all velas wallet related logic.
 */
export class VelasService {
    private readonly velasWallet: VelasClient;

    /** velasService contains http implementation of velas API  */
    public constructor(velasWallet: VelasClient) {
        this.velasWallet = velasWallet;
    }

    /** Gets transaction from api */
    public async getTransaction(signature: TransactionIdentificators): Promise<any> {
        await this.velasWallet.getTransaction(signature);
    }

    /** sends data to register user with velas wallet */
    public async register(walletAddress: string, accessToken: string, expiresAt: any, velasData: string): Promise<void> {
        await this.velasWallet.register(walletAddress, accessToken, expiresAt, velasData);
    }

    /** sends address to get nonce to login user */
    public async nonce(address: string): Promise<string> {
        return await this.velasWallet.nonce(address);
    }

    /** sends data to login user with velas wallet */
    public async login(nonce: string, walletAddress: string, accessToken: string, expiresAt: any): Promise<void> {
        await this.velasWallet.login(nonce, walletAddress, accessToken, expiresAt);
    }

    /** gets token to login user with velas wallet */
    public async csrfToken(): Promise<string> {
        return await this.velasWallet.csrfToken();
    }

    /** gets creds to fill velas vaclient */
    public async vaclientCreds(): Promise<any> {
        return await this.velasWallet.vaclientCreds();
    }

    /** gets creds to fill velas vaclient */
    public async vaclientData(userId:string): Promise<VelasData> {
        return await this.velasWallet.vaclientData(userId);
    }
}
