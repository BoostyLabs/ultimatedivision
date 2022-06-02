// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { User } from "@/users";
import { APIClient } from ".";

/**
 * UsersClient is a http implementation of users API.
 * Exposes all users-related functionality.
 */
export class UsersClient extends APIClient {
    private readonly ROOT_PATH: string = "/api/v0/auth";
    /** exposes user registration logic */
    public async register(user: User): Promise<void> {
        const path = `${this.ROOT_PATH}/register`;
        const response = await this.http.post(path, JSON.stringify(user));

        if (!response.ok) {
            await this.handleError(response);
        }
    }
    /** exposes user login logic */
    public async login(email: string, password: string): Promise<void> {
        const path = `${this.ROOT_PATH}/login`;
        const response = await this.http.post(
            path,
            JSON.stringify({
                email,
                password,
            })
        );

        if (!response.ok) {
            await this.handleError(response);
        }
    }
    /** changes user password */
    public async changePassword(password: string, newPassword: string): Promise<void> {
        const path = `${this.ROOT_PATH}/change-password`;
        const response = await this.http.post(
            path,
            JSON.stringify({
                password,
                newPassword,
            })
        );

        if (!response.ok) {
            await this.handleError(response);
        }
    }
    /** checks user token by email confirmation */
    public async checkEmailToken(token: string | null): Promise<void> {
        const path = `${this.ROOT_PATH}/email/confirm/${token}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        }
    }
    /** checks user recover token */
    public async checkRecoverToken(token: string | null): Promise<void> {
        const path = `${this.ROOT_PATH}/reset-password/${token}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        }
    }
    /** recovers user password */
    public async recoverPassword(newPassword: string): Promise<void> {
        const path = `${this.ROOT_PATH}/reset-password`;
        const response = await this.http.patch(path, JSON.stringify({ newPassword }));

        if (!response.ok) {
            await this.handleError(response);
        }
    }
    /** resets user password by email confirmation */
    public async sendEmailForPasswordReset(email: string): Promise<void> {
        const path = `${this.ROOT_PATH}/password/${email}`;
        const response = await this.http.get(path, JSON.stringify(email));

        if (!response.ok) {
            await this.handleError(response);
        }
    }
    /** sends data to register user with velas wallet */
    public async velasRegister(eth_wallet_address: string, access_token: string, expires_at: any): Promise<void> {
        const path = `${this.ROOT_PATH}/velas/register`;
        const response = await this.http.post(path, JSON.stringify({ eth_wallet_address, access_token, expires_at }));

        if (!response.ok) {
            await this.handleError(response);
        }
    }
    /** sends address to get nonce to login user */
    public async velasNonce(address: string): Promise<string> {
        const path = `${this.ROOT_PATH}/velas/nonce?address=${address}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        }
        const status = response.json();

        return status;
    }
    /** sends data to login user with velas wallet */
    public async velasLogin(nonce: string, authResult: any): Promise<void> {
        const path = `${this.ROOT_PATH}/velas/login`;
        const response = await this.http.post(path, JSON.stringify({ ...authResult, nonce }));

        if (!response.ok) {
            await this.handleError(response);
        }
    }

    /** gets token to login user with velas wallet */
    public async velasCsrfToken(): Promise<string> {
        const path = "http://localhost:3002/csrf";
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        }
        const result = await response.json();

        return result.token;
    }

    /** gets creds to fill velas vaclient */
    public async velasVaclientCreds(): Promise<any> {
        const path = `${this.ROOT_PATH}/velas/vaclient`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        }
        const result = await response.json();

        return result;
    }
}
