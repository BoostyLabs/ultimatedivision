// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '@/api/index';

/** ClubClient base implementation */
export class CardClient extends APIClient {
    /** To do: add exist endpoint for club */
    private readonly ROOT_PATH: string = '/api/v0/cards';

    /** method calls get method from APIClient */
    public async get() {
        const response = await this.http.get(this.ROOT_PATH);

        return await response.json();
    }
    /** method post for implementing buying cards */
    public async buy(id: string) {
        const response = await this.http.post(this.ROOT_PATH, id);

        return await response.json();
    }
}
