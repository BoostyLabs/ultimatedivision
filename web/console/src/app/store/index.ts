// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.


import { TypedUseSelectorHook, useDispatch, useSelector } from 'react-redux'
import { getDefaultMiddleware } from '@reduxjs/toolkit';

import { configureStore } from '@reduxjs/toolkit'
import cardsReducer from '@/app/store/reducers/cards';
import clubsReducer from '@/app/store/reducers/clubs';
import lootboxReducer from '@/app/store/reducers/lootboxes';
import marketplaceReducer from '@/app/store/reducers/marketplace';
import divisionsReducer from '@/app/store/reducers/divisions';
import matchesReducer from '@/app/store/reducers/matches';
import usersReducer from '@/app/store/reducers/users';

export const store = configureStore({
    reducer: {
        cards: cardsReducer,
        clubs: clubsReducer,
        lootboxes: lootboxReducer,
        marketplace: marketplaceReducer,
        user: usersReducer,
        divisions: divisionsReducer,
        matches: matchesReducer,
    },
    middleware:(getDefaultMiddleware) => getDefaultMiddleware({
        serializableCheck: false,
    })
});

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>;
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch;


export const useAppDispatch = () => useDispatch<AppDispatch>();
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector;

export default store;