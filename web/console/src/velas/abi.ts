// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

export const ABI = [
    {
        'inputs': [
            {
                'internalType': 'address',
                'name': '_nftAddress',
                'type': 'address',
            },
            {
                'internalType': 'address',
                'name': '_ethReceiver',
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
        'stateMutability': 'payable',
        'type': 'function',
    },
    {
        'inputs': [],
        'name': 'ethReceiver',
        'outputs': [
            {
                'internalType': 'address payable',
                'name': '',
                'type': 'address',
            },
        ],
        'stateMutability': 'view',
        'type': 'function',
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
        'name': 'mintWithSignature',
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
