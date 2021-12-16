// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import {immerable} from "immer";

import { Card } from '@/card';
import { createSlice } from '@reduxjs/toolkit';
import { openLootbox } from '../actions/lootboxes';

/** Lootbox state base implementation */
export class LootboxState {
    [immerable] = true;
    public cards: Card[] = [];
};

const lootboxState = new LootboxState();

export const lootboxesSlice = createSlice({
    name: 'lootboxes',
    initialState: new LootboxState(),
    reducers: {
    },
    extraReducers: (builder) => {
        builder.addCase(openLootbox.fulfilled, (state, action) => {
            state.cards = action.payload;
        })
    }
})

export default lootboxesSlice.reducer;