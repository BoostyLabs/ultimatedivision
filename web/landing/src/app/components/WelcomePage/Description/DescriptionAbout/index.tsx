// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import './index.scss';
import field from '@static/images/Description/field.png';
import fieldTablet from '@static/images/Description/field-834.png';
import fieldPhone from '@static/images/Description/field-414.png';
import fieldMirrored from '@static/images/Description/fieldMirrored.png';
import fieldMirroredTablet from '@static/images/Description/fieldMirrored-834.png';
import cardStats from '@static/images/Description/cardStats.svg';
import cardStatsTablet from '@static/images/Description/card-stats-834.png';
import cardStatsPhone from '@static/images/Description/card-stats-414.png';

export const DescriptionAbout = () => {
    return (
        <div className="description-about" id="about">
            <div className="description-about__image-area">
                <picture className="description-about__cards-field">
                    <source media="(max-width: 414px)" srcSet={fieldPhone} />
                    <source media="(max-width: 834px)" srcSet={fieldTablet} />
                    <source media="(min-width: 1440px)" srcSet={field}/>
                    <img src={field} alt="field image" />
                </picture>
                <picture className="description-about__mirrored">
                    <source media="(max-width: 834px)" srcSet={fieldMirroredTablet} />
                    <source media="(min-width: 1440px)" srcSet={fieldMirrored}/>
                    <img src={fieldMirrored} alt="field mirrored" />
                </picture>
                <picture className="description-about__stats-image">
                    <source media="(max-width: 414px)" srcSet={cardStatsPhone} />
                    <source media="(max-width: 834px)" srcSet={cardStatsTablet} />
                    <source media="(min-width: 1440px)" srcSet={cardStats}/>
                    <img src={cardStats} alt="card stats image" />
                </picture>
            </div>
            <div className="description-about__text-area">
                <h2 className="description-about__title">About the Game</h2>
                <p className="description-about__text">
                    Ultimate Division is football world simulator.
                    Players can own clubs, compete with each other in weekly
                    competitions and earn money by winning.
                    Other players can be hired as managers and coaches for your own club.
                    Each assets in the game is NFT and brings profit when put to use.
                </p>
            </div>
        </div>
    );
};
