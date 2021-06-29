/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import { createStore, combineReducers } from 'redux';

import { cardReducer } from './reducers/footballerCard';

const reducer = combineReducers({
    footballerCard: cardReducer,
});

export const store = createStore(reducer);
