// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { User } from '@/user';
import { APIClient } from '.';

/** Client for user controller of api */
export class UserClient extends APIClient {
    private readonly ROOT_PATH: string = '/api/v0/auth';
    /** Register new user  */
    public async register(user: User): Promise<void> {
        const path = `${this.ROOT_PATH}/register`;
        const response = await this.http.post(path, JSON.stringify(user));

        if (!response.ok) {
            await this.handleError(response);
        };
    };
    /** user login */
    public async login(email: string, password: string): Promise<void> {
        const path = `${this.ROOT_PATH}/login`;
        const response = await this.http.post(path, JSON.stringify({
            email, password,
        }));

        if (!response.ok) {
            await this.handleError(response);
        };
    };
    /** change user password implementation */
    public async changePassword(password: string, newPassword: string): Promise<void> {
        const path = `${this.ROOT_PATH}/change-password`;
        const response = await this.http.post(path, JSON.stringify({
            password, newPassword,
        }));

        if (!response.ok) {
            await this.handleError(response);
        };
    };
    /** check user email token */
    public async checkEmailToken(token: string | null): Promise<void> {
        const path = `${this.ROOT_PATH}/email/confirm/${token}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        };
    };
    /** check user recover token */
    public async checkRecoverToken(token: string | null): Promise<void> {
        const path = `${this.ROOT_PATH}/reset-password/${token}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        };
    };
    /** recover user password */
    public async recoverPassword(newPassword: string): Promise<void> {
        const path = `${this.ROOT_PATH}/reset-password`;
        const response = await this.http.patch(path, JSON.stringify({ newPassword }));

        if (!response.ok) {
            await this.handleError(response);
        };
    };
    /** reset user password by email confirmation */
    public async sendEmailForResetPassword(email: string): Promise<void> {
        const path = `${this.ROOT_PATH}/password/${email}`;
        const response = await this.http.get(path, JSON.stringify({ email }));

        if (!response.ok) {
            await this.handleError(response);
        };
    };
};
