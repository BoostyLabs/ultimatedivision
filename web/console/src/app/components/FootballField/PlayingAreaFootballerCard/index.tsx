// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { PlayerCard } from '@components/common/PlayerCard';

import { CardEditIdentificators } from '@/api/club';
import { deleteCard } from '@/app/store/actions/clubs';
import { RootState } from '@/app/store';
import { Card } from '@/card';

import './index.scss';

export const PlayingAreaFootballerCard: React.FC<{ card: Card; index?: number; place?: string }> = ({ card }) => {
    const dispatch = useDispatch();
    const squad = useSelector((state: RootState) => state.clubsReducer.squad);
    const [isVisibile, setIsVisibile] = useState(false);
    const visibilityBlock = isVisibile ? '-active' : '-inactive';

    /** show/hide delete block, preventing scroll to cardSelection */
    function handleVisibility(e: React.MouseEvent<HTMLInputElement>) {
        e.stopPropagation();
        setIsVisibile((prevVisability) => !prevVisability);
    }
    /** remove player card implementation */
    function handleDeletion(e: React.MouseEvent<HTMLInputElement>) {
        e.stopPropagation();
        e.preventDefault();
        dispatch(deleteCard(new CardEditIdentificators(squad.clubId, squad.id, card.id)));
    }

    return (
        <div onClick={handleVisibility} className="football-field-card">
            <div
                className={`football-field-card__wrapper${visibilityBlock}`}
            ></div>
            <PlayerCard card={card} parentClassName="football-field-card" />
            <div
                onClick={handleDeletion}
                className={`football-field-card__control${visibilityBlock}`}
            >
                &#10006; delete a player
            </div>
        </div>
    );
};
