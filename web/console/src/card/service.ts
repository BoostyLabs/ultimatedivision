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
    /** getting lot by id */
    public async getCardById(id: string): Promise<Response> {
        return await this.card.getCardById(id);
    }

    /** get filtered cards from api */
    public async getFilteredCards(filterParam: string) {
        const response = await this.card.getFilteredCards(filterParam);

        return await response.json();
    }
    /** get user cards from api */
    public async getCards({ page, limit }: { page: number; limit: number }): Promise<CardsResponse> {
        const response = await this.card.getCards({ page, limit });

        return await response.json();
    };
    /** create lot */
    public async createLot(lot: CreatedLot): Promise<Response> {
        return await this.card.createLot(lot);
    }

    /** getting lot by id */
    public async getLotById(id: string): Promise<Response> {
        return await this.card.getLotById(id);
    }
    /** get lots from api */
    public async getLots({ page, limit }: { page: number; limit: number }): Promise<MarkeplaceResponse> {
        const response = await this.card.getLots({ page, limit });

        return await response.json();
    };
    /** get filtered lots from api */
    public async getFilteredLots(filterParam: string) {
        const response = await this.card.getFilteredLots(filterParam);

        return await response.json();
    }
}
