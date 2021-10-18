// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.
import React from 'react';

import { AnimationImage } from '@components/common/AnimationImage';

import box from '@static/images/launchRoadmap/box.svg';

import animationImage_0 from '@static/images/launchRoadmap/animated-diagram/images/img_0.png';
import animationImage_1 from '@static/images/launchRoadmap/animated-diagram/images/img_1.png';
import animationImage_2 from '@static/images/launchRoadmap/animated-diagram/images/img_2.png';
import animationImage_3 from '@static/images/launchRoadmap/animated-diagram/images/img_3.png';
import animationImage_4 from '@static/images/launchRoadmap/animated-diagram/images/img_4.png';

import './index.scss';

export const Card: React.FC<{
    card: {
        id: string;
        title: string;
        subTitle: string;
        description: string;
        animation: any;
    };
}> = ({ card }) => {
    const animationImages: string[] = [
        animationImage_0,
        animationImage_1,
        animationImage_2,
        animationImage_3,
        animationImage_4,
    ];

    return (
        <div className="card">
            <div className="card__text-area">
                <h1 className="card__title">{card.title}</h1>
                <p className="card__description">{card.description}</p>
                <div className="card__box">
                    <img
                        className="card__box__present"
                        src={box}
                        alt="utlimate box"
                    />
                    <p className="card__box__subtitle">{card.subTitle}</p>
                </div>
            </div>
            <AnimationImage
                className={`card__image-${card.id}`}
                heightFrom={1000}
                heightTo={-200}
                loop={true}
                animationData={card.animation}
                animationImages={animationImages}
            />
        </div>
    );
};
