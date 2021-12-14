// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useMemo, useState } from 'react';

import { FilterByParameterWrapper } from '@/app/components/common/FilterField/FilterByParameterWrapper';

import { CardsQueryParametersField } from '@/card';

// TODO: rework functionality.
export const FilterByVersion: React.FC<{ 
    submitSearch: (queryParameters: CardsQueryParametersField[]) => void }> = ({ submitSearch }) => {
    /** Indicates if FilterByVersion component shown. */
    const [isFilterByVersionShown, setIsFilterByVersionShown] = useState(false);

    /** Describes quality version. */
    const [quality, setQuality] = useState<CardsQueryParametersField[]>([{ 'quality': '' }]);

    /** Shows and closes FilterByVersion component. */
    const showFilterByVersion = () => {
        setIsFilterByVersionShown(isFilterByVersionShown => !isFilterByVersionShown);
    };

    /** An array possibilities describes list of possible qualities version. */
    const possibleQalities: string[] = ['wood', 'silver', 'gold', 'diamond'];

    /** Changes quality version. */
    const changeQuality = (index: number) => {
        let updatedQuality = [{ 'quality': possibleQalities[index] }];
        setQuality(updatedQuality)
    };

    /** Submits query parameters by quality. */
    const handleSubmit = async () => {
        await submitSearch(quality);
        showFilterByVersion();
    };

    useEffect(() => {
        console.log(quality);
    }, [])

    return (
        <FilterByParameterWrapper
            showComponent={showFilterByVersion}
            isComponentShown={isFilterByVersionShown}
            title="Version"
        >
            {possibleQalities.map((quality, index: number) =>
                <>
                    <input
                        id={`checkbox-${quality}`}
                        className="filter-item__dropdown-active__checkbox"
                        type="checkbox"
                        onClick={() => changeQuality(index)}
                    />
                    <label
                        className="filter-item__dropdown-active__text"
                        htmlFor={`checkbox-${quality}`}
                    >
                        {quality}
                    </label>
                </>
            )}
            <input
                value="APPLY"
                type="submit"
                className="filter-item__dropdown-active__apply"
                onClick={handleSubmit}
            />
        </FilterByParameterWrapper>
    );
};