// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Card } from '@/card';

import { ADD_CARDS } from '@/app/store/actions/cards';

const FIRST_CARD_TYPE = 0;
const SECOND_CARD_TYPE = 1;
const THIRD_CARD_TYPE = 2;
const FOURTH_CARD_TYPE = 3;
const CARDS_AMOUNT = 20;


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


export const cardsReducer = (cardState = cardList, action: any = {}) =>
/** will be enabled when api will be returning list of cards */
// switch (action.type) {
// case ADD_CARDS:
//     cardState = action.action;
//     break;
// default:
//     break;
// }

    cardState.slice();

