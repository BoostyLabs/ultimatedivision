// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { ethers } from 'ethers';

import { EthersClient } from '@/api/ethers';
import { buildHash } from '../utils/ethers';

export class Service {
    private readonly provider;
    private readonly client = new EthersClient();

    public constructor(ethereumProvider: any) {
        this.provider = ethereumProvider;
    }

    /** Get nft adress and password. */
    public async getAddress(wallet: string, id: string) {
        return await this.client.getAddress(wallet, id);
    }

    /** Get current wallet address. */
    public async getWallet() {
        const signer = await this.provider.getSigner();

        return await signer.getAddress();
    }

    /** Send smart contract transaction. */
    public async sendTransaction(
        wallet: string,
        abi: any[],
        cardId: string,
    ) {
        const signer = await this.provider.getSigner();
        const address = await this.getAddress(wallet, cardId);
        const contract = await new ethers.Contract(
            address.contract.address,
            abi
        );
        const connect = await contract.connect(signer);
        const data = `${address.contract.addressMethod}${buildHash(40)}${buildHash(address.tokenId.toString(16))}${buildHash(60)}${buildHash(
            address.password.slice(-2)
        )}${address.password.slice(0, address.password.length - 2)}`;
        const gasLimit = await signer.estimateGas({
            to: address.contract.address,
            data,
        });

        await signer.sendTransaction({
            to: address.contract.address,
            data,
            gasLimit,
            chainId: 4,
        });
    }

    public async getBalance(id: string) {
        try {
            const balance = await this.provider.getBalance(id);

            return ethers.utils.formatEther(balance);
        } catch (error: any) {
            /* eslint-disable-next-line */
            console.log(error.message);
        }
    }
}
