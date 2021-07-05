/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import { NavLink } from 'react-router-dom';

import './MarketPlaceNavbar.scss';

import ultimate
    from '../../../img/MarketPlacePage/MarketPlaceNavbar/ultimate.png';
import { RouteConfig } from '../../../routes';

export const MarketPlaceNavbar: React.FC = () => {
    return (
        <div className="marketplace-navbar">
            <img className="marketplace-navbar__logo"
                src={ultimate}
                alt={ultimate} />
            <ul className="marketplace-navbar__list">
                <li className="marketplace-navbar__item">
                    <NavLink to="/ud"
                        className="marketplace-navbar__item__active">
                        HOME
                    </NavLink>
                </li>
                <li className="marketplace-navbar__item">
                    <NavLink
                        to={RouteConfig.MarketPlace.path}
                        exact={RouteConfig.MarketPlace.exact}
                        className="marketplace-navbar__item__active"
                    >
                        MARKETPLACE
                    </NavLink>
                </li>
                <li className="marketplace-navbar__item">
                    <NavLink
                        to={RouteConfig.Club.path}
                        className="marketplace-navbar__item__active"
                    >
                        CLUB
                    </NavLink>
                </li>
            </ul >
        </div>
    );
};
