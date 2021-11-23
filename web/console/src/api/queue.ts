// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { WebSocketAPICLient } from './webSocketClient';

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
    public async sendAction(action: string, squadId: string): Promise<void> {
        await this.ws.send(JSON.stringify({ action, squadId }));
    };

    /** starts searching match on first open webSocket connection. */
    public async startSearch(action: string, squadId: string): Promise<void> {
        this.ws.onopen = async() => {
            await this.sendAction(action, squadId);
        };
    };
};
