/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
*/

import { FootballerCardIllustrations } from '@FootballerCard/FootballerCardIllustrations';
import { FootballerCardPrice } from '@FootballerCard/FootballerCardPrice';
import { FootballerCardStatsArea } from '@FootballerCard/FootballerCardStatsArea';
import { FootballerCardInformation } from '@FootballerCard/FootballerCardInformation';

import './index.scss';

export const FootballerCard: React.FC = (props) => {
    // @ts-ignore
    const cardData = props.location.state.card;
    const FIRST_CARD_INDEX = 0;

    return (
        <div className="footballer-card">
            <div className="footballer-card__border">
                <div className="footballer-card__wrapper">
                    <div className="footballer-card__name-wrapper">
                        <h1 className="footballer-card__name">
                            {cardData.overalInfo[FIRST_CARD_INDEX].value}
                        </h1>
                    </div>
                    <FootballerCardIllustrations card={cardData} />
                    <div className="footballer-card__stats-area">
                        <FootballerCardPrice />
                        <FootballerCardStatsArea card={cardData} />
                        <FootballerCardInformation card={cardData} />
                    </div>
                </div>
            </div>
        </div>
    );
};
