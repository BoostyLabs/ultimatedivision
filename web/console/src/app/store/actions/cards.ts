// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

export const ADD_CARDS = 'AddCards';
import { Card } from '@/card';
import { Dispatch } from 'redux';

export const addCards = (cards: []) => ({
    type: ADD_CARDS,
    action: cards,
});

// thunk for creating cards list
export const createCardList = (cards: any) => async function(dispatch: Dispatch) {
    // @ts-ignore
    dispatch(addCards(cards.data.map(card => new Card(...card))));
};
