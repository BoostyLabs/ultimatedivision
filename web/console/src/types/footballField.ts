/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

export class FootballField {
    public options = {
        formation: '4-4-2',
        captain: '',
        tactics: '',
    };
}

export class FootballFieldInformationLine {
    constructor(
        public id: string ='',
        public title: string = '',
        public options: string[] = [],
        public action: (option: string) => {
            type: string,
            action: string
        }
    ) { }
}
