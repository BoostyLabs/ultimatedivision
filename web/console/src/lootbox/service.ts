// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '@/api/index';
import { LootboxClient } from '@/api/lootboxes';
import { Card } from '@/card';
import { Lootbox } from '.';

/** exposes all lootbox related logic */
export class LootboxService {
    private readonly lootboxes: LootboxClient;
    /** receives LootboxClient */
    public constructor(lootboxes: LootboxClient) {
        this.lootboxes = lootboxes;
    };

    /** handles lootbox buying */
    public async buy(lootbox: Lootbox): Promise<Card[]> {
        const responseCards = await this.lootboxes.buy(lootbox);

        return await responseCards.json();
    };
};

