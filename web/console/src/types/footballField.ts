/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

export class FootballField {
    public options = {
        formation: '4-4-2',
        captain: '',
        tactics: '',
        chosedCard: null
    };
    public cardsList = [
        {
            id: 0,
            cardData: {}
        },
        {
            id: 1,
            cardData: {}
        },
        {
            id: 2,
            cardData: {}
        },
        {
            id: 3,
            cardData: {}
        },
        {
            id: 4,
            cardData: {}
        },
        {
            id: 5,
            cardData: {}
        },
        {
            id: 6,
            cardData: {}
        },
        {
            id: 7,
            cardData: {}
        },
        {
            id: 8,
            cardData: {}
        },
        {
            id: 9,
            cardData: {}
        },
        {
            id: 10,
            cardData: {}
        },
    ]
}

export class FootballFieldInformationLine {
    constructor(
        public id: string = '',
        public title: string = '',
        public options: string[] = [],
    ) { }
}
