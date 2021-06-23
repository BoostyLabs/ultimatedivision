import { createStore, combineReducers } from 'redux';

import { cardStatsReducer } from './reducers/footballerCard';
import { cardPriceReducer } from './reducers/footballerCardPrice';
import { filterFieldTitlesReducer } from './reducers/filterFieldTitles';

const reducer = combineReducers({
    fotballerCardPrice: cardPriceReducer,
    footballerCardStats: cardStatsReducer,
    filterFieldTitles: filterFieldTitlesReducer
});

export const store = createStore(reducer);
