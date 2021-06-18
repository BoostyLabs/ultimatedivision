import React from 'react';
import { PropTypes } from 'prop-types';
import { NavLink } from 'react-router-dom';

import './MarketPlaceFootballerCard.scss';

export const MarketPlaceFootballerCard = ({ player }) => {

    /**
    * get only player's game propertys
    */
    let keysOfPlayer = Object.keys(player).slice(2);

    return (
        <div className="marketplace-playerCard">
            <img className="marketplace-playerCard__background-type"
                src={player.mainInfo.backgroundType}
                alt={player.mainInfo.backgroundType} />
            <img className="marketplace-playerCard__face-picture"
                src={player.mainInfo.facePicture}
                alt={player.mainInfo.facePicture} />
            <NavLink to="/marketplace/card">
                <span className="marketplace-playerCard__name">
                    {player.overalInfo['name']}
                </span>
            </NavLink>
            <ul className="marketplace-playerCard__list">
                {keysOfPlayer.map(
                    (playerKey, index) => {
                        return (
                            <li
                                className="marketplace-playerCard__list__item"
                                key={index}>
                                {
                                    /**
                                    * get only general value of player's game property
                                    */
                                    `${Object.values(player[playerKey])[0]} ${playerKey.slice(0, 3)}`
                                }
                            </li>
                        )
                    }
                )}
            </ul>
            <div className="marketplace-playerCard__price">
                <img className="marketplace-playerCard__price__picture"
                    src={player.mainInfo.pricePicture}
                    alt={player.mainInfo.pricePicture} />
                <span className="marketplace-playerCard__price__current">
                    {player.mainInfo.price}
                </span>
            </div>
        </div>
    );
};

MarketPlaceFootballerCard.propTypes = {
    player: PropTypes.object.isRequired
};
