// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { DragEvent, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { toast } from 'react-toastify';

import { FootballFieldCardSelection } from
    '@components/FootballField/FootballFieldCardSelection';
import { FootballFieldPlayingArea } from
    '@components/FootballField/FotballFieldPlayingArea';

import { RootState } from '@/app/store';
import { createClub, deleteCard, getClub } from '@/app/store/actions/club';

import './index.scss';
import { CardEditIdentificators } from '@/app/types/club';

const FootballField: React.FC = () => {
    const dispatch = useDispatch();
    useEffect(() => {
        /** TODO: refactor method call after backend response changes */
        (async function setClub() {
            try {
                await dispatch(getClub());
            } catch (error: any) {
                try {
                    await dispatch(createClub());
                } catch (error: any) {
                    toast.error('Something went wrong', {
                        position: toast.POSITION.TOP_RIGHT,
                        theme: 'colored',
                    });
                }
            }
        })();
    }, []);
    const dragStartIndex = useSelector(
        (state: RootState) => state.clubReducer.options.dragStart
    );

    const squad = useSelector((state: RootState) => state.clubReducer.squad);
    const club = useSelector((state: RootState) => state.clubReducer);
    const cardSelectionVisibility = useSelector((state: RootState) => state.clubReducer.options.showCardSeletion);

    /** prevent default user agent action */
    function dragOverHandler(e: DragEvent<HTMLDivElement>) {
        e.preventDefault();
    };

    /** TO DO: ADD TYPE FOR Event */
    function drop(e: any) {
        if (e.target.className === 'football-field__wrapper') {
            dragStartIndex &&
                dispatch(deleteCard(new CardEditIdentificators(squad.clubId, squad.id, club.squadCards[dragStartIndex].cardId)));
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
