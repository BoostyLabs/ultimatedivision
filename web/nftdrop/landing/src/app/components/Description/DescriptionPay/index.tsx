// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from "react";
import lottie from "lottie-web";
import playToEarnData from "@static/images/description/playToEarn/data.json";
import animationImage_0 from "@static/images/description/playToEarn/images/img_0.png";
import animationImage_1 from "@static/images/description/playToEarn/images/img_1.png";
import animationImage_2 from "@static/images/description/playToEarn/images/img_2.png";
import animationImage_3 from "@static/images/description/playToEarn/images/img_3.png";
import animationImage_4 from "@static/images/description/playToEarn/images/img_4.png";

import "./index.scss";

export const DescriptionPay = () => {
    const [isAnimation, setIsAnimation] = useState<boolean>(false);

    /** Reading and parsing JSON with data to animate playToEarn block. */
    const loadedImagesData = JSON.stringify(playToEarnData);
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
        (image, i: number) => (image.p = animationImages[i])
    );

    useEffect(() => {
        /** Scroll listener. */
        window.addEventListener("scroll", () => {
            const animationBlock = document?.querySelector(
                ".description-pay__radar"
            );

            /** Height of the page to the animated block. */
            const heightFromTop: number | undefined =
                animationBlock?.getBoundingClientRect().top;

            /** Set animation state to true when the user scrolls to the required block. */
            if (
                heightFromTop &&
                heightFromTop >= -500 &&
                heightFromTop <= 1000
            ) {
                setIsAnimation(true);

                return;
            }

            /** Set animation state to false when the user scrolls up or down from the animated block. */
            setIsAnimation(false);
        });

        /** Show animation if the animation state is true. */
        if (isAnimation) {
            lottie.loadAnimation({
                // @ts-ignore
                container: document.querySelector(".description-pay__radar"),
                animationData: parsedImagesData,
                loop: false,
                autoplay: true,
            });

            return;
        }

        /** Delete the picture when if the animation state is false. */
        lottie.destroy();
    }, [isAnimation, parsedImagesData]);

    return (
        <div className="description-pay">
            <div className="description-pay__radar"></div>
            <div className="description-pay__text-area">
                <h2 className="description-pay__title">Play-to-Earn</h2>
                <p className="description-pay__text">
                    Club Owners who hold a Founder Collection NFT will be
                    awarded the in-game title of UD Founder. The UD Founders
                    will receive exclusive airdrops and will start the game in
                    UD’s top division.
                </p>
            </div>
        </div>
    );
};
