// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { immerable } from "immer"
import { createSlice } from "@reduxjs/toolkit";

import { loginUser, changeUserPassword, recoverUserPassword } from "../actions/users";

/**
 * UsersState is a representation of users reducer state.
 */
export class UsersState {
    [immerable] = true;
    public email: string = '';
    public password: string = '';
};

export const userSlice = createSlice({
    name: 'user',
    initialState: new UsersState(),
    reducers: {
    },
    extraReducers: (builder) => {
        builder.addCase(loginUser.fulfilled, (state, action) => {
            state = Object.assign(state, action.payload)
        })
        builder.addCase(changeUserPassword.fulfilled, (state, action) => {
            state.password = action.payload;
        })
        builder.addCase(recoverUserPassword.fulfilled, (state, action) => {
            state.password = action.payload;
        })
    }
});

export default userSlice.reducer;