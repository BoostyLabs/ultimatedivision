// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

export interface ClubFromApi {
    clubs: {
        id: string,
        name: string,
        createdAt: string
    },
    squad: {
        id: string,
        clubId: string,
        formation: number,
        tactic: number,
        captainId: string
    },
    squadCards: null | SquadCard[]
}

export interface SquadCard {
    squadId: string,
    cardId: string,
    position: number
}