// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import {
    DivisionsState,
    CurrentSeasonsDivision,
    DivisionMatchesStatistics,
} from "@/divisions";
import {
    GET_CURRENT_SEASONS_DIVISIONS,
    GET_DIVISION_MATCHES_STATISTICS,
} from "@/app/store/actions/divisions";

export const divisionsReducer = (
    divisionsState: DivisionsState = new DivisionsState(),
    action: any = {}
) => {
    switch (action.type) {
        case GET_CURRENT_SEASONS_DIVISIONS:
            return {
                ...divisionsState,
                divisions: action.divisions,
            };
        case GET_DIVISION_MATCHES_STATISTICS:
            return {
                ...divisionsState,
                statistics: action.statistics,
            };
        default:
            return divisionsState;
    }
};
