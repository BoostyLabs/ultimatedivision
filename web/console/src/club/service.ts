// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { ClubClient } from '@/api/club';
import { ClubCardPath } from '@/app/types/club';
import { Club, Squad } from '.';

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
    public async getClub(): Promise<Club> {
        return await this.club.getClub();
    };

    /** creating squad in selected club */
    public async createSquad(clubId: string): Promise<string> {
        return await this.club.createSquad(clubId);
    };

    /** adding card to squad cards list */
    public async addCard(path: ClubCardPath): Promise<void> {
        return await this.club.addCard(path);
    };

    /** change position of existing card */
    public async changeCardPosition(path: ClubCardPath): Promise<void> {
        return await this.club.changeCardPosition(path);
    };

    /** delete card from squad cards list */
    public async deleteCard(path: ClubCardPath): Promise<void> {
        return await this.club.deleteCard(path);
    };

    /** updating squad tactic, formation or captain */
    public async updateSquad(squad: Squad): Promise<void> {
        return await this.club.updateSquad(squad);
    }
};
