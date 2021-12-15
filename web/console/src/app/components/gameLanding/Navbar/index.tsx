// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from "react";
import { NavLink } from "react-router-dom";

import { JoinButton } from "@components/common/JoinButton";

import { CloseDropdownIcon, DropdownIcon } from "@/app/static/img/Navbar";
import ultimate from "@static/img/Navbar/ultimate.svg";

import { RouteConfig } from "@/app/routes";

import "./index.scss";

export const Navbar: React.FC = () => {
    const [isDropdownActive, setIsDropdownActive] = useState<boolean>(false);

    /** Сlass visibility for navbar items. */
    const visibleClassName = isDropdownActive ? "-active" : "";

    /** TODO: DIVISIONS will be replaced with id parameter */
    const navbarItems: Array<{ name: string; path: string }> = [
        { name: "My Club", path: RouteConfig.Club.path },
        { name: "Store", path: RouteConfig.Store.path },
        { name: "Marketplace", path: RouteConfig.MarketPlace.path },
        { name: "FAQ", path: RouteConfig.Summary.path },
    ];

    /** Marketplace component index from navbar items. */
    const navbarMarketplaceIndex: number = 2;

    return (
        <div className="ultimatedivision-navbar">
            <a href={navbarItems[navbarMarketplaceIndex].path}>
                <img
                    className="ultimatedivision-navbar__logo"
                    src={ultimate}
                    alt="UltimateDivision logo"
                />
            </a>
            <div
                className="ultimatedivision-navbar__dropdown"
                onClick={() => setIsDropdownActive(!isDropdownActive)}
            >
                {isDropdownActive ? <CloseDropdownIcon /> : <DropdownIcon />}
            </div>
            <ul className={`ultimatedivision-navbar__list${visibleClassName}`}>
                {navbarItems.map((item, index) => (
                    <li
                        key={index}
                        className={`ultimatedivision-navbar__list${visibleClassName}__item`}
                    >
                        <NavLink
                            key={index}
                            to={item.path}
                            className={`ultimatedivision-navbar__list${visibleClassName}__item__active`}
                            onClick={() => setIsDropdownActive(false)}
                        >
                            {item.name}
                        </NavLink>
                    </li>
                ))}
                <JoinButton />
            </ul>
        </div>
    );
};
