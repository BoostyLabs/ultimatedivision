// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { MarketplaceClient } from '@/api/marketplace';
import { CreatedLot } from '@/app/types/marketplace';
import { CardsQueryParametersField } from '@/card';
import { Marketplaces } from '@/marketplace/service';
import { createAsyncThunk } from '@reduxjs/toolkit';

const marketplaceClient = new MarketplaceClient();
const marketplaces = new Marketplaces(marketplaceClient);

/** Creates lots query parameters and sets them to marketplace service. */
export const createLotsQueryParameters = (queryParameters: CardsQueryParametersField[]) => {
    marketplaces.changeLotsQueryParameters(queryParameters);
};

/** Returns lots */
export const listOfLots = createAsyncThunk(
    'marketplace/listOfLots',
    async function (selectedPage: number) {
        const marketplace = await marketplaces.list(selectedPage);
        const lots = marketplace.lots;
        const page = marketplace.page;

        return { lots, page }
    });

/** Creates lot from users card */
export const createLot = createAsyncThunk(
    'marketplace/createLot',
    async function (lot: CreatedLot) {
        await marketplaces.createLot(lot);
    });

/** Opens fotballerCardPage with reload possibility */
export const openMarketplaceCard = createAsyncThunk(
    'marketplace/openMarketplaceCard',
    async function (id: string) {
        const lot = await marketplaces.getLotById(id);

        return lot.card;
    });
