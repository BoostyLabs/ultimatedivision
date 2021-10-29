// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

//@ts-ignore
import cardsVideo from '@static/images/description/players-for-slider/slides.mp4';

import './index.scss';

export const DescriptionCards = () => {
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
                <div className="description-cards__animate-block">
                    <video autoPlay loop muted>
                        <source src={cardsVideo} type="video/mp4" />
                    </video>
                </div>
            </div>
        </div>
    );
};
