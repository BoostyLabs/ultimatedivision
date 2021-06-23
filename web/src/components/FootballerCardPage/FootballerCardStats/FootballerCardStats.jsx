// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React from 'react';
import './FootballerCardStats.scss';

/* eslint-disable */
export const FootballerCardStats = ({ props, title }) => {

    const fields = props.fields;

    return (
        <div className="footballer-card-stats-item">
            <div
                className="footballer-card-stats-item__heading"
            >
                <span className="footballer-card-stats-item__heading-name"
                >
                    {title}
                </span>
                <span
                    className="footballer-card-stats-item__heading-value"
                    style={{ color: props.color }}
                >
                    {props.average}
                </span>
            </div>
            <div className="footballer-card-stats-item__diagram">
                <div
                    className="footballer-card-stats-item__diagram-value"
                    style={{ background: props.color, width: `${props.average}%` }}
                ></div>
            </div>
            {fields.map((item, index )=> (
                <div
                    key={index}
                    className="footballer-card-stats-item__row"
                >
                    <span className="footballer-card-stats-item__row-name"
                    >
                        {item.label}
                    </span>
                    <span
                        className="footballer-card-stats-item__row-value"
                    >
                        {item.value}
                    </span>
                </div>
            ))}
        </div>
    );
};
