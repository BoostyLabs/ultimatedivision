// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Card } from '@/card';

const DEFAULT_VALUE = 0;
const ACTIVE_STATUS_VALUE = 1;

/** Squad defines squad entity. */
export class Squad {
    public id: string = '';
    public clubId: string = '';
    public formation: number = DEFAULT_VALUE;
    public tactic: number = DEFAULT_VALUE;
    public captainId: string = '';
}

/** SquadCard defines squad card entity.  */
export class SquadCard {
    public squadId: string = '';
    public card: Card = new Card();
    public position: number = DEFAULT_VALUE;
}

/** Club defines club entity. */
export class Club {
    /** Includes id, name cleatedAt, squad, squadCards and status fields  */
    constructor(
        public id: string = '',
        public name: string = '',
        public createdAt: string = '',
        public squad: Squad = new Squad(),
        public squadCards: SquadCard[] = [],
        public status: number = ACTIVE_STATUS_VALUE,
    ) { }
}

/** Class defines fields for drag and drop */
export class Options {
    /** chosedCard for adding card on field
     * showCardSelection for showing/hiding list of cards
     * dragStart and dragTarget for changing card position or swapping cards
    */
    constructor(
        public chosedCard: number = DEFAULT_VALUE,
        public showCardSeletion: boolean = false,
        public dragStart: number | null = null,
        public dragTarget: number | null = null
    ) { }
}

/** club reducer initial state  */
export class ClubState {
    public clubs: Club[] = [];
    public activeClub = new Club();
    public options: Options = new Options();
    public isSearchingMatch: boolean = false;
}

