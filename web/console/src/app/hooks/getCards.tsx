// @ts-nocheck
// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Club } from '@/club/service';
import { Marketplace } from '@/marketplace/service';
import { useEffect, useState } from 'react';

export const getCards = (client: Marketplace | Club) => {
    const [response, handleResponse] = useState({ data: null, isLoading: true });

    /** Calls metgod get from MarketplaceClient | ClubClient */
    async function getDataFromApi() {
        const data = await client.getCards();

        handleResponse({
            data: data,
            isLoading: false,
        });
    };

    useEffect(() => {
        getDataFromApi();
    }, []);

    return response;
};
