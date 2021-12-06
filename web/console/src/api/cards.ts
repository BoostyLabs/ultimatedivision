// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '@/api/index';
import { Card, CardsPage, CardsQueryParameters } from '@/card';

/** CardClient base implementation */
export class CardClient extends APIClient {
    private readonly ROOT_PATH: string = '/api/v0/cards';

    /** method calls get method from APIClient */
    public async list(selectedPage: number, queryParameters?: Partial<CardsQueryParameters>): Promise<CardsPage> {
        /** Variable limit is default limit value of cards on page. */
        const limit: number = 24;

        let path = `${this.ROOT_PATH}?page=${selectedPage}&limit=${limit}`;

        if (queryParameters && queryParameters.deffence) {
            path += `&defence_gte=${queryParameters.deffence.min}&defence_lt=${queryParameters.deffence.max}`;
        };

        if (queryParameters && queryParameters.goalkeeping) {
            path += `&goalkeeping_gte=${queryParameters.goalkeeping.min}&goalkeeping_lt=${queryParameters.goalkeeping.max}`;
        };

        if (queryParameters && queryParameters.offense) {
            path += `&offside_trap_gte=${queryParameters.offense.min}&offside_trap_lt=${queryParameters.offense.max}`;
        };

        if (queryParameters && queryParameters.physique) {
            path += `&physique_gte=${queryParameters.physique.min}&physique_lt=${queryParameters.physique.max}`;
        };

        if (queryParameters && queryParameters.quality) {
            path += `&quality=${queryParameters.quality}`;
        };

        if (queryParameters && queryParameters.tactic) {
            path += `&tactics_gte=${queryParameters.tactic.min}&tactics_lt=${queryParameters.tactic.max}`;
        };

        if (queryParameters && queryParameters.technique) {
            path += `&technique_gte=${queryParameters.technique.min}&technique_lt=${queryParameters.technique.max}`;
        };

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
