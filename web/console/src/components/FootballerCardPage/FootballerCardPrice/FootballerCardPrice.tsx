/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
*/

import React from 'react';
import './FootballerCardPrice.scss';
import { Doughnut } from 'react-chartjs-2';
import { useSelector } from 'react-redux';
import { RootState } from '../../../store';

import currency from '../../../img/FootballerCardPage/currency.png';

export const FootballerCardPrice: React.FC = () => {
    const PRICE_DATA = useSelector((state: RootState) => state.cardReducer[0].price);
    const PRP_VALUE: number = PRICE_DATA.prp.value;

    return (
        <div className="footballer-card-price">
            <div className="footballer-card-price__wrapper">
                <div className="footballer-card-price__diagram">
                    <p className="footballer-card-price__diagram-value">{`PRP: ${PRP_VALUE}%`}</p>
                    <Doughnut
                        type={Doughnut}
                        data={{
                            datasets: [{
                                data: [PRP_VALUE, (100 - PRP_VALUE)],
                                backgroundColor: [
                                    `${PRICE_DATA.color}`,
                                    '#5E5EAA'
                                ],
                                borderColor: [
                                    'transparent'
                                ],
                                cutout: '80%',
                                rotation: 90,
                                esponsive: true,
                                maintainAspectRatio: true
                            }],
                        }}
                    />
                </div>
                <div className="footballer-card-price__info-area">
                    <h2 className="footballer-card-price__price">
                        <>
                            {PRICE_DATA.price.value}
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
                                {PRICE_DATA.updated.value} mins ago
                            </span>
                        </div>
                        <div>
                            PR: <span
                                className="footballer-card-price__value"
                            >
                                {PRICE_DATA.pr.value}
                            </span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};
