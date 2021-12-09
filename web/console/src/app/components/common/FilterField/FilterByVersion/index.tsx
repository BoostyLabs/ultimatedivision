// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';

import { FilterByParameterWrapper } from '@/app/components/common/FilterField/FilterByParameterWrapper';

import { QueryParametersField } from '@/card';

// TODO: rework functionality.
export const FilterByVersion: React.FC<{ submitSearch: (queryParameters: QueryParametersField[]) => void }> = ({ submitSearch }) => {
    /** Indicates if FilterByVersion component shown. */
    const [isFilterByVersionShown, setIsFilterByVersionShown] = useState(false);

    /** Shows and closes FilterByVersion component. */
    const showFilterByVersion = () => {
        setIsFilterByVersionShown(isFilterByVersionShown => !isFilterByVersionShown);
    };

    /** Indicates if is choosed diamond quality of cards. */
    const [isDiamondQuality, setIsDiamondQuality] = useState<boolean>(false);
    /** Indicates if is choosed gold quality of cards. */
    const [isGoldQuality, setIsGoldQuality] = useState<boolean>(false);
    /** Indicates if is choosed silver quality of cards. */
    const [isSilverQuality, setIsSilverQuality] = useState<boolean>(false);
    /** Indicates if is choosed wood quality of cards. */
    const [isWoodQuality, setIsWoodQuality] = useState<boolean>(false);

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
    const changeQuality: () => string = () => {
        let quality: string = '';

        if (isDiamondQuality) {
            quality = 'diamond';
        };

        if (isGoldQuality) {
            quality = 'gold';
        };

        if (isSilverQuality) {
            quality = 'silver';
        };

        if (isWoodQuality) {
            quality = 'wood';
        };

        return quality;
    };

    /** Exposes default page number. */
    const DEFAULT_PAGE_INDEX: number = 1;

    /** Submits query parameters by quality. */
    const handleSubmit = async() => {
        await submitSearch([{ quality: changeQuality() }]);
        showFilterByVersion();
    };

    return (
        <FilterByParameterWrapper
            showComponent={showFilterByVersion}
            isComponentShown={isFilterByVersionShown}
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
                onClick={chooseSilverQuality}
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
                onClick={chooseDiamondQuality}
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
