// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Card } from '@/card';
import { SetStateAction, useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { RootState } from '../store';
import { createCardList } from '../store/actions/cards';

export const useCards = () => {
    const cardService = useSelector((state: RootState) => state.cardsReducer.cardService);

    type Data = {
        data: null | Card[];
        isLoading: boolean;
    };

    const dispatch = useDispatch();

    const [data, handleData] = useState<SetStateAction<Data>>({ data: null, isLoading: true });

    /** Calls method get from  ClubClient */
    async function getDataFromApi() {
        await dispatch(createCardList());

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
