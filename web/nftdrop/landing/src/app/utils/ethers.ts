// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

export const buildHash = (hashSignature: number | string) => {
    const hashBase = '0000000000000000000000000000000000000000000000000000000000000000';
    return hashBase.slice(0, hashBase.length - hashSignature.toString().length) + hashSignature;
};
