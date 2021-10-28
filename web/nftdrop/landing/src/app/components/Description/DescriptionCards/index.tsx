// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Swiper, SwiperSlide } from "swiper/react";

import desktopCard_1 from "@static/images/description/players-for-slider/desktop/1.avif";
import desktopCard_2 from "@static/images/description/players-for-slider/desktop/2.avif";
import desktopCard_3 from "@static/images/description/players-for-slider/desktop/3.avif";
import { useEffect } from "react";

import "./index.scss";

export const DescriptionCards = () => {
    const cards: string[] = [desktopCard_1, desktopCard_2, desktopCard_3];

    useEffect(() => {}, []);

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
                <picture>
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
                    <Swiper
                        spaceBetween={50}
                        slidesPerView={1}
                        effect="fade"
                        loop={true}
                        autoplay={{
                            delay: 1000,
                            disableOnInteraction: false,
                        }}
                    >
                        {cards.map((card, index) => {
                            return (
                                <SwiperSlide>
                                    <img
                                        key={index}
                                        className="description-cards__card"
                                        src={card}
                                        alt="cards"
                                        loading="lazy"
                                    />
                                </SwiperSlide>
                            );
                        })}
                    </Swiper>
                </picture>
            </div>
        </div>
    );
};
