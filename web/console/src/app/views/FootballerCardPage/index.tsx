// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect } from 'react';

import { useDispatch, useSelector } from 'react-redux';
import { RootState } from '@/app/store';
import { openMarketplaceCard, openUserCard } from '@/app/store/actions/cards';

import { FootballerCardIllustrations } from '@components/FootballerCard/FootballerCardIllustrations';
import { FootballerCardPrice } from '@components/FootballerCard/FootballerCardPrice';
import { FootballerCardStatsArea } from '@components/FootballerCard/FootballerCardStatsArea';
import { FootballerCardInformation } from '@components/FootballerCard/FootballerCardInformation';

import './index.scss';

const FootballerCard: React.FC = () => {
    const dispatch = useDispatch();
    const cardData = useSelector((state: RootState) => state.cardsReducer.openedCard);

    useEffect(() => {
        const LENGTH_INDEX_DIFFEFENCE = 1;
        const url = location.pathname.split('/');
        const id = url[url.length - LENGTH_INDEX_DIFFEFENCE];
        console.log(id)
        console.log(location.pathname.includes('lot'))
        location.pathname.includes('lot')?
        dispatch(openMarketplaceCard(id)):
        dispatch(openUserCard(id));
    }, []);

    return (
        cardData &&
        <div className="footballer-card">
            <div className="footballer-card__border">
                <div className="footballer-card__wrapper">
                    <div className="footballer-card__name-wrapper">
                        <h1 className="footballer-card__name">
                            {cardData.playerName}
                        </h1>
                    </div>
                    <FootballerCardIllustrations card={cardData} />
                    <div className="footballer-card__stats-area">
                        <FootballerCardPrice card={cardData} />
                        <FootballerCardStatsArea card={cardData} />
                        <FootballerCardInformation card={cardData} />
                    </div>
                </div>
            </div>
        </div>
    );
};

export default FootballerCard;
