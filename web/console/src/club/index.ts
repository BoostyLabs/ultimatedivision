// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

const DEFAULT_POSITION = 0;
/** backend card interface */
export class SquadCard {
    public squadId = '';
    public cardId = '';
    public position = DEFAULT_POSITION;
    data: any;
}

/** backend club interface */
export interface Club {
    clubs: {
        id: string;
        name: string;
        createdAt: string;
    };
    squad: {
        id: string;
        clubId: string;
        formation: number;
        tactic: number;
        captainId: string;
    };
    squadCards: SquadCard[];
}

