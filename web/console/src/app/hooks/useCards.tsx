// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Card } from '@/card/service';
import { useEffect, useState } from 'react';

export const useCards = () => {
    const [data, handleData] = useState({ data: null, isLoading: true });

    /** Calls metgod get from MarketplaceClient | ClubClient */
    async function getDataFromApi() {
        const cards = await new Card().getCards();

        handleData({
            // @ts-ignore
            data: cards,
            isLoading: false,
        });
    };

    useEffect(() => {
        getDataFromApi();
    }, []);

    return data;
};
