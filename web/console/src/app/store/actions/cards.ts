// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

export const ADD_CARDS = 'AddCards';
import { CardClient } from '@/api/cards';
import { CardService } from '@/card/service';
import { Dispatch } from 'redux';

export const addCards = (cards: []) => ({
    type: ADD_CARDS,
    action: cards,
});

/** thunk for creating cards list */
export const createCardList = () => async function(dispatch: Dispatch) {
    const response = await new CardService(new CardClient()).get();
    // @ts-ignore
    dispatch(addCards(response));
};
