// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useMemo, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { setFormation, startSearchingMatch } from '@/app/store/actions/clubs';
import { RootState } from '@/app/store';
import { amountColumnsElements, Control } from '@/app/types/club';
import { SquadCard } from '@/club';
import { DropdownStyle } from '@/app/internal/dropdownStyle';
import { FieldDropdown } from './FieldDropdown';

import arrowIcon from '@static/img/FieldPage/arrow.svg';
import arrowActiveIcon from '@static/img/FieldPage/arrow-active.svg';

import './index.scss';

export const FieldControlsArea: React.FC = () => {
    const dispatch = useDispatch();

    const [currentOption, setCurrentOption] = useState<null | Control>(null);
    const [optionVisibility, changeVisibility] = useState(false);

    const [activeClub, setActiveClub] = useState<string>('CLUB 1');
    const [activeComposition, setActiveComposition] = useState<string>('Composition 1');

    const [isPossibleToStartMatch, setIsPossibleToStartMatch] = useState<boolean>(true);
    const squadCards = useSelector((state: RootState) => state.clubsReducer.activeClub.squadCards);
    const formation = useSelector((state: RootState) => state.clubsReducer.activeClub.squad.formation);
    const EMPTY_CARD_ID = '00000000-0000-0000-0000-000000000000';

    const isDropdownActive = useMemo(
        () => currentOption !== null && optionVisibility,
        [currentOption, optionVisibility]
    );

    const checkActiveElement = (item: Control) => item.title === currentOption?.title && optionVisibility;

    useEffect(() => {
        /** Function checks field cards and compare it with player cards array */
        function isPossibleToStart() {
            const emptyCard = squadCards.find((squadCard: SquadCard) => squadCard.card.id === EMPTY_CARD_ID);
            emptyCard ? setIsPossibleToStartMatch(false) : setIsPossibleToStartMatch(true);
        }
        isPossibleToStart();
    });

    const CONTROLS_FIELDS = [
        new Control(
            '1',
            'club',
            setActiveClub,
            ['CLUB 1', 'CLUB 2', 'CLUB 3', 'CLUB 4', 'CLUB 5', 'CLUB 6'],
            amountColumnsElements['four-elements'],
            activeClub
        ),
        new Control(
            '2',
            'squad',
            setActiveComposition,
            ['Composition 1', 'Composition 2', 'Composition 3', 'Composition 4'],
            amountColumnsElements['four-elements'],
            activeComposition
        ),
        new Control(
            '3',
            'formation',
            setFormation,
            ['4-4-2', '4-2-4', '4-2-2-2', '4-3-1-2', '4-3-3', '4-2-3-1', '4-3-2-1', '4-1-3-2', '5-3-2', '4-5-1'],
            amountColumnsElements['five-elements'],
            formation
        ),
    ];

    const setCurrentControlsAreaOption = (item: Control) => {
        setCurrentOption(item);

        if (item.title !== currentOption?.title && optionVisibility) {
            changeVisibility(false);
        }
        changeVisibility((prev) => !prev);
    };

    /** shows matchFinder component */
    const showMatchFinder = () => {
        dispatch(startSearchingMatch(true));
        window.scrollTo({
            top: 0,
            behavior: 'smooth',
        });
    };

    return (
        <div className="field-controls">
            <div className="field-controls__wrapper">
                <div className="field-controls__settings">
                    {CONTROLS_FIELDS.map((item, index) =>
                        <div className="field-controls__settings__item" key={item.title}>
                            <div
                                className="field-controls__settings__item__heading"
                                onClick={() => setCurrentControlsAreaOption(item)}
                            >
                                <h4 className="field-controls__settings__item__title">{item.title}</h4>

                                <img
                                    className="field-controls__settings__item__image"
                                    src={checkActiveElement(item) ? arrowActiveIcon : arrowIcon}
                                    alt="triangle img"
                                    id={`triangle-${item.id}`}
                                    style={
                                        checkActiveElement(item)
                                            ? { transform: new DropdownStyle(true).triangleRotate }
                                            : {}
                                    }
                                />
                            </div>
                        </div>
                    )}
                </div>
                {isDropdownActive && <FieldDropdown option={currentOption} />}
            </div>
            <input
                type="button"
                value="Play"
                className="field-controls__play"
                onClick={showMatchFinder}
                disabled={isPossibleToStartMatch}
            />
        </div>
    );
};
