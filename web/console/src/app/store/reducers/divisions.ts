// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { DivisionsState, CurrentSeasonsDivision } from '@/divisions';
import {
    GET_CURRENT_SEASONS_DIVISIONS,
    GET_DIVISION_SEASONS_STATISTICS,
    SET_ACTIVE_DIVISION,
} from '@/app/store/actions/divisions';

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
                ? action.currentSeasonsDivisions[0]
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
