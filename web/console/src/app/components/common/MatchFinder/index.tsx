// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { toast } from 'react-toastify';

import { Timer } from './Timer';
import { AutoCloseTimer } from './AutoCloseTimer';

import { QueueClient } from '@/api/queue';
import { YOU_CONFIRM_PLAY_MESSAGE } from '@/api/webSocketClient';
import { RootState } from '@/app/store';
import { startSearchingMatch } from '@/app/store/actions/clubs';

import './index.scss';

const MatchFinder: React.FC = () => {
    const { squad } = useSelector((state: RootState) => state.clubsReducer.activeClub)
    const { isSearchingMatch } = useSelector((state: RootState) => state.clubsReducer);

    const dispatch = useDispatch();

    /** Indicates if match is found. */
    const [isMatchFound, setIsMatchFound] = useState<boolean>(false);

    /** closes MatchFinder component. */
    const closeMatchFinder = () => {
        dispatch(startSearchingMatch(false));
        setIsMatchFound(false);
        queueClient.ws.close();
    };

    const queueClient = new QueueClient();

    /** exposes confirm match logic */
    const confirmMatch = async () => {
        /** TODO: also here handles errors */
        await queueClient.sendAction('confirm', squad.id);
        queueClient.ws.onmessage = async ({ data }) => {
            /** TODO: dispatch match score here */
        };

        queueClient.ws.onerror = (event: Event) => {
            toast.error('Something wrong, please, try later.', {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
        };

        setIsMatchFound(false);

        /** TODO: here will be check for match score and after auto close connection. */
        dispatch(startSearchingMatch(false));
    };

    /** exposes start searching match logic. */
    const startSearchMatch = async () => {
        await queueClient.startSearch('startSearch', squad.id);
        queueClient.ws.onmessage = async ({ data }) => {
            const jsonEventMessage = JSON.parse(data);
            queueClient.changeEventMessage(jsonEventMessage.message);
            const eventMessage = queueClient.getEventMessage();
            setIsMatchFound(eventMessage === YOU_CONFIRM_PLAY_MESSAGE);
        };

        queueClient.ws.onerror = (event: Event) => {
            toast.error('Something wrong, please, try later.', {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
        };
    };

    useEffect(() => {
        !isMatchFound && startSearchMatch();
    }, [isSearchingMatch, isMatchFound]);

    return isSearchingMatch ? <section className={isMatchFound ? 'match-finder__wrapper' : ''}>
        <div className="match-finder">
            <h1 className="match-finder__title">
                {isMatchFound ? 'YOUR MATCH WAS FOUND' : 'LOOKING FOR A MATCH'}
            </h1>
            {isMatchFound ? <AutoCloseTimer /> : <Timer />}
            <div className="match-finder__form">
                {isMatchFound ? <input
                    className="match-finder__form__accept"
                    value="Accept"
                    type="button"
                    onClick={confirmMatch}
                /> : <input
                    className="match-finder__form__accept-not-allowed"
                    value="Accept"
                    type="button"
                />}
                <input
                    className="match-finder__form__cancel"
                    value="Cancel"
                    type="button"
                    onClick={closeMatchFinder}
                />
            </div>
        </div>
    </section> : <></>;
};

export default MatchFinder;
