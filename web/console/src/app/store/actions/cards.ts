// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch } from 'redux';

import { CardClient } from '@/api/cards';
import { Card, CardsPage, CardsQueryParametersField } from '@/card';
import { CardService } from '@/card/service';

export const GET_USER_CARDS = ' GET_USER_CARDS';
export const USER_CARD = 'OPEN_USER_CARD';

/** Exposes action type to search cards by query parameters. */
export const SEARCH_CARDS_BY_QUERY_PARAMETERS: string = 'SEARCH_CARDS_BY_QUERY_PARAMETERS';

const getCards = (cardsPage: CardsPage) => ({
    type: GET_USER_CARDS,
    cardsPage,
});
const userCard = (card: Card) => ({
    type: USER_CARD,
    card,
});

/** Exposes an action object that sets query parameters to cardsReducer. */
export const setCardsQueryParameters = (parameters: CardsQueryParametersField[]) => ({
    type: SEARCH_CARDS_BY_QUERY_PARAMETERS,
    payload: parameters,
});

const client = new CardClient();
const service = new CardService(client);

/** Creates cards query parameters and sets them to CardsService. */
export const createCardsQueryParameters = (parameters: CardsQueryParametersField[]) => async function (dispatch: Dispatch) {
    await service.changeCardsQueryParameters(parameters);
    dispatch(setCardsQueryParameters(parameters));
};

/** thunk for creating user cards list */
export const listOfCards = (selectedPage: number) => async function (dispatch: Dispatch) {
    const response = await service.list(selectedPage);
    const page = response.page;
    const cards = response.cards;

    dispatch(getCards({ cards, page }));
};
/** thunk for opening fotballerCardPage with reload possibility */
export const openUserCard = (id: string) => async function (dispatch: Dispatch) {
    const card = await service.getCardById(id);

    dispatch(userCard(card));
};
