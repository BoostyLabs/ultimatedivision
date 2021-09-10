// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { CardClient } from '@/api/cards';
import { CreatedLot, MarkeplaceResponse, CardsResponse } from '@/card';

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
    public async getLots(): Promise<MarkeplaceResponse> {
        const response = await this.card.getLots();
        return await response.json()
    }
    /** get user cards from api */
    public async getUserCards(): Promise<CardsResponse> {
        const response = await this.card.getUserCards()
        return await response.json();
    }
    /** create lot */
    public async createLot(lot: CreatedLot): Promise<Response> {
        return await this.card.createLot(lot);
    }

    /** getting lot by id */
    public async getLotById(id: string): Promise<Response> {
        return await this.card.getLotById(id);

    }
}
