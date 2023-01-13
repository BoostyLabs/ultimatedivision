// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import { toast } from 'react-toastify';
import { ethers } from 'ethers';

// @ts-ignore
import KeyStorageHandler from './keyStorageHandler';
// @ts-ignore
import StorageHandler from './storageHandler';
// @ts-ignore
import { VAClient } from '@velas/account-client';

import { UsersClient } from '@/api/users';
import { UsersService } from '@/users/service';
import { VelasClient } from '@/api/velas';
import { VelasService } from '@/velas/service';

import { TransactionIdentificators } from '@/ethers';
import { buildHash } from '@/app/internal/ethers';
import { ToastNotifications } from '@/notifications/service';

/** TODO: change to real data */
const GAS = 50000;
const GAS_PRICE = 2000000001;

const usersClient = new UsersClient();
const usersService = new UsersService(usersClient);

const velasClient = new VelasClient();
const velasService = new VelasService(velasClient);

const vaclientService = async() => {
    try {
        const creds = await velasService.vaclientCreds();

        const vaclient = new VAClient({
            mode: 'redirect',
            clientID: creds.clientId,
            redirectUri: creds.redirectUri,
            StorageHandler,
            KeyStorageHandler,
            accountProviderHost: creds.accountProviderHost,
            networkApiHost: creds.networkApiHost,
            transactionsSponsorApiHost: creds.transactionsSponsorApiHost,
            transactionsSponsorPubKey: creds.transactionsSponsorPubKey,
        });

        return vaclient;
    } catch (e) {
        toast.error('Something went wrong', {
            position: toast.POSITION.TOP_RIGHT,
            theme: 'colored',
        });
    }

    return null;
};

/** VelasTransactionService describes velas transaction entity. */
class VelasTransactionService {
    private readonly gas = GAS;
    private readonly gasPrice = GAS_PRICE;
    private readonly client = new VelasClient();
    public walletAddress: string = '';
    public provider: any = '';
    public from: any = '';
    public vaclient: any = '';
    public nonce: string = '';
    public storage: any;

    /** default VelasTransactionService implementation */
    constructor(from: string, vaclient: string) {
        this.from = from;
        this.vaclient = vaclient;
    }

    /** Sets provider. */
    async setProvider() {
        const vaclient = await vaclientService();
        await vaclient.defaultAccount(JSON.parse(this.vaclient));
        const provider = new ethers.providers.Web3Provider(vaclient.provider);

        this.provider = provider;
    }

    /** Gets wallet adress from api */
    async getWalletAdress() {
        const user = await usersService.getUser();
        this.walletAddress = user.wallet;
    }

    /** Gets transaction from api */
    async getTransaction(signature: TransactionIdentificators): Promise<any> {
        return await this.client.getTransaction(signature);
    }

    /** Sends smart contract transaction. */
    async sendTansaction(cardId: string) {
        try {
            await this.setProvider();
            await this.getWalletAdress();

            const csrfToken = await velasService.csrfToken();

            const nonce = await this.provider.getTransactionCount(this.from, 'latest');

            const transaction = await this.getTransaction(new TransactionIdentificators(this.walletAddress, cardId));

            // /* eslint-disable */
            // const data = `${transaction.nftCreateContract.mintWithSignatureSelector}${buildHash(40)}${buildHash(transaction.tokenId.toString(16))}${buildHash(60)}${buildHash(
            //     transaction.password.slice(-2))}${transaction.password.slice(0, transaction.password.length - 2)}`;

            const raw = {
                nonce: ethers.utils.hexlify(nonce),
                to: transaction.nftCreateContract.address,
                from: this.from,
                gas: ethers.utils.hexlify(this.gas),
                gasPrice: ethers.utils.hexlify(this.gasPrice),
                chainId: transaction.nftCreateContract.chainId,
                broadcast: true,
                /* eslint-disable */
                csrf_token: csrfToken,
            };

            await this.provider.send('eth_sendTransaction', raw);
        } catch (error: any) {
            ToastNotifications.velasTransactionError(error)
        }
    }
}

export default VelasTransactionService;
