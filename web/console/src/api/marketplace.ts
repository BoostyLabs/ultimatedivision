// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '@/api/index';
import { CreatedLot } from '@/app/types/marketplace';
import { Lot, LotsQueryParameters, LotsQueryParametersField, MarketPlacePage } from '@/marketplace';

/** client for marketplace of api */
export class MarketplaceClient extends APIClient {
    private readonly ROOT_PATH: string = '/api/v0/marketplace';

    private queryParameters: LotsQueryParameters = new LotsQueryParameters();

    /** Changes queryParameters object. */
    public changeCardsQueryParameters(queryParameters: LotsQueryParametersField[]) {
        queryParameters.forEach(queryParameter => {
            for (const queryProperty in queryParameter) {
                if (queryParameter) {
                    this.queryParameters[queryProperty] = queryParameter[queryProperty];
                }
            };
        });
    };

    /** returns marketplace domain entity with list of lots*/
    public async list(selectedPage: number): Promise<MarketPlacePage> {
        /** Variable limit is default limit value of lots on page. */
        const limit: number = 24;

        let queryParametersPath = '';

        for (const queryParameter in this.queryParameters) {
            if (this.queryParameters[queryParameter]) {
                queryParametersPath += `&${queryParameter}=${this.queryParameters[queryParameter]}`;
            }
        };

        const path = `${this.ROOT_PATH}?page=${selectedPage}&limit=${limit}${queryParametersPath}`;

        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        };

        const lotsPage = await response.json();

        return new MarketPlacePage(lotsPage.lots.map((lot: any) => new Lot(lot)), lotsPage.page);
    };

    /** implements opening lot */
    public async getLotById(id: string): Promise<Lot> {
        const path = `${this.ROOT_PATH}/${id}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        };

        const lot = await response.json();

        return new Lot(lot);
    };

    /** implements creating lot (selling card) */
    public async createLot(lot: CreatedLot): Promise<void> {
        const path = `${this.ROOT_PATH}`;
        const response = await this.http.post(path, JSON.stringify(lot));

        if (!response.ok) {
            await this.handleError(response);
        };
    };
};
