// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import cards from '@static/images/Description/cardsGroup.svg';
import tabletCards from '@static/images/Description/cards-834.png';
import phoneCards from '@static/images/Description/cards-414.png';

import './index.scss';

export const DescriptionCards = () => {
    return (
        <div className="description-cards">
            <div className="description-cards__text-area">
                <h2 className="description-cards__title">
                    The Player Cards - Become UD Founder
                </h2>
                <p className="description-cards__text">
                    Each football player on the field is controlled by Player Card
                    NFT - they are required to make a squad. We will be releasing a
                    limited founder collection of extremely rare and powerful
                    player cards on September 20. Each player card NFT has stats
                    that will determine how good the footballer is.
                </p>
            </div>
            <div className="description-cards__cards">
                <picture className="description-cards__card" >
                    <source media="(max-width: 414px)" srcSet={phoneCards} />
                    <source media="(max-width: 834px)" srcSet={tabletCards} />
                    <source media="(min-width: 1440px)" srcSet={cards}/>
                    <img src={cards} alt="cards" />
                </picture>
            </div>
        </div>
    );
};
