// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState, useEffect, useContext } from 'react';

import { FilterByParameterWrapper } from '@/app/components/common/FilterField/FilterByParameterWrapper';

import { CardsQueryParametersField } from '@/card';
import { FilterContext } from '../index';

// TODO: rework functionality.
export const FilterByVersion: React.FC<{
    submitSearch: (queryParameters: CardsQueryParametersField[]) => void;
}> = ({ submitSearch }) => {
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

    /** An array possibilities describes list of possible qualities version. */
    const possibleQalities: string[] = ['wood', 'silver', 'gold', 'diamond'];

    /** Changes quality of cards. */
    const changeQuality: () => string[] = () => {
        const qualities: string[] = [];
        return qualities;
    };

    /** Submits query parameters by quality. */
    const handleSubmit = async() => {
        await submitSearch([{ quality: changeQuality() }]);
        setIsFilterByVersionShown(false);
        setActiveFilterIndex(DEFAULT_FILTER_ITEM_INDEX);
    };

    useEffect(() => {
        FILTER_BY_VERSION_INDEX !== activeFilterIndex && setIsFilterByVersionShown(false);
    }, [activeFilterIndex]);

    return (
        <FilterByParameterWrapper
            showComponent={showFilterByVersion}
            isVisible={isVisible}
            title="Version"
        >
            {possibleQalities.map((quality, index: number) =>
                <>
                    <input
                        id={`checkbox-${quality}`}
                        className="filter-item__dropdown-active__checkbox"
                        type="checkbox"
                        onClick={changeQuality}
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