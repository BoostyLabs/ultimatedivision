/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import './FootballFieldInformationCaptain.scss';

import { FotballFieldInformationLine } from '../../../../types/fotballerCard';

import triangle from '../../../../img/FootballFieldPage/triangle.png'
import { useState } from 'react';
import { ListStyle } from '../../../../utils/footballField';
import { TriangleStyle } from '../../../../utils/footballField';
import { useDispatch } from 'react-redux';
import { handleCaptain } from '../../../../store/reducers/footballField';

export const FootballFieldInformationCaptain: React.FC<{ props: FotballFieldInformationLine }> = ({ props }) => {
    const [state, changeState] = useState(true);

    const ListHeight = new ListStyle(state);
    const TriangleRotate = new TriangleStyle(state)

    const dispatch = useDispatch();

    return (
        <div className="football-field-information-option">
            <div
                className="football-field-information-option__heading"
                onClick={() => changeState(prev => !prev)}
            >
                <h4 className="football-field-information-option__title">
                    {props.title}
                </h4>
                <img
                    className="football-field-information-option__image"
                    src={triangle}
                    style={{ transform: TriangleRotate.style }}
                    alt="triangle img"
                    id={`triangle-${props.id}`}
                />
            </div>
            <ul
                style={{ height: ListHeight.style }}
                className="football-field-information-option__list"
                id={props.id}
            >
                {props.options.map((item, index) => (
                    <li
                        key={index}
                        className="football-field-information-option__item"
                        onClick={() => dispatch(handleCaptain)}
                    >
                        {item}
                    </li>
                ))}
            </ul>
        </div>
    )
}
