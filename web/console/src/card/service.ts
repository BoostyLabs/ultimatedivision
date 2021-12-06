// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { CardClient } from '@/api/cards';
import { Card,CardsPage, CardsQueryParameters } from '@/card';

/**
 * exposes all bandwidth related logic
 */
export class CardService {
    protected readonly card: CardClient;

    /** sets CardClient into card field */
    public constructor(card: CardClient) {
        this.card = card;
    };

    /** gets list of cards by user */
    public async list(selectedPage: number, queryParameters?: Partial<CardsQueryParameters>): Promise<CardsPage> {
        return await this.card.list(selectedPage, queryParameters);
    };

    /** gets card by id from list of cards */
    public async getCardById(id: string): Promise<Card> {
        return await this.card.getCardById(id);
    };
};
