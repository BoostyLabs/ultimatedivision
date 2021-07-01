/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import { PropTypes } from 'prop-types';

import './MarketPlaceCardsGroup.scss';
import { MarketPlaceFootballerCard }
    from './MarketPlaceFootballerCard/MarketPlaceFootballerCard';

export const MarketPlaceCardsGroup = ({ cards }) => {
    return (
        <div className="marketplace-cards">
            <div className="marketplace-cards__wrapper">
                {cards.map((card, index) =>
                    <MarketPlaceFootballerCard
                        card={card}
                        key={index}
                    />
                )}
            </div>
        </div>
    );
};

MarketPlaceCardsGroup.propTypes = {
    cards: PropTypes.array.isRequired
};
