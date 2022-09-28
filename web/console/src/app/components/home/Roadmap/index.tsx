// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { RoadmapPoint } from './RoadmapPoint';

import roadmap from '@static/img/gameLanding/roadmap/roadmap.png';

import './index.scss';

export const Roadmap: React.FC = () => {
    const dataList = [
        {
            id: 1,
            step: 'start',
            points: [
                'Ultimate Division Metaverse MVP launch',
                'Play to Earn mechanics available to players',
                'UDT (Ultimate Division Token) introduced',
            ],
            done: true,
        },
        {
            id: 2,
            step: 'middle',
            points: [
                'Game in full swing',
                'Management roles available',
                'Club owners can tokenize their clubs and sell shares to other players',
                'UDT partnership with Top-5 Leagues',
            ],
            done: false,
        },
        {
            id: 3,
            step: 'finish',
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
            <h2 className="roadmap__title">Our <span className="roadmap__title__second-part">Roadmap</span></h2>
            {dataList.map((item) =>
                <RoadmapPoint key={item.id} item={item} />
            )}
            <img className="roadmap__image" src={roadmap} alt="roadmap"/>
        </section>
    );
};
