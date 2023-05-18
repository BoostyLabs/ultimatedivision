// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Card } from '@/card';

const DEFAULT_VALUE = 0;

/** Defines Marketplace lot entity. */
export class Lot {
    public cardId: string = '';
    public type: string = '';
    public userId: string = '';
    public shopperId: string = '';
    public status: string = '';
    public startPrice: number = DEFAULT_VALUE;
    public maxPrice: number = DEFAULT_VALUE;
    public currentPrice: number = DEFAULT_VALUE;
    public startTime: string = '';
    public endTime: string = '';
    public period: number = DEFAULT_VALUE;
    public card: Card = new Card();
    /** Tooks each field of parameter and assign it with exist fields */
    constructor(lot: Lot) {
        Object.assign(this, lot);
        this.card = new Card(lot.card);
    }
};

/** Class defines Marketplace response */
export class MarketPlacePage {
    /** Includes lots and page field */
    constructor(
        public lots: Lot[],
        public page: {
            offset: number;
            limit: number;
            currentPage: number;
            pageCount: number;
            totalCount: number;
        }
    ) { }
};


/** Class defines Marketplace response */
export class MarketCreateLotTransaction {
    /** Includes lots and page field */
    constructor(
        public address: string = '',
        public rpcNodeAddress: string = '',
        public tokenId: string = '',
        public minBidPrice: number = 200,
        public auctionDuration: number = 1,
        public price: number = 300,
    ) { }
};