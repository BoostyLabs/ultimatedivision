/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
*/

import React from 'react';
import { useSelector } from 'react-redux';
import './FootballerCardStatsArea.scss';
import { FootballerCardStats }
    from '../FootballerCardStats/FootballerCardStats';

export const FootballerCardStatsArea = () => {
    const stats = useSelector(state => state.footballerCardStats[0]);

    return (
        <div className="footballer-card-stats">
            {Object.keys(stats).map(key => (
                <FootballerCardStats
                    key={key}
                    title={key}
                    props={stats[key]}
                />
            ))}
        </div>
    );
};
