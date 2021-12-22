// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useMemo, useState } from 'react';
import { useSelector } from 'react-redux';
import MetaMaskOnboarding from '@metamask/onboarding';
import { toast } from 'react-toastify';

import coin from '@static/img/match/money.svg';

import { QueueClient } from '@/api/queue';
import { RootState } from '@/app/store';
import { ServicePlugin } from '@/app/plugins/service';
import { actionForbidAddress, getCurrentQueueClient, queueActionAllowAddress } from '@/queue/service';

import './index.scss';

export const MatchScore: React.FC = () => {
    const [queueClient, setQueueClient] = useState<QueueClient | null>(null);

    const [wallet, setWallet] = useState<string | null>(null);

    const onboarding = useMemo(() => new MetaMaskOnboarding(), []);
    const service = ServicePlugin.create();

    const { matchResults, transaction } = useSelector((state: RootState) => state.matchesReducer.gameResult);

    const { question } = useSelector((state: RootState) => state.matchesReducer);

    /** FIRST_TEAM_INDEX is variable that describes first team index in teams array. */
    const FIRST_TEAM_INDEX: number = 0;
    /** SECOND_TEAM_INDEX is variable that describes second team index in teams array. */
    const SECOND_TEAM_INDEX: number = 1;

    /** Variable describes that it needs alllow to add address or forbid add adress. */
    const CONFIRM_ADD_WALLET: string = 'you allow us to take your address?';

    /** Returns metamask wallet address for earning reward */
    const addWallet = async () => {
        /** Code which indicates that 'eth_requestAccounts' already processing */
        const METAMASK_RPC_ERROR_CODE = -32002;
        if (MetaMaskOnboarding.isMetaMaskInstalled()) {
            try {
                // @ts-ignore
                await window.ethereum.request({
                    method: 'eth_requestAccounts',
                });

                toast.success('Please, add your wallet to get coins!', {
                    position: toast.POSITION.TOP_RIGHT,
                    theme: 'light',
                });

                const wallet = await service.getWallet();

                const currentQueueClient = getCurrentQueueClient();

                setQueueClient(currentQueueClient);

                setWallet(wallet);
            } catch (error: any) {
                error.code === METAMASK_RPC_ERROR_CODE
                    ? toast.error('Please open metamask manually!', {
                        position: toast.POSITION.TOP_RIGHT,
                        theme: 'colored',
                    })
                    : toast.error('Something went wrong', {
                        position: toast.POSITION.TOP_RIGHT,
                        theme: 'colored',
                    });
            }
        } else {
            onboarding.startOnboarding();
        }
    };

    queueClient.ws.onmessage = ({ data }: MessageEvent) => {
        const messageEvent = JSON.parse(data);
        console.log(messageEvent);
        switch (messageEvent.message) {
            case CONFIRM_ADD_WALLET:
                wallet ? queueActionAllowAddress(wallet) : actionForbidAddress();
                return;
            default:
                return
        }
    };

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
                <div className="match__score__board__coins">
                    <img
                        className="match__score__board__coins-image"
                        src={coin}
                        alt="Coin"
                    />
                    <span className="match__score__board__coins-value">
                        {transaction.value}
                    </span>
                    {question === CONFIRM_ADD_WALLET && <button
                        className="match__score__board__coins__btn"
                        onClick={addWallet}
                    >
                        <span className="match__score__board__coins__btn-text">
                            GET
                        </span>
                    </button>}
                </div>
            </div>
        </div>
    );
};
