// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';
import { useDispatch } from 'react-redux';

import { FieldControl } from '@/app/types/club';

import { DropdownStyle } from '@/app/utils/dropdownStyle';

import triangle from '@static/img/FootballFieldPage/triangle.svg';

import './index.scss';

export const FootballFieldControl: React.FC<{ props: FieldControl }> = ({ props }) => {
    const [optionVisibility, changeVisibility] = useState(false);
    const optionStyle = new DropdownStyle(optionVisibility);

    const dispatch = useDispatch();

    return (
        <div className="football-field-control">
            <div
                className="football-field-control__heading"
                onClick={() => changeVisibility(prev => !prev)}
            >
                <h4 className="football-field-control__title">
                    {props.title}
                </h4>
                <img
                    className="ffootball-field-control__image"
                    src={triangle}
                    style={{ transform: optionStyle.triangleRotate }}
                    alt="triangle img"
                    id={`triangle-${props.id}`}
                />
            </div>
            <ul
                style={{ height: optionStyle.listHeight }}
                className="football-field-control__list"
                id="0"
            >
                {props.options.map((item, index) =>
                    <li
                        key={index}
                        className="football-field-control__item"
                        onClick={() => dispatch(props.action(item))}
                    >
                        {item}
                    </li>

                )}
            </ul>
        </div>
    );
};
