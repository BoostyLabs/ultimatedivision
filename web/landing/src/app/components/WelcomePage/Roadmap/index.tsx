// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React, { useEffect } from 'react';

import Aos from 'aos';

import { RoadmapCircle } from '@components/WelcomePage/Roadmap/RoadmapCircle';

import './index.scss';

export const Roadmap: React.FC = () => {
    useEffect(() => {
        Aos.init({
            duration: 1500,
        });
    }, []);

    const dataList = [
        {
            id: 1,
            date: '2021 Q4',
            points: [
                'Ultimate Division Metaverse MVP launch',
                'Play to Earn mechanics available to players',
                'UDT (Ultimate Division Token) introduced'
            ],
            done: true,
        },
        {
            id: 2,
            date: '2022 Q1',
            points: [
                'Game in full swing',
                'Management roles available',
                'Club owners can tokenize their clubs and sell shares to other players',
                'UD partnership with Top-5 Leagues',
            ],
            done: false,
        },
        {
            id: 3,
            date: '2022 Q2',
            points: [
                'Advanced gameplay introduced',
                'Local competitions launched',
                'DAO governance adopted',
            ],
            done: false,
        },
    ];

    return (
        <section className="roadmap">
            <h2 className="roadmap__title">
                Development Roadmap
            </h2>
            <div
                className="roadmap__road"
                data-aos="zoom-out-down">
                {dataList.map((item) => (
                    <RoadmapCircle
                        key={item.id} item={item} />
                ))}
            </div>
        </section>
    );
};
