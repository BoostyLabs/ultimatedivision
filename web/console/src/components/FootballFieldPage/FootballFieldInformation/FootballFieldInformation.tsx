/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import './FootballFieldInformation.scss'

import { FootballFieldInformationOption }
    from '../FootballFieldInformationOption/FootballFieldInformationOption';

import { FootballFieldInformationLine } from '../../../types/footballField';
import { handleFormations } from '../../../store/reducers/footballField';
import { handleTactics } from '../../../store/reducers/footballField';
import { handleCaptain } from '../../../store/reducers/footballField';

export const FootballFieldInformation: React.FC = () => {

    const InformationFIelds: FootballFieldInformationLine[] = [
        new FootballFieldInformationLine('0', 'formation', ['4-4-2', '4-2-4', '4-3-3'], handleFormations),
        new FootballFieldInformationLine('1', 'tactics', ['4-4-2', '4-2-4', '4-3-3'], handleTactics),
        new FootballFieldInformationLine('2', 'captain', ['4-4-2', '4-2-4', '4-3-3'], handleCaptain),
    ]

    return (
        <div className="football-field-information">
            <h2 className="football-field-information__title">
                information
            </h2>
            {
                InformationFIelds.map(item => (
                    <FootballFieldInformationOption
                        key={item.id}
                        props={item}
                    />
                ))
            }
        </div>
    )
}
