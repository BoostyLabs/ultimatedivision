/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import { useState } from 'react';
import { useDispatch } from 'react-redux';

import { FotballFieldInformationLine } from '@/app/types/fotballerCard';
import { handleCaptain } from '@/app/store/reducers/footballField';

import { DropdownStyle } from '@/app/utils/dropdownStyle';

import triangle from '@static/img/FootballFieldPage/triangle.svg';

import './index.scss';

export const FootballFieldInformationCaptain: React.FC<{ props: FotballFieldInformationLine }> = ({ props }) => {
    const [optionVisibility, changeVisibility] = useState(false);

    const optionStyle = new DropdownStyle(optionVisibility, 90);

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
                style={{ height: optionStyle.triangleRotate }}
                className="football-field-information-option__list"
                id={props.id}
            >
                {props.options.map((item, index) =>
                    <li
                        key={index}
                        className="football-field-information-option__item"
                        onClick={() => dispatch(handleCaptain)}
                    >
                        {item}
                    </li>,
                )}
            </ul>
        </div>
    );
};
