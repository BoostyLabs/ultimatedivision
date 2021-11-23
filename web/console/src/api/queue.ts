// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { WebSocketAPICLient } from './webSocketClient';

/**
 * QueueClient is a ws implementation of users API.
 * Exposes queue-related functionality.
 */
export class QueueClient extends WebSocketAPICLient {
    public eventMessage: string = 'wrong action';

    /** changes event Message after ws response */
    public changeEventMessage(eventMessage: string) {
        this.eventMessage = eventMessage;
    };

    /** returns current event message */
    public getEventMessage() {
        return this.eventMessage;
    };

    /** sends action to confirm match, finish search and others. */
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
