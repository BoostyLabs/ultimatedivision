/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import { NavLink } from 'react-router-dom';

import { Card } from '../../../../store/reducers/footballerCard';

import { RouteConfig } from '../../../../routes';

import './MyCard.scss';

export const MyCard: React.FC<{ card: Card, place?: string }> = ({ card, place }) => {
    return (
        <div
            className="marketplace-myCard"
        >
            <img
                className="marketplace-myCard__background-type"
                src={card.mainInfo.backgroundType}
                alt="Player background type"
            />
            <img
                className="marketplace-myCard__face-picture"
                src={card.mainInfo.playerFace}
                alt="Player face"
            />
            <NavLink to={RouteConfig.FootballerCard.path} >
                <span className="marketplace-myCard__name">
                    {card.mainInfo.lastName}
                </span>
            </NavLink>
            <img
                className="marketplace-myCard__confirm-icon"
                src={card.mainInfo.confirmIcon}
                alt="Confirm icon"
            />
            <img
                className="marketplace-myCard__price-gold"
                src={card.mainInfo.priceGoldIcon}
                alt="Price icon"
            />
            <ul className="marketplace-myCard__list">
                {card.stats.map(
                    (stat, index) => {
                        return (
                            <li
                                className="marketplace-myCard__list__item"
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
        </div>
    );
};
