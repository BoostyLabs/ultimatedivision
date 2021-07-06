/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';

import './MarketPlaceCardsGroup.scss';
import { MarketPlaceFootballerCard }
    from './MarketPlaceFootballerCard/MarketPlaceFootballerCard';
import { useSelector } from 'react-redux';
import { Card }
    from '../../../store/reducers/footballerCard';

export const MarketPlaceCardsGroup: React.FC<{ cards:Card[], place?: string }> = ({ cards, place }) => {
    return (
        <div className="marketplace-cards">
            <div className="marketplace-cards__wrapper">
                {cards.map((card, index) => (
                    <MarketPlaceFootballerCard
                        card={card}
                        key={index}
                        place={place}
                    />
                ))}
            </div>
        </div>
    );
};

