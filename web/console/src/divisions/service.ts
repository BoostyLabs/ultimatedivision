// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { DivisionsClient } from '@/api/divisions';
import { CurrentDivisionSeasons, DivisionSeasonsStatistics } from '.';

/** exposes all divisions related logic */
export class DivisionsService {
    private readonly divisions: DivisionsClient;

    /** receives DivisionsClient */
    public constructor(divisions: DivisionsClient) {
        this.divisions = divisions;
    }

    /** handles gets current seasons divisions */
    public async getCurrentDivisionSeasons(): Promise<CurrentDivisionSeasons[]> {
        const divisions = await this.divisions.getCurrentDivisionSeasons();

        return divisions;
    }

    /** handles gets divisions matches statistics */
    public async getDivisionSeasonsStatistics(id: string): Promise<DivisionSeasonsStatistics> {
        const divisionsSeasonsStatistics =
            await this.divisions.getDivisionSeasonsStatistics(id);

        return divisionsSeasonsStatistics;
    }

    /** handles gets seasons reward */
    public async getDivisionSeasonsReward(): Promise<any> {
        const divisionsSeasonsReward =
            await this.divisions.getDivisionSeasonsReward();

        return divisionsSeasonsReward;
    }

    /** handles request divisions seasons rewards status */
    public async seasonsRewardStatus(): Promise<number> {
        const seasonRewardStatus = await this.divisions.seasonsRewardStatus();

        return seasonRewardStatus;
    }
}
