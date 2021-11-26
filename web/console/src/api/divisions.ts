// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from "@/api/index";

import { CurrentSeasonsDivision, DivisionSeasonsStatistics } from "@/divisions";

/** DivisionsClient base implementation */
export class DivisionsClient extends APIClient {
    private readonly ROOT_PATH: string = "/api/v0";

    /** gets divisions of current seasons */
    public async getCurrentSeasonsDivisions(): Promise<CurrentSeasonsDivision[]> {
        const response = await this.http.get(
            `${this.ROOT_PATH}/seasons/current`
        );

        if (!response.ok) {
            await this.handleError(response);
        }

        const currentSeasonsDivisions = await response.json();

        if (!currentSeasonsDivisions) {
            this.handleError(currentSeasonsDivisions);
        }

        return currentSeasonsDivisions;
    }

    /** gets division seasons statistics */
    public async getDivisionSeasonsStatistics(
        id: string
    ): Promise<DivisionSeasonsStatistics> {
        const response = await this.http.get(
            `${this.ROOT_PATH}/seasons/statistics/division/${id}`
        );

        console.log(
            "path",
            `${this.ROOT_PATH}/seasons/statistics/division/${id}`
        );

        if (!response.ok) {
            await this.handleError(response);
        }

        const divisionsSeasonsStatistics = await response.json();
        console.log("api", divisionsSeasonsStatistics);

        if (!divisionsSeasonsStatistics) {
            this.handleError(divisionsSeasonsStatistics);
        }

        return divisionsSeasonsStatistics;
    }
}
