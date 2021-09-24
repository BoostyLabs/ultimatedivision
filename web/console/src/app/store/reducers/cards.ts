// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import {
    GET_SELLING_CARDS, GET_USER_CARDS,
    MARKETPLACE_CARD, USER_CARD,
} from '@/app/store/actions/cards';
import { CardService } from '@/card/service';
import { CardClient } from '@/api/cards';
import { Card, MarketplaceLot } from '@/card';

const DEFAULT_CURRENT_PAGE_NUMBER: number = 1;
const DEFAULT_PAGES_COUNT: number = 1;
/** Card state base implementation */
class CardSetup {
    /** class implementation */
    constructor(
        public cardService: CardService,
        public marketplace: MarketplaceLot[],
        public club: Card[],
        public openedCard: Card,
        public marketplacePagesCount: number = DEFAULT_PAGES_COUNT,
        public clubPagesCount: number = DEFAULT_PAGES_COUNT,
        public marketplaceCurrentPage: number = DEFAULT_CURRENT_PAGE_NUMBER,
        public clubCurrentPage: number = DEFAULT_CURRENT_PAGE_NUMBER,
    ) { };
};

const cardClient = new CardClient();
const cardService = new CardService(cardClient);
export const cardSetup = new CardSetup(cardService, [], [], new Card());

export const cardsReducer = (cardState = cardSetup, action: any = {}) => {
    switch (action.type) {
        case GET_USER_CARDS:
            cardState.club = action.cards.cards;
            cardState.clubPagesCount = action.cards.pagesCount;
            cardState.clubCurrentPage = action.cards.currentPage;
            break;
        case GET_SELLING_CARDS:
            cardState.marketplace = action.lots.lots;
            cardState.marketplacePagesCount = action.lots.pagesCount;
            cardState.marketplaceCurrentPage = action.lots.currentPage;
            break;
        case MARKETPLACE_CARD:
            cardState.openedCard = action.card;
            break;
        case USER_CARD:
            cardState.openedCard = action.card;
            break;
        default:
            break;
    };

    return { ...cardState };
};
