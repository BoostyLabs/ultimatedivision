// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import coin from '@static/img/MarketPlacePage/MyCard/goldPrice.svg';

import './index.scss';

export const MintingArea: React.FC<{ isInactive: boolean; time: string }> = ({
    isInactive,
    time,
}) =>
    <div className="minting-area">
        <button className="minting-area__button" disabled={isInactive}>
            <span className="minting-area__button-text">MINT</span>
            <span className="minting-area__button-value">
                <img src={coin} alt="coin" />
        100
            </span>
        </button>
        {isInactive &&
            <div className="minting-area__timer">
                 <p className="minting-area__timer__text">RESTOCK IN</p>
                <span className="minting-area__timer__time">{time}</span>
            </div>
        }
    </div>;

