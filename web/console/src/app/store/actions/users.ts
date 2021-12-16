// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { ChangePasswordFields, LoginFields, User } from '@/users';
import { UsersClient } from '@/api/users';
import { UsersService } from '@/users/service';
import { createAsyncThunk } from '@reduxjs/toolkit';

const usersClient = new UsersClient();
const usersService = new UsersService(usersClient);

/** thunk that implements user registration */
export const registerUser = createAsyncThunk(
    'user/registerUser',
    async function (user: User) {
        await usersService.register(user);
    });

/** thunk that implements user login */
export const loginUser = createAsyncThunk(
    'user/loginUser',
    async function (fields: LoginFields) {
        await usersService.login(fields.email, fields.password);
        return fields;
    });

/** thunk that implements changes user password */
export const changeUserPassword = createAsyncThunk(
    'user/changePassword',
    async function (fieds: ChangePasswordFields) {
        await usersService.changePassword(fieds.password, fieds.newPassword);
        return fieds.newPassword;
    });

/** thunk that implements resets user password */
export const recoverUserPassword = createAsyncThunk(
    'user/recoverUserPassword',
    async function (password: string) {
        await usersService.recoverPassword(password);
        return password
    });
