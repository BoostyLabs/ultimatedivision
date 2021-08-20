// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

export const ADD_CARDS = 'AddCards';
import { useCards } from '@/app/hooks/useCards';
import { CardDev } from '@/card/indexDev';
import { Dispatch } from 'redux';

export const addCards = (cards: []) => ({
    type: ADD_CARDS,
    action: cards,
});

// thunk for creating cards list
export const createCardList = () => async function(dispatch: Dispatch) {
    const cardsRequest = await useCards();
    // @ts-ignore
    const listOfCards = cardsRequest.data.map(card => new CardDev(...card));
    dispatch(addCards(listOfCards));
};
