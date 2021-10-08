// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React, { useEffect } from 'react';
import Aos from 'aos';

import footballer from '@static/images/home/footballer.webp';
import footballerMobile from '@static/images/home/footballer-mobile.webp';
import footballerTablet from '@static/images/home/footballer-tablet.webp';
import { ScrollTop } from '../ScrollTop';

import discord from '@static/images/home/discord.svg';
import twitter from '@static/images/home/twitter.svg';

import './index.scss';

export const Home: React.FC = () => {
    useEffect(() => {
        Aos.init({
            duration: 1500,
        });
    });

    return (
        <section className="home" id="home">
            <div className="home__wrapper">
                <div className="home__text-area"
                    data-aos="fade-right"
                    data-aos-duration="900"
                    data-aos-easing="ease-in-out-cubic"
                >
                    <h2 className="home__value">10 000</h2>
                    <h3 className="home__title">Unique Collectible Player Cards.</h3>
                    <p className="home__description">
                        Get one to become UD founder and join the Play-to-Earn game.
                        Build your club in the metaverse.
                    </p>
                    <div className="home__buttons-wrapper">
                        <a
                            className="home__discord"
                            href="https://discord.com/invite/ultimatedivision"
                            target="_blank"
                            rel="noreferrer"
                        >
                            <img
                                className="home__discord__logo"
                                src={discord}
                                alt="discord logo"
                            />
                            <span className="home__discord__text">
                                Join Discord
                            </span>
                        </a>
                        <a
                            className="home__twitter"
                            href="https://twitter.com/UltimateDivnft"
                            target="_blank"
                            rel="noreferrer"
                        >
                            <img
                                className="home__twitter__logo"
                                src={twitter}
                                alt="twitter logo"
                            />
                        </a>
                    </div>
                </div>
                <picture
                >
                    <source media="(max-width: 600px)" srcSet={footballerMobile} />
                    <source media="(max-width: 800px)" srcSet={footballerTablet} />
                    <source media="(min-width: 800px)" srcSet={footballer} />

                    <img
                        className="home__player-image"
                        data-aos="fade-left"
                        data-aos-duration="900"
                        data-aos-easing="ease-in-out-cubic"
                        src={footballer}
                        alt="Player Illustration"
                    />
                </picture>
                <div className="home__mobile-wrapper">
                    <div className="home__description-mobile">
                        Get one to become UD founder and join the Play-to-Earn game.
                        Build your club in the metaverse.
                        <div className="home__buttons-wrapper__mobile">
                            <a
                                className="home__discord"
                                href="https://discord.com/invite/ultimatedivision"
                                target="_blank"
                                rel="noreferrer"
                            >
                                <img
                                    className="home__discord__logo"
                                    src={discord}
                                    alt="discord logo"
                                />
                                <span className="home__discord__text">
                                    Join Discord
                                </span>
                            </a>
                            <a
                                className="home__twitter"
                                href="https://twitter.com/UltimateDivnft"
                                target="_blank"
                                rel="noreferrer"
                            >
                                <img
                                    className="home__twitter__logo"
                                    src={twitter}
                                    alt="twitter logo"
                                />
                            </a>
                        </div>
                    </div>
                </div>
                <ScrollTop />
            </div>
        </section >
    );
};
