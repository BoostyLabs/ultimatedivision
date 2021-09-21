// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { ClubClient } from '@/api/club';

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
    public async createClub(): Promise<string> {
        return await this.club.createClub();
    };
    /** returning club with existing squads */
    public async getClub(): Promise<string> {
        return await this.club.getClub();
    };
    /** creating squad in selected club */
    public async createSquad(clubId: string): Promise<string> {
        return await this.club.createSquad(clubId);
    };

    /** adding card to squad cards list */
    public async addCard(clubId: string, squadId: string, cardId: string, position: number): Promise<Response> {
        return await this.club.addCard({ clubId, squadId, cardId, position });
    };

    /** change position of existing card */
    public async changeCardPosition(clubId: string, squadId: string, cardId: string, position: number): Promise<Response> {
        return await this.club.changeCardPosition({ clubId, squadId, cardId, position });
    };

    /** delete card from squad cards list */
    public async deleteCard(clubId: string, squadId: string, cardId: string): Promise<Response> {
        return await this.club.deleteCard({ clubId, squadId, cardId });
    };
};
