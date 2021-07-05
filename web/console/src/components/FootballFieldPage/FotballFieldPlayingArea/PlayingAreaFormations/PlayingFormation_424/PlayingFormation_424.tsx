/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import './PlayingFormation_424.scss';
import { FootballField } from '../../../../../types/footballField';
import { useDispatch } from 'react-redux';
import { ChoseCardPosition }
    from '../../../../../store/reducers/footballField';

export const PlayingFormation_424: React.FC<{ props: FootballField }> = ({ props }) => {
    const Dispatch = useDispatch();

    return (
        <div className="playing-formation-424">
            {props.cardsList.map((card, index) => (
                <div
                    onClick={() => Dispatch(ChoseCardPosition(index.toString()))}
                    key={index}
                    className="playing-formation-424__card"
                >
                </div>
            ))}
        </div>
    )
}
