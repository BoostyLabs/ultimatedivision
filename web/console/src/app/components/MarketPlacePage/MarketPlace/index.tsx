/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import { useSelector } from 'react-redux';

import { MarketPlaceCardsGroup } from '@MarketPlace/MarketPlaceCardsGroup';
import { MarketPlaceFilterField } from '@MarketPlace/MarketPlaceFilterField';
import { MarketPlaceFootballerCard } from '@MarketPlace/MarketPlaceCardsGroup/MarketPlaceFootballerCard';
import { MyCard } from '@MarketPlace/MyCard';
import { Paginator } from '@Paginator';

import { RouteConfig } from '@Routes';
import { RootState } from '@Store';

import './index.scss';

export const MarketPlace = ({ ...children }) => {
    const cards = useSelector((state: RootState) => state.cardReducer);

    let Component = MarketPlaceFootballerCard;
    let title = 'MARKETPLACE';
    if (children.path === RouteConfig.MyCards.path) {
        Component = MyCard;
        title = 'MY CARDS';
    };

    return (
        <section className="marketplace">
            <MarketPlaceFilterField
                title={title}
            />
            <MarketPlaceCardsGroup
                cards={cards}
                Component={Component}
            />
            <Paginator
                itemCount={cards.length} />
        </section>
    );
};
