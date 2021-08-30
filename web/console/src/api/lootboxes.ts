// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '@/api/index';
import { Lootbox } from '@/lootbox';

/** LootboxClient is a lootbox api client */
export class LootboxClient extends APIClient {
    private readonly ROOT_PATH: string = '/lootboxes';
    /** buys lootbox */
    public async buy(lootbox: Lootbox) {
        return await this.http.post(this.ROOT_PATH, JSON.stringify(lootbox));
    };
    /** opens lootbox*/
    public async open(lootbox: Lootbox) {
        return await this.http.delete(this.ROOT_PATH, JSON.stringify(lootbox));
    };
};
