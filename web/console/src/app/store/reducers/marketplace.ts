// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import {immerable} from "immer";

import { Card } from '@/card';
import { MarketPlacePage } from '@/marketplace';
import { createSlice } from '@reduxjs/toolkit';
import { listOfLots, openMarketplaceCard } from '../actions/marketplace';
import { matchesSlice } from './matches';

/** Markeplace state base implementation */
class MarketplaceState {
    [immerable] = true;
    /** default state implementation */
    constructor(
        public marketplacePage: MarketPlacePage,
        public card: Card,
    ) { };
};

const DEFAULT_OFFSET_VALUE: number = 0;
const DEFAULT_LIMIT_VALUE: number = 24;
const FIRST_PAGE: number = 1;
const PAGES_COUNT: number = 1;
const LOTS_TOTAL_COUNT: number = 1;

const page = {
    offset: DEFAULT_OFFSET_VALUE,
    limit: DEFAULT_LIMIT_VALUE,
    currentPage: FIRST_PAGE,
    pageCount: PAGES_COUNT,
    totalCount: LOTS_TOTAL_COUNT,
};

const marketplacePage = new MarketPlacePage([], page);
const card = new Card();

export const marketplaceSlice = createSlice({
    name: 'marketplace',
    initialState: new MarketplaceState(marketplacePage, card),
    reducers: {
    },
    extraReducers: (builder) => {
        builder.addCase(listOfLots.fulfilled, (state, action) => {
            state.marketplacePage = action.payload;
        })
        builder.addCase(openMarketplaceCard.fulfilled, (state, action) => {
            state.card = action.payload;
        })
    }
})

export default marketplaceSlice.reducer;