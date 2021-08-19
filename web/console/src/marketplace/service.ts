// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { MarketplaceClient } from '@/api/marketplace';

/**
 * exposes all bandwidth related logic
 */
export class Marketplace {
    protected readonly marketplace: MarketplaceClient;
    /** sets MarketplaceClient into marketplace field */
    public constructor(marketplace: MarketplaceClient = new MarketplaceClient()) {
        this.marketplace = marketplace;
    }
    /** get catds from api */
    public async getCards() {
        return await this.marketplace.getCards();
    }
}
