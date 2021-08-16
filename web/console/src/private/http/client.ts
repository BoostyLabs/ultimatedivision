// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/**
 * HttpClient is a custom wrapper around fetch api.
 * Exposes get, post and delete methods for JSON strings.
 */
export class HttpClient {

    public async post(path: string, body: string | null): Promise<Response> {
        return this.do('POST', path, body);
    }

    public async patch(path: string, body: string | null): Promise<Response> {
        return this.do('PATCH', path, body);
    }

    public async put(path: string, body: string | null, _auth = true): Promise<Response> {
        return this.do('PUT', path, body);
    }

    public async get(path: string, _auth = true): Promise<Response> {
        return this.do('GET', path, null);
    }

    public async delete(path: string, _auth = true): Promise<Response> {
        return this.do('DELETE', path, null);
    }

    private async do(method: string, path: string, body: string | null): Promise<Response> {
        const request: RequestInit = {
            method: method,
            body: body,
        };

        request.headers = {
            'Content-Type': 'application/json',
        };

        return await fetch(path, request);
    }
}