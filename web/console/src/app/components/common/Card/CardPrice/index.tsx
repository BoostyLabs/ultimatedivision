// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';

import { PlaceBid } from '../popUps/PlaceBid';
import { Sell } from '../popUps/Sell';

import currency from '@static/img/FootballerCardPage/currency.svg';

import { Card } from '@/card';

import './index.scss';

/** Initial bid value. */
const INITIAL_BID: number = 0;
/** Initial TEST current bid value. */
// TODO: Waiting for backend.
const INITIAL_CURRENT_BID: number = 400;

export const FootballerCardPrice: React.FC<{ card: Card }> = ({ card }) => {
    const [isOpenPlaceBidPopup, setIsOpenPlaceBidPopup] = useState<boolean>(false);
    const [isOpenSellPopup, setIsOpenSellPopup] = useState<boolean>(false);
    const [cardBid, setCardBid] = useState<number>(INITIAL_BID);
    const [currentCardBid, setCurrentCardBid] = useState<number>(INITIAL_CURRENT_BID);

    /** Handle opening of a place bids pop-up. */
    const handleOpenPlaceBidPopup = () => {
        setIsOpenPlaceBidPopup(true);
    };

    /** Handle opening of a selles pop-up. */
    const handleOpenSellPopup = () => {
        setIsOpenSellPopup(true);
    };

    //TODO: Mock data (auction not yet completed).
    const isCardSold: boolean = false;

    return (
        <>
            {isOpenSellPopup && <Sell setIsOpenSellPopup={setIsOpenSellPopup} />}
            {isCardSold
                ?
                <div className="footballer-card-price">
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
                                    <span className="bid__title">Current bid</span>
                                    <span className="bid__value">{currentCardBid}</span>
                                </div>
                                <div className="footballer-card-price__bid__coin">
                                    <img
                                        className="footballer-card-price__price-currency"
                                        src={currency}
                                        alt="currency img"
                                    />
                                </div>
                            </div>
                            <div className="footballer-card-price__auction">
                                <span className="auction-title">
                                Auction expires in
                                </span>
                                <span className="auction-expire-time">22:12:03</span>
                            </div>
                        </div>
                        <div className="footballer-card-price__buttons">
                            <button className="place-bid" onClick={handleOpenPlaceBidPopup}>
                                <span className="place-bid__text">Plase a bid</span>
                            </button>
                            <button className="buy-now">
                                <span className="buy-now__text">Buy now price</span>
                                <img
                                    className="footballer-card-price__price-currency"
                                    src={currency}
                                    alt="currency img"
                                />
                                <span className="buy-now__value">1000</span>
                            </button>
                        </div>
                        <div className="footballer-card-price__last">
                            <span className="title">Last buy price:</span>
                            <span className="value">228'000 </span>
                            <img
                                className="footballer-card-price__price-currency"
                                src={currency}
                                alt="currency img"
                            />
                        </div>
                    </div>
                </div>
                :
                <button
                    className="card__sell-btn"
                    // TODO: Waiting for logic from backend.
                    // onClick={handleOpenSellPopup}
                >
                    <span className="card__sell-btn__text">SELL</span>
                </button>
            }
        </>
    );
};
