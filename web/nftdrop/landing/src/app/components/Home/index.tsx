// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React from 'react';

import { ScrollTop } from '../ScrollTop';
import { AnimationImage } from '@components/common/AnimationImage';

import footballerAnimation from '@static/images/home/animation-player/data.json';

import ball from '@static/images/home/animation-player/images/ball.svg';
import card from '@static/images/home/animation-player/images/card.svg';
import head from '@static/images/home/animation-player/images/head.svg';
import leftArm from '@static/images/home/animation-player/images/leftArm.svg';
import leftLeg from '@static/images/home/animation-player/images/leftLeg.svg';
import rightArm from '@static/images/home/animation-player/images/rightArm.svg';
import rightLeg from '@static/images/home/animation-player/images/rightLeg.svg';
import discord from '@static/images/home/discord.svg';
import twitter from '@static/images/home/twitter.svg';

import './index.scss';

export const Home: React.FC = () => {
    const animationImages: string[] = [
        leftArm,
        leftLeg,
        rightLeg,
        rightArm,
        head,
        ball,
        card,
    ];

    return (
        <section className="home" id="home">
            <div className="home__wrapper">
                <div className="home__text-area">
                    <h1 className="home__value">10 000</h1>
                    <h3 className="home__title">
                        Unique Collectible Player Cards.
                    </h3>
                    <p className="home__description">
                        Get one to become UD founder and join the Play-to-Earn
                        game. Build your club in the metaverse.
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
                <AnimationImage
                    className={'home__player-image'}
                    heightFrom={4000}
                    heightTo={-500}
                    loop={true}
                    animationData={footballerAnimation}
                    animationImages={animationImages}
                    isNeedScrollListener={false}
                />
                <div className="home__mobile-wrapper">
                    <div className="home__description-mobile">
                        Get one to become UD founder and join the Play-to-Earn
                        game. Build your club in the metaverse.
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
        </section>
    );
};
