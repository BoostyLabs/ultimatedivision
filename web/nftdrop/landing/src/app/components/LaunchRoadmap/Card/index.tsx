// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.
import React, { useEffect, useState } from 'react';
import lottie from 'lottie-web';
import box from '@static/images/launchRoadmap/box1.svg';
import roadmapDiagramData from '@static/images/launchRoadmap/animated-diagram/data.json';
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
        percent: string;
    };
}> = ({ card }) => {
    const [isAnimation, setIsAnimation] = useState<boolean>(false);

    /** Reading and parsing JSON with data to animate playToEarn block. */
    const loadedImagesData = JSON.stringify(roadmapDiagramData);
    const parsedImagesData = JSON.parse(loadedImagesData);
    const animationImages: string[] = [
        animationImage_0,
        animationImage_1,
        animationImage_2,
        animationImage_3,
        animationImage_4,
    ];

    /** Adding the path to the pictures in JSON. */
    parsedImagesData.assets.forEach(
        //@ts-ignore
        (image, i) => {
            //@ts-ignore
            image.p = animationImages[i];
        }
    );

    const autoAnimation = () => {
        const animationBlock = document?.querySelector(
            `.card__image-${card.id}`
        );

        /** Height of the page to the animated block. */
        const heightFromTop: number | undefined
            = animationBlock?.getBoundingClientRect().top;

        /** Set animation state to true when the user scrolls to the required block. */
        if (heightFromTop && heightFromTop <= 800 && heightFromTop >= 0) {
            setIsAnimation(true);

            return;
        }

        /** Set animation state to false when the user scrolls up or down from the animated block. */
        setIsAnimation(false);
    };

    useEffect(() => {
        /** Scroll listener. */
        window.addEventListener('scroll', autoAnimation);

        /** Show animation if the animation state is true. */
        if (isAnimation) {
            lottie.loadAnimation({
                // @ts-ignore
                container: document.querySelector(`.card__image-${card.id}`),
                animationData: parsedImagesData,
                loop: false,
                autoplay: true,
            });

            return;
        }

        /** Delete the picture when animation state is false. */
        const animationSvg = document?.querySelector(`.card__image-${card.id}`);

        if (animationSvg?.hasChildNodes()) {
            animationSvg?.removeChild(animationSvg?.childNodes[0]);
        }

        return () => {
            window.removeEventListener('scroll', autoAnimation);
        };
    }, [isAnimation, parsedImagesData, card.id, autoAnimation]);

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
            <div className={`card__image-${card.id}`}></div>
        </div>
    );
};
