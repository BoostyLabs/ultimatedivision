import React from 'react';
import { useSelector } from 'react-redux';
import './FootballerCardInformation.scss'

export const FootballerCardInformation = () => {

    const overalInfo = useSelector(state => Object.entries(state.footballerCard[0].overalInfo));

    return (
        <div className='footballer-card-information'>
            {overalInfo.map(item => (
                <div className="footballer-card-information__item"
                    key={overalInfo.indexOf(item)}
                >
                    <div className="footballer-card-information__item-title">{item[0]}</div>
                    <div className="footballer-card-information__item-value">{item[1]}</div>
                </div>
            ))}
        </div>
    );
};
