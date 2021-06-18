import React from 'react';
import './FootballerCardPrice.scss';
import { Doughnut } from 'react-chartjs-2';
import { diagramColor } from '../../../utils/fotballerCard';

/* eslint-disable */
export const FootballerCardPrice = ({
    mainPrice, priceArea, PRP, updated, PR, logo
}) => {
    return (
        <div className="footballer-card-price">
            <img src={logo} alt="" className="footballer-card-price__logo" />
            <h2 className="footballer-card-price__main-price">{mainPrice}</h2>
            <div className="footballer-card-price__price-area">
                {priceArea.map(item => (
                    <span key={priceArea.indexOf(item)}>{item}</span>
                ))}
            </div>
            <div className="footballer-card-price__diagram">
                <p className="footballer-card-price__diagram-value">{`PRP: ${PRP}%`}</p>
                <Doughnut
                    data={{
                        datasets: [{
                            data: [+PRP, (100 - PRP)],
                            backgroundColor: [
                                `${diagramColor(+PRP, 'priceArea')}`,
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
            <div className="footballer-card-price__info">
                <div className="footballer-card-price__info-line">
                    <span className="footballer-card-price__info-title">
                        {'Price updated: '}
                    </span>
                    <span className="footballer-card-price__info-value">
                        {updated} mins ago
                    </span>
                </div>
                <div className="footballer-card-price__info-line">
                    <span className="footballer-card-price__info-title">
                        {'PR: '}
                    </span>
                    <span className="footballer-card-price__info-value">
                        {PR}
                    </span>
                </div>
            </div>
        </div>
    );
};
