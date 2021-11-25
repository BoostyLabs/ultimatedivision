// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/**
 * QueueClient is a ws implementation of users API.
 * Exposes queue-related functionality.
 */
export class QueueClient {
    public ws: WebSocket;
    /** The WebSocket provides the API for creating and managing
     * a websocket connection to a server and for sending and
     * receiving data on the connection. */
    public constructor(path: string) {
        this.ws = new WebSocket(`ws://${path}/api/v0/queue`);
    }
    /** sends action to confirm and reject match, finish search */
    public sendAction(action: string, squadId: string) {
        this.ws.send(JSON.stringify({ action, squadId }));
    };

    /** starts searching match on first open webSocket connection. */
    public startSearch(action: string, squadId: string) {
        this.ws.onopen = () => {
            this.sendAction(action, squadId);
        };
    };
};
