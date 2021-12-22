// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '@/api/index';
import { Card, CardsPage, CardsQueryParameters, CardsQueryParametersField } from '@/card';

/** CardClient base implementation */
export class CardsClient extends APIClient {
    private readonly ROOT_PATH: string = '/api/v0/cards';

    private queryParameters: CardsQueryParameters = new CardsQueryParameters();

    public getCurrentQueryParameters(): CardsQueryParameters {
        return this.queryParameters;
    };

    /** Clears queryParameters object. */
    public clearCardsQueryParameters() {
        this.queryParameters = new CardsQueryParameters();
    };

    /** Returns current cards queryParameters object. */
    public getCardsQueryParameters() {
        return this.queryParameters;
    };

    /** Changes queryParameters object. */
    public changeCardsQueryParameters(queryParameters: CardsQueryParametersField[]) {
        queryParameters.forEach(queryParameter => {
            for (const queryProperty in queryParameter) {
                if (queryParameter) {
                    this.queryParameters[queryProperty] = queryParameter[queryProperty];
                }
            };
        });
    };

    /** Clears concret cards query Parameters. */
    public clearConcretCardsQueryParameters(queryParameters: CardsQueryParametersField[]) {
        queryParameters.forEach(queryParameter => {
            for (const queryProperty in queryParameter) {
                if (queryParameter) {
                    this.queryParameters[queryProperty] = undefined;
                }
            };
        });
    };

    /** method calls get method from APIClient */
    public async list(selectedPage: number): Promise<CardsPage> {
        /** Variable limit is default limit value of cards on page. */
        const limit: number = 24;

        let queryParametersPath = '';

        /** Adds qualities query parameters to query path. */
        const addQualitiesQueryParameters = (queryParameter: string, quality: string) => {
            queryParametersPath += `&${queryParameter}=${quality}`;
        };

        for (const queryParameter in this.queryParameters) {
            if (this.queryParameters[queryParameter]) {
                queryParameter === 'quality' ? this.queryParameters[queryParameter].
                    forEach((quality: string) => addQualitiesQueryParameters(queryParameter, quality)) :
                    queryParametersPath += `&${queryParameter}=${this.queryParameters[queryParameter]}`;
            }
        };

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
