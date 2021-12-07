// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '@/api/index';
import { Card, CardsPage, CardsQueryParameters, CardsQueryParametersField } from '@/card';

/** CardClient base implementation */
export class CardClient extends APIClient {
    private readonly ROOT_PATH: string = '/api/v0/cards';

    private queryParameters: CardsQueryParameters | null = new CardsQueryParameters();

    /** Changes queryParameters object. */
    public changeCardsQueryParameters(parameters: CardsQueryParametersField[]) {
        parameters.forEach(parameter => {
            if (this.queryParameters) {
                this.queryParameters[parameter.key] = parameter.value;
            };
        });
    };

    /** method calls get method from APIClient */
    public async list(selectedPage: number): Promise<CardsPage> {
        /** Variable limit is default limit value of cards on page. */
        const limit: number = 24;

        let queryParametersPath = '';

        for (let queryParameter in this.queryParameters) {
            if (this.queryParameters[queryParameter]) {
                queryParametersPath += `&${queryParameter}=${this.queryParameters[queryParameter]}`
            }
        }

        const path = `${this.ROOT_PATH}?page=${selectedPage}&limit=${limit}${queryParametersPath}`;

        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        };

        const cardsPage = await response.json();

        return new CardsPage(cardsPage.cards.map((card: any) => new Card(card)), cardsPage.page);
    };

    /** method calls get method from APIClient */
    public async getCardById(id: string): Promise<Card> {
        const path = `${this.ROOT_PATH}/${id}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        };

        const card = await response.json();

        return new Card(card);
    };
};
