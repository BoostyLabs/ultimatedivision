// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/** Transaction props from api */
export interface Transaction {
    password: string;
    tokenId: number;
    contract: {
        address: string;
        addressMethod: string;
    };
}

/** Abi input interface */
interface Input {
    internalType: string;
    name: string;
    type: string;
}

/** Abi block interface */
interface Abi {
    inputs: Input[];
    outputs: Input[];
    name: string;
    stateMutability: string;
    type: string;
}

/** Smart conract interface */
export const GAME_ABI: Partial<Abi>[] = [
    {
        'inputs': [
            {
                'internalType': 'address',
                'name': '_nftAddress',
                'type': 'address',
            },
        ],
        'stateMutability': 'nonpayable',
        'type': 'constructor',
    },
    {
        'inputs': [
            {
                'internalType': 'bytes',
                'name': '_signature',
                'type': 'bytes',
            },
            {
                'internalType': 'uint256',
                'name': 'tokenID',
                'type': 'uint256',
            },
        ],
        'name': 'buyWithSignature',
        'outputs': [],
        'stateMutability': 'nonpayable',
        'type': 'function',
    },
    {
        'inputs': [],
        'name': 'nft',
        'outputs': [
            {
                'internalType': 'contract INFT',
                'name': '',
                'type': 'address',
            },
        ],
        'stateMutability': 'view',
        'type': 'function',
    },
    {
        'inputs': [],
        'name': 'verifyAddress',
        'outputs': [
            {
                'internalType': 'address',
                'name': '',
                'type': 'address',
            },
        ],
        'stateMutability': 'view',
        'type': 'function',
    },
];
