import React from 'react';
import { useSelector } from 'react-redux';

import { FootballerCardIllustrations} from
    '../FootballerCardIllustrations/FootballerCardIllustrations';
import { FootballerCardPriceArea } from
    '../FootballerCardPriceArea/FootballerCardPriceArea';
import { FootballerCardStatsArea } from
    '../FootballerCardStatsArea/FootballerCardStatsArea';
import { FootballerCardInformation } from
    '../FootballerCardInformation/FootballerCardInformation';

import './FootballerCard.scss';

export const FootballerCard = () => {

    const cardData = useSelector(state => state.footballerCard[0].overalInfo);

    return (
        <div className="footballer-card">
            <div className="footballer-card__wrapper">
                <div className="footballer-card__name-wrapper">
                    <h1 className="footballer-card__name">
                        {cardData.Name}
                    </h1>
                </div>
                <FootballerCardIllustrations />
                <div className="footballer-card__stats-area">
                    <FootballerCardPriceArea />
                    <FootballerCardStatsArea />
                    <FootballerCardInformation />
                </div>
            </div>
        </div>
    );
};
