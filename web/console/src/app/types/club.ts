// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

const defaultNumber = 0;
const fourElementsColumn = 4;
const fiveElementsColumn = 5;

export enum amountColumnsElements {
    'default' = defaultNumber,
    'four-elements' = fourElementsColumn,
    'five-elements' = fiveElementsColumn,
}

/** class for each control in option selection on field */
export class Control {
    /** includes id, title and options parameters */
    constructor(
        public id: string = '',
        public title: string = '',
        public action: any = {},
        public options: string[] = [],
        public columnElements: number = defaultNumber
    ) {}
}
