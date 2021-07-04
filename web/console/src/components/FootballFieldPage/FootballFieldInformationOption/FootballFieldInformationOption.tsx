/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import './FootballFieldInformationOption.scss';

import { FotballFieldInformationLine } from '../../../types/fotballerCard';

import triangle from '../../../img/FootballFieldPage/triangle.png'
import { useState } from 'react';
import { Styles } from '../../../utils/footballField';

export const FootballFieldInformationOption: React.FC<{ props: FotballFieldInformationLine }> = ({ props }) => {
    const [state, changeState] = useState(true);

    const styles = new Styles(state)

    console.log(props)

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
                    style={{ transform: styles.triangleStyle }}
                    alt="triangle img"
                    id={`triangle-${props.id}`}
                />
            </div>
            <ul
                style={{ height: styles.listStyle }}
                className="football-field-information-option__list"
                id={props.id}
            >
                {props.options.map((item, index) => (
                    <li
                        key={index}
                        className="football-field-information-option__item"
                        onClick={() => { }}
                    >
                        {item}
                    </li>
                ))}
            </ul>
        </div>
    )
}
