// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/** Domain entity Division implementation */
export class Division {
    constructor(
        public id: string,
        public name: number,
        public passingPercent: number,
        public createdAt: Date
    ) {}
}

/** Divisions of current season entity. */
export class CurrentSeasonsDivision {
    constructor(
        public id: string,
        public divisionId: string,
        public startedAt: Date,
        public endedAt: Date
    ) {}
}

/** Division matches statistics entity. */
//TODO: statistics need rewrite (wait backend).
export class DivisionMatchesStatistics {
    public division: Division = new Division("0", 0, 0, new Date());
    public statistics: null = null;
}

/** divisions reducer initial state  */
//TODO: Can be changed (waiting backend)
export class DivisionsState {
    public currentSeasonsDivisions: CurrentSeasonsDivision[] = [];
    public matchesStatistics: DivisionMatchesStatistics =
        new DivisionMatchesStatistics();
}
