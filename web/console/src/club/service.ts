// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { ClubClient } from '@/api/club';
import { CardPath } from '@/app/types/club';

/**
 * exposes all bandwidth related logic
 */
export class ClubService {
    protected readonly club: ClubClient;

    /** sets ClubClient into club field */
    public constructor(club: ClubClient) {
        this.club = club;
    };

    /** creating club */
    public async createClub(): Promise<Response> {
        const response = await this.createClub();

        return await response.json();
    };
    /** returning club with existing squads */
    public async getClub(): Promise<Response> {
        const response = await this.getClub();

        return await response.json();
    };
    /** creating squad in selected club */
    public async createSquad(clubId: string): Promise<Response> {
        const response = await this.createSquad(clubId);

        return await response.json();
    };

    /** adding card to squad cards list */
    public async addCard(clubId: string, squadId: string, cardId: string, position: number): Promise<Response> {
        return await this.addCard(clubId, squadId, cardId, position);
    };

    /** change position of existing card */
    public async changeCardPosition(clubId: string, squadId: string, cardId: string, position: number): Promise<Response> {
        return await this.changeCardPosition(clubId, squadId, cardId, position);
    };

    /** delete card from squad cards list */
    public async deleteCard(clubId: string, squadId: string, cardId: string): Promise<Response> {
        return await this.deleteCard(clubId, squadId, cardId);
    };
}
