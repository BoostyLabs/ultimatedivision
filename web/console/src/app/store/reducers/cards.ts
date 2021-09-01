// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { ADD_CARDS } from '@/app/store/actions/cards';
import { CardService } from '@/card/service';
import { CardClient } from '@/api/cards';
import { Card } from '@/card';

/** class for data from backent (test) */
class CardSetup {
    /** class implementation */
    constructor(
        public cardService: CardService,
        public cards: Card[]
    ) {}
}

/** To do: replace cards by empty array */
export const cardSetup = new CardSetup(new CardService(new CardClient()), []);

export const cardsReducer = (cardState = cardSetup, action: any = {}) => {
    /** will be enabled when api will be returning list of cards */
    switch (action.type) {
    case ADD_CARDS:
        cardState.cards = action.action;
        break;
    default:
        break;
    }

    return { ...cardState };
};
