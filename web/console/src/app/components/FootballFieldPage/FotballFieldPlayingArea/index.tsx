//Copyright (C) 2021 Creditor Corp. Group.
//See LICENSE for copying information.

import { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { FootballFieldInformation } from '@components/FootballFieldPage/FootballFieldInformation';
import { PlayingAreaFootballerCard } from '@components/FootballFieldPage/PlayingAreaFootballerCard';

import { RootState } from '@/app/store';
import { cardSelectionVisibility, choseCardPosition, exchangeCards, removeCard, setDragStart, setDragTarget }
    from '@/app/store/reducers/footballField';

import './index.scss';
import { FootballFieldCard } from '@/app/types/footballField';

export const FootballFieldPlayingArea: React.FC = () => {
    const formation = useSelector((state: RootState) => state.fieldReducer.options.formation);
    const dragStartIndex = useSelector((state: RootState) => state.fieldReducer.options.dragStart);

    const dispatch = useDispatch();
    const fieldSetup = useSelector((state: RootState) => state.fieldReducer);

    /** MouseMove event Position */
    const [mousePosition, setMousePosition] = useState({ x: 0, y: 0 });
    /** This var created to not allow mouseUpEvent without Dragging before it */
    const [isDragging, handleDrag] = useState(false);
    /** outer padding of playingArea */
    const [outerOffset, handleOffset] = useState({ x: 0, y: 0 });

    const DEFAULT_VALUE = 0;
    const OFFSET_TOP = 330;

    /** with getBoundingClientRect() we gettins outer padding of playingArea on any screen width and scale */
    useEffect(() => {
        const playingArea = document.getElementById('playingArea');
        if (playingArea) {
            const position = playingArea.getBoundingClientRect();
            const HALF_OF_CARD_WIDTH = 60;
            const HALF_OF_CARD_HEIGHT = 100;

            handleOffset({
                x: position.x + HALF_OF_CARD_WIDTH,
                y: position.y + HALF_OF_CARD_HEIGHT,
            });
        }
    }, []);
    const useMousePosition = (ev: any) => {
        setMousePosition({ x: ev.pageX, y: ev.pageY });
    };

    /** Add card position, and shows card selection */
    function handleClick(index: number) {
        dispatch(choseCardPosition(index));
        dispatch(cardSelectionVisibility(true));
        setTimeout(() => {
            window.scroll(0, 1200);
        }, 100);
    };

    /** getting dragged card index and changing state to allow mouseUp */
    function dragStart(e: any, index: number = DEFAULT_VALUE): void {
        handleDrag(true);
        dispatch(setDragStart(index));
    }
    /** eslint-disable */
    /** getting second drag index  and exchanging with first index*/
    function onMouseUp(e: any, index: number = DEFAULT_VALUE): void {
        e.stopPropagation();

        if (isDragging && dragStartIndex !== null) {
            dispatch(setDragTarget(index));
            dispatch(exchangeCards(dragStartIndex, fieldSetup.options.dragTarget));
        }

        dispatch(setDragTarget());
        dispatch(setDragStart());
        handleDrag(false);
    }

    /** when we release card not on target it just brings it on start position*/
    function mouseUpOnArea(e: any) {
        e.stopPropagation();
        dispatch(setDragStart());
    }

    /** deleting card when release beyond playing area */
    function removeFromArea() {
        if (isDragging) {
            dispatch(removeCard(dragStartIndex));
            dispatch(setDragStart());
            handleDrag(false);
        }
    }

    const { x, y } = mousePosition;

    return (
        <div
            className="playing-area__wrapper"
            onMouseMove={(ev) => useMousePosition(ev)}
            onMouseUp={removeFromArea}
            style={isDragging ? { cursor: 'not-allowed' } : {}}
        >
            <div className="playing-area" id="playingArea">
                <div
                    style={dragStartIndex ? { cursor: 'grabbing' } : {}}
                    className={`playing-area__${formation}`}
                    onMouseUp={mouseUpOnArea}
                >
                    {fieldSetup.cardsList.map((card, index) => {
                        const data = card.cardData;
                        const equality = dragStartIndex === index;
                        // TO DO: change style by some class to change style in card

                        return (
                            <div
                                style={
                                    equality
                                        ? { left: x - outerOffset.x, top: y - OFFSET_TOP, zIndex: 5, pointerEvents: 'none' }
                                        : undefined
                                }
                                key={index}
                                className={`playing-area__${formation}__${data ? 'card' : 'empty-card'}`}
                                onClick={(e) => handleClick(index)}
                                onDragStart={(e) => dragStart(e, index)}
                                onMouseUp={(e) => onMouseUp(e, index)}
                                draggable={true}
                            >
                                {
                                    data && <PlayingAreaFootballerCard card={data} index={index} place={'PlayingArea'} />
                                }
                            </div>
                        );
                    })}
                </div>
                <div className={`playing-area__${formation}-shadows`}>
                    {fieldSetup.cardsList.map((card: FootballFieldCard, index: number) => {
                        const data = card.cardData;

                        return (
                            <div
                                className={`playing-area__${formation}-shadows__card`}
                                key={index}
                            >
                                {data &&
                                    <img
                                        // If data exist it has maininfo, but TS do not let me use it even with check
                                        // @ts-ignore
                                        src={data.mainInfo.shadowType}
                                        alt="card shadow"
                                        className={`playing-area__${formation}-shadows__shadow`}
                                    />
                                }
                            </div>
                        );
                    })}
                </div>
            </div>
            <FootballFieldInformation />
        </div>
    );
};

