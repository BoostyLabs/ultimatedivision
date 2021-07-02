/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import { useSelector } from 'react-redux';
import { RootState } from '../../../store';

import { MarketPlaceNavbar } from '../MarketPlaceNavbar/MarketPlaceNavbar';
import { MarketPlaceFilterField }
    from '../MarketPlaceFilterField/MarketPlaceFilterField';
import { MarketPlaceCardsGroup }
    from '../MarketPlaceCardsGroup/MarketPlaceCardsGroup';
import './MarketPlace.scss';
import { MarketPlacePaginator }
    from '../MarketPlacePaginator/MarketPlacePaginator';

export const MarketPlace: React.FC = () => {
    const cards = useSelector((state: RootState) => state.footballerCard);

    return (
        <section className="marketplace">
            <MarketPlaceNavbar />
            <MarketPlaceFilterField />
            <MarketPlaceCardsGroup
                cards={cards} />
            <MarketPlacePaginator
                countCards={cards.length} />
        </section>
    );
};
