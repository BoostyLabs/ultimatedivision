// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch } from "redux";
import { DivisionsClient } from "@/api/divisions";
import { CurrentSeasonsDivision, DivisionMatchesStatistics } from "@/divisions";
import { DivisionsService } from "@/divisions/service";

export const GET_CURRENT_SEASONS_DIVISIONS = "GET_CURRENT_SEASONS_DIVISIONS";
export const GET_DIVISION_MATCHES_STATISTICS =
    "GET_DIVISION_MATCHES_STATISTICS";

/** handles gets current seasons divisions */
export const getCurrentSeasonsDivisions = (
    divisions: CurrentSeasonsDivision[]
) => ({
    type: GET_CURRENT_SEASONS_DIVISIONS,
    divisions,
});

/** handles gets divisions matches statistics */
export const getDivisionMatchesStatistics = (
    statistics: DivisionMatchesStatistics
) => ({
    type: GET_DIVISION_MATCHES_STATISTICS,
    statistics,
});

const client = new DivisionsClient();
const service = new DivisionsService(client);

/** thunk that handles gets current seasons divisions */
export const listOfCurrentSeasonsDivisions = () =>
    async function (dispatch: Dispatch) {
        const currentSeasonsDivisions =
            await service.getCurrentSeasonsDivisions();

        currentSeasonsDivisions &&
            dispatch(
                getCurrentSeasonsDivisions(
                    currentSeasonsDivisions.map(
                        (division) =>
                            new CurrentSeasonsDivision(
                                division.id,
                                division.divisionId,
                                division.startedAt,
                                division.endedAt
                            )
                    )
                )
            );
    };

// /** thunk that handles gets current seasons divisions */
// export const listOfCurrentSeasonsDivisions = () =>
//     async function (dispatch: Dispatch) {
//         const currentSeasonsDivisions =
//             await service.getCurrentSeasonsDivisions();

//         currentSeasonsDivisions &&
//             dispatch(
//                 getCurrentSeasonsDivisions(
//                     response.map(
//                         (division) => new CurrentSeasonsDivision(division)
//                     )
//                 )
//             );
//     };
