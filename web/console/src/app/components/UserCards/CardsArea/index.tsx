// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useSelector } from 'react-redux';

import { UserCard } from './UserCard';

import { RootState } from '@/app/store';
import { Card } from '@/card';

import './index.scss';
import { useState } from 'react';

export const CardsArea: React.FC = () => {
    const { cards } =
        useSelector((state: RootState) => state.cardsReducer.cardsPage);

    const [sellCardIndex, setIndex] = useState<number>(-1)

    return <div className="cards-area">
        <div className="cards-area__wrapper">
            {cards.map((card: Card, index: number) =>
                <UserCard
                    key={index}
                    card={card}
                    position={index}
                    sellingCardIndex={sellCardIndex}
                    setIndex={setIndex}
                />,
            )}
        </div>
    </div>;
};
