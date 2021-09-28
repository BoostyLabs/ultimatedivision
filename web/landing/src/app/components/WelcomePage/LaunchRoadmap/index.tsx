// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Card } from './Card';

import './index.scss';

export const LaunchRoadmap: React.FC = () => {
    const roadmap = [
        {
            title: 'Equipment skyfall',
            subTitle: 'The Treasury now contains 10 ETH',
            description: `15 unique items will enter the metaverse and drop from the sky.
            Will you be the lucky one to receive unique boots NFT
            which will guarantee bonuses in game?`,
            value: 15,
        },
        {
            title: 'What’s inside?',
            subTitle: 'The Treasury now contains 10 ETH ',
            description: `20 mysterious lootboxes will be sent to lucky owners
            of UD founder player cards. You can sell it for a minimum listing price
            of 0.05 ETH or wait for the game launch to see what’s inside.`,
            value: 50,
        },
        {
            title: 'Team Spirit',
            subTitle: '50 ETH have been deposited to the treasury',
            description: `It’s almost time to get to the field.
            Get one of 20 founder kits for your club
            that your fans will absolutely love. `,
            value: 80,
        },
        {
            title: 'Game On',
            subTitle: '50 ETH have been deposited to the treasury',
            description: `It’s almost time to get to the field. 
            Get one of 20 founder kits for your club that your fans will absolutely love. `,
            value: 100,
        },
    ];

    return <section className="launch-roadmap" id="roadmap">
        <div className="wrapper"> 
            <h1 className="launch-roadmap__title"
                data-aos="fade-right"
                data-aos-duration="600"
                data-aos-easing="ease-in-out-cubic"
            >
                Launch Roadmap
            </h1>
            {roadmap.map((card, index) => (
                <Card card={card} key={index} />
            ))}
        </div>
    </section>;
};
