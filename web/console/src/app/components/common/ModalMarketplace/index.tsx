// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch, SetStateAction, useState } from 'react';

import { MarketplaceClient } from '@/api/marketplace';
import { Marketplaces } from '@/marketplace/service';
import { Lot } from '@/marketplace';
import { PlayerCard } from '../PlayerCard';
import { MarketplaceTimer } from '../../MarketPlace/MarketplaceTimer';

import CloseModal from '@/app/static/img/MarketPlacePage/marketplaceModal/close.svg';

import './index.scss';

const ONE_COIN = 1;

export const ModalMarketPlace: React.FC<{ lot: Lot; setShowModal: Dispatch<SetStateAction<boolean>> }> =
    ({ lot, setShowModal }) => {
        const [cardBid, setCardBid] = useState<number>(lot.currentPrice);

        const marketplaceClient = new MarketplaceClient();
        const marketplaceService = new Marketplaces(marketplaceClient);

        const bidButton = () => {
            marketplaceService.placeBid(lot.cardId, cardBid);
        };

        /** TODO: add function entity */
        const buyNowButton = () => {};

        const onChangeCardBid = (e: React.ChangeEvent<HTMLInputElement>) => {
            Number(e.target.value) > lot.maxPrice?
                setCardBid(lot.maxPrice)
                :
                setCardBid(Number(e.target.value));
        };

        return <div className="marketplace-modal">
            <div className="marketplace-modal__wrapper">
                <div>
                    <button onClick={() => setShowModal(false)} className="marketplace-modal__close-button">
                        <img src={CloseModal} alt="modal close button" className="marketplace-modal__close-button__img" />
                    </button>
                </div>
                <h2 className="marketplace-modal__title">{lot.card.playerName}</h2>
                <div className="marketplace-modal__content">
                    <PlayerCard id={lot.card.id} className="marketplace-modal__card" />
                    <div className="marketplace-modal__lot">
                        <div className="marketplace-modal__bid">
                            <input className="marketplace-modal__bid__input"
                                placeholder="place a bid|"
                                type="number" max={lot.maxPrice}
                                onChange={onChangeCardBid}
                                value={cardBid}
                            />

                            <button onClick={() => bidButton()} className="marketplace-modal__button marketplace-modal__button__bid">
                                bid
                            </button>
                            <span className="marketplace-modal__bid__label">or</span>
                        </div>
                        <div className="marketplace-modal__buy-now" >
                            <MarketplaceTimer lot={lot} className="marketplace-modal__timer"/>
                            <div className="marketplace-modal__buy-now__label__mobile">
                                    for
                                <span>
                                    {lot.currentPrice} {lot.currentPrice > ONE_COIN ? 'coins' : 'coin'}
                                </span>
                            </div>
                            <button onClick={() => buyNowButton()} className="marketplace-modal__button marketplace-modal__button__buy-now">
                                buy now
                                <div className="marketplace-modal__buy-now__label">
                                    for
                                    <span>
                                        {lot.currentPrice} {lot.currentPrice > ONE_COIN ? 'coins' : 'coin'}
                                    </span>
                                </div>
                            </button>
                        </div>
                        <div className="marketplace-modal__timer__mobile">
                            <MarketplaceTimer lot={lot} className="marketplace-modal__timer"/>
                        </div>
                    </div>
                </div>
            </div>
        </div>;
    };


