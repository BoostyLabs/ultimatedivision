// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { ethers } from 'ethers';

export class Service {
    private readonly provider;

    public constructor(ethereumProvider: any) {
        this.provider = ethereumProvider;
    }

    public async totalSupply (adress: string, abi: any[]) {
        const contract = await new ethers.Contract(adress, abi);
        const signer = await this.provider.getSigner();
        const connect = await contract.connect(signer);
        const token = await connect.functions.totalSupply();
        return parseInt(token[0]._hex, 16);
    };

    public async contract(adress: string, abi: any[]) {
        const contractSale = await new ethers.Contract(adress, abi);
        const signer = await this.provider.getSigner();
        const connectSale = await contractSale.connect(signer).getCurrentPrice();
        return parseInt(connectSale[0]._hex, 16);
    };

    public async sendTransaction(adress: string, amount: string) {
        try {
            const signer = this.provider.getSigner();

            //throws error when adress is not valid
            ethers.utils.getAddress(adress);

            const transaction = await signer.sendTransaction({
                to: adress,
                value: ethers.utils.parseEther(amount)
            });
        } catch (error: any) {
            /* eslint-disable-next-line */
            console.log(error.message)
        }
    }

    public async getBalance(id: string) {
        try {
            const balance = await this.provider.getBalance(id);

            return ethers.utils.formatEther(balance);
        } catch (error: any) {
            /* eslint-disable-next-line */
            console.log(error.message)
        }
    }
};
