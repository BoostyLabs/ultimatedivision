// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React from 'react';
import { useSelector } from 'react-redux';
import './FootballerCardStatsArea.scss';
import { FootballerCardStats }
    from '../FootballerCardStats/FootballerCardStats';

import { RootState } from '../../../store';

export const FootballerCardStatsArea = () => {
    const stats = useSelector((state: RootState )=> state.footballerCard[0].stats);

    return (
        <div className="footballer-card-stats">
            {stats.map(item => (
                <FootballerCardStats key={item.title} props={item} />
            ))}
        </div>
    );
};
