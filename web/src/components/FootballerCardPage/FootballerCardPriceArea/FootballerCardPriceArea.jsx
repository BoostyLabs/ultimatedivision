import React from 'react';
import './FootballerCardPriceArea.scss';
import { FootballerCardPrice }
    from '../FootballerCardPrice/FootballerCardPrice.jsx';

import PS from '../../../img/FootballerCardPage/PS.png';
import xBox from '../../../img/FootballerCardPage/xBox.png';
import logo from '../../../img/FootballerCardPage/logo3.png';

export const FootballerCardPriceArea = () => {

    const priceData = [
        {
            id: '1',
            mainPrice: '11,400,00',
            priceArea: ['11,700,00', '12,000,00', '11,999,00', '12,900,00'],
            PRP: 80,
            updated: '16',
            PR: '1,142,000 - 15,000,000',
            logo: PS,
        },
        {
            id: '2',
            mainPrice: '10,868,000',
            priceArea: ['10,999,000', '12,499,000', '11,999,00', '12,900,00'],
            PRP: 75,
            updated: '16',
            PR: '1,142,000 - 15,000,000',
            logo: xBox,
        },
        {
            id: '3',
            mainPrice: '10,868,000',
            priceArea: ['10,999,000', '12,499,000', '11,999,00', '12,900,00'],
            PRP: 53,
            updated: '16',
            PR: '1,142,000 - 15,000,000',
            logo: logo,
        },
    ];

    return (
        <div className="football-card-price-area">
            {priceData.map(item => (
                <FootballerCardPrice key={item.id} {...item} />
            ))
            }
        </div>
    );
};
