// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Card } from '@/card';
import { SquadCard } from '@/club';

const DEFAULT_INDEX = 0;
/** class created to get card data from exist cardList */
export class FootballFieldCard {
    public squadId: string = '';
    public cardId: string = '';
    public position: number = DEFAULT_INDEX;
    /** class which implements fotballField card */
    constructor(
        public cardFromApi: SquadCard
    ) {
        Object.assign(this, cardFromApi);
    }

    /** get card data from cards */
    public cardData(cards: Card[]) {
        return cards.find(card => card.id === this.cardId);
    }
}

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
        new FootballFieldCard(new InitialCard(1)),
        new FootballFieldCard(new InitialCard(2)),
        new FootballFieldCard(new InitialCard(3)),
        new FootballFieldCard(new InitialCard(4)),
        new FootballFieldCard(new InitialCard(5)),
        new FootballFieldCard(new InitialCard(6)),
        new FootballFieldCard(new InitialCard(7)),
        new FootballFieldCard(new InitialCard(8)),
        new FootballFieldCard(new InitialCard(9)),
        new FootballFieldCard(new InitialCard(10)),
        new FootballFieldCard(new InitialCard(11)),
    ];
}

/** initial card before fetch, will be replaced by FootballFieldCard */
export class InitialCard {
    constructor(public position: number) { }
    public squadId = '';
    public cardId = '';

}
