// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Card } from './Card';

import './index.scss';

export const LaunchRoadmap: React.FC = () => {
    const roadmap = [
        {
            id: '0',
            title: 'Equipment skyfall',
            subTitle: 'The Treasury now contains 10 ETH',
            description: `15 unique items will enter the metaverse and drop from the sky.
            Will you be the lucky one to receive unique boots NFT
            which will guarantee bonuses in game?`,
            percent: '15%',
        },
        {
            id: '1',
            title: 'What’s inside?',
            subTitle: 'The Treasury now contains 20 ETH ',
            description: `20 mysterious lootboxes will be sent to lucky owners of UD
            founder player cards. You can sell it or wait for the
            game launch to see what's inside.`,
            percent: '50%',
        },
        {
            id: '2',
            title: 'Team Spirit',
            subTitle: 'The Treasury now contains 30 ETH',
            description: `It’s almost time to get to the field.
            Get one of 20 founder kits for your club
            that your fans will absolutely love. `,
            percent: '80%',
        },
        {
            id: '3',
            title: 'Game On',
            subTitle: 'The Treasury now contains 40 ETH',
            description: `The UD Metaverse is now unstoppable and the first
            competition will soon begin. Prepare your NFTs to become 1 of 10000 UD
            founders and join the game before anyone else. Will your club reach
            the top of Ultimate Division?`,
            percent: '100%',
        },
    ];

    return (
        <div className="launch-roadmap" id="roadmap">
            <div className="launch-roadmap__wrapper">
                <h1 className="launch-roadmap__title">Launch Roadmap</h1>
                {roadmap.map((card, index) => (
                    <Card card={card} key={index} />
                ))}
            </div>
        </div>
    );
};
