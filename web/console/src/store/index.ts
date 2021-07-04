import { createStore, combineReducers } from 'redux';

import { cardReducer } from './reducers/footballerCard';
import { footballField } from './reducers/footballField';

const reducer = combineReducers({
    footballerCard: cardReducer,
    footballField: footballField
});

export const store = createStore(reducer);

export type RootState = ReturnType<typeof store.getState>
