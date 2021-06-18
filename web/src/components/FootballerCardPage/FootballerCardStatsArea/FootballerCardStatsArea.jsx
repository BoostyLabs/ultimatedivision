import React from 'react';
import { useSelector } from 'react-redux';
import './FootballerCardStatsArea.scss';
import { FootballerCardStats }
    from '../FootballerCardStats/FootballerCardStats';

export const FootballerCardStatsArea = () => {

    const stats = useSelector(state => state.footballerCard[0]);

    return (
        <div className="footballer-card-stats">
            {Object.keys(stats).map(key => (
                key === 'overalInfo'
                    ? null
                    : <FootballerCardStats key={key} props={stats[key]} />
            ))}
        </div>
    );
};
