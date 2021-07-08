/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React, { DragEvent } from 'react';
import './PlayingFormation_424.scss';
import { FootballField } from '../../../../../types/footballField';
import { useDispatch, useSelector } from 'react-redux';
import { choseCardPosition, setDragStart, setDragTarget }
    from '../../../../../store/reducers/footballField';
import { PlayingAreaFootballerCard }
    from '../../../FootballFieldCardSelection/PlayingAreaFootballerCard/PlayingAreaFootballerCard';
import { exchangeCards }
    from '../../../../../store/reducers/footballField';
import { RootState } from '../../../../../store';

export const PlayingFormation_424: React.FC<{ props: FootballField }> = ({ props }) => {
    const dispatch = useDispatch();
    const fieldSetup = useSelector((state: RootState) => state.fieldReducer.options);
    
    function dragOverHandler(e: any) {
        e.preventDefault();
    };
    
    function dropHandler(e: DragEvent<HTMLDivElement>, index: number) {
        dispatch(setDragTarget(index));
        dispatch(exchangeCards(fieldSetup.dragStart, fieldSetup.dragTarget));
    };

    return (
        <div className="playing-formation-424">
            {props.cardsList.map((card, index) => {
                const data = card.cardData;
                return (
                    <div
                        key={index}
                        className="playing-formation-424__card box"
                        draggable={true}
                        onDragOver={e => dragOverHandler(e)}
                        onMouseDown={() => dispatch(setDragStart(index))}
                        onDrop={e => dropHandler(e, index)}
                    >
                        {
                            data
                                ? <PlayingAreaFootballerCard card={data} index={index} place={'PlayingArea'} />
                                : <a
                                    onClick={() => dispatch(choseCardPosition(index))}
                                    href="#cardList"
                                    className="playing-formation-424__link"
                                >
                                </a>
                        }
                    </div>
                )
            })}
        </div>
    )
}
