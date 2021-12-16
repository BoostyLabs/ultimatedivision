// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.
import {immerable} from "immer"
import { Card, CardsPage } from '@/card';
import { getCards, getCard } from '@/app/store/actions/cards';

import { createSlice } from '@reduxjs/toolkit';

const DEFAULT_OFFSET_VALUE: number = 0;
const DEFAULT_LIMIT_VALUE: number = 24;
const FIRST_PAGE: number = 1;
const PAGES_COUNT: number = 1;
const CARDS_TOTAL_COUNT: number = 1;

/** class for data from backent (test) */
class CardsState {
    [immerable] = true;
    /** class implementation */
    constructor(
        public cardsPage: CardsPage,
        public card: Card,
    ) { };
};

const page = {
    offset: DEFAULT_OFFSET_VALUE,
    limit: DEFAULT_LIMIT_VALUE,
    currentPage: FIRST_PAGE,
    pageCount: PAGES_COUNT,
    totalCount: CARDS_TOTAL_COUNT,
};

const cardsPage = new CardsPage([], page);
const openedCard = new Card();

export const cardsSlice = createSlice({
    name: 'cards',
    initialState: new CardsState(cardsPage, openedCard),
    reducers: {
    },
    extraReducers: (builder) => {
        builder.addCase(getCards.fulfilled, (state, action) => {
            state.cardsPage = action.payload;
        })
        builder.addCase(getCard.fulfilled, (state, action) => {
            state.card = action.payload;
        })
    }
})

export default cardsSlice.reducer;