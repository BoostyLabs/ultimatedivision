// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { ethers } from 'ethers';

import { Service } from '@/app/ethers/service';

/** Class for creating ethers service */
export class ServicePlugin {
    /** Creates ethers provider instance */
    public static create() {
        let ethereumProvider = null;
        try {
            //@ts-ignore
            ethereumProvider = new ethers.providers.Web3Provider(window.ethereum);

        } catch (e) {
            return new Service(null);
        }

        return new Service(ethereumProvider);
    }
}
