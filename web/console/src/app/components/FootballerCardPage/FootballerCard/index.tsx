/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
*/
import { useLocation } from 'react-router-dom';

import { FootballerCardIllustrations } from '@components/FootballerCardPage/FootballerCardIllustrations';
import { FootballerCardPrice } from '@components/FootballerCardPage/FootballerCardPrice';
import { FootballerCardStatsArea } from '@components/FootballerCardPage/FootballerCardStatsArea';
import { FootballerCardInformation } from '@components/FootballerCardPage/FootballerCardInformation';

import { Card } from '@/app/store/reducers/footballerCard';

import './index.scss';

const FootballerCard: React.FC = () => {
    // @ts-ignore
    const FIRST_CARD_INDEX = 0;

    const { state } = useLocation<{ card: Card }>();
    const { card } = state;

    return (
        <div className="footballer-card">
            <div className="footballer-card__border">
                <div className="footballer-card__wrapper">
                    <div className="footballer-card__name-wrapper">
                        <h1 className="footballer-card__name">
                            {card.overalInfo[FIRST_CARD_INDEX].value}
                        </h1>
                    </div>
                    <FootballerCardIllustrations card={card} />
                    <div className="footballer-card__stats-area">
                        <FootballerCardPrice />
                        <FootballerCardStatsArea card={card} />
                        <FootballerCardInformation card={card} />
                    </div>
                </div>
            </div>
        </div>
    );
};

export default FootballerCard;
