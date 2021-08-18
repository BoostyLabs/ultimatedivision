//@ts-nocheck
// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { ClubClient } from '@/api/club';
import { MarketplaceClient } from '@/api/marketplace';
import { useEffect, useState } from 'react'

export const getCards = ( client: MarketplaceClient | ClubClient ) => {
    const [response, handleResponse] = useState({ data: null, isLoading: true});

    async function getDataFromApi() {
        const data = await client.getCards();

        handleResponse({
            data: data,
            isLoading: false,
        })
    };

    useEffect(() => {
        getDataFromApi();
    }, []);

    return response;
}
