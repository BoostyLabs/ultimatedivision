/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import './PlayingFormation_442.scss';
import { FootballField } from '../../../../../types/footballField';
import { useDispatch } from 'react-redux';
import { choseCardPosition }
    from '../../../../../store/reducers/footballField';
import { PlayingAreaFootballerCard }
    from '../../../FootballFieldCardSelection/PlayingAreaFootballerCard/PlayingAreaFootballerCard';

export const PlayingFormation_442: React.FC<{ props: FootballField }> = ({ props }) => {
    const dispatch = useDispatch();

    return (
        <div className="playing-formation-442">
            {props.cardsList.map((card, index) => {
                const data = card.cardData;
                return (
                    <div
                        key={index}
                        className="playing-formation-442__card"
                    >
                        {
                            data
                                ? <PlayingAreaFootballerCard card={data} index={index} place={'PlayingArea'} />
                                : <a
                                    onClick={() => dispatch(choseCardPosition(index))}
                                    href="#cardList"
                                    className="playing-formation-433__link"
                                >
                                </a>
                        }
                    </div>
                )
            })}
        </div>
    )
}
