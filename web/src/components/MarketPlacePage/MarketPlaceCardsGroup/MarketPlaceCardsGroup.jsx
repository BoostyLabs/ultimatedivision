import React from 'react';
import { PropTypes } from 'prop-types';

import './MarketPlaceCardsGroup.scss';
import { MarketPlaceFootballerCard }
    from './MarketPlaceFootballerCard/MarketPlaceFootballerCard';

export const MarketPlaceCardsGroup = ({ listOfFootballerCards }) => {

    return (
        <div className="marketplace-cards">
            {listOfFootballerCards.map((player, index) =>
                <MarketPlaceFootballerCard
                    player={player}
                    key={index}
                />
            )}
        </div>
    );

};

MarketPlaceCardsGroup.propTypes = {
    listOfFootballerCards: PropTypes.array.isRequired
};

