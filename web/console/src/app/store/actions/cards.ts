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

export const getCards = ({ cards, pagesCount, currentPage }: {
    cards: Card[],
    pagesCount: number,
    currentPage: number,
}) => ({
    type: GET_USER_CARDS,
    cards: {
        cards,
        pagesCount,
        currentPage,
    },
});
export const getLots = ({ lots, pagesCount, currentPage }: {
    lots: Array<Partial<MarketplaceLot>>,
    pagesCount: number,
    currentPage: number,
}) => ({
    type: GET_SELLING_CARDS,
    lots: {
        lots,
        pagesCount,
        currentPage,
    },
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
export const userCards = (page: number) => async function (dispatch: Dispatch) {
    const response = await service.getCards(page);
    const pagesCount = response.page.pageCount;
    const currentPage = response.page.currentPage;
    const cards = response.cards.
        map((card: Partial<CardInterface>) => new Card(card));
    dispatch(getCards({ cards, pagesCount, currentPage }));
};
/** thunk for opening fotballerCardPage with reload possibility */
export const openUserCard = (id: string) => async function (dispatch: any) {
    const response = await service.getCardById(id);
    const card = await response.json();
    dispatch(userCard(new Card(card)));
};
/** thunk for creating user cards list */
export const marketplaceLots = (page: number) => async function (dispatch: Dispatch) {
    const response = await service.getLots(page);
    const pagesCount = response.page.pageCount;
    const currentPage = response.page.currentPage;
    const lots = response.lots.
        map((lot: Partial<MarketplaceLot>) => ({ ...lot, card: new Card(lot.card) }));
    dispatch(getLots({ lots, pagesCount, currentPage }));
};

export const createLot = (lot: CreatedLot) => async function (dispatch: Dispatch) {
    await service.createLot(lot);
};

/** thunk for opening fotballerCardPage with reload possibility */
export const openMarketplaceCard = (id: string) => async function (dispatch: Dispatch) {
    const response = await service.getLotById(id);
    const lot = await response.json();
    dispatch(marketplaceCard(new Card(lot.card)));
};

/** thunk returns filtered cards */
export const filteredCards = (lowRange: string, topRange: string) => async function (dispatch: Dispatch) {
    const filterParam = `${lowRange}&${topRange}`;
    const response = await service.getFilteredCards(filterParam);
    const cards = await response.json();
    dispatch(getCards(cards.map((card: Partial<CardInterface>) => new Card(card))));
};

/** thunk returns filtered lots */
export const filteredLots = (lowRange: string, topRange: string) => async function (dispatch: Dispatch) {
    const filterParam = `${lowRange}&${topRange}`;
    const response = await service.getFilteredLots(filterParam);
    const lots = await response.json();
    dispatch(getLots(lots.map((lot: Partial<MarketplaceLot>) => ({ ...lot, card: new Card(lot.card) }))));
};
