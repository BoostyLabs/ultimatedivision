// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { CardClient } from '@/api/cards';
import { CardsResponse, CreatedLot, MarkeplaceResponse } from '@/card';

/**
 * exposes all bandwidth related logic
 */
export class CardService {
    protected readonly card: CardClient;
    /** sets CardClient into card field */
    public constructor(card: CardClient) {
        this.card = card;
    }
    /** get marketplace cards from api */
    public async getSellingCards(): Promise<MarkeplaceResponse> {
        const response = await this.card.getSellingCards();

        return await response.json();
    }
    /** get user cards from api */
    public async getUserCards(): Promise<CardsResponse> {
        const response = await this.card.getUserCards();

        return await response.json();
    }
    /** sell card */
    public async sellCard(lot: CreatedLot): Promise<Response> {
        return await this.card.sellCard(lot);
    }
}
