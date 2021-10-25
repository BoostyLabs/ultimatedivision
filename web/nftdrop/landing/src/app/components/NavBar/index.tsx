// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React, { useState } from "react";

import { MintButton } from "@components/common/MintButton";

import NavBarLogo from "@static/images/navbar/navbar-logo.png";
import { Cross, DropdownNavBar } from "@static/images/navbar/svg";

import "./index.scss";

export const Navbar: React.FC = () => {
    const [dropdownMenu, setDropdownMenu] = useState<boolean>(false);

    const navBarItems: Array<string> = [
        "Metaverse",
        "About",
        "Cards",
        "Roadmap",
    ];

    return (
        <div className="ultimatedivision-navbar">
            <div className="wrapper">
                <picture className="ultimatedivision-navbar__logo">
                    <a href="/">
                        <img src={NavBarLogo} alt="Ultimate-division logo" />
                    </a>
                </picture>
                <div
                    className="ultimatedivision-navbar__dropdown"
                    onClick={() => setDropdownMenu(!dropdownMenu)}
                >
                    {dropdownMenu ? <Cross /> : <DropdownNavBar />}
                </div>
                <ul
                    className={`ultimatedivision-navbar__items${
                        dropdownMenu ? "-active" : ""
                    }`}
                >
                    {navBarItems.map((item, index) => (
                        <li
                            key={index}
                            className={`ultimatedivision-navbar__items${
                                dropdownMenu ? "-active" : ""
                            }__item`}
                        >
                            <a
                                href={`#${item.toLowerCase()}`}
                                className={`ultimatedivision-navbar__items${
                                    dropdownMenu ? "-active" : ""
                                }__item__icon`}
                                onClick={() => setDropdownMenu(false)}
                            >
                                {item}
                            </a>
                        </li>
                    ))}
                </ul>
                <MintButton />
            </div>
        </div>
    );
};
