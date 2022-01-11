// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { FieldControlsArea } from '@components/Field/FieldControlsArea';
import { FootballerCard } from '@components/Field/FootballerCard';

import { CardEditIdentificators } from '@/api/club';
import { RootState } from '@/app/store';
import { Card } from '@/card';
import { SquadCard } from '@/club';
import {
    cardSelectionVisibility,
    changeCardPosition,
    choosePosition,
    deleteCard,
    setDragStart,
    swapCards,
} from '@/app/store/actions/clubs';

import './index.scss';

export const FieldPlayingArea: React.FC = () => {
    const dispatch = useDispatch();

    const cards = useSelector(
        (state: RootState) => state.cardsReducer.cardsPage.cards
    );
    const formation = useSelector(
        (state: RootState) => state.clubsReducer.activeClub.squad.formation
    );
    const dragStartIndex = useSelector(
        (state: RootState) => state.clubsReducer.options.dragStart
    );
    const club = useSelector(
        (state: RootState) => state.clubsReducer.activeClub
    );
    const squad = useSelector(
        (state: RootState) => state.clubsReducer.activeClub.squad
    );

    const [targerCard, setTargetCard] = useState<Element | null>(null);
    /** MouseMove event Position */
    const [mousePosition, setMousePosition] = useState({ x: 0, y: 0 });
    /** This var created to not allow mouseUpEvent without Dragging before it */
    const [isDragging, handleDrag] = useState(false);
    /** Playing area position */
    const [playingAreaPosition, setplayingAreaPosition] = useState({
        x: 0,
        y: 0,
    });

    const DEFAULT_VALUE = 0;
    const X_SCROLL_POINT = 0;
    const Y_SCROLL_POINT = 1200;
    const DELAY = 100;

    /** Gets playing area position */
    useEffect(() => {
        const playingArea = document.getElementById('playingArea');
        playingArea &&
            setplayingAreaPosition({
                x: playingArea.offsetLeft,
                y: playingArea.offsetTop,
            });
    }, []);
    const useMousePosition = (ev: any) => {
        setMousePosition({ x: ev.pageX, y: ev.pageY });
    };

    /** Compares card id with default id */
    function isCardDefined(id: string) {
        const defaultId = '00000000-0000-0000-0000-000000000000';

        return id !== defaultId;
    }

    /** Add card position, and shows card selection */
    function handleClick(index: number, e: React.MouseEvent<HTMLDivElement>) {
        if ((e.target as Element).className.includes('empty')) {
            dispatch(choosePosition(index));
            dispatch(cardSelectionVisibility(true));
            setTimeout(() => {
                window.scroll(X_SCROLL_POINT, Y_SCROLL_POINT);
            }, DELAY);
        }
    }

    /** getting dragged card index and changing state to allow mouseUp */
    function dragStart(
        e: React.MouseEvent<HTMLDivElement>,
        index: number = DEFAULT_VALUE
    ): void {
        handleDrag(true);
        dispatch(setDragStart(index));
    }
    /** getting second drag index  and exchanging with first index*/
    function onMouseUp(
        e: React.MouseEvent<HTMLDivElement>,
        index: number = DEFAULT_VALUE
    ): void {
        e.stopPropagation();
        if (isDragging && dragStartIndex !== null) {
            const cards = club.squadCards;
            isCardDefined(cards[index].card.id)
                ? dispatch(
                    swapCards(
                        new CardEditIdentificators(
                            squad.clubId,
                            squad.id,
                            cards[dragStartIndex].card.id,
                            index
                        ),
                        new CardEditIdentificators(
                            squad.clubId,
                            squad.id,
                            cards[index].card.id,
                            dragStartIndex
                        )
                    )
                )
                : dispatch(
                    changeCardPosition(
                        new CardEditIdentificators(
                            squad.clubId,
                            squad.id,
                            cards[dragStartIndex].card.id,
                            index
                        )
                    )
                );
        }

        dispatch(setDragStart());
        handleDrag(false);
    }

    /** when we release card not on target it just brings it on start position*/
    function mouseUpOnArea(e: React.MouseEvent<HTMLInputElement>) {
        e.stopPropagation();
        dispatch(setDragStart());
    }

    /** deleting card when release beyond playing area */
    function removeFromArea() {
        if (isDragging && dragStartIndex !== null) {
            dispatch(
                deleteCard(
                    new CardEditIdentificators(
                        squad.clubId,
                        squad.id,
                        club.squadCards[dragStartIndex].card.id,
                        dragStartIndex
                    )
                )
            );
        }
        dispatch(setDragStart());
        handleDrag(false);
    }

    /** Show/hide delete block, preventing scroll to cardSelection. */
    const handleVisibility = (e: React.MouseEvent<HTMLInputElement>) => {
        e.stopPropagation();

        const target = e.target as Element;

        target && target.id
            ? targerCard && target.id === targerCard.id
                ? setTargetCard(null)
                : setTargetCard(target)
            : setTargetCard(null);
    };

    return (
        <div
            className="playing-area__wrapper"
            onMouseMove={(ev) => useMousePosition(ev)}
            onMouseUp={removeFromArea}
            style={isDragging ? { cursor: 'not-allowed' } : {}}
            onClick={handleVisibility}
        >
            <div className="playing-area" id="playingArea">
                <div
                    style={dragStartIndex ? { cursor: 'grabbing' } : {}}
                    className={`playing-area__${formation}`}
                    onMouseUp={mouseUpOnArea}
                >
                    {club.squadCards.map(
                        (fieldCard: SquadCard, index: number) => {
                            const isDefined = isCardDefined(fieldCard.card.id);
                            const isDragging = dragStartIndex === index;

                            return (
                                <div
                                    style={
                                        isDragging
                                            ? {
                                                left:
                                                      mousePosition.x -
                                                      playingAreaPosition.x,
                                                top:
                                                      mousePosition.y -
                                                      playingAreaPosition.y,
                                                transform:
                                                      'translate(-55%, -50%)',
                                                zIndex: 5,
                                                pointerEvents: 'none',
                                            }
                                            : undefined
                                    }
                                    key={index}
                                    className={`playing-area__${formation}__${
                                        isDefined ? 'card' : 'empty-card'
                                    }`}
                                    onClick={(e) => handleClick(index, e)}
                                    onDragStart={(e) => dragStart(e, index)}
                                    onMouseUp={(e) => onMouseUp(e, index)}
                                    draggable={true}
                                >
                                    {isDefined && 
                                        <FootballerCard
                                            card={fieldCard.card}
                                            index={index}
                                            place={"PlayingArea"}
                                            setTargetCard={setTargetCard}
                                            targerCard={targerCard}
                                        />
                                    }
                                </div>
                            );
                        }
                    )}
                </div>
                <div className={`playing-area__${formation}-shadows`}>
                    {club.squadCards.map(
                        (fieldCard: SquadCard, index: number) => {
                            const isDefined = isCardDefined(fieldCard.card.id);

                            return (
                                <div
                                    className={`playing-area__${formation}-shadows__card`}
                                    key={index}
                                >
                                    {isDefined && 
                                        <img
                                            src={fieldCard.card.shadow}
                                            alt="card shadow"
                                            className={`playing-area__${formation}-shadows__shadow`}
                                        />
                                    }
                                </div>
                            );
                        }
                    )}
                </div>
            </div>
            <FieldControlsArea />
        </div>
    );
};
