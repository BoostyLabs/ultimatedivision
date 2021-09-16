// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { PlayerCard } from '@components/common/PlayerCard';

import { RouteConfig } from '@/app/router';
import { Link } from 'react-router-dom';

import { Card, MarketplaceLot } from '@/card';

import './index.scss';

export const MarketPlaceFootballerCard: React.FC<{ lot: MarketplaceLot; place?: string }> = ({ lot }) =>
    <div
        className="marketplace-playerCard"
    >
        <Link
            className="marketplace-playerCard__link"
            to={`${RouteConfig.Lot.path}${lot.id}`}
        >
            <PlayerCard
                card={lot.card}
                parentClassName={'marketplace-playerCard'}
            />
            <div className="marketplace-playerCard__price">
                {/* <img className="marketplace-playerCard__price__picture"
                    src={card.mainInfo.priceIcon}
                    alt="Player price" />
                <span className="marketplace-playerCard__price__current">
                    {card.mainInfo.price}
                </span>
                <img className="marketplace-playerCard__price__status"
                    src={card.mainInfo.priceStatus}
                    alt="Price status" /> */}
            </div>
        </Link>
    </div >;
