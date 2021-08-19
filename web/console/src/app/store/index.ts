// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { applyMiddleware, combineReducers, createStore } from 'redux';
import thunk from 'redux-thunk';

import { marketplaceReducer } from '@/app/store/reducers/marketplace';
import { fieldReducer } from '@/app/store/reducers/footballField';
import { ClubReducer } from '@/app/store/reducers/club';

const reducer = combineReducers({
    marketplaceReducer,
    fieldReducer,
    ClubReducer,
});

export const store = createStore(reducer, applyMiddleware(thunk));

export type RootState = ReturnType<typeof store.getState>;
