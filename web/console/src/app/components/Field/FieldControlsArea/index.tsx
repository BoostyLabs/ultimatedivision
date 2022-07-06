// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import Clubs from '@/app/components/Field/FieldControlsArea/Clubs';

import { setCaptain, setFormation, setTactic, startSearchingMatch } from '@/app/store/actions/clubs';
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
    const [isPossibleToStartMatch, setIsPossibleToStartMatch] = useState<boolean>(true);
    const squadCards = useSelector((state: RootState) => state.clubsReducer.activeClub.squadCards);
    const EMPTY_CARD_ID = '00000000-0000-0000-0000-000000000000';

    const [currentOption, setCurrentOption] = useState<null | Control>(null);
    const [optionVisibility, changeVisibility] = useState(false);
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
            '0',
            'formation',
            setFormation,
            ['4-4-2', '4-2-4', '4-2-2-2', '4-3-1-2', '4-3-3', '4-2-3-1', '4-3-2-1', '4-1-3-2', '5-3-2', '4-5-1'],
            amountColumnsElements['five-elements']
        ),
        new Control(
            '1',
            'tactics',
            setTactic,
            ['attack', 'defence', 'balanced'],
            amountColumnsElements['four-elements']
        ),
        new Control(
            '2',
            'captain',
            setCaptain,
            ['Captain 1', 'Captain 2', 'Captain 3'],
            amountColumnsElements['four-elements']
        ),
    ];

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
                {/* <Clubs /> */}

                <div className="field-controls__settings">
                    {CONTROLS_FIELDS.map((item, index) =>
                        <div className="field-controls__settings__item" key={item.title}>
                            <div
                                className="field-controls__settings__item__heading"
                                onClick={() => {
                                    setCurrentOption(item);

                                    if (item.title !== currentOption?.title && optionVisibility) {
                                        changeVisibility(false);
                                        changeVisibility((prev) => !prev);
                                    } else {
                                        changeVisibility((prev) => !prev);
                                    }
                                }}
                            >
                                <h4 className="field-controls__settings__item__title">{item.title}</h4>
                                {item.title === currentOption?.title && optionVisibility ?
                                    <img
                                        className="field-controls__settings__item__image"
                                        src={arrowActiveIcon}
                                        alt="triangle img"
                                        id={`triangle-${item.id}`}
                                        style={{ transform: new DropdownStyle(true).triangleRotate }}
                                    />
                                    :
                                    <img
                                        className="field-controls__settings__item__image"
                                        src={arrowIcon}
                                        alt="triangle img"
                                        id={`triangle-${item.id}`}
                                    />
                                }
                            </div>
                        </div>
                    )}
                </div>
                {currentOption !== null && optionVisibility && <FieldDropdown option={currentOption} />}
            </div>
            <input
                type="button"
                value="Play"
                className="field-controls__play"
                onClick={showMatchFinder}
                disabled={!isPossibleToStartMatch}
            />
        </div>
    );
};
