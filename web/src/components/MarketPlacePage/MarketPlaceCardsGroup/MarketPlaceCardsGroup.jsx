import React from 'react';
import { PropTypes } from 'prop-types';

import './MarketPlaceCardsGroup.scss';
import { MarketPlaceFootballerCard }
    from './MarketPlaceFootballerCard/MarketPlaceFootballerCard';

export const MarketPlaceCardsGroup = ({ listOfFootballerCards }) => {

    return (
        <div className="marketplace-cards">
            {listOfFootballerCards.map(player =>
                <MarketPlaceFootballerCard
                    player={player}
                    key={listOfFootballerCards.indexOf(player)}
                />
            )}
        </div>
    );

};

MarketPlaceCardsGroup.propTypes = {
    listOfFootballerCards: PropTypes.array.isRequired
};
