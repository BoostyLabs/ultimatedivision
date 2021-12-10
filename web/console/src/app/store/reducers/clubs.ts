// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { createSlice } from '@reduxjs/toolkit';

import { Club, ClubState } from '@/club';
import {
    createClub, getClubs, setFormation,
    setCaptain, setTactic, addCard,
    deleteCard, changePosition, swapCards,
    changeActiveClub
} from '@/app/store/actions/clubs'


const ACTIVE_STATUS_VALUE = 1;


export const cardSelectionVisibility = (isVisible: boolean) => ({
    type: SELECTION_VISIBILITY,
    isVisible,
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

/** shows MatchFinder component that exposes logic of searching matches */
export const startSearchingMatch = (isSearchingMatch: boolean) => ({
    type: START_SEARCHING_MATCH,
    isSearchingMatch,
});

export const exchangeCards = (
    previous: dragParamType,
    current: dragParamType
) => ({
    type: EXCHANGE_CARDS,
    position: {
        previous,
        current,
    },
});
switch (action.type) {
    case SET_CLUBS:
        return {
            ...state,
            clubs: action.clubs,
            activeClub: action.clubs.find((club:Club) => club.status === ACTIVE_STATUS_VALUE),
        };
    case SELECTION_VISIBILITY:
        state.options.showCardSeletion = action.isVisible;
        break;
    case CARD_POSITION:
        state.options.chosedCard = action.index;
        break;
    case DRAG_START:
        state.options.dragStart = action.index;
        break;
    case START_SEARCHING_MATCH:
        return {
            ...state,
            isSearchingMatch: action.isSearchingMatch,
        };
    default:
        break;
    }
const clubState = new ClubState();

export const cardsSlice = createSlice({
    name: 'cards',
    initialState: new ClubState(),
    reducers: {
        cardSelectionVisibility: (state, )
    },
    extraReducers: (builder) => {
        builder.addCase(createClub.fulfilled, (state, action) => {
            state.clubs = action.payload;
            state.activeClub = state.clubs.find((club:Club) => club.status === ACTIVE_STATUS_VALUE) || new Club()
        })
        builder.addCase(getClubs.fulfilled, (state, action) => {
            state.clubs = action.payload;
            state.activeClub = state.clubs.find((club:Club) => club.status === ACTIVE_STATUS_VALUE) || new Club()
        })
        builder.addCase(setFormation.fulfilled, (state, action) => {
            state.clubs = action.payload;
            state.activeClub = state.clubs.find((club:Club) => club.status === ACTIVE_STATUS_VALUE) || new Club()
        })
        builder.addCase(setCaptain.fulfilled, (state, action) => {
            state.clubs = action.payload;
            state.activeClub = state.clubs.find((club:Club) => club.status === ACTIVE_STATUS_VALUE) || new Club()
        })
        builder.addCase(setTactic.fulfilled, (state, action) => {
            state.clubs = action.payload;
            state.activeClub = state.clubs.find((club:Club) => club.status === ACTIVE_STATUS_VALUE) || new Club()
        })
        builder.addCase(addCard.fulfilled, (state, action) => {
            state.clubs = action.payload;
            state.activeClub = state.clubs.find((club:Club) => club.status === ACTIVE_STATUS_VALUE) || new Club()
        })
        builder.addCase(deleteCard.fulfilled, (state, action) => {
            state.clubs = action.payload;
            state.activeClub = state.clubs.find((club:Club) => club.status === ACTIVE_STATUS_VALUE) || new Club()
        })
        builder.addCase(changePosition.fulfilled, (state, action) => {
            state.clubs = action.payload;
            state.activeClub = state.clubs.find((club:Club) => club.status === ACTIVE_STATUS_VALUE) || new Club()
        })
        builder.addCase(swapCards.fulfilled, (state, action) => {
            state.clubs = action.payload;
            state.activeClub = state.clubs.find((club:Club) => club.status === ACTIVE_STATUS_VALUE) || new Club()
        })
    }
})
