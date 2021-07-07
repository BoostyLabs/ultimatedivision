/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';

import { UltimateDIvisionPlayerCard } from '../../../UltimateDIvisionPlayerCard/UltimateDIvisionPlayerCard';

import { Card }
    from '../../../../store/reducers/footballerCard';
import { RootState } from '../../../../store';
import { useDispatch, useSelector } from 'react-redux';
import { handleCard }
    from '../../../../store/reducers/footballField';

import './PlayingAreaFootballerCard.scss';

export const PlayingAreaFootballerCard: React.FC<{ card: Card, place?: string }> = ({ card, place }) => {

    const dispatch = useDispatch();
    const chosenCard = useSelector((state: RootState) => state.fieldReducer.options.chosedCard);

    return (
        <div
            onClick={place ? () => { } : () => dispatch(handleCard(card, chosenCard))}
            className="football-field-card"
            data-background={card.mainInfo.backgroundType}
        >
            <UltimateDIvisionPlayerCard
                card={card}
                parentClassName={"football-field-card"}
            />
        </div>
    );
};

