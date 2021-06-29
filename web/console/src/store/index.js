/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
*/

import { createStore, combineReducers } from 'redux';

import { cardStatsReducer } from './reducers/footballerCard';
import { cardPriceReducer } from './reducers/footballerCardPrice';
import { cardInfoReducer } from './reducers/footballerCardOveralInfo';
import { filterFieldTitlesReducer } from './reducers/filterFieldTitles';

const reducer = combineReducers({
    fotballerCardPrice: cardPriceReducer,
    footballerCardStats: cardStatsReducer,
    footballerCardOveralInfo: cardInfoReducer,
    filterFieldTitles: filterFieldTitlesReducer
});

export const store = createStore(reducer);
