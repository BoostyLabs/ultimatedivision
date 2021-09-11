// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { CardClient } from '@/api/cards';
import { Card, MarketplaceLot } from '@/card';
import { CardService } from '@/card/service';
import { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { RootState } from '../store';
import { marketplaceLots } from '../store/actions/cards';

const client = new CardClient();
const service = new CardService(client);

export const useMarketplace = () => {
    const dispatch = useDispatch();

    /** Calls method get from  ClubClient */
    async function getCards() {
        dispatch(marketplaceLots());
    };

    useEffect(() => {
        getCards();
    }, []);

    return useSelector((state: RootState) => state.cardsReducer.marketplace);
};


    export async function getLotFromApi(id: string): Promise<Card> {
        const response = await service.getLotById(id);
        const lot = await response.json();
        return await new Card(lot.card);
    }