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
import { exchangeCards }
    from '../../../../../store/reducers/footballField';
import { useState } from 'react';

export const PlayingFormation_442: React.FC<{ props: FootballField }> = ({ props }) => {
    const dispatch = useDispatch();

    const [currentPosition, handleDrag] = useState(-1);
    const [dragTarget, handleDragTarget] = useState(-1);

    function dragOverHandler(e: any, index: number) {
        e.preventDefault();
        e.target.style.boxShadow = '0 0 10px white';
        handleDragTarget(index);
    }
    function dragleaveHandler(e: any) {
        e.target.style.boxShadow = 'none';

    }
    function dragEndHandler(e: any) {
        e.target.style.boxShadow = 'none';

    }
    function dropHandler(e: any) {
        e.target.style.boxShadow = 'none';
        dispatch(exchangeCards(currentPosition, dragTarget));
    }

    return (
        <div className="playing-formation-442">
            {props.cardsList.map((card, index) => {
                const data = card.cardData;
                return (
                    <div
                        id={index.toString()}
                        key={index}
                        className="playing-formation-442__card box"
                        draggable={true}
                        onDragOver={e => dragOverHandler(e, index)}
                        onDragLeave={e => dragleaveHandler(e)}
                        onMouseDown={(e: any) => handleDrag(index)}
                        onDragEnd={e => dragEndHandler(e)}
                        onDrop={e => dropHandler(e)}
                    >
                        {
                            data
                                ? <PlayingAreaFootballerCard card={data} index={index} place={'PlayingArea'} />
                                : <a
                                    id={index.toString()}
                                    onClick={() => dispatch(choseCardPosition(index))}
                                    href="#cardList"
                                    className="playing-formation-442__link"
                                    style={{ color: 'red' }}
                                >
                                </a>
                        }
                    </div>
                )
            })}
        </div>
    )
}
