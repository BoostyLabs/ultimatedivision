import React from 'react';
import './FootballerCardStats.scss';

/* eslint-disable */
export const FootballerCardStats = ({ props }) => {

    const data = Object.entries(props).slice(0, -1)

    return (
        <div className="footballer-card-stats-item">
            <div
                className="footballer-card-stats-item__heading"
            >
                <span className="footballer-card-stats-item__heading-name"
                >
                    {data[0][0]}
                </span>
                <span
                    className="footballer-card-stats-item__heading-value"
                    style={{ color: props.color}}
                >
                    {data[0][1]}
                </span>
            </div>
            <div className="footballer-card-stats-item__diagram">
                <div
                    className="footballer-card-stats-item__diagram-value"
                    style={{background: props.color, width: `${data[0][1]}%`}}
                ></div>
            </div>
            {data.slice(1).map(item => (
                <div
                    key={data.indexOf(item)}
                    className="footballer-card-stats-item__row"
                >
                    <span className="footballer-card-stats-item__row-name"
                    >
                        {item[0]}
                    </span>
                    <span
                        className="footballer-card-stats-item__row-value"
                    >
                        {item[1]}
                    </span>
                </div>
            ))}
        </div>
    );
};
