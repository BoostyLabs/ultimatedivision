// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.
import React, { useEffect } from "react";
import lottie from "lottie-web";
import box from "@static/images/launchRoadmap/box1.svg";
import roadmapDiagram from "@static/images/launchRoadmap/animated-diagram/data.json";
import images_0 from "@static/images/launchRoadmap/animated-diagram/images/img_0.png";
import images_1 from "@static/images/launchRoadmap/animated-diagram/images/img_1.png";
import images_2 from "@static/images/launchRoadmap/animated-diagram/images/img_2.png";
import images_3 from "@static/images/launchRoadmap/animated-diagram/images/img_3.png";
import images_4 from "@static/images/launchRoadmap/animated-diagram/images/img_4.png";

import "./index.scss";

export const Card: React.FC<{
    card: {
        id: string;
        title: string;
        subTitle: string;
        description: string;
        image: string;
        percent: string;
    };
}> = ({ card }) => {
    const loadedImagesData = JSON.stringify(roadmapDiagram);
    const parsedImagesData = JSON.parse(loadedImagesData);
    const images: string[] = [images_0, images_1, images_2, images_3, images_4];

    // const percents: string[] = ["15%", "50%", "80%", "100%"];

    const changeImagesPath = parsedImagesData.assets.map(
        //@ts-ignore
        (img: string, i: number) => {
            //@ts-ignore
            img.p = images[i];
            //@ts-ignore
            img[5] = card.percent;
            //@ts-ignore
            return;
        }
    );

    parsedImagesData.layers.map(
        //@ts-ignore
        (p: string, i: number) => {
            //@ts-ignore
            p[0] = card.percent;
            //@ts-ignore
            return;
        }
    );

    // parsedImagesData.assets[5].nm = card.percent;
    // parsedImagesData.layers[0].nm = card.percent;
    // console.log(parsedImagesData);
    // console.log(card.percent);

    useEffect(() => {
        lottie.loadAnimation({
            //@ts-ignore
            container: document.querySelector(`.card__image-${card.id}`),
            animationData: parsedImagesData,
        });
    }, []);

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
            {/* <img
                src={card.image}
                alt="diagram"
                className="card__image"
            /> */}
            <div className={`card__image-${card.id}`}></div>
        </div>
    );
};
