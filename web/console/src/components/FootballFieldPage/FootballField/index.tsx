/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import { useDispatch, useSelector } from 'react-redux';

import { DragEvent } from 'react';

import { FootballFieldCardSelection } from '../FootballFieldCardSelection';
import { FootballFieldInformation } from '../FootballFieldInformation';
import { FootballFieldPlayingArea } from '../FotballFieldPlayingArea';

import { RootState } from '../../../store';
import { removeCard } from '../../../store/reducers/footballField';

import './index.scss';

export const FootballField: React.FC = () => {
    const dispatch = useDispatch();
    const dragItemPosition = useSelector((state: RootState) => state.fieldReducer.options.dragStart);
    /** prevent default user agent action */
    function dragOverHandler(e: DragEvent<HTMLDivElement>) {
        e.preventDefault();
    };

    /** TO DO: ADD TYPE FOR Event */
    function drop(e: any) {
        if (e.target.className === 'football-field__wrapper') {
            dispatch(removeCard(dragItemPosition));
        }
    };

    return (
        <div className="football-field"
            onDrop={e => drop(e)}
            onDragOver={e => dragOverHandler(e)}
        >
            <h1 className="football-field__title">Football Field</h1>
            <div className="football-field__wrapper"
            >
                <FootballFieldPlayingArea />
                <FootballFieldInformation />
            </div>
            <FootballFieldCardSelection />
        </div>
    );
};
