/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import { useSelector } from 'react-redux';
import './FootballFieldPlayingArea.scss';

import { RootState } from '../../../store';

import { PlayingFormation_442 }
    from './PlayingAreaFormations/PlayingFormation_442/PlayingFormation_442';
import { PlayingFormation_424 }
    from './PlayingAreaFormations/PlayingFormation_424/PlayingFormation_424';
import { PlayingFormation_433 }
    from './PlayingAreaFormations/PlayingFormation_433/PlayingFormation_433';

export const FootballFieldPlayingArea: React.FC = () => {
    const param = useSelector((state: RootState) => state.footballField.options.formation);

    const formation = (param: string) => {
        switch (param) {
            case '4-4-2':
                return <PlayingFormation_442 />
            case '4-2-4':
                return <PlayingFormation_424 />;
            case '4-3-3':
                return <PlayingFormation_433 />
        };
    }

    return (
        <div className="football-field-playing-area">
            {formation(param)}
        </div>
    )
}
