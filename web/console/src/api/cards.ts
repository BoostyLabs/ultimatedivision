// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Card, CardsPage } from '@/card';

import { APIClient } from '@/api/index';

import { Pagination } from '@/app/types/pagination';

/** CardClient base implementation */
export class CardClient extends APIClient {
    private readonly ROOT_PATH: string = '/api/v0/cards';

    /** method calls get method from APIClient */
    public async list({ selectedPage, limit }: Pagination): Promise<CardsPage> {
        const path = `${this.ROOT_PATH}?page=${selectedPage}&limit=${limit}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        };

        const cardsJSON = await response.json();

        return new CardsPage(
            cardsJSON.cards.map((card: Partial<Card>) => new Card(card)),
            cardsJSON.page,
        );
    };

    /** method calls get method from APIClient */
    public async getCardById(id: string): Promise<Card> {
        const path = `${this.ROOT_PATH}/${id}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        };

        const cardJSON = await response.json();

        return new Card(cardJSON);
    };

    /** method returns filtered card list */
    public async filteredList(filterParam: string): Promise<CardsPage> {
        const path = `${this.ROOT_PATH}/?${filterParam}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        };

        const cardsJSON = await response.json();

        return new CardsPage(
            cardsJSON.cards.map((card: Partial<Card>) => new Card(card)),
            cardsJSON.page,
        );
    };
};
