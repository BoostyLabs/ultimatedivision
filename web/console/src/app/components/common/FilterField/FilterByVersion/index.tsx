// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState, useEffect, useContext } from 'react';

import { FilterByParameterWrapper } from '@/app/components/common/FilterField/FilterByParameterWrapper';

import { CardsQueryParametersField, CardsQueryParameters } from '@/card';
import { useCardsQueryParameters } from '@/app/hooks/useCardsQueryParameters';
import { FilterContext } from '../index';

// TODO: rework functionality.
export const FilterByVersion: React.FC<{
    submitSearch: (queryParameters: CardsQueryParametersField[]) => void;
    cardsQueryParameters: CardsQueryParameters;
}> = ({ submitSearch, cardsQueryParameters }) => {
    const { activeFilterIndex, setActiveFilterIndex }: {
        activeFilterIndex: number;
        setActiveFilterIndex: React.Dispatch<React.SetStateAction<number>>;
    } = useContext(FilterContext);
    /** Exposes default index which does not exist in array. */
    const DEFAULT_FILTER_ITEM_INDEX = -1;
    const FILTER_BY_VERSION_INDEX = 1;
    /** Indicates if FilterByVersion component shown. */
    const [isFilterByVersionShown, setIsFilterByVersionShown] = useState(false);

    const isVisible = FILTER_BY_VERSION_INDEX === activeFilterIndex && isFilterByVersionShown;

    /** Shows and closes FilterByVersion component. */
    const showFilterByVersion = () => {
        setActiveFilterIndex(FILTER_BY_VERSION_INDEX);
        setIsFilterByVersionShown(isFilterByVersionShown => !isFilterByVersionShown);
    };

    /** Describes all statistics required fields. */
    const cardsQueryParametersFields = useCardsQueryParameters([
        'quality',
    ], cardsQueryParameters);

    const [quality, setQuality] = useState<CardsQueryParametersField[]>(cardsQueryParametersFields);

    /** Indicates if is choosed diamond quality of cards. */
    const [isDiamondQuality, setIsDiamondQuality] = useState<boolean>(!!(quality[0] && quality[0].quality?.includes('diamond')));
    /** Indicates if is choosed gold quality of cards. */
    const [isGoldQuality, setIsGoldQuality] = useState<boolean>(false);
    /** Indicates if is choosed silver quality of cards. */
    const [isSilverQuality, setIsSilverQuality] = useState<boolean>(false);
    /** Indicates if is choosed wood quality of cards. */
    const [isWoodQuality, setIsWoodQuality] = useState<boolean>(false);

    console.log(!!(quality[0] && quality[0].quality?.includes('diamond')))
    /** Chooses diamond quality of cards. */
    const chooseDiamondQuality = () => {
        setIsDiamondQuality(isDiamondQuality => !isDiamondQuality);
    };

    /** Chooses gold quality of cards. */
    const chooseGoldQuality = () => {
        setIsGoldQuality(isGoldQuality => !isGoldQuality);
    };

    /** Chooses silver quality of cards. */
    const chooseSilverQuality = () => {
        setIsSilverQuality(isSilverQuality => !isSilverQuality);
    };

    /** Chooses wood quality of cards. */
    const chooseWoodQuality = () => {
        setIsWoodQuality(isWoodQuality => !isWoodQuality);
    };
    
    /** Changes quality of cards. */
    const changeQuality = () => {
        const qualities = [];

        if (isDiamondQuality) {
            qualities.push('diamond');
        };

        if (isGoldQuality) {
            qualities.push('gold');
        };

        if (isSilverQuality) {
            qualities.push('silver');
        };

        if (isWoodQuality) {
            qualities.push('wood');
        };

        const updatedQuality = [...quality];

        updatedQuality[0].quality = qualities;

        setQuality(updatedQuality);
    };

    /** Submits query parameters by quality. */
    const handleSubmit = async () => {
        changeQuality();
        await submitSearch(quality);
        setIsFilterByVersionShown(false);
        setActiveFilterIndex(DEFAULT_FILTER_ITEM_INDEX);
    };

    useEffect(() => {
        FILTER_BY_VERSION_INDEX !== activeFilterIndex && setIsFilterByVersionShown(false);
        setQuality(cardsQueryParametersFields);
    }, [activeFilterIndex, cardsQueryParametersFields, quality]);

    return (
        <FilterByParameterWrapper
            showComponent={showFilterByVersion}
            isVisible={isVisible}
            title="Version"
        >
            <input
                id="division-checkbox-wood"
                className="filter-item__dropdown-active__checkbox"
                type="checkbox"
                onClick={chooseWoodQuality}
            />
            <label
                className="filter-item__dropdown-active__text"
                htmlFor="division-checkbox-wood"
            >
                wood
            </label>
            <input
                id="checkbox-silver"
                className="filter-item__dropdown-active__checkbox"
                type="checkbox"
                onChange={chooseSilverQuality}
            />
            <label
                className="filter-item__dropdown-active__text"
                htmlFor="checkbox-silver"
            >
                silver
            </label>
            <input
                id="checkbox-gold"
                className="filter-item__dropdown-active__checkbox"
                type="checkbox"
                onClick={chooseGoldQuality}
            />
            <label
                className="filter-item__dropdown-active__text"
                htmlFor="checkbox-gold"
            >
                gold
            </label>
            <input
                id="checkbox-diamond"
                className="filter-item__dropdown-active__checkbox"
                type="checkbox"
                checked={isDiamondQuality}
                onChange={chooseDiamondQuality}
            />
            <label
                className="filter-item__dropdown-active__text"
                htmlFor="checkbox-diamond"
            >
                diamond
            </label>
            <input
                value="APPLY"
                type="submit"
                className="filter-item__dropdown-active__apply"
                onClick={handleSubmit}
            />
        </FilterByParameterWrapper>
    );
};