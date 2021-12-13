// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';

import { FilterByParameterWrapper } from '@/app/components/common/FilterField/FilterByParameterWrapper';

import { CardsQueryParametersField } from '@/card';

// TODO: rework functionality.
export const FilterByVersion: React.FC<{ submitSearch: (queryParameters: CardsQueryParametersField[]) => void }> = ({ submitSearch }) => {
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

    const [quality, setQuality] = useState<CardsQueryParametersField[]>([{ 'quality': '' }]);

    const possibleQalities: string[] = ['wood, silver, gold, diamond'];

    const changeQuality = (index: number) => {
        let updatedQuality = [...quality, { 'quality': possibleQalities[index] }];
        setQuality(updatedQuality)
    };

    /** Submits query parameters by quality. */
    const handleSubmit = async () => {
        await submitSearch(quality);
        showFilterByVersion();
    };

    return (
        <FilterByParameterWrapper
            showComponent={showFilterByVersion}
            isComponentShown={isFilterByVersionShown}
            title="Version"
        >
            {possibleQalities.map((quality, index: number) => {
                return <>
                    <input
                        id={`division-checkbox-${quality}`}
                        className="filter-item__dropdown-active__checkbox"
                        type="checkbox"
                        onClick={() => changeQuality(index)}
                    />
                    <label
                        className="filter-item__dropdown-active__text"
                        htmlFor={`division-checkbox-${quality}`}
                    >
                        {quality}
                    </label>
                </>
            })}
            <input
                value="APPLY"
                type="submit"
                className="filter-item__dropdown-active__apply"
                onClick={handleSubmit}
            />
        </FilterByParameterWrapper>
    );
};
