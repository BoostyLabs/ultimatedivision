// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Card } from '@/cards';
import { MarketplaceClient } from '@/api/marketplace';
import { getCards } from '@/app/hooks/getCards';
import { CardDev } from '@/cards/indexDev';
import { Dispatch } from 'redux';

const FIRST_CARD_TYPE = 0;
const SECOND_CARD_TYPE = 1;
const THIRD_CARD_TYPE = 2;
const FOURTH_CARD_TYPE = 3;
const CARDS_AMOUNT = 20;

const GET_CARDS = 'GetCards';

const addCards = (cards: []) => ({
    type: GET_CARDS,
    action: cards,
});


/** create list of player cards (implementation for test)*/
function cardListCreator(count: number): Card[] {
    const list: Card[] = [];
    while (count) {
        list.push(
            new Card(FIRST_CARD_TYPE),
            new Card(SECOND_CARD_TYPE),
            new Card(THIRD_CARD_TYPE),
            new Card(FOURTH_CARD_TYPE)
        );
        count--;
    }

    return list;
}

/** for testing, will be replaced by empty array */
const cardList = cardListCreator(CARDS_AMOUNT);
// const cardList: CardDev[] = [];

// thunk for creating cards list
export const createCardList = () => async function (dispatch: Dispatch) {
    const cardsRequest = await getCards(new MarketplaceClient());
    // @ts-ignore
    const listOfCards = cardsRequest.data.map(card => new CardDev(...card));
    dispatch(addCards(listOfCards));
};

// to do: add switch case for reducer when api will be done
export const marketplaceReducer = (cardState = cardList, action: any = {}) => {
    /** will be enabled when api will be returning list of cards */
    // switch (action.type) {
    // case GET_CARDS:
    //     cardState = action.action;
    //     break;
    // default:
    //     break;
    // }

    // @ts-ignore
    return [...cardState];
}
