/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import './PlayingFormation_433.scss';
import { FootballField } from '../../../../../types/footballField';
import { ChoseCardPosition }
    from '../../../../../store/reducers/footballField';
import { useDispatch } from 'react-redux';

export const PlayingFormation_433: React.FC<{ props: FootballField }> = ({ props }) => {
    const Dispatch = useDispatch();

    return (
        <div className="playing-formation-433">
            {props.cardsList.map((card, index) => (
                <div
                    onClick={() => Dispatch(ChoseCardPosition(index.toString()))}
                    key={index}
                    className="playing-formation-433__card"
                >
                </div>
            ))}
        </div>
    )
}
