// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect } from 'react';

import Aos from 'aos';

import { DescriptionAbout } from '@components/Description/DescriptionAbout';
import { DescriptionCards } from '@components/Description/DescriptionCards';
import { DescriptionPay } from '@components/Description/DescriptionPay';

import './index.scss';

export const Description = () => {
    useEffect(() => {
        Aos.init({
            duration: 1500,
        });
    }, []);
    return (
        <section className="description">
            <div className="description__wrapper">
                <div className="description__container"
                    data-aos="fade-right"
                    data-aos-duration="900"
                    data-aos-easing="ease-in-out-cubic"
                >
                    <DescriptionAbout />
                </div>
                <div className="description__container"
                    data-aos="fade-left"
                    data-aos-duration="900"
                    data-aos-easing="ease-in-out-cubic"
                >
                    <DescriptionCards />
                </div>
                <div className="description__container"
                    data-aos="fade-right"
                    data-aos-duration="900"
                    data-aos-easing="ease-in-out-cubic"
                >
                    <DescriptionPay />
                </div>
            </div>
        </section>
    );
};
