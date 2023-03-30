// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';

import { Lot } from '@/marketplace';
import { PlaceBid } from '@components/common/Card/popUps/PlaceBid';

import './index.scss';

/** Initial bid value. */
const INITIAL_BID: number = 0;

export const BidArea: React.FC<{ lot: Lot }> = (lot) => {
    const [isOpenPlaceBidPopup, setIsOpenPlaceBidPopup] = useState<boolean>(false);
    const [cardBid, setCardBid] = useState<number>(INITIAL_BID);
    const [currentCardBid, setCurrentCardBid] = useState<number>(lot.lot.currentPrice);

    /** Handle opening of a place bids pop-up. */
    const handleOpenPlaceBidPopup = () => {
        setIsOpenPlaceBidPopup(true);
    };

    return <div className="footballer-card-price">
        {isOpenPlaceBidPopup &&
            <PlaceBid
                setCurrentCardBid={setCurrentCardBid}
                setIsOpenPlaceBidPopup={setIsOpenPlaceBidPopup}
                setCardBid={setCardBid}
                cardBid={cardBid}
            />
        }
        <div className="footballer-card-price__wrapper">
            <div className="footballer-card-price__info-area">
                <div className="footballer-card-price__bid">
                    <div className="bid">
                        <span className="bid__title">Current bid:</span>
                        <span className="bid__value">{currentCardBid}</span>
                    </div>
                </div>
                <div className="footballer-card-price__auction">
                    <span className="auction-title">
                        Auction expires in:
                    </span>
                    <span className="auction-expire-time">22:12:03</span>
                </div>
            </div>
            <div className="footballer-card-price__buttons">
                <button className="place-bid" onClick={handleOpenPlaceBidPopup}>
                    <span className="place-bid__text">Plase a bid</span>
                </button>
                <button className="buy-now">
                    <span className="buy-now__text">Buy now</span>
                    <span className="buy-now__value">{currentCardBid} COIN</span>
                </button>
            </div>
        </div>
    </div>;
};
