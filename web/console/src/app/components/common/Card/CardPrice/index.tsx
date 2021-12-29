// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Card } from '@/card';

import currency from '@static/img/FootballerCardPage/currency.svg';

import './index.scss';

export const FootballerCardPrice: React.FC<{ card: Card }> = ({ card }) => {
    const priceData = card.cardPrice;

    <h2 className="footballer-card-price__price">
        <>
            {priceData.price}
            <img
                className="footballer-card-price__price-currency"
                src={currency}
                alt="currency img"
            />
        </>
    </h2>;

    return (
        <div className="footballer-card-price">
            <div className="footballer-card-price__wrapper">
                <div className="footballer-card-price__info-area">
                    <div className="footballer-card-price__bid">
                        <div className="bid">
                            <span className="bid__title">Current bid</span>
                            <span className="bid__value">{priceData.price}</span>
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
                    <button className="place-bid">
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
    );
};
