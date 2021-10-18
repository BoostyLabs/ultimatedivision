// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.
import React, { useEffect, useState } from "react";
import lottie from "lottie-web";

export const AnimationImage: React.FC<{
    classData: any;
    className: string;
    animationData: any;
    animationImages: string[];
    heightFrom: number;
    heightTo: number;
    loop: boolean;
}> = ({
    className,
    classData,
    animationData,
    animationImages,
    heightFrom,
    heightTo,
    loop,
}) => {
    const [isAnimation, setIsAnimation] = useState<boolean>(false);

    /** Reading and parsing JSON with data to animate playToEarn block. */
    const loadedImagesData = JSON.stringify(animationData);
    const parsedImagesData = JSON.parse(loadedImagesData);

    /** Adding the path to the pictures in JSON. */
    parsedImagesData.assets.forEach(
        //@ts-ignore
        (image, i) => {
            //@ts-ignore
            image.p = animationImages[i];
        }
    );

    const autoAnimation = () => {
        /** Get id card block. */
        const animationBlock = document?.querySelector(
            `.${className}${classData ? classData.id : ""}`
        );

        /** Height of the page to the animated block. */
        const heightFromTop: number | undefined =
            animationBlock?.getBoundingClientRect().top;

        /** Set animation state to true when the user scrolls
         * to the required block. */
        if (
            heightFromTop &&
            heightFromTop <= heightFrom &&
            heightFromTop >= heightTo
        ) {
            if (isAnimation) {
                return null;
            }
            setIsAnimation(true);

            return null;
        }

        /** Set animation state to false when the user scrolls up
         * or down from the animated block. */
        if (!isAnimation) {
            return null;
        }
        
        setIsAnimation(false);
    };

    useEffect(() => {
        /** Scroll listener. */
        window.addEventListener("scroll", autoAnimation);

        /** Show animation if the animation state is true. */
        if (isAnimation) {
            lottie.loadAnimation({
                // @ts-ignore
                container: document.querySelector(
                    `.${className}${classData ? classData.id : ""}`
                ),
                animationData: parsedImagesData,
                loop: loop,
                autoplay: true,
            });

            return;
        }

        /** Delete the picture when animation state is false. */
        const animationSvg = document?.querySelector(
            `.${className}${classData ? classData.id : ""}`
        );

        if (animationSvg?.hasChildNodes()) {
            animationSvg?.removeChild(animationSvg?.childNodes[0]);
        }

        return () => {
            window.removeEventListener("scroll", autoAnimation);
        };
    }, [isAnimation, parsedImagesData]);

    return (
        <div className={`${className}${classData ? classData.id : ""}`}></div>
    );
};
