// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from "@/api/index";

import { CurrentSeasonsDivision, DivisionSeasonsStatistics } from "@/divisions";

/** DivisionsClient base implementation */
export class DivisionsClient extends APIClient {
    private readonly ROOT_PATH: string = "/api/v0";

    /** gets divisions of current seasons from api */
    public async getCurrentSeasonsDivisions(): Promise<
        CurrentSeasonsDivision[]
    > {
        const response = await this.http.get(
            `${this.ROOT_PATH}/seasons/current`
        );

        if (!response.ok) {
            await this.handleError(response);
        }

        return await response.json();
    }

    /** gets division seasons statistics */
    public async getDivisionSeasonsStatistics(): Promise<DivisionSeasonsStatistics> {
        const response = await this.http.get(
            `${this.ROOT_PATH}/seasons/statistics`
        );

        if (!response.ok) {
            await this.handleError(response);
        }

        return await response.json();
    }
}
