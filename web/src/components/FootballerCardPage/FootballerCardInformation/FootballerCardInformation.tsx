// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React from 'react';
import { useSelector } from 'react-redux';
import { RootState } from '../../../store';
import './FootballerCardInformation.scss'

export const FootballerCardInformation: React.FC = () => {
    const overalInfo = useSelector((state: RootState )=> state.footballerCard[0].overalInfo);

    return (
        <div className='footballer-card-information'>
            {overalInfo.map((item, index) => (
                <div className="footballer-card-information__item"
                    key={index}
                >
                    <div className="footballer-card-information__item-title">{item.label}</div>
                    <div className="footballer-card-information__item-value">
                        <>
                            {item.value}
                            <img
                                src={item.icon}
                                alt=""
                            />
                        </>
                    </div>
                </div>
            ))}
        </div>
    );
};
