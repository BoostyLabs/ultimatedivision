import React from 'react';
import { useState } from 'react';

import { DropdownStyle } from '@/app/utils/dropdownStyle';

import ultimate from '@static/img/Navbar/ultimate.png';
import triangle from '@static/img/FootballFieldPage/triangle.svg';

import './index.scss';

export const InfoMenu = () => {

    const [whitePaperVisibility, changeWhitePaperVisibility] = useState(false);
    const [tokenomicsVisibility, changeTokenomicsVisibility] = useState(false);

    const whitePaperStyle = new DropdownStyle(whitePaperVisibility, 130);
    const tokenomicsStyle = new DropdownStyle(tokenomicsVisibility, 130);

    const menuFields = {
        whitepaper: ['Game Mechanics', 'Play to Earn and Economy', 'Technology', 'Team'],
        tokenomics: ['UDT Spending', 'Play to Earn', 'Staking', 'UD DAO Fund']
    }

    return (
        <div className="info-menu">
            <div className="info-menu__logo-wrapper">
                <img src={ultimate} alt="ultimate logo" />
            </div>
            <div
                className="info-menu__whitepaper"
                onClick={() => changeWhitePaperVisibility(prev => !prev)}
            >
                <h2>Whitepaper</h2>
                <img
                    className="info-menu__whitepaper-image"
                    src={triangle}
                    style={{ transform: whitePaperStyle.triangleRotate }}
                    alt="triangle img"
                />
            </div>
            <ul
                className="info-menu__whitepaper-list"
                style={{ height: whitePaperStyle.listHeight }}
            >
                {
                    menuFields.whitepaper.map((item, index) => (
                        <li
                            key={index}
                            className="info-menu__whitepaper-item"
                        >
                            {item}
                        </li>
                    ))
                }
            </ul>

            <div
                className="info-menu__tokenomics"
                onClick={() => changeTokenomicsVisibility(prev => !prev)}
            >
                <h2>Tokenomics</h2>
                <img
                    className="info-menu__whitepaper-image"
                    src={triangle}
                    style={{ transform: tokenomicsStyle.triangleRotate }}
                    alt="triangle img"
                />
            </div>
            <ul
            className="info-menu__tokenomics-list"
            style={{height: tokenomicsStyle.listHeight}}
            >
                {
                    menuFields.tokenomics.map((item, index) => (
                        <li
                            key={index}
                            className="info-menu__tokenomics-item"
                        >
                            {item}
                        </li>
                    ))
                }
            </ul>
        </div>
    )
}
