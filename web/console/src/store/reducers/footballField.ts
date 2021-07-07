/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import { FootballField } from '../../types/footballField';
import { Card } from '../../store/reducers/footballerCard'

const FieldSetup = new FootballField()

const FORMATION = 'Formation';
const TACTICS = 'Cactics';
const CAPTAIN = 'Captain';
const CHOSE = 'ChoseCard';
const ADD = 'AddCard';
const REMOVE = 'RemoveCard';



//Chose type of cards positioning on football field
export const handleFormations = (option: string) => {
    return {
        type: FORMATION,
        action: option
    }
};

//Adding into cardList in reducer
export const addCard = (card: Card, index: string | null) => {
    return {
        type: ADD,
        action: [card, index]
    }
}

export const removeCard = (index: number = -1) => {
    return {
        type: REMOVE,
        action: index
    }
}

//Selection position of card which should be added
export const choseCardPosition = (index: number) => {
    return {
        type: CHOSE,
        action: index
    }
}

export const handleTactics = (option: string) => {
    return {
        type: TACTICS,
        action: option
    }
};

export const handleCaptain = (option: string) => {
    return {
        type: CAPTAIN,
        action: option
    }
};


export const fieldReducer = (cardState = FieldSetup, action: any) => {

    switch (action.type) {
        case FORMATION:
            cardState.options.formation = action.action;
            return cardState;
        case CHOSE:
            cardState.options.chosedCard = action.action
            return cardState;
        case ADD:
            cardState.cardsList[action.action[1]].cardData = action.action[0];
            return {...cardState}
        case REMOVE:
            cardState.cardsList[action.action].cardData = null;
            return {...cardState}
        default:
            return cardState;
    }
};
