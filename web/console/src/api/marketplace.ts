// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '@/api/index';
import { CreatedLot } from '@/app/types/marketplace';
import { CardsQueryParameters, CardsQueryParametersField } from '@/card';
import { Lot, MarketPlacePage } from '@/marketplace';

/** client for marketplace of api */
export class MarketplaceClient extends APIClient {
    private readonly ROOT_PATH: string = '/api/v0/marketplace';

    public queryParameters: CardsQueryParameters = new CardsQueryParameters();

    /** Changes queryParameters object. */
    public changeLotsQueryParameters(queryParameters: CardsQueryParametersField[]): void {
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

    /** places a bid */
    public async placeBid(lotId: string, amount: number): Promise<void> {
        const path = '/api/v0/bids';
        const response = await this.http.post(path, JSON.stringify({ lotId, amount }));

        if (!response.ok) {
            await this.handleError(response);
        };
    };

    /** marketplace lot end time */
    public async endTime(lotId: string): Promise<boolean> {
        const path = `${this.ROOT_PATH}/end-time/${lotId}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        };

        const lotEndTime = await response.json();

        return lotEndTime;
    };

    /** marketplace lot data for casper */
    public async lotData(cardId: string): Promise<any> {
        const path = `${this.ROOT_PATH}/lot-data/${cardId}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        };

        const lotData = await response.json();

        return lotData;
    };

    /** make offer casper */
    public async makeOffer(cardId: string): Promise<any> {
        const path = `/api/v0/bids/make-offer/${cardId}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        };

        const makeOfferData = await response.json();

        return makeOfferData;
    };

    /** approve transaction casper */
    public async approve(cardId?: string): Promise<any> {
        let path;

        if (cardId) {
            path = `${this.ROOT_PATH}/approve?card_id=${cardId}`;
        }
        else {
            path = `${this.ROOT_PATH}/approve?card_id=`;
        }

        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        };

        const approveData = await response.json();

        return approveData;
    };
};
