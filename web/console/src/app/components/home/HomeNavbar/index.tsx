// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';
import { NavLink } from 'react-router-dom';

import { JoinButton } from '@components/common/JoinButton';

import { CloseDropdownIcon, DropdownIcon } from '@/app/static/img/Navbar';

import { RouteConfig } from '@/app/routes';

import './index.scss';

const HomeNavbar: React.FC = () => {
    const [isDropdownActive, setIsDropdownActive] = useState<boolean>(false);

    /** Сlass visibility for navbar items. */
    const visibleClassName = isDropdownActive ? '-active' : '';

    /** TODO: DIVISIONS will be replaced with id parameter */
    const navbarItems: Array<{ name: string; path: string }> = [
        { name: 'Home', path: RouteConfig.Home.path },
        { name: 'My Club', path: RouteConfig.Field.path },
        { name: 'Store', path: RouteConfig.Store.path },
        { name: 'FAQ', path: RouteConfig.Summary.path },
    ];

    return (
        <nav className="ultimatedivision-home-navbar">
            <div className={`wrapper ${isDropdownActive?'wrapper--active':''}`}>
                <div className="ultimatedivision-home-navbar__wrapper">
                    {isDropdownActive && <p className="ultimatedivision-home-navbar__text">Menu</p>}
                    <div
                        className="ultimatedivision-home-navbar__dropdown"
                        onClick={() => setIsDropdownActive(!isDropdownActive)}
                    >
                        {isDropdownActive ?
                            <>
                                <CloseDropdownIcon />
                            </>
                            :
                            <DropdownIcon />
                        }
                    </div>
                </div>

                <ul
                    className={`ultimatedivision-home-navbar__list${visibleClassName}`}
                >
                    {navbarItems.map((item, index) =>
                        <li
                            key={index}
                            className={`ultimatedivision-home-navbar__list${visibleClassName}__item`}
                        >
                            <NavLink
                                key={index}
                                to={item.path}
                                className={`ultimatedivision-home-navbar__list${visibleClassName}__item__active`}
                                onClick={() => setIsDropdownActive(false)}
                            >
                                {item.name}
                            </NavLink>
                        </li>
                    )}
                </ul>
            </div>
        </nav>
    );
};

export default HomeNavbar;
