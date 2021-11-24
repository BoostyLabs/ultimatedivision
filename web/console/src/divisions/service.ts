// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { DivisionsClient } from "@/api/divisions";
import { CurrentSeasonsDivision, DivisionMatchesStatistics } from ".";

/** exposes all divisions related logic */
export class DivisionsService {
    private readonly divisions: DivisionsClient;

    /** receives DivisionsClient */
    public constructor(divisions: DivisionsClient) {
        this.divisions = divisions;
    }

    /** handles gets current seasons divisions */
    public async getCurrentSeasonsDivisions(): Promise<
        CurrentSeasonsDivision[]
    > {
        const divisions = await this.divisions.getCurrentSeasonsDivisions();

        return divisions.map(
            (division: CurrentSeasonsDivision) =>
                new CurrentSeasonsDivision(
                    division.id,
                    division.divisionId,
                    division.startedAt,
                    division.endedAt
                )
        );
    }

    /** handles gets divisions matches statistics */
    public async getDivisionMatchesStatistics(): Promise<DivisionMatchesStatistics> {
        const divisionsMatchesStatistics =
            await this.divisions.getDivisionMatchesStatistics();

        return divisionsMatchesStatistics;
    }
}
