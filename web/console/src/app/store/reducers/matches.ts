// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Goal, Match, Team } from '@/matches';
import { createSlice, PayloadAction } from '@reduxjs/toolkit';

/** Describes default summary goals scored by first team. */
const DEFAULT_FIRST_TEAM_GOALS: number = 0;
/** Describes default summary goals scored by second team. */
const DEFAULT_SECOND_TEAM_GOALS: number = 0;

/** Describes default goal scorers by first team. */
const DEFAULT_FIRST_TEAM_GOAL_SCORERS: Goal[] = [];
/** Describes default goal scorers by second team. */
const DEFAULT_SECOND_TEAM_GOAL_SCORERS: Goal[] = [];

/** Describes default userId value of first player. */
const DEFAULT_FIRST_USER_ID: string = '';

/** Describes default userId valuew of second player. */
const DEFAULT_SECOND_USER_ID: string = '';

const firstTeam = new Team(DEFAULT_FIRST_TEAM_GOALS, DEFAULT_FIRST_TEAM_GOAL_SCORERS, DEFAULT_FIRST_USER_ID);
const secondTeam = new Team(DEFAULT_SECOND_TEAM_GOALS, DEFAULT_SECOND_TEAM_GOAL_SCORERS, DEFAULT_SECOND_USER_ID);

export const matchesSlice = createSlice({
    name: 'matches',
    initialState: new Match(firstTeam, secondTeam),
    reducers: {
        getMatchScore: (state, action) => {
            state = Object.assign(state, action.payload);
        }
    }
})

export const { getMatchScore } = matchesSlice.actions;
export default matchesSlice.reducer;
