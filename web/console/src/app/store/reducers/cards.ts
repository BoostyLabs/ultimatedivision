// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { GET_USER_CARDS, USER_CARD } from '@/app/store/actions/cards';
import { Card, CardsPage } from '@/card';
import { GET_FIELD_CARDS } from '../actions/cards';

const DEFAULT_OFFSET_VALUE: number = 0;
const DEFAULT_LIMIT_VALUE: number = 24;
const FIRST_PAGE: number = 1;
const PAGES_COUNT: number = 1;
const CARDS_TOTAL_COUNT: number = 1;

/** class for data from backent (test) */
class CardsState {
    /** class implementation */
    constructor(
        public cardsPage: CardsPage,
        public card: Card,
        public currentCardsPage: number,
        public currentFieldCardsPage: number,
    ) { };
}

export const page = {
    offset: DEFAULT_OFFSET_VALUE,
    limit: DEFAULT_LIMIT_VALUE,
    currentPage: FIRST_PAGE,
    pageCount: PAGES_COUNT,
    totalCount: CARDS_TOTAL_COUNT,
};

const cardsPage = new CardsPage([], page);
const openedCard = new Card();

export const cardsReducer = (cardsState: CardsState = new CardsState(cardsPage, openedCard, FIRST_PAGE, FIRST_PAGE), action: any = {}) => {
    switch (action.type) {
    case GET_USER_CARDS:
        return {
            ...cardsState,
            cardsPage: action.payload.cardsPage,
            currentCardsPage: action.payload.currentPage,
        };
    case GET_FIELD_CARDS:
        return {
            ...cardsState,
            cardsPage: action.payload.cardsPage,
            currentFieldCardsPage: action.payload.currentPage,
        };
    case USER_CARD:
        return {
            ...cardsState,
            card: action.card,
        };
    default:
        return cardsState;
    };
};
