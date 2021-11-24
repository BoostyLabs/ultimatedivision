// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { ethers } from 'ethers';

import { Service } from '@/app/ethers/service';

/** Class for creating ethers service */
export class ServicePlugin {
    /** Creates ethers provider instance */
    public static create() {
        const provider = window.ethereum ? new ethers.providers.Web3Provider(window.ethereum) : null;

        return new Service(provider);
    }
}
