// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { DivisionsClient } from '@/api/divisions';
import { CurrentSeasonsDivision, DivisionSeasonsStatistics } from '.';

/** exposes all divisions related logic */
export class DivisionsService {
    private readonly divisions: DivisionsClient;

    /** receives DivisionsClient */
    public constructor(divisions: DivisionsClient) {
        this.divisions = divisions;
    }

    /** handles gets current seasons divisions */
    public async getCurrentSeasonsDivisions(): Promise<CurrentSeasonsDivision[]> {
        const divisions = await this.divisions.getCurrentSeasonsDivisions();

        return divisions;
    }

    /** handles gets divisions matches statistics */
    public async getDivisionSeasonsStatistics(): Promise<DivisionSeasonsStatistics> {
        const divisionsSeasonsStatistics =
            await this.divisions.getDivisionSeasonsStatistics();

        return divisionsSeasonsStatistics;
    }
}
