/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import './FootballFieldInformation.scss'

import { FootballFieldInformationOption }
    from '../FootballFieldInformationOption/FootballFieldInformationOption';

import { FotballFieldInformationLine } from '../../../types/fotballerCard';

export const FootballFieldInformation = () => {

    const InformationField: FotballFieldInformationLine[] = [
        new FotballFieldInformationLine('0', 'formation', ['4-4-2', '4-2-4', '4-3-3']),
        new FotballFieldInformationLine('1', 'tactics', ['4-4-2', '4-2-4', '4-3-3']),
        new FotballFieldInformationLine('2', 'captain', ['4-4-2', '4-2-4', '4-3-3']),
    ]

    return (
        <div className="football-field-information">
            <h2 className="football-field-information__title">
                information
            </h2>
            {
                InformationField.map(item => (
                    <FootballFieldInformationOption
                        key={item.id}
                        props={item}
                    />
                ))
            }
        </div>
    )
}
