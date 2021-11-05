// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.
import { createSlice, PayloadAction } from '@reduxjs/toolkit';

import { AppThunk, AppDispatch } from '@/app/store/indexToolkit';
import { LootboxClient } from '@/api/lootboxes';
import { LootboxService } from '@/lootbox/service';
import { Card } from '@/card';
import { Lootbox } from '@/lootbox';

export class LootboxState {
    public readonly lootboxService: LootboxService;
    public lootbox = [];
    /** receives lootbox service */
    public constructor(lootboxService: LootboxService) {
        this.lootboxService = lootboxService;
    };
};

const client = new LootboxClient();
const service = new LootboxService(client);

const lootboxClient = new LootboxClient();
const lootboxService = new LootboxService(lootboxClient);
const initialState = new LootboxState(lootboxService);

export const lootboxSlice = createSlice({
    name: 'lootbox',
    initialState,
    reducers: {
        buyLootbox(state, action: PayloadAction<any>) {
            state.lootbox = action.payload;
        }
    }
});

export const openLootbox = (
    lootbox: Lootbox
): AppThunk => async (dispatch: AppDispatch) => {
    const opennedLootbox = await service.buy(lootbox);

    opennedLootbox &&
    dispatch(lootboxSlice.actions.buyLootbox(opennedLootbox.map(card => new Card(card))))
}

export const { buyLootbox } = lootboxSlice.actions;

export default lootboxSlice.reducer;  