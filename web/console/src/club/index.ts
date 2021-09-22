// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/** backend card interface */
export interface SquadCard {
    squadId: string;
    cardId: string;
    position: number;
}

/** backend club interface */
export interface ClubFromApi {
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

