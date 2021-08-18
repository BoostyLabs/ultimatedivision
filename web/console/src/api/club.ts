// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '@/api/index';

export class ClubClient extends APIClient {
    // To do: add exist endpoint for club
    private readonly ROOT_PATH: string = '/api/v0/cards';

    public async getCards() {
        return this.http.get(this.ROOT_PATH);
    }
    public async buyCard(param: string) {
        return this.http.post(this.ROOT_PATH, param);
    }
}


