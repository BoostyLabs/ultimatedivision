// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Squad } from '@/club';
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
export class ExactCardPath {
    /** class implementation */
    constructor(
        public squad: Squad,
        public cardId: string,
        public position: number = DEFAULT_INDEX
    ) { }
}
