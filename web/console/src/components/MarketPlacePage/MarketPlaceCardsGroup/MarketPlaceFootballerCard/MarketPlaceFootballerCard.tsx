/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import { NavLink } from 'react-router-dom';

import { Card } from '../../../../store/reducers/footballerCard';

import './MarketPlaceFootballerCard.scss';

export const MarketPlaceFootballerCard: React.FC<{ card: Card }> = ({ card }) => {
    /* get player card name */
    const getPlayerName = () => {
        let playerName;
        card.overalInfo.forEach(element => {
            return element.label === 'name' ? playerName = element.value
                : null;
        });
        return playerName;
    };
    return (
        <div className="marketplace-playerCard">
            <img className="marketplace-playerCard__background-type"
                src={card.mainInfo.backgroundType}
                alt="Player background type" />
            <img className="marketplace-playerCard__face-picture"
                src={card.mainInfo.facePicture}
                alt="Player face" />
            <NavLink to="/ud/marketplace/card">
                <span className="marketplace-playerCard__name">
                    {getPlayerName()}
                </span>
            </NavLink>
            <ul className="marketplace-playerCard__list">
                {card.stats.map(
                    (stat, index) => {
                        return (
                            <li
                                className="marketplace-playerCard__list__item"
                                key={index}>
                                {
                                    /**
                                    * get only average value of player's game property
                                    */
                                    `${stat.abbreviated} ${stat.average}`
                                }
                            </li>
                        );
                    }
                )}
            </ul>
            <div className="marketplace-playerCard__price">
                <img className="marketplace-playerCard__price__picture"
                    src={card.mainInfo.pricePicture}
                    alt="Player price" />
                <span className="marketplace-playerCard__price__current">
                    {card.mainInfo.price}
                </span>
                <img className="marketplace-playerCard__price__status"
                    src={card.mainInfo.priceStatus}
                    alt="Price status" />
            </div>
        </div>
    );
};
