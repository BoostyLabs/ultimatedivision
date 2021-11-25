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
        public id: string = '',
        public divisionId: string = '',
        public startedAt: Date = new Date(),
        public endedAt: Date = new Date()
    ) {}
}

/** Division matches statistics entity. */
// TODO: statistics need rewrite (waiting for backend).
export class DivisionSeasonsStatistics {
    public division: Division = new Division('0', 0, 0, new Date());
    public statistics: null = null;
}

/** divisions reducer initial state  */
// TODO: Can be changed (waiting for backend)
export class DivisionsState {
    constructor(
        public currentSeasonsDivisions: CurrentSeasonsDivision[] = [],
        public seasonsStatistics: DivisionSeasonsStatistics = new DivisionSeasonsStatistics(),
        public activeDivision: CurrentSeasonsDivision = new CurrentSeasonsDivision()
    ) {}
}
