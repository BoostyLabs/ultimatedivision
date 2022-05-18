// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { UsersClient } from '@/api/users';
import { User } from '.';

/**
 * Exposes all users related logic.
 */
export class UsersService {
    private readonly users: UsersClient;
    /** UsersService contains http implementation of users API  */
    public constructor(users: UsersClient) {
        this.users = users;
    }
    /** handles user registration */
    public async register(user: User): Promise<void> {
        await this.users.register(user);
    }
    /** handles user login */
    public async login(email: string, password: string): Promise<void> {
        await this.users.login(email, password);
    }
    /** changes user password */
    public async changePassword(password: string, newPassword: string): Promise<void> {
        await this.users.changePassword(password, newPassword);
    }
    /** checks user email token */
    public async checkEmailToken(token: string | null): Promise<void> {
        await this.users.checkEmailToken(token);
    }
    /** checks recover token */
    public async checkRecoverToken(token: string | null): Promise<void> {
        await this.users.checkRecoverToken(token);
    }
    /** recovers user password */
    public async recoverPassword(password: string): Promise<void> {
        await this.users.recoverPassword(password);
    }
    /** resets user password by email confirmation */
    public async sendEmailForPasswordReset(email: string): Promise<void> {
        await this.users.sendEmailForPasswordReset(email);
    }
    /** sends data to register user with velas wallet */
    public async velasRegister(authResult: any): Promise<void> {
        await this.users.velasRegister(authResult);
    }
    /** sends address to get nonce to login user */
    public async velasNonce(address: string): Promise<string> {
        return await this.users.velasNonce(address);
    }
    /** sends data to login user with velas wallet */
    public async velasLogin(nonce: string, authResult: any): Promise<void> {
        await this.users.velasLogin(nonce, authResult);
    }
}
