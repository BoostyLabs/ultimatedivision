// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { LootboxClient } from "@/api/lootboxes";
import { Lootboxes } from "@/lootbox/servise";

import { BUY_LOOTBOX, OPEN_LOOTBOX } from "../actions/lootboxes";
export class LootboxState {
    public readonly lootboxes: Lootboxes;
    public constructor(lootboxes: Lootboxes) {
        this.lootboxes = lootboxes;
    };
};

export const lootboxReducer = (
    lootboxState = new LootboxState(new Lootboxes(new LootboxClient())),
    action: any = {}
) => {
    const lootboxes = lootboxState.lootboxes;

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
