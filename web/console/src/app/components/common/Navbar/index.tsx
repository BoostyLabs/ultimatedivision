// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.
import { useState } from "react";
import { NavLink } from "react-router-dom";
import { RouteConfig } from "@/app/router";
import { DropDownIcon } from "@/app/static/img/Navbar";
import ultimate from "@static/img/Navbar/ultimate.svg";

import "./index.scss";

export const Navbar: React.FC = () => {
  const [isDropdownOpen, setIsDropdownOpen] = useState<boolean>(false);

  const dropdownClass = isDropdownOpen ? "-active" : "";

  return (
    <div className="ultimatedivision-navbar">
      <img
        className="ultimatedivision-navbar__logo"
        src={ultimate}
        alt="UltimateDivision logo"
      />
      <ul className={`ultimatedivision-navbar__list${dropdownClass}`}>
        <li className="ultimatedivision-navbar__item">
          <NavLink
            to={RouteConfig.Summary.path}
            className="ultimatedivision-navbar__item__active"
          >
            HOME
          </NavLink>
        </li>
        <li className="ultimatedivision-navbar__item">
          <NavLink
            to={RouteConfig.Store.path}
            className="ultimatedivision-navbar__item__active"
          >
            Store
          </NavLink>
        </li>
        <li className="ultimatedivision-navbar__item">
          <NavLink
            to={RouteConfig.MarketPlace.path}
            className="ultimatedivision-navbar__item__active"
          >
            MARKETPLACE
          </NavLink>
        </li>
        <li className="ultimatedivision-navbar__item">
          <NavLink
            to={RouteConfig.Club.path}
            className="ultimatedivision-navbar__item__active"
          >
            CLUB
          </NavLink>
        </li>
        <li className="ultimatedivision-navbar__item">
          <NavLink
            to={RouteConfig.FootballField.path}
            className="ultimatedivision-navbar__item__active"
          >
            FIELD
          </NavLink>
        </li>
      </ul>
      <div
        className="ultimatedivision-navbar__dropdown"
        onClick={() => setIsDropdownOpen(!isDropdownOpen)}
      >
        <DropDownIcon />
      </div>
    </div>
  );
};
