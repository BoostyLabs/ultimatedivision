// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { PlayerCard } from '@components/common/PlayerCard';

import { Card } from '@/card';
import { CardEditIdentificators } from '@/api/club';
import { deleteCard } from '@/app/store/actions/clubs';
import { RootState } from '@/app/store';

import './index.scss';

type FootballerCardProps = {
    card: Card;
    index?: number;
    place?: string;
    setCurrentCard: (currentCard: Element | null) => void;
    currentCard: Element | null;
};

export const FootballerCard: React.FC<FootballerCardProps> = ({ card, index, setCurrentCard, currentCard }) => {
    const dispatch = useDispatch();
    const squad = useSelector((state: RootState) => state.clubsReducer.activeClub.squad);

    /** Remove player card implementation. */
    function handleDeletion(e: React.MouseEvent<HTMLInputElement>) {
        e.stopPropagation();
        e.preventDefault();

        dispatch(deleteCard(new CardEditIdentificators(squad.clubId, squad.id, card.id)));
    }

    /** Changing the state of a card class. */
    const visibilityBlock = currentCard && parseInt(currentCard.id) === index ? '-active' : '-inactive';

    return (
        <div className="footballer-card">
            <div className={`football-field-card__wrapper${visibilityBlock}`}></div>
            <PlayerCard id={card.id} className="footballer-card" index={index} />
            <div onClick={handleDeletion} className={`footballer-card__control${visibilityBlock}`}></div>
        </div>
    );
};
