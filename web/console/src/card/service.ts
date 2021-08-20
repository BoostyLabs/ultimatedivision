// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { CardClient } from '@/api/cards';

/**
 * exposes all bandwidth related logic
 */
export class Card {
    protected readonly club: CardClient;
    /** sets ClubClient into club field */
    public constructor(club: CardClient = new CardClient()) {
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
