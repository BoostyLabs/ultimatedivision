// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/** variable describes that user added to gueue */
export const YOU_ADDED_MESSAGE: string = 'you added!';
/** variable describes that user already added to gueue */
export const YOU_HAVE_ALREADY_BEEN_ADDED_MESSAGE: string = 'you have already been added!';
/** variable describes that it needs confirm game from user */
export const YOU_CONFIRM_PLAY_MESSAGE: string = 'you confirm play?';
/** variable describes that was send wrong action from user */
export const WRONG_ACTION_MESSAGE: string = 'wrong action';

/**
 * WebSocketAPICLient is base client that holds webSocket.
 */
export class WebSocketAPICLient {
    private readonly ROOT_PATH: string = 'ws://localhost:8088/api/v0/queue';

    /** The WebSocket object provides the API for creating and managing
     * a WebSocket connection to a server, as well as for sending and 
     * receiving data on the connection. */
    public readonly ws: WebSocket = new WebSocket(this.ROOT_PATH);
};
