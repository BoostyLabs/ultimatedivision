// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { ClubClient } from '@/api/club';

/**
 * exposes all bandwidth related logic
 */
export class Club {
    protected readonly club: ClubClient;
    /** sets ClubClient into club field */
    public constructor(club: ClubClient = new ClubClient()) {
        this.club = club;
    }
    /** get catds from api */
    public async getCards() {
        return await this.club.getCards();
    }
    /** post cards into buyed cardlist */
    public async buyCard(param: string) {
        return await this.club.buyCard(param);
    }
}
