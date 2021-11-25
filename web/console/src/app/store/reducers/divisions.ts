// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { DivisionsState, CurrentSeasonsDivision } from '@/divisions';
import {
    GET_CURRENT_SEASONS_DIVISIONS,
    GET_DIVISION_SEASONS_STATISTICS,
    SET_ACTIVE_DIVISION,
} from '@/app/store/actions/divisions';

/** First divisions index from list. */
const FIRST_DIVISIONS_INDEX: number = 0;

export const divisionsReducer = (
    divisionsState: DivisionsState = new DivisionsState(),
    action: any = {}
) => {
    switch (action.type) {
    case GET_CURRENT_SEASONS_DIVISIONS:
        return {
            ...divisionsState,
            currentSeasonsDivisions: action.currentSeasonsDivisions,
            activeDivision: action.currentSeasonsDivisions.length
                ? action.currentSeasonsDivisions[FIRST_DIVISIONS_INDEX]
                : new CurrentSeasonsDivision(),
        };
    case GET_DIVISION_SEASONS_STATISTICS:
        return {
            ...divisionsState,
            seasonsStatistics: action.seasonsStatistics,
        };
    case SET_ACTIVE_DIVISION:
        return {
            ...divisionsState,
            activeDivision: action.activeDivision,
        };
    default:
        return divisionsState;
    }
};
