//@ts-nocheck
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

export const clubsSlice = createSlice({
    name: 'clubs',
    initialState: new ClubState(),
    reducers: {
        startSearchingMatch: (state, action) => {
            state.isSearchingMatch = action.payload;
        },
        cardSelectionVisibility: (state, action) => {
            state.options.showCardSeletion = action.payload;
        },
        choosePosition: (state, action) => {
            state.options.chosedCard = action.payload;
        },
        setDragStart: (state, action) => {
            state.options.dragStart = action.payload;
        }

    },
    extraReducers: (builder) => {
        builder.addCase(createClub.fulfilled, (state, action) => {
            state.clubs = action.payload;
            state.activeClub = state.clubs.find((club: Club) => club.status === ACTIVE_STATUS_VALUE)
        })
        builder.addCase(getClubs.fulfilled, (state, action) => {
            state.clubs = action.payload;
            state.activeClub = state.clubs.find((club: Club) => club.status === ACTIVE_STATUS_VALUE)
        })
        builder.addCase(setFormation.fulfilled, (state, action) => {
            state.clubs = action.payload;
            state.activeClub = state.clubs.find((club: Club) => club.status === ACTIVE_STATUS_VALUE) || new Club()
        })
        builder.addCase(setCaptain.fulfilled, (state, action) => {
            state.clubs = action.payload;
            state.activeClub = state.clubs.find((club: Club) => club.status === ACTIVE_STATUS_VALUE) || new Club()
        })
        builder.addCase(setTactic.fulfilled, (state, action) => {
            state.clubs = action.payload;
            state.activeClub = state.clubs.find((club: Club) => club.status === ACTIVE_STATUS_VALUE) || new Club()
        })
        builder.addCase(addCard.fulfilled, (state, action) => {
            state.clubs = action.payload;
            state.activeClub = state.clubs.find((club: Club) => club.status === ACTIVE_STATUS_VALUE) || new Club()
        })
        builder.addCase(deleteCard.fulfilled, (state, action) => {
            state.clubs = action.payload;
            state.activeClub = state.clubs.find((club: Club) => club.status === ACTIVE_STATUS_VALUE) || new Club()
        })
        builder.addCase(changePosition.fulfilled, (state, action) => {
            state.clubs = action.payload;
            state.activeClub = state.clubs.find((club: Club) => club.status === ACTIVE_STATUS_VALUE) || new Club()
        })
        builder.addCase(swapCards.fulfilled, (state, action) => {
            state.clubs = action.payload;
            state.activeClub = state.clubs.find((club: Club) => club.status === ACTIVE_STATUS_VALUE) || new Club()
        })
    }
})

export const { startSearchingMatch, cardSelectionVisibility, setDragStart, choosePosition } = clubsSlice.actions;

export default clubsSlice.reducer;