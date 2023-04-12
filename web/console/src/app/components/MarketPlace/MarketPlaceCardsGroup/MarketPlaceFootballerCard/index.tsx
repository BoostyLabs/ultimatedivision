// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Link } from 'react-router-dom';

import { PlayerCard } from '@components/common/PlayerCard';
import { Lot } from '@/marketplace';
import { MarketplaceTimer } from '@components/MarketPlace/MarketplaceTimer';

import './index.scss';

const ONE_COIN = 1;

export const MarketPlaceFootballerCard: React.FC<{ lot: Lot; handleShowModal: (lot: Lot) => void; place?: string }> =
    ({ lot, handleShowModal }) => {
        /** TODO: add function entity */
        const buyNowButton = () => { };
        const bidButton = () => { };

        return <div className="marketplace-playerCard" onClick={() => bidButton()}>
            <Link
                className="marketplace-playerCard__link"
                to={`/lot/${lot.cardId}`}
            >
                <PlayerCard
                    id={lot.cardId}
                    className={'marketplace-playerCard__image'}
                />
            </Link>
            <div className="marketplace-playerCard__info">
                <div className="marketplace-playerCard__text">
                    <p className="marketplace-playerCard__text__info"> Max Bid</p>
                    <span className="marketplace-playerCard__price">
                        {lot.maxPrice} {lot.maxPrice > ONE_COIN ? 'coins' : 'coin'}
                    </span>
                </div>
                <button className="marketplace-playerCard__button" onClick={() => handleShowModal(lot)}>
                    Bid
                </button>
            </div>
            <div className="marketplace-playerCard__info">
                <div className="marketplace-playerCard__text">
                    <p className="marketplace-playerCard__text__info">Current bid</p>
                    <span className="marketplace-playerCard__price">
                        {lot.currentPrice} {lot.currentPrice > ONE_COIN ? 'coins' : 'coin'}
                    </span>
                </div>
                <button className="marketplace-playerCard__button" onClick={() => buyNowButton()}>
                    Buy now
                </button>
            </div>
            <MarketplaceTimer lot={lot} className="marketplace-playerCard__timer" />
        </div>;
    };


