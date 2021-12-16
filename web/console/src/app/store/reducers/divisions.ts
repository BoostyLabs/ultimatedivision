// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.



import { DivisionsState, CurrentDivisionSeasons } from '@/divisions';
import { createSlice } from '@reduxjs/toolkit';
import {
    getCurrentDivisionSeasons, getDivisionSeasonsStatistics
} from '@/app/store/actions/divisions';

/** First divisions index from list. */
const FIRST_DIVISIONS_INDEX: number = 0;

export const divisionsSlice = createSlice({
    name: 'divisions',
    initialState: new DivisionsState(),
    reducers: {
        setActiveDivision: (state, action) => {
            state.activeDivision = action.payload
        }
    },
    extraReducers: (builder) => {
        builder.addCase(getDivisionSeasonsStatistics.fulfilled, (state, action) => {
            state.seasonsStatistics = action.payload;
        })
        builder.addCase(getCurrentDivisionSeasons.fulfilled, (state, action) => {
            state.currentDivisionsSeasons = action.payload;
            state.activeDivision = action.payload.length
                ? action.payload[FIRST_DIVISIONS_INDEX]
                : new CurrentDivisionSeasons()
        })
    }
})

export const { setActiveDivision } = divisionsSlice.actions;
export default divisionsSlice.reducer;
