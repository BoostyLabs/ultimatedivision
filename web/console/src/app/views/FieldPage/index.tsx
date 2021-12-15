// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { DragEvent, useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { toast } from 'react-toastify';

import { FieldCardSelection } from '@/app/components/Field/FieldCardSelection';
import { FieldPlayingArea } from '@/app/components/Field/FieldPlayingArea';
import { RegistrationPopup } from '@/app/components/common/Registration/Registration';

import { NotFoundError, UnauthorizedError } from '@/api';
import { RootState } from '@/app/store';
import { createClubs, deleteCard, getClubs } from '@/app/store/actions/clubs';
import { CardEditIdentificators } from '@/api/club';

import './index.scss';

const FootballField: React.FC = () => {
    const dispatch = useDispatch();
    /** Indicates if registration required. */
    const [isRegistrationRequired, setIsRegistrationRequired] = useState(false);

    /** Closes RegistrationPopup componnet. */
    const closeRegistrationPopup = () => {
        setIsRegistrationRequired(false);
    };

    useEffect(() => {
        (async function setClub() {
            try {
                await dispatch(getClubs());
            } catch (error: any) {
                if (error instanceof UnauthorizedError) {
                    setIsRegistrationRequired(true);

                    window.localStorage.setItem('IS_LOGGINED', 'false');

                    return;
                }

                if (!(error instanceof NotFoundError)) {
                    toast.error('Something went wrong', {
                        position: toast.POSITION.TOP_RIGHT,
                        theme: 'colored',
                    });

                    return;
                }
                try {
                    await dispatch(createClubs());
                } catch (error: any) {
                    toast.error('Something went wrong', {
                        position: toast.POSITION.TOP_RIGHT,
                        theme: 'colored',
                    });
                }
            }
        })();
    }, []);

    const dragStartIndex = useSelector((state: RootState) => state.clubsReducer.options.dragStart);

    const squad = useSelector((state: RootState) => state.clubsReducer.activeClub.squad);
    const club = useSelector((state: RootState) => state.clubsReducer.activeClub);
    const cardSelectionVisibility = useSelector((state: RootState) => state.clubsReducer.options.showCardSeletion);

    /** prevent default user agent action */
    function dragOverHandler(e: DragEvent<HTMLDivElement>) {
        e.preventDefault();
    }

    /** TO DO: ADD TYPE FOR Event */
    function drop(e: any) {
        if (e.target.className === 'football-field__wrapper') {
            dragStartIndex &&
                dispatch(
                    deleteCard(
                        new CardEditIdentificators(
                            squad.clubId,
                            squad.id,
                            club.squadCards[dragStartIndex].card.id
                        )
                    )
                );
        }
    }

    return (
        <>
            {isRegistrationRequired &&
                <RegistrationPopup
                    closeRegistrationPopup={closeRegistrationPopup}
                />
            }
            <div
                className="football-field"
                onDrop={(e) => drop(e)}
                onDragOver={(e) => dragOverHandler(e)}
            >
                <h1 className="football-field__title">Football Field</h1>
                <FieldPlayingArea />
                <div className="football-field__wrapper">
                    {cardSelectionVisibility && <FieldCardSelection />}
                </div>
            </div>
        </>
    );
};

export default FootballField;
