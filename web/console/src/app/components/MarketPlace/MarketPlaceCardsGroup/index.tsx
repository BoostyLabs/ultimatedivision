// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useSelector } from 'react-redux';

import { RootState } from '@/app/store';

import { MarketPlaceFootballerCard } from '@components/MarketPlace/MarketPlaceCardsGroup/MarketPlaceFootballerCard';

import './index.scss';

export const MarketPlaceCardsGroup: React.FC = () => {
    const lots =
        useSelector((state: RootState) => state.cardsReducer.marketplace);

    return <div className="marketplace-cards">
        <div className="marketplace-cards__wrapper">
            {lots.map((lot, index) =>
                <MarketPlaceFootballerCard lot={lot} key={index} />
            )}
        </div>
    </div>;
};
