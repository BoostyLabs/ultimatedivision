import { createStore, combineReducers } from 'redux';

import { cardStatsReducer } from './reducers/footballerCard';
import { filterFieldTitlesReducer } from './reducers/filterFieldTitles';

const reducer = combineReducers({
    footballerCard: cardStatsReducer,
    filterFieldTitles: filterFieldTitlesReducer
});

export const store = createStore(reducer);

export type RootState = ReturnType<typeof reducer>