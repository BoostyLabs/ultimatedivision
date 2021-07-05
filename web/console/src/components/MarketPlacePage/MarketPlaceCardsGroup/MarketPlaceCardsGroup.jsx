/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import { PropTypes } from 'prop-types';

import './MarketPlaceCardsGroup.scss';
import { MarketPlaceFootballerCard }
    from './MarketPlaceFootballerCard/MarketPlaceFootballerCard';
import { useSelector } from 'react-redux';

export const MarketPlaceCardsGroup = ({ cards, place }) => {
    const cardsf = useSelector(state => state.footballField)
    console.log(cardsf)
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

MarketPlaceCardsGroup.propTypes = {
    cards: PropTypes.array.isRequired
};
