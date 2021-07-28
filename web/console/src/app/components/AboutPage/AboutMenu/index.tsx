import React from 'react';
import { useState } from 'react';

import { DropdownStyle } from '@/app/utils/dropdownStyle';

import ultimate from '@static/img/Navbar/ultimate.png';
import triangle from '@static/img/FootballFieldPage/triangle.svg';

import './index.scss';
import { Link } from 'react-router-dom';

export const AboutMenu = () => {

    const [whitePaperVisibility, changeWhitePaperVisibility] = useState(false);
    const [tokenomicsVisibility, changeTokenomicsVisibility] = useState(false);

    const whitePaperStyle = new DropdownStyle(whitePaperVisibility, 130);
    const tokenomicsStyle = new DropdownStyle(tokenomicsVisibility, 130);

    const menuFields = {
        whitepaper:
            [
                {
                    path: '/summary',
                    name: 'Summary',
                },
                {
                    path: '/game-mechanicks',
                    name: 'Game Mechanics',
                },
                {
                    path: '/pay-to-earn-and-economy',
                    name: 'Play to Earn and Economy'
                },
                {
                    path: '/technology',
                    name: 'Technology'
                },
                {
                    path: '/team',
                    name: 'Team'
                }
            ],
        tokenomics:
            [
                {
                    path: '/udt-spending',
                    name: 'UDT Spending',
                },
                {
                    path: '/pay-to-earn',
                    name: 'Play to Earn'
                },
                {
                    path: '/staking',
                    name: 'Staking'
                },
                {
                    path: '/ud-dao-fund',
                    name: 'UD DAO Fund'
                }
            ],
    }

    return (
        <div className="about-menu">
            <div className="about-menu__logo-wrapper">
                <img src={ultimate} alt="ultimate logo" />
            </div>
            <div
                className="about-menu__whitepaper"
                onClick={() => changeWhitePaperVisibility(prev => !prev)}
            >
                <h2>Whitepaper</h2>
                <img
                    className="about-menu__whitepaper-image"
                    src={triangle}
                    style={{ transform: whitePaperStyle.triangleRotate }}
                    alt="triangle img"
                />
            </div>
            <ul
                className="about-menu__whitepaper-list"
                style={{ height: whitePaperStyle.listHeight }}
            >
                {
                    menuFields.whitepaper.map((item, index) => (
                        <li
                            key={index}
                            className="about-menu__whitepaper-item"
                        >
                            <Link
                                to={`/test/whitepaper${item.path}`}
                                className="about-menu__whitepaper-link"
                            >
                                {item.name}

                            </Link>
                        </li>
                    ))
                }
            </ul>

            <div
                className="about-menu__tokenomics"
                onClick={() => changeTokenomicsVisibility(prev => !prev)}
            >
                <h2>Tokenomics</h2>
                <img
                    className="about-menu__whitepaper-image"
                    src={triangle}
                    style={{ transform: tokenomicsStyle.triangleRotate }}
                    alt="triangle img"
                />
            </div>
            <ul
                className="about-menu__tokenomics-list"
                style={{ height: tokenomicsStyle.listHeight }}
            >
                {
                    menuFields.tokenomics.map((item, index) => (
                        <li
                            key={index}
                            className="about-menu__tokenomics-item"
                        >
                            <Link
                                to={`/test/tokenomics${item.path}`}
                                className="about-menu__tokenomics-link"
                            >
                                {item.name}

                            </Link>
                        </li>
                    ))
                }
            </ul>
        </div>
    )
}
