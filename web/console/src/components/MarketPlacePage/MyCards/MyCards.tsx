/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import { MarketPlaceCardsGroup } from '../MarketPlaceCardsGroup/MarketPlaceCardsGroup';

import { MarketPlaceFilterField } from
    '../MarketPlaceFilterField/MarketPlaceFilterField';

import { MyCard } from
    './MyCard/MyCard';
import { UltimateDivisionPaginator } from '../../UltimateDivisionPaginator/UltimateDivisionPaginator';

import { useSelector } from 'react-redux';
import { RootState } from '../../../store';

export const MyCards: React.FC = () => {
    let cards = useSelector((state: RootState) => state.cardReducer);

    return (
        <section className="marketplace">
            <MarketPlaceFilterField
                title="MY CARDS"
            />
            <MarketPlaceCardsGroup
                Component={MyCard}
                cards={cards}
            />
            <UltimateDivisionPaginator
                itemCount={cards.length}
            />
        </section>
    );
}
