// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '@/api/index';
import { LootboxClient } from '@/api/lootboxes';
import { Card } from '@/card';
import { Lootbox } from '.';

/** exposes all lootbox related logic */
export class LootboxService extends APIClient {
    private readonly lootboxes: LootboxClient;
    /** receives LootboxClient */
    public constructor(lootboxes: LootboxClient) {
        super();
        this.lootboxes = lootboxes;
    };

    /** handles lootbox buying */
    public async buy(lootbox: Lootbox): Promise<Card[]> {
        const responseCards = await this.lootboxes.buy(lootbox);

        if (!responseCards.ok) {
            this.handleError(responseCards);
        }

        const cards = await responseCards.json();

        if(!cards) {
            this.handleError(cards);
        }

        return cards;
    };
};
