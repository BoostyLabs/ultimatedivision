// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch } from "redux";
import { DivisionsClient } from "@/api/divisions";
import { CurrentSeasonsDivision, DivisionSeasonsStatistics } from "@/divisions";
import { DivisionsService } from "@/divisions/service";

export const GET_CURRENT_SEASONS_DIVISIONS = "GET_CURRENT_SEASONS_DIVISIONS";
export const GET_DIVISION_SEASONS_STATISTICS =
    "GET_DIVISION_SEASONS_STATISTICS";
export const SET_ACTIVE_DIVISION = "SET_ACTIVE_DIVISION";

/** handles gets current seasons divisions */
export const getCurrentSeasonsDivisions = (
    currentSeasonsDivisions: CurrentSeasonsDivision[]
) => ({
    type: GET_CURRENT_SEASONS_DIVISIONS,
    currentSeasonsDivisions,
});

/** handles gets divisions matches statistics */
export const getDivisionSeasonsStatistics = (
    seasonsStatistics: DivisionSeasonsStatistics
) => ({
    type: GET_DIVISION_SEASONS_STATISTICS,
    seasonsStatistics,
});

/** handles sets active division */
export const setActiveDivision = (activeDivision: CurrentSeasonsDivision) => ({
    type: SET_ACTIVE_DIVISION,
    activeDivision,
});

const client = new DivisionsClient();
const service = new DivisionsService(client);

/** thunk that handles gets current seasons divisions */
export const listOfCurrentSeasonsDivisions = () =>
    async function (dispatch: Dispatch) {
        const currentSeasonsDivisions =
            await service.getCurrentSeasonsDivisions();

        currentSeasonsDivisions &&
            dispatch(getCurrentSeasonsDivisions(currentSeasonsDivisions));
    };

/** thunk that handles gets seasons statistics */
export const divisionSeasonsStatistics = () =>
    async function (dispatch: Dispatch) {
        const seasonsStatistics = await service.getDivisionSeasonsStatistics();

        dispatch(getDivisionSeasonsStatistics(seasonsStatistics));
    };
