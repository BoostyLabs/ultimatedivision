// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { WSConnectionClient } from '../api/connection';

/** Exposes all queue related logic. */
export class WSConnectionService {
    public wsConnectionClient: WSConnectionClient = new WSConnectionClient();

    /** Changes current conection client. */
    public changeWSConnectionClient() {
        this.wsConnectionClient = new WSConnectionClient();
    };

    /** Sends action that indicates that the client allows to add address of wallet. */
    public actionAllowAddress(wallet: string, nonce: number): void {
        this.wsConnectionClient.actionAllowAddress(wallet, nonce);
    };

    /** Sends action that indicates that the client allows to add address of wallet. */
    public casperActionAllowAddress(wallet: string, walletType: string, squadId: string): void {
        this.wsConnectionClient.casperActionAllowAddress(wallet, walletType, squadId);
    };

    /** Sends action that indicates that the client is forbidden to add wallet address. */
    public actionForbidAddress(): void {
        this.wsConnectionClient.actionForbidAddress();
    };

    /** Sends action to confirm or reject match. */
    public sendAction(action: string, squadId: string): void {
        this.wsConnectionClient.sendAction(action, squadId);
    };

    /** Sends action, i.e 'startSearch', 'finishSearch', on open webSocket connection. */
    public onOpenConnectionSendAction(action: string, squadId: string): void {
        this.wsConnectionClient.onOpenConnectionSendAction(action, squadId);
    };
    /** Closes ws connection. */
    public close() {
        this.wsConnectionClient.close();
    };
    /** Opens ws connection. */
    public openConnection() {
        this.wsConnectionClient.openConnection();
    };

    /** Sets match queue */
    public matchQueue() {
        this.wsConnectionClient.matchQueue();
    };
};

const queueService = new WSConnectionService();

/** Sends action to confirm or reject match. */
export const queueSendAction = (action: string, squadId: string) => {
    queueService.sendAction(action, squadId);
};

/** Changes current queue client, and after sends action,
 * i.e 'startSearch', 'finishSearch', on open webSocket connection. */
export const onOpenConnectionSendAction = (action: string, squadId: string) => {
    queueService.changeWSConnectionClient();
    queueService.onOpenConnectionSendAction(action, squadId);
};

/** Sends action that indicates that the client allows to add address of wallet. */
export const queueActionAllowAddress = (wallet: string, nonce: number) => {
    queueService.actionAllowAddress(wallet, nonce);
};

/** Sets match queue */
export const setMatchQueue = () => {
    queueService.matchQueue();
};

/** Sends action that indicates that the client allows to add address of wallet. */
export const queueCasperActionAllowAddress = (wallet: string, walletType: string, squadId: string) => {
    queueService.casperActionAllowAddress(wallet, walletType, squadId);
};

/** Sends action that indicates that the client allows to add address of wallet. */
export const actionForbidAddress = () => {
    queueService.actionForbidAddress();
};

/** Returns current queue client. */
export const getCurrentQueueClient = () => queueService.wsConnectionClient;

/** Opens ws connection. */
export const onOpenConnection = () => queueService.openConnection();
