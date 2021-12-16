// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useAppSelector } from '@/app/store';

import { MyCard } from './MyCard';

import { RootState } from '@/app/store';
import { Card } from '@/card';

import './index.scss';

export const ClubCardsArea: React.FC = () => {
    const { cards } =
        useAppSelector((state: RootState) => state.cards.cardsPage);

    return <div className="club-cards">
        <div className="club-cards__wrapper">
            {cards.map((card: Card, index: number) =>
                <MyCard
                    card={card}
                    key={index}
                />,
            )}
        </div>
    </div>;
};
