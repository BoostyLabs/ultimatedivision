// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '@/api/index';

export class MarketplaceClient extends APIClient {
    private readonly ROOT_PATH: string = '/api/v0/cards';

    public async getCards() {
        return this.http.get(this.ROOT_PATH);
    }
}


