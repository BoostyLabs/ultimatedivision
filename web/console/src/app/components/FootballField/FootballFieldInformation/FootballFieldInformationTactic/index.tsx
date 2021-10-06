// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';
import { useDispatch } from 'react-redux';

import { FootballFieldInformationLine } from '@/app/types/footballField';
import { setTactic } from '@/app/store/actions/footballField';

import { DropdownStyle } from '@/app/utils/dropdownStyle';

import triangle from '@static/img/FootballFieldPage/triangle.svg';

import './index.scss';

export const FootballFieldInformationTactic: React.FC<{ props: FootballFieldInformationLine }> = ({ props }) => {
    const [optionVisibility, changeVisibility] = useState(false);
    const LIST_HEIGHT = 90;
    const optionStyle = new DropdownStyle(optionVisibility, LIST_HEIGHT);

    const dispatch = useDispatch();

    return (
        <div className="football-field-information-option">
            <div
                className="football-field-information-option__heading"
                onClick={() => changeVisibility(prev => !prev)}
            >
                <h4 className="football-field-information-option__title">
                    {props.title}
                </h4>
                <img
                    className="football-field-information-option__image"
                    src={triangle}
                    style={{ transform: optionStyle.triangleRotate }}
                    alt="triangle img"
                    id={`triangle-${props.id}`}
                />
            </div>
            <ul
                style={{ height: optionStyle.listHeight }}
                className="football-field-information-option__list"
                id={props.id}
            >
                {props.options.map((item, index) =>
                    <li
                        key={index}
                        className="football-field-information-option__item"
                        onClick={() => dispatch(setTactic)}
                    >
                        {item}
                    </li>,
                )}
            </ul>
        </div>
    );
};