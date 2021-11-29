// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Card } from '@/card';

/** Class defines Marketplace lot entity */
export class Lot {
    public id: string = '';
    public itemId: string = '';
    public type: string = '';
    public userId: string = '';
    public shopperId: string = '';
    public status: string = '';
    public startPrice: number = 0;
    public maxPrice: number = 0;
    public currentPrice: number = 0;
    public startTime: string = '';
    public endTime: string = '';
    public period: number = 0;
    public card: Card = new Card()
    constructor(lot: Lot) {
        Object.assign(this, lot);
        this.card = new Card(lot.card);
    }
};

/** Class defines Marketplace response */
export class MarketPlacePage {
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
