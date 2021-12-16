// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';
import { useAppDispatch, useAppSelector } from '@/app/store';

import { FieldControlsArea } from '@components/Field/FieldControlsArea';
import { FootballerCard } from '@components/Field/FootballerCard';

import { CardEditIdentificators } from '@/api/club';
import { RootState } from '@/app/store';
import { Card } from '@/card';
import { SquadCard } from '@/club';
import {
    deleteCard,
    swapCards,
    changePosition,
} from '@/app/store/actions/clubs';
import {
    cardSelectionVisibility,
    choosePosition,
    setDragStart,
} from '@/app/store/reducers/clubs';

import './index.scss';

export const FieldPlayingArea: React.FC = () => {
    const dispatch = useAppDispatch();

    const cards = useAppSelector((state: RootState) => state.cards.cardsPage.cards);
    const formation = useAppSelector((state: RootState) => state.clubs.activeClub.squad.formation);
    const dragStartIndex = useAppSelector((state: RootState) => state.clubs.options.dragStart);
    const club = useAppSelector((state: RootState) => state.clubs.activeClub);
    const squad = useAppSelector((state: RootState) => state.clubs.activeClub.squad);

    /** MouseMove event Position */
    const [mousePosition, setMousePosition] = useState({ x: 0, y: 0 });
    /** This var created to not allow mouseUpEvent without Dragging before it */
    const [isDragging, handleDrag] = useState(false);
    /** Playing area position */
    const [playingAreaPosition, setplayingAreaPosition] = useState({ x: 0, y: 0 });

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
    function handleClick(index: number) {
        dispatch(choosePosition(index));
        dispatch(cardSelectionVisibility(true));
        setTimeout(() => {
            window.scroll(X_SCROLL_POINT, Y_SCROLL_POINT);
        }, DELAY);
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
            isCardDefined(cards[index].card.id) ?
                dispatch(swapCards({
                    currentCard: new CardEditIdentificators(squad.clubId, squad.id, cards[dragStartIndex].card.id, index),
                    existCard: new CardEditIdentificators(squad.clubId, squad.id, cards[index].card.id, dragStartIndex)
                }))
                :
                dispatch(changePosition(
                    new CardEditIdentificators(squad.clubId, squad.id, cards[dragStartIndex].card.id, index),
                ));
        }

        dispatch(setDragStart(null));
        handleDrag(false);
    }

    /** when we release card not on target it just brings it on start position*/
    function mouseUpOnArea(e: React.MouseEvent<HTMLInputElement>) {
        e.stopPropagation();
        dispatch(setDragStart(null));
    }

    /** deleting card when release beyond playing area */
    function removeFromArea() {
        if (isDragging && dragStartIndex) {
            dispatch(deleteCard(
                new CardEditIdentificators(squad.clubId, squad.id, club.squadCards[dragStartIndex].card.id, dragStartIndex))
            );
        }
        dispatch(setDragStart(null));
        handleDrag(false);
    }

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
                    {club.squadCards.map(
                        (fieldCard: SquadCard, index: number) => {
                            const isDefined = isCardDefined(fieldCard.card.id);
                            const isDragging = dragStartIndex === index;

                            return (
                                <div
                                    style={
                                        isDragging ? {
                                            left: mousePosition.x - playingAreaPosition.x,
                                            top: mousePosition.y - playingAreaPosition.y,
                                            transform: 'translate(-55%, -50%)',
                                            zIndex: 5,
                                            pointerEvents: 'none',
                                        }
                                            : undefined
                                    }
                                    key={index}
                                    className={`playing-area__${formation}__${isDefined ? 'card' : 'empty-card'
                                        }`}
                                    onClick={() => handleClick(index)}
                                    onDragStart={(e) => dragStart(e, index)}
                                    onMouseUp={(e) => onMouseUp(e, index)}
                                    draggable={true}
                                >
                                    {isDefined &&
                                        <FootballerCard
                                            card={fieldCard.card}
                                            index={index}
                                            place={'PlayingArea'}
                                        />
                                    }
                                </div>
                            );
                        })}
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
