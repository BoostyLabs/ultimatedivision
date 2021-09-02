// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '@/api/index';

/** ClubClient base implementation */
export class CardClient extends APIClient {
    /** To do: add exist endpoint for club */
    private readonly ROOT_PATH: string = '/api/v0';

    /** method calls get method from APIClient */
    public async getUserCards() {
        return await this.http.get(`${this.ROOT_PATH}/cards`);
    }
    /** method post for implementing buying cards */
    public async sellCard(id: string) {
        return await this.http.post(`${this.ROOT_PATH}/marketplace/bet`, id);
    }
    /** method calls get method from APIClient */
    public async getSellingCards() {
        return await this.http.get(`${this.ROOT_PATH}/marketplace`);
    }
}
