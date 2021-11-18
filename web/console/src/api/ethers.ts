// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '.';

export class EthersClient extends APIClient {
    private readonly ROOT_PATH = '/api/v0/nft-waitlist'

    public async getAddress(wallet: string, id: string): Promise<{
        "password": string,
        "tokenId": number,
        "contract": {
            "address": string,
            "addressMethod": string
        }
    }> {
        const response = await this.http.post(`${this.ROOT_PATH}`, JSON.stringify({
            'cardId': id,
            'walletAddress': wallet
        }));

        if (!response.ok) {
            await this.handleError(response);
        }

        return await response.json();
    }
}