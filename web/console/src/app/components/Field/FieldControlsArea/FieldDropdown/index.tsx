// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import { useDispatch, useSelector } from 'react-redux';
import { useMemo } from 'react';

import { RootState } from '@/app/store';

import addNewIcon from '@static/img/FieldPage/add-new.png';

import './index.scss';

export const FieldDropdown: React.FC<{ option: any }> = ({ option }) => {
    const dispatch = useDispatch();

    const squad = useSelector((state: RootState) => state.clubsReducer.activeClub.squad);

    const ADD_NEW_BUTTON = 1;
    const NO_ADD_NEW_BUTTON = 0;

    const AddNewElement = option.title !== 'formation' ? ADD_NEW_BUTTON : NO_ADD_NEW_BUTTON;

    const columnsAmount = useMemo(
        () => Math.ceil((option.options.length + AddNewElement) / option.columnElements),
        [option.options.length, option.columnElements]
    );

    const sendCheckedOption = (event?: any) => {
        if (event) {
            dispatch(option.action(squad, event.target.value));
        } else {
            event.target.value = null;
        }
    };

    /** TODO: add new field button */
    const addNewElement = () => {};

    return (
        <ul
            className={`field-dropdown field-dropdown__${columnsAmount}--columns__${option.columnElements}--rows field-dropdown__${option.title} `}
        >
            {option.options.map((item: any, index: number) =>
                <li key={`${option.title}-${index}`} className={'field-dropdown__item '}>
                    <input
                        type="radio"
                        className="field-dropdown__item__input"
                        name={option.title}
                        id={item}
                        value={item}
                        defaultChecked={item === option.currentValue}
                        onChange={sendCheckedOption}
                    />
                    <label htmlFor={item} className="field-dropdown__item__label">
                        <span className="field-dropdown__item__text"> {item}</span>
                        <span className="field-dropdown__item__radio"></span>
                    </label>
                </li>
            )}
            {option.title !== 'formation' &&
                <li className={'field-dropdown__item'}>
                    <button className="field-dropdown__item__button" onClick={addNewElement}>
                        <span className="field-dropdown__item__button__text">Add new</span>
                        <span className="field-dropdown__item__button__icon">
                            <img src={addNewIcon} alt="add-new" />
                        </span>
                    </button>
                </li>
            }
        </ul>
    );
};
