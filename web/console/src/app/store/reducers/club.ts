// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Club } from '@/club/service';
import { getCards } from '@/app/hooks/getCards';
import { CardDev } from '@/cards/indexDev';
import { Dispatch } from 'redux';

const GET_CARDS = 'GetCards';

const addCards = (cards: CardDev[]) => ({
    type: GET_CARDS,
    action: cards,
});

const cardList: CardDev[] = [];

// thunk for creating cards list
export const createCardList = () => async function(dispatch: Dispatch) {
    const cardsRequest = await getCards(new Club());
    // @ts-ignore
    const listOfCards = cardsRequest.data.map(card => new CardDev(...card));
    dispatch(addCards(listOfCards));
};

export const ClubReducer = (cardState = cardList, action: any = {}) => {
    switch (action.type) {
    case GET_CARDS:
        cardState = action.action;
        break;
    default:
        break;
    }

    // @ts-ignore
    return [...cardState];
};
