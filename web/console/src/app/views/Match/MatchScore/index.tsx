// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useMemo, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import MetaMaskOnboarding from '@metamask/onboarding';

import { WebSocketClient } from '@/api/websockets';
import { UDT_ABI } from '@/ethers';
import { RootState } from '@/app/store';
import { ServicePlugin } from '@/app/plugins/service';
import { actionAllowAddress, casperActionAllowAddress, getCurrentWebSocketClient } from '@/webSockets/service';
import { setCurrentUser } from '@/app/store/actions/users';
import WalletService from '@/wallet/service';
import { walletTypes } from '@/wallet';
import { ToastNotifications } from '@/notifications/service';
import { MarketplaceClient } from '@/api/marketplace';
import { Marketplaces } from '@/marketplace/service';

import coin from '@static/img/match/money.svg';

import './index.scss';
import { CasperTransactionApprove } from '@/casper/types';
import { CasperNetworkClient } from '@/api/casper';
import CasperTransactionService from '@/casper';
import { CasperNetworkService } from '@/casper/service';

export const MatchScore: React.FC = () => {
    const dispatch = useDispatch();

    const [webSocketClient, setWebSocketClient] = useState<WebSocketClient | null>(null);

    const onboarding = useMemo(() => new MetaMaskOnboarding(), []);

    const service = ServicePlugin.create();

    const { squad } = useSelector((state: RootState) => state.clubsReducer.activeClub);

    const { matchResults, transaction } = useSelector((state: RootState) => state.matchesReducer);

    const { question } = useSelector((state: RootState) => state.matchesReducer);

    const user = useSelector((state: RootState) => state.usersReducer.user);

    /** FIRST_TEAM_INDEX is variable that describes first team index in teams array. */
    const FIRST_TEAM_INDEX: number = 0;
    /** SECOND_TEAM_INDEX is variable that describes second team index in teams array. */
    const SECOND_TEAM_INDEX: number = 1;

    /** Variable describes that it needs alllow to add address or forbid add adress. */
    const CONFIRM_ADD_WALLET: string = 'do you allow us to take your address?';

    const casperClient = new CasperNetworkClient();
    const casperService = new CasperNetworkService(casperClient);

    /** sets user info */
    async function setUser() {
        try {
            await dispatch(setCurrentUser());
        } catch (error: any) {
            ToastNotifications.couldNotGetUser();
        }
    }

    /** Adds metamask wallet address for earning reward. */
    const addMetamaskWallet = async() => {
        /** Code which indicates that 'eth_requestAccounts' already processing */
        const METAMASK_RPC_ERROR_CODE = -32002;
        if (MetaMaskOnboarding.isMetaMaskInstalled()) {
            try {
                // @ts-ignore .
                await window.ethereum.request({
                    method: 'eth_requestAccounts',
                });

                const wallet = await service.getWallet();

                const currentQueueClient = getCurrentWebSocketClient();

                const nonce = await service.getNonce(transaction.udtContract.address, UDT_ABI);

                setWebSocketClient(currentQueueClient);

                actionAllowAddress(wallet, nonce);
            } catch (error: any) {
                ToastNotifications.metamaskError(error);
            }
        } else {
            onboarding.startOnboarding();
        }
    };

    /** Adds casper wallet address for earning reward. */
    const addCasperWallet = () => {
        try {
            const currentWebSocketClient = getCurrentWebSocketClient();

            setWebSocketClient(currentWebSocketClient);

            casperActionAllowAddress(user.casperWalletHash, user.walletType, squad.id);
        }
        catch (error: any) {
            ToastNotifications.couldNotAddCasperWallet();
        }
    };

    /** Adds velas wallet address for earning reward. */
    const addVelasWallet = async() => { };

    /** Adds wallets addresses for earning reward. */
    const addWallet = () => {
        const mintingTokens = new Map();

        const mintingTokenTypes = [
            {
                walletType: walletTypes.VELAS_WALLET_TYPE,
                mint: addVelasWallet,
            },
            {
                walletType: walletTypes.CASPER_WALLET_TYPE,
                mint: addCasperWallet,
            },
            {
                walletType: walletTypes.METAMASK_WALLET_TYPE,
                mint: addMetamaskWallet,
            },
        ];

        mintingTokenTypes.forEach(mintingTokenType =>
            mintingTokens.set(mintingTokenType.walletType, mintingTokenType.mint));

        mintingTokens.get(user.walletType)();
    };

    const approve = async() => {
        const approveData = await casperService.approve();

        const walletService = new WalletService(user);
        await walletService.approveTokenReward(approveData);
    };

    if (webSocketClient) {
        webSocketClient.ws.onmessage = async({ data }: MessageEvent) => {
            const messageEvent = JSON.parse(data);

            const walletService = new WalletService(user);
            await walletService.mintToken(messageEvent);
        };
    }

    useEffect(() => {
        setUser();
    }, []);

    return (
        <div className="match__score">
            <div className="match__score__board">
                <div className="match__score__board__gradient"></div>
                <div className="match__score__board__timer">90:00</div>
                <div className="match__score__board__result">
                    <div className="match__score__board-team-1">
                        {matchResults[FIRST_TEAM_INDEX].quantityGoals}
                    </div>
                    <div className="match__score__board-dash">-</div>
                    <div className="match__score__board-team-2">
                        {matchResults[SECOND_TEAM_INDEX].quantityGoals}
                    </div>
                </div>
                {question === CONFIRM_ADD_WALLET && <div className="match__score__board__coins">
                    <img
                        className="match__score__board__coins-image"
                        src={coin}
                        alt="Coin"
                    />
                    <span className="match__score__board__coins-value">
                        {transaction.value}
                    </span>
                    <button
                        className="match__score__board__coins__btn"
                        onClick={addWallet}
                    >
                        <span className="match__score__board__coins__btn-text">
                            GET
                        </span>
                    </button>
                    <button
                        className="match__score__board__coins__btn"
                        onClick={() => approve()}
                    >
                        <span className="match__score__board__coins__btn-text">
                            approve
                        </span>
                    </button>
                </div>
                }
            </div>
        </div>
    );
};
