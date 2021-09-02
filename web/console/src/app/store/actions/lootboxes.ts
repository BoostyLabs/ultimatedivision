// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch } from 'redux';

import { Lootbox } from '@/lootbox';
import { LootboxClient } from '@/api/lootboxes';
import { LootboxService } from '@/lootbox/servise';

export const BUY_LOOTBOX = 'BUY_LOOTBOX';
export const OPEN_LOOTBOX = 'OPEN_LOOTBOX';
/** handles buying new lootbox */
export const buyLootbox = (lootbox: Lootbox) => ({
    type: BUY_LOOTBOX,
    lootbox,
});

const client = new LootboxClient();
const service = new LootboxService(client);
/** thunk that handles opening new lootbox */
export const openLootbox = (lootbox: Lootbox) => async function(dispatch: Dispatch) {
    const response = await service.buy(lootbox);
    const openedLootbox = await response.json();
    dispatch(buyLootbox(openedLootbox));
};
