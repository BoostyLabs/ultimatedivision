// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch } from 'redux';

import { CardClient } from '@/api/cards';
import { Card, CardsPage, CardsQueryParameters, NumericCardsQueryParameters } from '@/card';
import { CardService } from '@/card/service';

export const GET_USER_CARDS = ' GET_USER_CARDS';
export const USER_CARD = 'OPEN_USER_CARD';
/** Exposes action type to search cards by defence skills. */
export const SEARCH_CARDS_BY_DEFENCE: string = 'SEARCH_CARDS_BY_DEFENCE';
/** Exposes action type to search cards by goalkeeping skills. */
export const SEARCH_CARDS_BY_GOALKEEPING: string = 'SEARCH_CARDS_BY_GOALKEEPING';
/** Exposes action type to search cards by offence skills. */
export const SEARCH_CARDS_BY_OFFENSE: string = 'SEARCH_CARDS_BY_OFFENSE';
/** Exposes action type to search cards by physique skills. */
export const SEACH_CARDS_BY_PHYSIQUE: string = 'SEACH_CARDS_BY_PHYSIQUE';
/** Exposes action type to search cards by card quality. */
export const SEARCH_CARDS_BY_QUALITY: string = 'SEARCH_CARDS_BY_QUALITY';
/** Exposes action type to search cards by card status. */
export const SEARCH_CARDS_BY_STATUS: string = 'SEARCH_CARDS_BY_STATUS';
/** Exposes action type to search cards by tactic skills. */
export const SEARCH_CARDS_BY_TACTIC: string = 'SEARCH_CARDS_BY_TACTIC';
/** Exposes action type to search cards by technique skills. */
export const SEACH_CARDS_BY_TECHNIQUE: string = 'SEACH_CARDS_BY_TECHNIQUE';

const getCards = (cardsPage: CardsPage) => ({
    type: GET_USER_CARDS,
    cardsPage,
});
const userCard = (card: Card) => ({
    type: USER_CARD,
    card,
});

/** Exposes an action object that sets defence field to query parameters. */
export const searchCardsByDefence = (defence: NumericCardsQueryParameters) => ({
    type: SEARCH_CARDS_BY_DEFENCE,
    payload: defence,
});

/** Exposes an action object that sets goalkeeping field to query parameters. */
export const searchCardsByGoalkeeping = (goalkeeping: NumericCardsQueryParameters) => ({
    type: SEARCH_CARDS_BY_GOALKEEPING,
    payload: goalkeeping,
});

/** Exposes an action object that sets offence field to query parameters. */
export const searchCardsByOffense = (offence: NumericCardsQueryParameters) => ({
    type: SEARCH_CARDS_BY_OFFENSE,
    payload: offence,
});

/** Exposes an action object that sets physique field to query parameters. */
export const searchCardsByPhysique = (physique: NumericCardsQueryParameters) => ({
    type: SEACH_CARDS_BY_PHYSIQUE,
    payload: physique,
});

/** Exposes an action object that sets quality field i.e. wood, silver, gold and diamond to query parameters. */
export const searchCardsByQuality = (quality: string) => ({
    type: SEARCH_CARDS_BY_QUALITY,
    payload: quality,
});

/** Exposes an action object that sets physique status i.e. locked and unlocked. to query parameters. */
export const searchCardsByStatus = (status: number) => ({
    type: SEARCH_CARDS_BY_STATUS,
    payload: status,
});

/** Exposes an action object that sets tactic field to query parameters. */
export const searchCardsByTactic = (tactic: NumericCardsQueryParameters) => ({
    type: SEARCH_CARDS_BY_TACTIC,
    payload: tactic,
});

/** Exposes an action object that sets technique field to query parameters. */
export const searchCardsByTechnique = (technique: NumericCardsQueryParameters) => ({
    type: SEACH_CARDS_BY_TECHNIQUE,
    payload: technique,
});

const client = new CardClient();
const service = new CardService(client);

/** thunk for creating user cards list */
export const listOfCards = (selectedPage: number, queryParameters?: Partial<CardsQueryParameters>) => async function(dispatch: Dispatch) {
    const response = await service.list(selectedPage, queryParameters);
    const page = response.page;
    const cards = response.cards;

    dispatch(getCards({ cards, page }));
};
/** thunk for opening fotballerCardPage with reload possibility */
export const openUserCard = (id: string) => async function(dispatch: Dispatch) {
    const card = await service.getCardById(id);

    dispatch(userCard(card));
};
