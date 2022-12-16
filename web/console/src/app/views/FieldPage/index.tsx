// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { DragEvent, useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { FieldCardSelection } from '@/app/components/Field/FieldCardSelection';
import { FieldPlayingArea } from '@/app/components/Field/FieldPlayingArea';
import { RegistrationPopup } from '@/app/components/common/Registration';

import { NotFoundError, UnauthorizedError } from '@/api';
import { useLocalStorage } from '@/app/hooks/useLocalStorage';
import { RootState } from '@/app/store';
import { createClubs, deleteCard, getClubs } from '@/app/store/actions/clubs';
import { CardEditIdentificators } from '@/api/club';
import { ToastNotifications } from '@/notifications/service';

import './index.scss';

const FootballField: React.FC = () => {
    const dispatch = useDispatch();
    /** Indicates if registration required. */
    const [isRegistrationRequired, setIsRegistrationRequired] = useState(false);

    const [setLocalStorageItem, getLocalStorageItem] = useLocalStorage();

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

                    setLocalStorageItem('IS_LOGGINED', false);

                    return;
                }

                if (!(error instanceof NotFoundError)) {
                    ToastNotifications.notFound();

                    return;
                }
                try {
                    await dispatch(createClubs());
                } catch (error) {
                    ToastNotifications.couldNotCreateClub();
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
            dragStartIndex !== null &&
                dispatch(
                    deleteCard(
                        new CardEditIdentificators(squad.clubId, squad.id, club.squadCards[dragStartIndex].card.id)
                    )
                );
        }
    }

    return (
        <>
            {isRegistrationRequired && <RegistrationPopup closeRegistrationPopup={closeRegistrationPopup} />}
            <div className="football-field" onDrop={(e) => drop(e)} onDragOver={(e) => dragOverHandler(e)}>
                <FieldPlayingArea />
                <div className="football-field__wrapper">{cardSelectionVisibility && <FieldCardSelection />}</div>
            </div>
        </>
    );
};

export default FootballField;
