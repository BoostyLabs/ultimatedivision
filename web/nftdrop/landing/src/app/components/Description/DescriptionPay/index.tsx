// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect } from "react";
import lottie from "lottie-web";
import Aos from "aos";

import playToEarn from "@static/images/description/playToEarn/data.json";
import images_0 from "@static/images/description/playToEarn/images/img_0.png";
import images_1 from "@static/images/description/playToEarn/images/img_1.png";
import images_2 from "@static/images/description/playToEarn/images/img_2.png";
import images_3 from "@static/images/description/playToEarn/images/img_3.png";
import images_4 from "@static/images/description/playToEarn/images/img_4.png";

import "./index.scss";

export const DescriptionPay = () => {
    const loadedImagesData = JSON.stringify(playToEarn);
    const parsedImagesData = JSON.parse(loadedImagesData);
    const images: string[] = [images_0, images_1, images_2, images_3, images_4];

    parsedImagesData.assets.forEach(
        //@ts-ignore
        (img: string, i: number) => (img.p = images[i])
    );

    const parentBlock = document.querySelector(
        ".description-pay__radar aos-init aos-animate"
    );
    console.log(parentBlock);
    
    useEffect(() => {
        Aos.init({
            duration: 3000,
        });

        if (parentBlock) {
            lottie.loadAnimation({
                // @ts-ignore
                container: document.querySelector(".description-pay__radar"),
                animationData: parsedImagesData,
                loop: false,
                autoplay: true,
            });
        }

        lottie.destroy();
    }, []);

    return (
        <div className="description-pay">
            {/* <img
                className="description-pay__radar"
                src={radar}
                alt="radar diagram"
                loading="lazy"
            /> */}

            <div
                className="description-pay__radar"
                data-aos="fade-up"
                data-aos-duration="700"
                data-aos-delay="800"
            ></div>
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
