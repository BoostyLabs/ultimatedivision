// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React from 'react';
import { useSelector } from 'react-redux';
import { RootState } from '../../../store';
import './FootballerCardInformation.scss';

export const FootballerCardInformation: React.FC = () => {
    const OVERAL_INFO = useSelector((state: RootState )=> state.cardReducer[0].overalInfo);

    return (
        <div className="footballer-card-information">
            {OVERAL_INFO.map((item, index) => (
                <div className="footballer-card-information__item"
                    key={index}
                >
                    <div className="footballer-card-information__item-title">{item.label}</div>
                    <div className="footballer-card-information__item-value">
                        <>
                            {item.value}
                            <img
                                className="footballer-card-information__item-icon"
                                src={item.icon}
                                alt={item.icon && 'item icon'}
                            />
                        </>
                    </div>
                </div>
            ))}
        </div>
    );
};
