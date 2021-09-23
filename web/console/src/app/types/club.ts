// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { SquadCard } from '@/club';

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


/** initial value for clubReducer */
export class InitialClubValue {
    /** local options for dNd functionality */
    public options = {
        chosedCard: 0,
        showCardSeletion: false,
        dragStart: null,
        dragTarget: null,
    };
    public clubs = {
        id: '',
        name: '',
        createdAt: '',
    };
    public squad = {
        id: '',
        clubId: '',
        formation: '4-4-2',
        tactic: 0,
        captainId: '',
    };
    /* eslint-disable */
    squadCards = [
        new SquadCard(),
        new SquadCard(),
        new SquadCard(),
        new SquadCard(),
        new SquadCard(),
        new SquadCard(),
        new SquadCard(),
        new SquadCard(),
        new SquadCard(),
        new SquadCard(),
        new SquadCard(),
    ];
}

