// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React from 'react';

import twitter from '@static/images/footer/twitter.svg';
import discord from '@static/images/footer/discord.svg';

import './index.scss';

export const Footer: React.FC = () => {
    const socialList = [
        {
            id: 1,
            path: 'https://discord.com/invite/ultimatedivision',
            img: discord,
            text: 'Discord',
        },
        {
            id: 2,
            path: 'https://twitter.com/UltimateDivnft',
            img: twitter,
            text: 'Twitter',
        },
    ];

    return (
        <footer className="footer">
            <div className="footer__wrapper">
                <div className="footer__links">
                    <a
                        className="footer__link"
                        href="https://ultimatedivision.com/ud/whitepaper/summary"
                    >
                        Whitepaper
                    </a>
                    <a
                        className="footer__link"
                        href="https://ultimatedivision.com/ud/whitepaper/summary"
                    >
                        FAQ
                    </a>
                </div>
                <div className="footer__social">
                    <ul className="footer__list">
                        {socialList.map((social) => (
                            <a
                                key={social.id}
                                className="footer__social-item"
                                href={social.path}
                                target="_blank"
                                rel="noreferrer"
                            >
                                <div className="footer__social-link">
                                    <img
                                        className="footer__social-img"
                                        src={social.img}
                                        alt="social logo"
                                    />
                                    <span className="footer__social-text">
                                        {social.text}
                                    </span>
                                </div>
                            </a>
                        ))}
                    </ul>
                </div>
            </div>
        </footer>
    );
};
