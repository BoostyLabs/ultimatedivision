// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import MetaMaskOnboarding from '@metamask/onboarding';

import { metamaskNotifications } from '@/app/internal/notifications';
import { ServicePlugin } from '@/app/plugins/service';

import CasperTransactionService from '@/app/casper';
import { User } from '@/users';
import { walletTypes } from './';

/**
 * Exposes all minting service related logic.
 */
class MintingService {
    public user: User = new User();
    public metamaskService = ServicePlugin.create();
    public onboarding = new MetaMaskOnboarding();

    /** default MintingService implementation */
    constructor(user: User, onboarding: MetaMaskOnboarding) {
        this.user = user;
        this.onboarding = onboarding;
    }

    /** Mints chosed card with metamask */
    private async metamaskMint(id: string) {
        if (MetaMaskOnboarding.isMetaMaskInstalled()) {
            try {
                // @ts-ignore .
                await window.ethereum.request({
                    method: 'eth_requestAccounts',
                });
                await this.metamaskService.sendTransaction(id);
            } catch (error: any) {
                metamaskNotifications(error);
            }
        } else {
            this.onboarding.startOnboarding();
        }
    };

    /** Mints chosed card with casper */
    private casperMint(id: string) {
        const casperTransactionService = new CasperTransactionService(this.user.casperWalletId);

        casperTransactionService.mint(id);
    };

    /** Mints chosed card with velas */
    private static velasMint() { };

    /** Mints chosed card. */
    public mintNft(id: string) {
        switch (this.user.walletType) {
        case walletTypes.VELAS_WALLET_TYPE:
            MintingService.velasMint();
            break;
        case walletTypes.CASPER_WALLET_TYPE:
            this.casperMint(id);
            break;
        case walletTypes.METAMASK_WALLET_TYPE:
            this.metamaskMint(id);
            break;
        default:
            break;
        }
    }

    /** Mints token with metamask wallet. */
    private metamaskMintToken(messageEvent: any) {
        this.metamaskService.mintUDT(messageEvent.message.transaction);
    };

    /** Mints token with casper wallet. */
    private casperMintToken(messageEvent: any) {
        const casperTransactionService = new CasperTransactionService(this.user.casperWalletId);

        casperTransactionService.mintUDT(messageEvent.message.casperTransaction, messageEvent.message.rpcNodeAddress);
    };

    /** Mints token with velas wallet. */
    private static velasMintToken() { };

    /** Mints token. */
    public mintToken(messageEvent: any) {
        switch (this.user.walletType) {
        case walletTypes.VELAS_WALLET_TYPE:
            MintingService.velasMintToken();
            break;
        case walletTypes.CASPER_WALLET_TYPE:
            this.casperMintToken(messageEvent);
            break;
        case walletTypes.METAMASK_WALLET_TYPE:
            this.metamaskMintToken(messageEvent);
            break;
        default:
            break;
        }
    };
}

export default MintingService;
