//@ts-nocheck
// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.
import MetaMaskOnboarding from '@metamask/onboarding';
import React, { useEffect } from 'react';

import './index.scss';

type MintButtonProps = {
    text: string,
}

export const MintButton: React.FC<MintButtonProps> = (
    { text }) => {
    const [accounts, setAccounts] = React.useState([]);
    const onboarding = React.useRef<MetaMaskOnboarding>();

    useEffect(() => {
        if (!onboarding.current) {
            onboarding.current = new MetaMaskOnboarding();
        }
    }, []);

    const handleClick = () => {
        if (MetaMaskOnboarding.isMetaMaskInstalled()) {
            window.ethereum
                .request({ method: 'eth_requestAccounts' })
                .then((newAccounts) => setAccounts(newAccounts));
        } else {
            onboarding.current.startOnboarding();
        }
    };

    return (
        <button className="ultimatedivision-mint-btn"
            data-aos="fade-right"
            data-aos-duration="600"
            data-aos-easing="ease-in-out-cubic"
            onClick={handleClick}
        >
            <span className="ultimatedivision-mint-btn__text">{text}</span>
        </button>
    );
};