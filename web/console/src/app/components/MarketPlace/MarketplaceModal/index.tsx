// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch, SetStateAction, useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { MarketplaceClient } from '@/api/marketplace';
import { Marketplaces } from '@/marketplace/service';
import { BidsMakeOfferTransaction } from '@/casper/types';
import { Lot } from '@/marketplace';
import { PlayerCard } from '@components/common/PlayerCard';
import { MarketplaceTimer } from '@components/MarketPlace/MarketplaceTimer';
import { setCurrentUser } from '@/app/store/actions/users';
import { ToastNotifications } from '@/notifications/service';
import WalletService from '@/wallet/service';
import { RootState } from '@/app/store';

import CloseModal from '@/app/static/img/MarketPlacePage/marketplaceModal/close.svg';

import './index.scss';

const ONE_COIN = 1;

export const MarketPlaceModal: React.FC<{ lot: Lot; setShowModal: Dispatch<SetStateAction<boolean>> }> =
    ({ lot, setShowModal }) => {
        const dispatch = useDispatch();

        const [cardBid, setCardBid] = useState<number>(lot.currentPrice);
        const [currentBid, setCurrentBid] = useState<number>(lot.currentPrice);
        const [isEndTime, setIsEndTime] = useState(false);

        const user = useSelector((state: RootState) => state.usersReducer.user);

        const marketplaceClient = new MarketplaceClient();
        const marketplaceService = new Marketplaces(marketplaceClient);

        const onChangeCardBid = (e: React.ChangeEvent<HTMLInputElement>) => {
            Number(e.target.value) > lot.maxPrice ?
                setCardBid(lot.maxPrice)
                :
                setCardBid(Number(e.target.value));
        };

        const bidButton = async() => {
            try {
                await marketplaceService.placeBid(lot.cardId, cardBid);

                const makeOfferData = await marketplaceService.makeOffer(lot.cardId);

                const walletService = new WalletService(user);

                const marketplaceMakeOfferTransaction = new BidsMakeOfferTransaction(
                    makeOfferData.address,
                    makeOfferData.addressNodeServer,
                    makeOfferData.tokenId,
                    makeOfferData.contractHash,
                    cardBid
                );

                await walletService.makeOffer(marketplaceMakeOfferTransaction);

                setCurrentBid(cardBid);
            } catch (e: any) {
                ToastNotifications.somethingWentsWrong();
            }
        };

        /** TODO: add function entity */
        const buyNowButton = async() => {
            const walletService = new WalletService(user);

            await walletService.buyListing({});
        };


        /** sets user info */
        async function setUser() {
            try {
                await dispatch(setCurrentUser());
            } catch (error: any) {
                ToastNotifications.couldNotGetUser();
            }
        }

        useEffect(() => {
            setUser();
        }, []);


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
                            {isEndTime ? <div className="marketplace-modal__timer">0 : 0 : 0 </div> :
                                <MarketplaceTimer setIsEndTime={setIsEndTime} isEndTime={isEndTime} lot={lot} className="marketplace-modal__timer" />
                            }
                            <div className="marketplace-modal__buy-now__label__mobile">
                                for
                                <span>
                                    {currentBid} {currentBid > ONE_COIN ? 'coins' : 'coin'}
                                </span>
                            </div>
                            <button onClick={() => buyNowButton()} className="marketplace-modal__button marketplace-modal__button__buy-now">
                                buy now
                                <div className="marketplace-modal__buy-now__label">
                                    for
                                    <span>
                                        {currentBid} {currentBid > ONE_COIN ? 'coins' : 'coin'}
                                    </span>
                                </div>
                            </button>
                        </div>
                        {isEndTime ? <div className="marketplace-modal__timer__mobile">0 : 0 : 0 </div> :
                            <MarketplaceTimer setIsEndTime={setIsEndTime} isEndTime={isEndTime} lot={lot} className="marketplace-modal__timer__mobile" />}
                    </div>
                </div>
            </div>
        </div>;
    };


