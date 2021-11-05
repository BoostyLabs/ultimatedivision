import { combineReducers, configureStore, Action } from "@reduxjs/toolkit";
import { ThunkAction } from 'redux-thunk';
import  cards from '@/app/store/reducers/cardsToolkit';
import  lootbox from '@/app/store/reducers/lootboxToolkit';

const rootReducer = combineReducers({
    cards,
    lootbox
})

const store = configureStore({
    reducer: rootReducer
})

export type RootState = ReturnType<typeof rootReducer>;
export type AppDispatch = typeof store.dispatch;
export type AppThunk = ThunkAction<void, RootState, null, Action<string>>;

export default store;