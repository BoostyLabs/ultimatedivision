// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { PlayerCard } from '@components/common/PlayerCard';

import { Card, DeleteCardVisability } from '@/card';
import { CardEditIdentificators } from '@/api/club';
import { deleteCard } from '@/app/store/actions/clubs';
import { RootState } from '@/app/store';

import './index.scss';

type FootballerCardProps = {
    card: Card;
    index?: number;
    place?: string;
    setDeleteCardVisability: (
        deleteCardVisability: DeleteCardVisability
    ) => void;
    deleteCardVisability: DeleteCardVisability;
};

export const FootballerCard: React.FC<FootballerCardProps> = ({
    card,
    setDeleteCardVisability,
    deleteCardVisability,
}) => {
    const dispatch = useDispatch();
    const squad = useSelector(
        (state: RootState) => state.clubsReducer.activeClub.squad
    );

    /** Show/hide delete block, preventing scroll to cardSelection. */
    const handleVisibility = (e: React.MouseEvent<HTMLInputElement>) => {
        e.stopPropagation();

        setDeleteCardVisability(
            deleteCardVisability.id === card.id
                ? new DeleteCardVisability(
                    deleteCardVisability.id,
                    !deleteCardVisability.isVisible
                )
                : new DeleteCardVisability(card.id, true)
        );
    };

    /** Remove player card implementation. */
    function handleDeletion(e: React.MouseEvent<HTMLInputElement>) {
        e.stopPropagation();
        e.preventDefault();

        setDeleteCardVisability(new DeleteCardVisability('', false));

        dispatch(
            deleteCard(
                new CardEditIdentificators(squad.clubId, squad.id, card.id)
            )
        );
    }

    /** Changing the state of a card class. */
    const visibilityBlock =
        card.id === deleteCardVisability.id && deleteCardVisability.isVisible
            ? '-active'
            : '-inactive';

    return (
        <div onClick={handleVisibility} className="footballer-card">
            <div
                className={`football-field-card__wrapper${visibilityBlock}`}
            ></div>
            <PlayerCard id={card.id} className="footballer-card" />
            <div
                onClick={handleDeletion}
                className={`footballer-card__control${visibilityBlock}`}
            >
                &#10006; delete a player
            </div>
        </div>
    );
};
