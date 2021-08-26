// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { FootballField } from '@/app/types/footballField';

import {
    ADD_CARD,
    CAPTAIN_TYPE,
    CARD_SELECTION_VISIBILITY,
    CHOSE_CARD_POSITION,
    DRAG_START,
    DRAG_TARGET,
    EXCHANGE_CARDS,
    FORMATION_TYPE,
    REMOVE_CARD,
    TACTICS_TYPE,
} from '@/app/store/actions/footballField';

const FieldSetup = new FootballField();


const FITST_ACTION_PARAM = 0;
const SECOND_ACTION_PARAM = 1;


export const fieldReducer = (cardState = FieldSetup, action: any = {}) => {
    const options = cardState.options;
    const cardsList = cardState.cardsList;

    switch (action.type) {
    case FORMATION_TYPE:
        options.formation = action.action;
        break;
    case CARD_SELECTION_VISIBILITY:
        options.showCardSeletion = action.action;
        break;
    case CHOSE_CARD_POSITION:
        options.chosedCard = action.action;
        break;
    case ADD_CARD:
        cardsList[action.action[SECOND_ACTION_PARAM]].cardData = action.action[FITST_ACTION_PARAM];
        break;
    case REMOVE_CARD:
        cardsList[action.action].cardData = null;
        break;
    case DRAG_START:
        options.dragStart = action.action;
        break;
    case DRAG_TARGET:
        options.dragTarget = action.action;
        break;
    case EXCHANGE_CARDS:
        const prevValue = cardsList[action.action[FITST_ACTION_PARAM]];
        cardsList[action.action[FITST_ACTION_PARAM]] = cardsList[action.action[SECOND_ACTION_PARAM]];
        cardsList[action.action[SECOND_ACTION_PARAM]] = prevValue;
        break;
    default:
        break;
    }

    return { ...cardState };
};
