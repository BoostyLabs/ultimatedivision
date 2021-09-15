// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '@/api/index';

/** ClubClient base implementation */
export class ClubClient extends APIClient {
    private readonly ROOT_PATH: string = '/api/v0';

    /** method calls get method from APIClient */
    public async createClub(): Promise<Response> {
        return await this.http.post(`${this.ROOT_PATH}/clubs`);
    }
    /** method calls get method from APIClient */
    public async getClub(): Promise<Response> {
        return await this.http.get(`${this.ROOT_PATH}/clubs`);
    }
    /** method calls get method from APIClient */
    public async createSquad(clubId: string): Promise<Response> {
        return await this.http.post(`${this.ROOT_PATH}/clubs/${clubId}/squads`);
    }
    /** method calls get method from APIClient */
    public async addCard(path: string, position: {position: number}): Promise<Response> {
        return await this.http.post(`${this.ROOT_PATH}/${path}`, JSON.stringify(position));
    }
    /** method calls get method from APIClient */
    public async changeCardPosition(path: string, position: {position: number}): Promise<Response> {
        return await this.http.patch(`${this.ROOT_PATH}/${path}`, JSON.stringify(position));
    }
    /** method calls get method from APIClient */
    public async deleteCard(path: string): Promise<Response> {
        return await this.http.delete(`${this.ROOT_PATH}/${path}`);
    }
}
