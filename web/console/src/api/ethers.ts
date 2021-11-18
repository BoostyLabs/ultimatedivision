// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '.';
import { Transaction } from '@/app/ethers';

/** Ethers api client */
export class EthersClient extends APIClient {
    private readonly ROOT_PATH = '/api/v0/nft-waitlist';

    /** Gets transaction from api  */
    public async getTransaction(wallet: string, id: string): Promise<Transaction> {
        const response = await this.http.post(`${this.ROOT_PATH}`, JSON.stringify({
            'cardId': id,
            'walletAddress': wallet,
        }));

        if (!response.ok) {
            await this.handleError(response);
        }

        return await response.json();
    }
}
