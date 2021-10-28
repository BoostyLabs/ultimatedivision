// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { User } from '@/user';
import { UserClient } from '@/api/user';
import { UserService } from '@/user/service';

/** action types implementation */
export const REGISTER = 'REGISTER';
export const LOGIN = 'LOGIN';
export const CHANGE_PASSWORD = 'CHANGE_PASSWORD';
export const RECOVER_PASSWORD = 'RECOVER_PASSWORD';
/** implement registration of new user */
export const register = (user: User) => ({
    type: REGISTER,
    user,
});
/** get registred user by id */
export const login = (email: string, password: string) => ({
    type: LOGIN,
    user: {
        email,
        password,
    },
});
/** changing user password */
export const changePassword = (password: string, newPassword: string) => ({
    type: CHANGE_PASSWORD,
    passwords: {
        password,
        newPassword,
    },
});
/** recover user password */
export const recoverPassword = (password: string) => ({
    type: RECOVER_PASSWORD,
    password,
});
