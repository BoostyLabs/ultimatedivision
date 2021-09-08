// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch } from 'redux';

import { CardClient } from '@/api/cards';
import { Card, CardInterface, CreatedLot, MarketplaceLot } from '@/card';
import { CardService } from '@/card/service';

export const GET_USER_CARDS = ' GET_USER_CARDS';
export const GET_SELLING_CARDS = ' GET_SELLING_CARDS';

export const getUserCards = (cards: []) => ({
    type: GET_USER_CARDS,
    cards,
});
export const getSellingCards = (cards: []) => ({
    type: GET_SELLING_CARDS,
    cards,
});

const client = new CardClient();
const service = new CardService(client);

/** thunk for creating user cards list */
export const userCards = () => async function(dispatch: Dispatch) {
    const response = await service.getUserCards();
    const cards = await response.json();
    dispatch(getUserCards(cards.map((card: Partial<CardInterface>) => new Card(card))));
};
/** thunk for creating user cards list */
export const marketplaceCards = () => async function(dispatch: Dispatch) {
    const response = await service.getSellingCards();
    const lots = await response.json();

    dispatch(getSellingCards(lots.map((lot: Partial<MarketplaceLot>) => ({ ...lot, card: new Card(lot.card) }))));
};
export const sellCard = (lot: CreatedLot) => async function(dispatch: any) {
    await service.sellCard(lot);
    dispatch(userCards());
    dispatch(marketplaceCards());
};
