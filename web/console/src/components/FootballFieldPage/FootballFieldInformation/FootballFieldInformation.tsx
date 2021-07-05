/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import './FootballFieldInformation.scss'

import { FootballFieldInformationFormation }
    from './FootballFieldInformationFormation/FootballFieldInformationFormation';
import { FootballFieldInformationTactic }
    from './FootballFieldInformationTactic/FootballFieldInformationTactic';
import { FootballFieldInformationCaptain }
    from './FootballFieldInformationCaptain/FootballFieldInformationCaptain';

import { FootballFieldInformationLine } from '../../../types/footballField';

export const FootballFieldInformation: React.FC = () => {

    const InformationFIelds: FootballFieldInformationLine[] = [
        new FootballFieldInformationLine('0', 'formation', ['4-4-2', '4-2-4', '4-3-3']),
        new FootballFieldInformationLine('1', 'tactics', ['attack', 'defence', 'regular']),
        new FootballFieldInformationLine('2', 'captain', ['4-4-2', '4-2-4', '4-3-3']),
    ]

    return (
        <div className="football-field-information">
            <h2 className="football-field-information__title">
                information
            </h2>
            <FootballFieldInformationFormation
                key={InformationFIelds[0].id}
                props={InformationFIelds[0]}
            />
            <FootballFieldInformationTactic
                key={InformationFIelds[1].id}
                props={InformationFIelds[1]}
            />
            <FootballFieldInformationCaptain
                key={InformationFIelds[2].id}
                props={InformationFIelds[2]}
            />
        </div>
    )
}
