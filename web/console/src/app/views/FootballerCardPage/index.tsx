// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { FootballerCardIllustrations } from '@components/FootballerCard/FootballerCardIllustrations';
import { FootballerCardPrice } from '@components/FootballerCard/FootballerCardPrice';
import { FootballerCardStatsArea } from '@components/FootballerCard/FootballerCardStatsArea';
import { FootballerCardInformation } from '@components/FootballerCard/FootballerCardInformation';

import './index.scss';
import { Card, MarketplaceLot } from '@/card';
import { getLotFromApi } from '@/app/hooks/marketplace';
import { useEffect, useState } from 'react';
import { RouteConfig } from '@/app/router';

const FootballerCard: React.FC = () => {
    // @ts-ignore
    const [cardData, handleData] = useState<SetStateAction<null | Card>>(null);
    const id = location.pathname.split(`/ud${RouteConfig.FootballerCard.path}/`).join('');
    useEffect(() => {
        getLotFromApi(id).then(data => handleData(data));
    }, []);
    // TODO: replace ud to needed path


    return (
        cardData &&
        <div className="footballer-card">
            <div className="footballer-card__border">
                <div className="footballer-card__wrapper">
                    <div className="footballer-card__name-wrapper">
                        <h1 className="footballer-card__name">
                            {cardData.playerName}
                        </h1>
                    </div>
                    <FootballerCardIllustrations card={cardData} />
                    <div className="footballer-card__stats-area">
                        <FootballerCardPrice card={cardData} />
                        <FootballerCardStatsArea card={cardData} />
                        <FootballerCardInformation card={cardData} />
                    </div>
                </div>
            </div>
        </div>
    );
};

export default FootballerCard;
