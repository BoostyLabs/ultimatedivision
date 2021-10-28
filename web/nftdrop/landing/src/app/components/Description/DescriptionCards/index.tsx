// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState, useEffect, useRef } from "react";

import desktopCard_1 from "@static/images/description/players-for-slider/desktop/1.avif";
import desktopCard_2 from "@static/images/description/players-for-slider/desktop/2.avif";
import desktopCard_3 from "@static/images/description/players-for-slider/desktop/3.avif";
import desktopCard_4 from "@static/images/description/players-for-slider/desktop/4.avif";
import desktopCard_5 from "@static/images/description/players-for-slider/desktop/5.avif";
import desktopCard_6 from "@static/images/description/players-for-slider/desktop/6.avif";
import desktopCard_7 from "@static/images/description/players-for-slider/desktop/7.avif";
import tabletCard_1 from "@static/images/description/players-for-slider/tablet/1.avif";
import tabletCard_2 from "@static/images/description/players-for-slider/tablet/2.avif";
import tabletCard_3 from "@static/images/description/players-for-slider/tablet/3.avif";
import tabletCard_4 from "@static/images/description/players-for-slider/tablet/4.avif";
import tabletCard_5 from "@static/images/description/players-for-slider/tablet/5.avif";
import tabletCard_6 from "@static/images/description/players-for-slider/tablet/6.avif";
import tabletCard_7 from "@static/images/description/players-for-slider/tablet/7.avif";

import "./index.scss";

const cards = [
    {
        desktop: desktopCard_1,
        tablet: tabletCard_1,
    },
    {
        desktop: desktopCard_2,
        tablet: tabletCard_2,
    },
    {
        desktop: desktopCard_3,
        tablet: tabletCard_3,
    },
    {
        desktop: desktopCard_4,
        tablet: tabletCard_4,
    },
    {
        desktop: desktopCard_5,
        tablet: tabletCard_5,
    },
    {
        desktop: desktopCard_6,
        tablet: tabletCard_6,
    },
    {
        desktop: desktopCard_7,
        tablet: tabletCard_7,
    },
];

const delay = 1500;

export const DescriptionCards = () => {
    const [index, setIndex] = useState<number>(0);
    // const [card, setCard] = useState<string>(desktopCard_1);
    const timeoutRef = useRef<any>(null);

    function resetTimeout() {
        if (timeoutRef.current) {
            clearTimeout(timeoutRef.current);
        }
    }

    function sets() {
        setIndex((prevIndex) =>
            prevIndex === cards.length - 1 ? 0 : prevIndex + 1
        );
        timeoutRef.current = setTimeout(sets, delay);
    }

    useEffect(() => {
        resetTimeout();

        //@ts-ignore
        timeoutRef.current = setTimeout(sets, delay);

        return () => {
            resetTimeout();
        };
    }, [index]);

    return (
        <div className="description-cards" id="cards">
            <div className="description-cards__text-area">
                <h2 className="description-cards__title">
                    The Player Cards - Become UD Founder
                </h2>
                <p className="description-cards__text">
                    Each player in your club is an NFT - build a squad of 11
                    NFTs, for your team to compete week in, week out. These
                    NFT’s have the player’s stats, which determine how strong
                    each player is. If you want to score a limited Founder
                    Collection NFT, then make sure you are ready for date.
                </p>
            </div>
            <div className="description-cards__wrapper">
                <div
                    className={`description-cards__animate-block-${index}`}
                    // style={{ transform: `translate3d(${-index * 100}%, 0, 0)` }}
                >
                    {/* <picture> */}
                    {cards.map((card, i) => (
                        <img
                            className="description-cards__card"
                            src={card.desktop}
                            alt="cards"
                            loading="lazy"
                        />
                    ))}

                    {/* <img
                        className="description-cards__card"
                        src={card}
                        alt="cards"
                        // loading="lazy"
                    /> */}

                    {/* <source
                                key={i}
                                media="(min-width: 801px)"
                                srcSet={card.desktop}
                                type="image/avif"
                            />
                            <source
                                key={i}
                                media="(max-width: 800px)"
                                srcSet={card.tablet}
                                type="image/avif"
                            /> */}
                    {/* </picture> */}
                    {/* <source
                        media="(min-width: 601px)"
                        srcSet={desktopCard_1}
                        type="image/avif"
                    />
                    <source
                        media="(max-width: 600px)"
                        srcSet={desktopCard_1}
                        type="image/avif"
                    />
                    <img
                        className="description-cards__card"
                        src={desktopCard_1}
                        alt="cards"
                        loading="lazy"
                    /> */}
                </div>
            </div>
        </div>
    );
};
