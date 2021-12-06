// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.
import { Card, CardsPage, CardsQueryParameters } from '@/card';
import {
    GET_USER_CARDS,
    SEACH_CARDS_BY_PHYSIQUE,
    SEARCH_CARDS_BY_DEFENCE,
    SEARCH_CARDS_BY_GOALKEEPING,
    SEARCH_CARDS_BY_OFFENSE,
    SEARCH_CARDS_BY_STATUS,
    SEARCH_CARDS_BY_TACTIC,
    SEARCH_CARDS_BY_QUALITY,
    SEACH_CARDS_BY_TECHNIQUE,
    USER_CARD,
} from '@/app/store/actions/cards';

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
        case SEARCH_CARDS_BY_DEFENCE:
            return {
                ...cardsState,
                queryParameters: {
                    ...cardsState.queryParameters,
                    deffence: action.payload,
                },
            };
        case SEARCH_CARDS_BY_GOALKEEPING:
            return {
                ...cardsState,
                queryParameters: {
                    ...cardsState.queryParameters,
                    goalkeeping: action.payload,
                },
            };
        case SEARCH_CARDS_BY_OFFENSE:
            return {
                ...cardsState,
                queryParameters: {
                    ...cardsState.queryParameters,
                    offense: action.payload,
                },
            };
        case SEACH_CARDS_BY_PHYSIQUE:
            return {
                ...cardsState,
                queryParameters: {
                    ...cardsState.queryParameters,
                    physique: action.payload,
                },
            };
        case SEARCH_CARDS_BY_QUALITY:
            return {
                ...cardsState,
                queryParameters: {
                    ...cardsState.queryParameters,
                    quality: action.payload,
                },
            };
        case SEARCH_CARDS_BY_STATUS:
            return {
                ...cardsState,
                searchParams: {
                    ...cardsState.queryParameters,
                    status: action.payload,
                },
            };
        case SEARCH_CARDS_BY_TACTIC:
            return {
                ...cardsState,
                queryParameters: {
                    ...cardsState.queryParameters,
                    tactic: action.payload,
                },
            };
        case SEACH_CARDS_BY_TECHNIQUE:
            return {
                ...cardsState,
                queryParameters: {
                    ...cardsState.queryParameters,
                    technique: action.payload,
                },
            };
        case GET_USER_CARDS:
            return {
                ...cardsState,
                cardsPage: action.cardsPage,
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
