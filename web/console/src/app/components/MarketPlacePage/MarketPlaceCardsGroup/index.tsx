// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Card } from '@/cards';
import { MarketPlaceFootballerCard } from '@components/MarketPlacePage/MarketPlaceCardsGroup/MarketPlaceFootballerCard';

import './index.scss';

export const MarketPlaceCardsGroup: React.FC<{ cards: Card[]}> = ({ cards }) =>
    <div className="marketplace-cards">
        <div className="marketplace-cards__wrapper">
            {cards.map((card, index) =>
                <MarketPlaceFootballerCard card={card} key={index} />
            )}
        </div>
    </div>;
