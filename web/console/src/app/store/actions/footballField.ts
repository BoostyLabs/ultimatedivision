// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Card } from '@/card';

export const FORMATION_TYPE = 'Formation';
export const CARD_SELECTION_VISIBILITY = 'SelectionVisibility';
export const TACTICS_TYPE = 'Cactics';
export const CAPTAIN_TYPE = 'Captain';
export const CHOSE_CARD_POSITION = 'ChoseCard';
export const ADD_CARD = 'AddCard';
export const REMOVE_CARD = 'RemoveCard';
export const DRAG_START = 'CurrentPossition';
export const DRAG_TARGET = 'DragTarget';
export const EXCHANGE_CARDS = 'ReplaceCard';

type dragParamType = number | null;
const DEFAULT_CARD_INDEX = null;

/** Chose type of cards positioning on football field */
export const setFormation = (option: string) => ({
    type: FORMATION_TYPE,
    action: option,
});

export const cardSelectionVisibility = (option: boolean) => ({
    type: CARD_SELECTION_VISIBILITY,
    action: option,
});

/** Adding into cardList in reducer */
export const addCard = (card: Card, index: number) => ({
    type: ADD_CARD,
    action: [card, index],
});

export const removeCard = (index: dragParamType = DEFAULT_CARD_INDEX) => ({
    type: REMOVE_CARD,
    action: index,
});

/** Selection position of card which should be added */
export const choseCardPosition = (index: number) => ({
    type: CHOSE_CARD_POSITION,
    action: index,
});

export const setDragStart = (index: dragParamType = DEFAULT_CARD_INDEX) => ({
    type: DRAG_START,
    action: index,
});

export const setDragTarget = (index: dragParamType = DEFAULT_CARD_INDEX) => ({
    type: DRAG_TARGET,
    action: index,
});

export const exchangeCards = (prevPosition: dragParamType, currentPosition: dragParamType) => ({
    type: EXCHANGE_CARDS,
    action: [prevPosition, currentPosition],
});

export const setTactic = (option: string) => ({
    type: TACTICS_TYPE,
    action: option,
});

export const setCaptain = (option: string) => ({
    type: CAPTAIN_TYPE,
    action: option,
});
