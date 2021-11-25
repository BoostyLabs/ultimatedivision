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

// TODO: statistics need rewrite (waiting for backend).
const INITIAL_DIVISION_NAME: number = 0;
const INITIAL_DIVISION_PERSENT: number = 0;
/** Division matches statistics entity. */
export class DivisionSeasonsStatistics {
    public division: Division = new Division(
        '0',
        INITIAL_DIVISION_NAME,
        INITIAL_DIVISION_PERSENT,
        new Date()
    );
    public statistics: null = null;
}

// TODO: Can be changed (waiting for backend)
/** divisions reducer initial state. */
export class DivisionsState {
    constructor(
        public currentSeasonsDivisions: CurrentSeasonsDivision[] = [],
        public seasonsStatistics: DivisionSeasonsStatistics = new DivisionSeasonsStatistics(),
        public activeDivision: CurrentSeasonsDivision = new CurrentSeasonsDivision()
    ) {}
}
