// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { FootballField } from '@/app/types/footballField';

import {
    ADD_CARD,
    CAPTAIN,
    CARD_POSITION,
    DRAG_START,
    DRAG_TARGET,
    EXCHANGE_CARDS,
    FORMATION,
    REMOVE_CARD,
    SELECTION_VISIBILITY,
    TACTICS,
} from '@/app/store/actions/footballField';

const FieldSetup = new FootballField();

export const fieldReducer = (cardState = FieldSetup, action: any = {}) => {
    const options = cardState.options;
    const cardsList = cardState.cardsList;

    switch (action.type) {
        case FORMATION:
            options.formation = action.formation;
            break;
        case SELECTION_VISIBILITY:
            options.showCardSeletion = action.isVisible;
            break;
        case CARD_POSITION:
            options.chosedCard = action.position;
            break;
        case ADD_CARD:
            cardsList[action.player.index].card =
                action.player.card;
            break;
        case REMOVE_CARD:
            cardsList[action.index].card = null;
            break;
        case DRAG_START:
            options.dragStart = action.index;
            break;
        case DRAG_TARGET:
            options.dragTarget = action.index;
            break;
        case EXCHANGE_CARDS:
            const prevCard = cardsList[action.position.previous];
            cardsList[action.position.previous] =
                cardsList[action.position.current];
            cardsList[action.position.current] = prevCard;
            break;
        default:
            break;
    }

    return { ...cardState };
};
