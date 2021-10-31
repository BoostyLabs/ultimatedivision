// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

const DEFAULT_INDEX = 0;
/** class for each control in option selection on field */
export class FieldControl {
    /** includes id, title and options parameters */
    constructor(
        public id: string = '',
        public title: string = '',
        public action: any = {},
        public options: string[] = [],
    ) { }
};

/** class for api methods to declare full path of card inside of club */
export class ClubCardPathModel {
    /** includes club id, squad idm card id and position parameters */
    constructor(
        public clubId: string,
        public squadId: string,
        public cardId: string,
        public position: number = DEFAULT_INDEX
    ) { }
}
