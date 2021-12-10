// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { CardEditIdentificators, ClubsClient } from '@/api/club';
import { Club, Squad } from '.';

/**
 * Exposes all bandwidth related logic
 */
export class ClubService {
    protected readonly club: ClubsClient;

    /** Sets ClubClient into club field */
    public constructor(club: ClubsClient) {
        this.club = club;
    };

    /** Creates club */
    public async createClub(): Promise<string> {
        return await this.club.createClub();
    };

    /** Returns club with existing squads */
    public async getClubs(): Promise<Club[]> {
        return await this.club.getClubs();
    };
    /** Returns active club */
    public async getActiveClub(): Promise<Club> {
        const clubs = await this.getClubs();

        return clubs.find((club: Club) => club.status !== 0) || new Club();
    }

    /** Creates squad in selected club */
    public async createSquad(clubId: string): Promise<string> {
        return await this.club.createSquad(clubId);
    };

    /** Adds card to squad cards list */
    public async addCard(path: CardEditIdentificators): Promise<void> {
        return await this.club.addCard(path);
    };

    /** Changes position of existing card */
    public async changeCardPosition(path: CardEditIdentificators): Promise<void> {
        return await this.club.changeCardPosition(path);
    };

    /** Deletes card from squad cards list */
    public async deleteCard(path: CardEditIdentificators): Promise<void> {
        return await this.club.deleteCard(path);
    };

    /** Updates squad tactic */
    public async updateTactic(squad: Squad, tactic: number): Promise<void> {
        return await this.club.updateTactic(squad, tactic);
    }

    /** Updates squad captain */
    public async updateCaptain(squad: Squad, captainId: string): Promise<void> {
        return await this.club.updateCaptain(squad, captainId);
    }

    /** Updates squad formation */
    public async updateFormation(squad: Squad, formation: number): Promise<void> {
        return await this.club.updateFormation(squad, formation);
    }
    /** Chandes active club */
    public async changeActiveClub(id: string): Promise<void> {
        return await this.club.changeActiveClub(id);
    }
};
