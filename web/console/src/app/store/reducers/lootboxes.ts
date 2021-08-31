// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { LootboxClient } from '@/api/lootboxes';
import { Lootboxes } from '@/lootbox/servise';

import { BUY_LOOTBOX, OPEN_LOOTBOX } from '../actions/lootboxes';

/** Lootbox state base implementation */
export class LootboxState {
    public readonly lootboxes: Lootboxes;
    /** receives lootbox service */
    public constructor(lootboxes: Lootboxes) {
        this.lootboxes = lootboxes;
    };
};

const lootboxClient = new LootboxClient();
const lootboxService = new Lootboxes(lootboxClient);
const lootboxState = new LootboxState(lootboxService);

export const lootboxReducer = (
    state = lootboxState,
    action: any = {}
) => {
    const lootboxes = state.lootboxes;

    switch (action.type) {
        case BUY_LOOTBOX:
            lootboxes.buy(action.lootbox);
            break;
        case OPEN_LOOTBOX:
            lootboxes.open(action.lootbox);
            break;
        default:
            break;
    };

    return { ...lootboxState };
};
