// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { DragEvent } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { FootballFieldCardSelection } from
    '@components/FootballField/FootballFieldCardSelection';
import { FootballFieldPlayingArea } from
    '@components/FootballField/FotballFieldPlayingArea';
import { RootState } from '@/app/store';
import { deleteCard, getClub } from '@/app/store/actions/club';

import './index.scss';
import { ExactCardPath } from '@/app/types/club';

const FootballField: React.FC = () => {
    const dispatch = useDispatch();
    dispatch(getClub());
    const dragStartIndex = useSelector(
        (state: RootState) => state.clubReducer.options.dragStart
    );

    const fieldSetup = useSelector((state: RootState) => state.clubReducer);
    const cardSelectionVisibility = useSelector((state: RootState) => state.clubReducer.options.showCardSeletion);

    /** prevent default user agent action */
    function dragOverHandler(e: DragEvent<HTMLDivElement>) {
        e.preventDefault();
    };

    /** TO DO: ADD TYPE FOR Event */
    function drop(e: any) {
        if (e.target.className === 'football-field__wrapper') {
            dragStartIndex &&
            dispatch(deleteCard(new ExactCardPath(fieldSetup.squad, fieldSetup.squadCards[dragStartIndex].cardId)));
        }
    };

    return (
        <div className="football-field"
            onDrop={e => drop(e)}
            onDragOver={e => dragOverHandler(e)}
        >
            <h1 className="football-field__title">Football Field</h1>
            <FootballFieldPlayingArea />
            <div
                style={{ height: cardSelectionVisibility ? 'unset' : '0' }}
                className="football-field__wrapper"
            >
                < FootballFieldCardSelection />
            </div>
        </div>
    );
};

export default FootballField;
