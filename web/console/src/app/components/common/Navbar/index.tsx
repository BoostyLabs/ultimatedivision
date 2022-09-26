// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState, useEffect } from 'react';
import { NavLink, useLocation } from 'react-router-dom';

import HomeNavbar from '@components/home/HomeNavbar';

import { CloseDropdownIcon, DropdownIcon } from '@/app/static/img/Navbar';
import { setScrollAble } from '@/app/internal/setScrollAble';

import { RouteConfig } from '@/app/routes';

import './index.scss';

const Navbar: React.FC = () => {
    const [isDropdownActive, setIsDropdownActive] = useState<boolean>(false);

    const [isHomePath, setIsHomePath] = useState<boolean>(false);

    const location = useLocation();

    /** Сlass visibility for navbar items. */
    const navbarListClassName = isDropdownActive ? 'ultimatedivision-navbar__list-active' : '';
    const navbarWrapperClassName = isDropdownActive ? 'ultimatedivision-navbar--active' : '';

    /** TODO: DIVISIONS will be replaced with id parameter */
    const navbarItems: Array<{ name: string; path: string }> = [
        { name: 'HOME', path: RouteConfig.Home.path },
        { name: 'STORE', path: RouteConfig.Store.path },
        { name: 'CARDS', path: RouteConfig.Cards.path },
        { name: 'FIELD', path: RouteConfig.Field.path },
    ];

    const setDropdownNavbarActivity = () => {
        setScrollAble();
        setIsDropdownActive(!isDropdownActive);
    };

    useEffect(() => {
        location.pathname === '/home' ? setIsHomePath(true) : setIsHomePath(false);
    }, [location]);

    return (
        <>
            {isHomePath ?
                <HomeNavbar />
                :
                <div className={`ultimatedivision-navbar ${navbarWrapperClassName}`}>
                    <div className="ultimatedivision-navbar__dropdown">
                        {isDropdownActive ?
                            <p className="ultimatedivision-navbar__dropdown__menu">Menu</p>
                            :
                            <a className="ultimatedivision-navbar__dropdown__logo" href="/">
                                <span className="ultimatedivision-navbar__dropdown__logo__first-part">Ultimate </span>
                                division
                            </a>
                        }
                        <button onClick={() => setDropdownNavbarActivity() }
                            className="ultimatedivision-navbar__dropdown__button">
                            {isDropdownActive ? <CloseDropdownIcon /> : <DropdownIcon />}
                        </button>
                    </div>
                    <ul className={`ultimatedivision-navbar__list ${navbarListClassName}`}>
                        {navbarItems.map((item, index) =>
                            <li key={index} className="ultimatedivision-navbar__list__item">
                                <NavLink
                                    key={index}
                                    to={item.path}
                                    className="ultimatedivision-navbar__list__item__active"
                                    onClick={() => setDropdownNavbarActivity()}
                                >
                                    {item.name}
                                </NavLink>
                            </li>
                        )}
                    </ul>
                </div>
            }
        </>
    );
};

export default Navbar;
