// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { CardService } from '@/card/service';
import { Card } from '@/card';
import { SetStateAction, useEffect, useState } from 'react';

export const useCards = () => {
    type Data = {
        data: null | Card[],
        isLoading: boolean
    }

    const [data, handleData] = useState<SetStateAction<Data>>({ data: null, isLoading: true });

    /** Calls metgod get from MarketplaceClient | ClubClient */
    async function getDataFromApi() {
        const cards = await new CardService().get();

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
