// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Lootbox } from '@/lootbox';

export const BUY_LOOTBOX = 'BUY_LOOTBOX';
export const OPEN_LOOTBOX = 'OPEN_LOOTBOX';
/** handles buying new lootbox */
export const buyLootbox = (lootbox: Lootbox) => ({
    type: BUY_LOOTBOX,
    lootbox,
});
/** handles opening new lootbox */
export const openLootbox = (lootbox: Lootbox) => ({
    type: OPEN_LOOTBOX,
    lootbox,
});
