// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch } from 'redux';

import { CardClient } from '@/api/cards';
import { Card, CardInterface, CreatedLot, MarketplaceLot } from '@/card';
import { CardService } from '@/card/service';

export const GET_USER_CARDS = ' GET_USER_CARDS';
export const GET_SELLING_CARDS = ' GET_SELLING_CARDS';
export const MARKETPLACE_CARD = 'OPEN_MARKETPLACE_CARD';
export const USER_CARD = 'OPEN_USER_CARD';

export const getUserCards = (cards: Card[]) => ({
    type: GET_USER_CARDS,
    cards,
});
export const getSellingCards = (cards: Array<Partial<MarketplaceLot>>) => ({
    type: GET_SELLING_CARDS,
    cards,
});
export const marketplaceCard = (card: Card) => ({
    type: MARKETPLACE_CARD,
    card,
});
export const userCard = (card: Card) => ({
    type: USER_CARD,
    card,
});

const client = new CardClient();
const service = new CardService(client);

/** thunk for creating user cards list */
export const userCards = () => async function(dispatch: Dispatch) {
    const response = await service.getUserCards();
    const cards = response.cards;
    dispatch(getUserCards(cards.map((card: Partial<CardInterface>) => new Card(card))));
};
/** thunk for opening fotballerCardPage with reload possibility */
export const openUserCard = (id: string) => async function(dispatch: any) {
    const response = await service.getCardById(id);
    console.log(response)
    const card = await response;
    console.log(card)
    // dispatch(userCard(new Card(card)));
};
/** thunk for creating user cards list */
export const marketplaceLots = () => async function(dispatch: Dispatch) {
    const response = await service.getLots();
    const lots = response.lots;
    dispatch(getSellingCards(lots.map((lot: Partial<MarketplaceLot>) => ({ ...lot, card: new Card(lot.card) }))));
};

export const createLot = (lot: CreatedLot) => async function(dispatch: any) {
    await service.createLot(lot);
    dispatch(userCards());
};

/** thunk for opening fotballerCardPage with reload possibility */
export const openMarketplaceCard = (id: string) => async function(dispatch: any) {
    const response = await service.getLotById(id);
    const lot = await response.json();
    dispatch(marketplaceCard(new Card(lot.card)));
};
