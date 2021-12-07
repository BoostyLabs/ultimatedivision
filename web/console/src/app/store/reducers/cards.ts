// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { GET_USER_CARDS, SEARCH_CARDS_BY_QUERY_PARAMETERS, USER_CARD } from '@/app/store/actions/cards';
import { Card, CardsPage, CardsQueryParameters } from '@/card';

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
        public queryParameters?: Partial<CardsQueryParameters>,
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

export const cardsReducer = (cardsState: CardsState = new CardsState(cardsPage, openedCard), action: any = {}) => {
    switch (action.type) {
    case GET_USER_CARDS:
        return {
            ...cardsState,
            cardsPage: action.cardsPage,
        };
    case SEARCH_CARDS_BY_QUERY_PARAMETERS:
        return {
            ...cardsState,
            queryParameters: action.payload,
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
