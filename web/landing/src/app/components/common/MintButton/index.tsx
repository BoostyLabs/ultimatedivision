// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React, { useState } from 'react';

import MetaMaskOnboarding from '@metamask/onboarding';

import './index.scss';

declare global {
    interface Window {
        ethereum: any;
    }
}

export const MintButton: React.FC = () => {
    const onboarding = React.useRef<MetaMaskOnboarding>();
    const [connectError, handleError] = useState(false);
    const [text, setButtonText] = useState('Mint');

    React.useEffect(() => {
        if (!onboarding.current) {
            onboarding.current = new MetaMaskOnboarding();
        }
    }, []);

    const connect = async () => {

        if (MetaMaskOnboarding.isMetaMaskInstalled()) {
            try {
                const request = await window.ethereum.request({ method: 'eth_requestAccounts' });
                setButtonText('Connected');
                return request
            } catch (error: any) {
                handleError(true);
                console.log(error.message);
                setTimeout(() => {
                    handleError(false)
                }, 5000);
            }

        } else {
            onboarding.current?.startOnboarding();
        }
    }

    return (
        <button className="ultimatedivision-mint-btn"
            data-aos="fade-right"
            data-aos-duration="600"
            data-aos-easing="ease-in-out-cubic"
            onClick={connect}
        >
            {
                connectError &&
                <span className="error">Please open metamask manually!</span>
            }
            <span className="ultimatedivision-mint-btn__text">{text}</span>
        </button>
    );
};