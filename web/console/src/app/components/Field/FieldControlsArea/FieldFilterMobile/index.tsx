// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from "react";
import { useDispatch, useSelector } from "react-redux";

import { FieldDropdown } from "../FieldDropdown"

import { DropdownStyle } from "@/app/internal/dropdownStyle"
import { Control, MobileControl } from "@/app/types/club"
import { changeActiveClub, setFormation } from "@/app/store/actions/clubs";
import { RootState } from "@/app/store";

import { Formations } from "@/club";

import arrowLeftIcon from '@static/img/FilterField/arrow-left.svg';
import arrowIcon from '@static/img/FieldPage/arrow.svg';
import arrowActiveIcon from '@static/img/FieldPage/arrow-active.svg';

import './index.scss'

export const FieldFilterMobile: React.FC<{
    isMobileFilterActive: boolean,
    returnToFilter: () => void,
    currentOption: null | Control | MobileControl,
    checkActiveElement: (item: Control | MobileControl) => boolean,
    isDropdownActive: boolean,
    setCurrentControlsAreaOption: (item: Control | MobileControl) => void,
    setActiveComposition: React.Dispatch<React.SetStateAction<string>>
    activeComposition: string
}> = ({
    isMobileFilterActive,
    returnToFilter,
    currentOption,
    checkActiveElement,
    isDropdownActive,
    setCurrentControlsAreaOption,
    setActiveComposition,
    activeComposition
}) => {
        const dispatch = useDispatch()

        const formation = useSelector((state: RootState) => state.clubsReducer.activeClub.squad.formation);
        const clubs = useSelector((state: RootState) => state.clubsReducer.clubs);
        const activeClubId = useSelector((state: RootState) => state.clubsReducer.activeClub.id);
        const squad = useSelector((state: RootState) => state.clubsReducer.activeClub.squad);

        const getCurrentFormationName = (currentFormation: string) => {
            if (typeof currentFormation == "string" && currentFormation.indexOf('-') > -1) {

                return currentFormation
            }
            /* @ts-ignore */
            return Formations[currentFormation]
        }

        const [activeClub, setActiveClub] = useState(activeClubId)
        const [mobileComposition, setMobileComposition] = useState(activeComposition)
        const [activeFormation, setActiveFormation] = useState(getCurrentFormationName(formation))

        const saveChanges = () => {
            dispatch(changeActiveClub(activeClub))
            dispatch(setFormation(squad, activeFormation))
            setActiveComposition(mobileComposition)
            returnToFilter()
        }

        const getClubName = (item: any) => {
            const currentClub = clubs.filter((club: any) => club.id === item.currentValue);
            const clubName = `club ${currentClub?.name ? currentClub.name : ''}`

            return clubName;
        }


        const isDropdownFieldActive = (item: MobileControl) => {
            return  isDropdownActive && currentOption?.title === item.title
        }

        const CONTROLS_FIELDS = [
            new MobileControl(
                '1',
                'club',
                setActiveClub,
                clubs,
                activeClub,
                'id',
                'name',
                'club'
            ),
            new MobileControl(
                '2',
                'squad',
                setMobileComposition,
                ['Composition 1', 'Composition 2', 'Composition 3', 'Composition 4', 'Composition 5'],
                mobileComposition
            ),
            new MobileControl(
                '3',
                'formation',
                setActiveFormation,
                ['4-4-2', '4-2-4', '4-2-2-2', '4-3-1-2', '4-3-3', '4-2-3-1', '4-3-2-1', '4-1-3-2', '5-3-2', '4-5-1'],
                activeFormation
            ),
        ];

        return (
            <div className="field-filter">
                <div className="field-filter__content">
                    <div>
                            <div className='field-filter__top-side'>
                                <span onClick={() => returnToFilter()}
                                    className='field-filter__top-side__arrow-left'>
                                    <img src={arrowLeftIcon} alt='arrow-left' />
                                </span>
                                <h2 className='field-filter__top-side__title'>
                                    Filter
                                </h2>
                            </div>
                        
                        {CONTROLS_FIELDS.map((item, _) => {
                            const currentName = item.title === 'club' ? getClubName(item) : item.currentValue

                            return (
                                <div key={item.title}>
                                    <div
                                        className={`field-filter__settings__item ${isDropdownFieldActive(item) ? 'field-filter__settings__item--active' : ''}`}
                                        onClick={() => setCurrentControlsAreaOption(item)}>
                                        <div className="field-filter__settings__item__heading" >
                                            <h4 className="field-filter__settings__item__title">
                                                {item?.currentValue ? currentName : item.title}
                                            </h4>

                                            <img
                                                className="field-filter__settings__item__image"
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
                                    {isMobileFilterActive &&isDropdownFieldActive(item) && <FieldDropdown option={currentOption} isMobile={true} />}
                                </div>
                            )
                        })}
                    </div>
                     <button className='field-filter__save-button' onClick={() => saveChanges()}>Save Changes</button>
                </div>
            </div>
        )
    }