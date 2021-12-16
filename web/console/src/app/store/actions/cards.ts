// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { createAsyncThunk } from '@reduxjs/toolkit';

import { CardsQueryParametersField } from '@/card';
import { CardService } from '@/card/service';
import { CardsClient } from '@/api/cards';

const cardsClient = new CardsClient();
const cardsService = new CardService(cardsClient);

/** Clears cards query parameters. */
export const clearCardsQueryParameters = () => {
    cardsService.clearCardsQueryParameters();
};

/** Creates cards query parameters and sets them to CardsService. */
export const createCardsQueryParameters = (queryParameters: CardsQueryParametersField[]) => {
    cardsService.changeCardsQueryParameters(queryParameters);
};

/** thunk for creating user cards list */
export const getCards = createAsyncThunk(
    'cards/getCards',
    async function (selectedPage: number) {
        return await cardsService.list(selectedPage);
    }
)

/** thunk for opening fotballerCardPage with reload possibility */
export const getCard = createAsyncThunk(
    'cards/getCard',
    async function (id: string) {
        return await cardsService.getCardById(id);
    }
)