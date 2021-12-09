// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { MarketplaceClient } from '@/api/marketplace';
import { Lot, MarketPlacePage } from '@/marketplace';
import { CreatedLot } from '@/app/types/marketplace';

/**
 * exposes all arketplace domain entity related logic
 */
export class Marketplaces {
    protected readonly marketplace: MarketplaceClient;

    /** default marketplaceClient implementation */
    public constructor(marketplace: MarketplaceClient) {
        this.marketplace = marketplace;
    };

    /** returns marketplace domain entity with list of lots */
    public async list(selectedPage: number): Promise<MarketPlacePage> {
        return await this.marketplace.list(selectedPage);
    };

    /** creates lot */
    public async createLot(lot: CreatedLot): Promise<void> {
        await this.marketplace.createLot(lot);
    };

    /** returns lot by id */
    public async getLotById(id: string): Promise<Lot> {
        return await this.marketplace.getLotById(id);
    };
};
