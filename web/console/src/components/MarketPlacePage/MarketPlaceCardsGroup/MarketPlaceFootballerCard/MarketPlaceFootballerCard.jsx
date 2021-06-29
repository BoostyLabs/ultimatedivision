/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React, { useEffect, useState } from 'react';
import { PropTypes } from 'prop-types';
import { NavLink } from 'react-router-dom';

import './MarketPlaceFootballerCard.scss';

export const MarketPlaceFootballerCard = ({ card }) => {
    let [mainCardProperties, setMainCardProperties] = useState([]);
    useEffect(() => {
        getCardStatsProperties();
    }, []);
    /**
    * get only card's stats properties
    */
    let getCardStatsProperties = () => {
        let statsProperties = [];
        Object.keys(card).forEach(field => {
            /* only stats data has property 'fields' */
            if (card[field].hasOwnProperty('fields')) {
                statsProperties.push({
                    field,
                    abbreviated: field.slice(0, 3)
                });
            }
        });
        setMainCardProperties(statsProperties);
    };

    return (
        <div className="marketplace-playerCard">
            <img className="marketplace-playerCard__background-type"
                src={card.mainInfo.backgroundType}
                alt="Player background type" />
            <img className="marketplace-playerCard__face-picture"
                src={card.mainInfo.facePicture}
                alt="Player face" />
            <NavLink to="/marketplace/card">
                <span className="marketplace-playerCard__name">
                    {card.overalInfo.name}
                </span>
            </NavLink>
            <ul className="marketplace-playerCard__list">
                {mainCardProperties.map(
                    (property, index) => {
                        return (
                            <li
                                className="marketplace-playerCard__list__item"
                                key={index}>
                                {
                                    /**
                                    * get only average value of player's game property
                                    */
                                    `${card[property.field].average} ${property.abbreviated}`
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

MarketPlaceFootballerCard.propTypes = {
    card: PropTypes.object.isRequired
};
