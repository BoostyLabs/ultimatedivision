/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import { useSelector } from 'react-redux';


import { MarketPlaceCardsGroup }
    from '../MarketPlaceCardsGroup/MarketPlaceCardsGroup';
import { MarketPlaceFilterField }
    from '../MarketPlaceFilterField/MarketPlaceFilterField';
import { MarketPlaceFootballerCard }
    from '../MarketPlaceCardsGroup/MarketPlaceFootballerCard/MarketPlaceFootballerCard';
import { MyCard } from '../MyCard/MyCard';
import { Paginator }
    from '../../Paginator/Paginator';

import { RouteConfig } from '../../../routes';

import { RootState } from '../../../store';

import './MarketPlace.scss';

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
