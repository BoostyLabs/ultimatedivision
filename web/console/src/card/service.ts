// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { CardClient } from '@/api/cards';
import { CreatedLot } from '@/card';

/**
 * exposes all bandwidth related logic
 */
export class CardService {
    protected readonly card: CardClient;
    /** sets ClubClient into club field */
    public constructor(club: CardClient) {
        this.card = club;
    }
    /** get marketplace cards from api */
    public async getSellingCards(): Promise<Response> {
        return await this.card.getSellingCards();
    }
    /** get user cards from api */
    public async getUserCards(): Promise<Response> {
        return await this.card.getUserCards();
    }
    /** sell card */
    public async sellCard(lot: CreatedLot): Promise<Response> {
        return await this.card.sellCard(lot);
    }
}
