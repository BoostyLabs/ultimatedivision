// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Card } from '@/card';

const DEFAULT_VALUE = 0;

/** LotsQueryParametersField is an interface for cards query parameters field, that consist of key and value. */
export interface LotsQueryParametersField {
    [key: string]: string | number;
};

/** LotsQueryParameters is class that uses for filtering lots by queries. */
export class LotsQueryParameters {
    [key: string]: number | string | null;
    public 'defence_gte': number | null = null;
    public 'defence_lt': number | null = null;
    public 'goalkeeping_gte': number | null = null;
    public 'goalkeeping_lt': number | null = null;
    public 'offside_trap_gte': number | null = null;
    public 'offside_trap_lt': number | null = null;
    public 'physique_gte': number | null = null;
    public 'physique_lt': number | null = null;
    public 'tactics_gte': string | null = null;
    public 'tactics_lt': string | null = null;
    public 'quality': string | null = null;
    public 'status': number | null = null;
    public 'technique_gte': number | null = null;
    public 'technique_lt': number| null = null;
};

/** Class defines Marketplace lot entity */
export class Lot {
    public id: string = '';
    public itemId: string = '';
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
