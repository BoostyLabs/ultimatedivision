import { VAClient }  from '@velas/account-client';

import StorageHandler    from './storageHandler';
import KeyStorageHandler from './keyStorageHandler';

export const agent = {}

export const vaclient = new VAClient({
    mode:        'redirect',
    clientID:    '48yTQiBWjiyifDp6fesNj72gosALuTvxLQ8Rqzy6sRwh',
    redirectUri: 'http://localhost:8088/auth-velas',
    StorageHandler,
    KeyStorageHandler,
    accountProviderHost:        'account.testnet.velas.com',
    networkApiHost:             'https://api.testnet.velas.com',
    transactionsSponsorApiHost: 'http://localhost:8088',
    transactionsSponsorPubKey:  '3WZuAbAnD8eNpmPaUjeS7ZP6aBFmNeeHA3xv9XcLszvq1o4dQU8evRCMfmD6ei1gz6y6ZCayfmiYZXL7sezTsbLN',

});

export const vaclient_wrong = new VAClient({
    mode:        'redirect',
    clientID:    'wrong',
    redirectUri: 'http://localhost:8088/auth-velas/',
    StorageHandler,
    KeyStorageHandler,
    accountProviderHost:       'account.testnet.velas.com',
    networkApiHost:            'https://api.testnet.velas.com',
    transactionsSponsorApiHost: 'http://localhost:8088',
    transactionsSponsorPubKey:  '3WZuAbAnD8eNpmPaUjeS7ZP6aBFmNeeHA3xv9XcLszvq1o4dQU8evRCMfmD6ei1gz6y6ZCayfmiYZXL7sezTsbLN',
});

