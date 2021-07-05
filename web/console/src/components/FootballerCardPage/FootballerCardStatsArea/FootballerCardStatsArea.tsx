/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
*/

import React from 'react';
import { useSelector } from 'react-redux';
import { RootState } from '../../../store';
import './FootballerCardStatsArea.scss';
import { FootballerCardStats }
    from '../FootballerCardStats/FootballerCardStats';


export const FootballerCardStatsArea: React.FC = () => {
    const STATS = useSelector((state: RootState )=> state.cardReducer[0].stats);

    return (
        <div className="footballer-card-stats">
            {STATS.map((item, index) => (
                <FootballerCardStats
                    key={index}
                    props={item}
                />
            ))}
        </div>
    );
};
