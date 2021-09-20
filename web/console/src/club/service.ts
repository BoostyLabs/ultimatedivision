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
    public async addCard(path: CardPath, position: number): Promise<Response> {
        return await this.addCard(path, position);
    };

    /** change position of existing card */
    public async changeCardPosition(path: CardPath, position: number): Promise<Response> {
        return await this.changeCardPosition(path, position);
    };

    /** delete card from squad cards list */
    public async deleteCard(path: CardPath): Promise<Response> {
        return await this.deleteCard(path);
    };
}
