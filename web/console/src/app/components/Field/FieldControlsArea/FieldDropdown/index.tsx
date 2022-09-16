// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import { useDispatch, useSelector } from 'react-redux';
import { useMemo } from 'react';

import { RootState } from '@/app/store';

import addNewIcon from '@static/img/FieldPage/add-new.png';

import './index.scss';

export const FieldDropdown: React.FC<{ option: any,isMobile?:boolean }> = ({ option, isMobile }) => {
    const dispatch = useDispatch();

    const squad = useSelector((state: RootState) => state.clubsReducer.activeClub.squad);

    const ADD_NEW_BUTTON = 1;
    const NO_ADD_NEW_BUTTON = 0;

    const AddNewElement = option.title !== 'formation' ? ADD_NEW_BUTTON : NO_ADD_NEW_BUTTON;

    const columnsAmount = useMemo(
        () => Math.ceil((option.options.length + AddNewElement) / option.columnElements),
        [option.options.length, option.columnElements]
    );

    const sendDesktopOptions = (event?: any) => {
        if (event) {
            switch (option.title) {
            case 'formation':
                dispatch(option.action(squad, event.target.value));
                break;
            case 'club':
                dispatch(option.action(event.target.value));
                break;
            case 'squad':
                option.action(event.target.value);
                break;
            default:
                break;
            }
        } else {
            // @ts-ignore
            document.querySelector(`input[name=${option.title}]:checked`).checked = false;
        }
    }
    const sendCheckedOption = (event?: any) => {
        if (isMobile) {
            option.action(event.target.value)
        }
        else {
            sendDesktopOptions(event)
        }
    };

    /** TODO: add new field button */
    const addNewElement = () => {
        sendCheckedOption();
    };

    return (
        <ul
            className={`field-dropdown field-dropdown__${columnsAmount}--columns__${option.columnElements}--rows field-dropdown__${option.title} `}
        >
            {option.options.map((item: any, index: number) => {
              
                const fieldName = item.hasOwnProperty(option.fieldName) ? item[option.fieldName] : item;
                const fieldId = item.hasOwnProperty(option.fieldId) ? item[option.fieldId] : item;
                const defaultChecked = item[option.fieldId]
                    ? item[option.fieldId] === option.currentValue
                    : item === option.currentValue;
                return (
                    <li key={`${option.title}-${index}`} className={'field-dropdown__item'}>
                        <input
                            type="radio"
                            className="field-dropdown__item__input"
                            name={option.title}
                            id={fieldId}
                            value={fieldId}
                            defaultChecked={defaultChecked}
                            onChange={sendCheckedOption}
                        />
                        <label htmlFor={fieldId} className="field-dropdown__item__label">
                            <span className="field-dropdown__item__text">
                                {option.fieldText ? `${option.fieldText} ${fieldName}` : fieldName}
                            </span>
                            <span className="field-dropdown__item__radio"></span>
                        </label>
                    </li>
                );
            })}
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
