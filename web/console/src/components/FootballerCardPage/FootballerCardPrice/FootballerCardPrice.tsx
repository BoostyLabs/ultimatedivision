/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
*/
import { useSelector } from 'react-redux';

import { Doughnut } from 'react-chartjs-2';

import { RootState } from '../../../store';

import currency from '../../../img/FootballerCardPage/currency.png';

import './FootballerCardPrice.scss';

export const FootballerCardPrice: React.FC = () => {
    const FIRST_CARD_INDEX = 0;
    const FULL_VALUE_STATISTIC_SCALE = 100;

    const priceData = useSelector((state: RootState) =>
        state.cardReducer[FIRST_CARD_INDEX].price);
    const prpValue: number = priceData.prp.value;

    return (
        <div className="footballer-card-price">
            <div className="footballer-card-price__wrapper">
                <div className="footballer-card-price__diagram">
                    <p className="footballer-card-price__diagram-value">{`PRP: ${prpValue}%`}</p>
                    <Doughnut
                        type={Doughnut}
                        data={{
                            datasets: [{
                                data: [prpValue, FULL_VALUE_STATISTIC_SCALE - prpValue],
                                backgroundColor: [
                                    `${priceData.color}`,
                                    '#5E5EAA',
                                ],
                                borderColor: [
                                    'transparent',
                                ],
                                cutout: '80%',
                                rotation: 90,
                                esponsive: true,
                                maintainAspectRatio: true,
                            }],
                        }}
                    />
                </div>
                <div className="footballer-card-price__info-area">
                    <h2 className="footballer-card-price__price">
                        <>
                            {priceData.price.value}
                            <img
                                className="footballer-card-price__price-currency"
                                src={currency}
                                alt="currency img"
                            />
                        </>
                    </h2>
                    <div className="footballer-card-price__additional-info">
                        <div>
                            Price updated: <span
                                className="footballer-card-price__value"
                            >
                                {priceData.updated.value} mins ago
                            </span>
                        </div>
                        <div>
                            PR: <span
                                className="footballer-card-price__value"
                            >
                                {priceData.pr.value}
                            </span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};
