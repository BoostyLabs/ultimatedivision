// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.


/** implementation for each field in
 * FootballFieldInformation component
 */
export class FootballFieldInformationLine {
    /** includes id, title and options parameters */
    constructor(
        public id: string = '',
        public title: string = '',
        public options: string[] = [],
    ) {
        this.id = id;
        this.title = title;
        this.options = options;
    }
};
