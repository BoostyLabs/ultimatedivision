// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

export class HttpClient {
    /* do sends an HTTP request and returns
    * an HTTP response as configured on the client. */
    private async do(
        method: string,
        path: string,
        body: string | null
    ): Promise<Response> {
        const request: RequestInit = {
            method: method,
            body: body,
        };

        request.headers = {
            'Accept': 'application/json',
            'Content-type': 'application/json',
        };

        return await fetch(path, request);
    };
    public async post(path: string, body: string | null): Promise<Response> {
        return await this.do('POST', path, body);
    };
    public async get(path: string): Promise<Response> {
        return await this.do('GET', path, null);
    };
    public async put(path: string, body: string | null) {
        return await this.do('PUT', path, '');
    };
    public async delete(path: string): Promise<Response> {
        return await this.do('DELETE', path, '');
    };
    public async patch(path: string, body: string): Promise<Response> {
        return await this.do('PATCH', path, body);
    }
};