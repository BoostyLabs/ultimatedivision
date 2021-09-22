// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { ClubClient } from '@/api/club';
import { Card } from '@/card';
import { ClubFromApi } from '@/club';
import { ClubService } from '@/club/service';
import { Dispatch } from 'redux';

export const CREATE_CLUB = 'CREATE_CLUB';
export const FORMATION = 'FORMATION';
export const SELECTION_VISIBILITY = 'SELECTION_VISIBILITY';
export const TACTICS = 'TACTICS';
export const CAPTAIN = 'CAPTAIN';
export const CARD_POSITION = 'CARD_POSITION';
export const ADD_CARD = 'ADD_CARD';
export const REMOVE_CARD = 'REMOVE_CARD';
export const DRAG_START = 'DRAG_START';
export const DRAG_TARGET = 'DRAG_TARGET';
export const EXCHANGE_CARDS = 'EXCHANGE_CARDS';

type dragParamType = number | null;
const DEFAULT_CARD_INDEX = null;

const client = new ClubClient();
const service = new ClubService(client);

export const createClub = (club: ClubFromApi) => ({
    type: CREATE_CLUB,
    club,
});

/** Chose type of cards positioning on football field */
export const setFormation = (formation: string) => ({
    type: FORMATION,
    formation,
});

export const cardSelectionVisibility = (isVisible: boolean) => ({
    type: SELECTION_VISIBILITY,
    isVisible,
});

/** Adding into cardList in reducer */
export const addCard = (card: Card, index: number) => ({
    type: ADD_CARD,
    fieldCard: {
        card,
        index,
    },
});

export const removeCard = (index: dragParamType = DEFAULT_CARD_INDEX) => ({
    type: REMOVE_CARD,
    index,
});

/** Selection position of card which should be added */
export const choosePosition = (index: number) => ({
    type: CARD_POSITION,
    index,
});

export const setDragStart = (index: dragParamType = DEFAULT_CARD_INDEX) => ({
    type: DRAG_START,
    index,
});

export const setDragTarget = (index: dragParamType = DEFAULT_CARD_INDEX) => ({
    type: DRAG_TARGET,
    index,
});

export const exchangeCards = (previous: dragParamType, current: dragParamType) => ({
    type: EXCHANGE_CARDS,
    position: {
        previous,
        current,
    },
});

export const setTactic = (tactic: string) => ({
    type: TACTICS,
    tactic,
});

export const setCaptain = (captain: string) => ({
    type: CAPTAIN,
    captain,
});
