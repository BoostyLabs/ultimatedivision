// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { PlayerCard } from '@/app/components/common/PlayerCard';

import { RouteConfig } from '@/app/router';
import { Link } from 'react-router-dom';

import { Card } from '@/cards';

import './index.scss';

export const MarketPlaceFootballerCard: React.FC<{ card: Card; place?: string }> = ({ card }) =>
    <div
        className="marketplace-playerCard"
    >
        <Link
            className="marketplace-playerCard__link"
            to={{
                pathname: RouteConfig.FootballerCard.path,
                state: {
                    card,
                },
            }}
        >
            <PlayerCard
                card={card}
                parentClassName={'marketplace-playerCard'}
            />
            <div className="marketplace-playerCard__price">
                <img className="marketplace-playerCard__price__picture"
                    src={card.mainInfo.priceIcon}
                    alt="Player price" />
                <span className="marketplace-playerCard__price__current">
                    {card.mainInfo.price}
                </span>
                <img className="marketplace-playerCard__price__status"
                    src={card.mainInfo.priceStatus}
                    alt="Price status" />
            </div>
        </Link>
    </div >;
