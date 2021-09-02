// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Card } from '@/card';
import { SetStateAction, useEffect, useState } from 'react';
import { useDispatch } from 'react-redux';

export const useCards = (thunk: any) => {
    type Data = {
        data: null | Card[];
        isLoading: boolean;
    };

    const dispatch = useDispatch();

    const [data, handleData] = useState<SetStateAction<Data>>({ data: null, isLoading: true });

    /** Calls method get from  ClubClient */
    async function getDataFromApi() {
        await dispatch(thunk());

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
