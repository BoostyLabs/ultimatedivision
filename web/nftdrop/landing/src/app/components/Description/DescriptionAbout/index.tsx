// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import field from '@static/images/description/field.webp';
import fieldMobile from '@static/images/description/field-mobile.webp';
import fieldTablet from '@static/images/description/field-tablet.webp';

import './index.scss';


export const DescriptionAbout = () => {
    return (
        <div className="description-about" id="about">
            <div className="description-about__image-area">
                <picture>
                    <source media="(max-width: 400px)" srcSet={fieldMobile} />
                    <source media="(min-width: 400px)" srcSet={field} />
                    <img
                        className="description-about__field"
                        src={field}
                        alt="field image"
                        loading="lazy"
                    />
                </picture>
            </div>
            <div className="description-about__text-area">
                <h2 className="description-about__title">About the Game</h2>
                <p className="description-about__text">
                    Ultimate Division is a world football simulator.
                    UD players will own clubs, players and face each other
                    in weekly competitions to win cash prizes!
                    Other players can be hired as managers or coaches
                    for your Club!
                </p>
            </div>
        </div>
    );
};
