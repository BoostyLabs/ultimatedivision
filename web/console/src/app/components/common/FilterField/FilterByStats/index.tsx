// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';

import { FilterByParameterWrapper } from '@/app/components/common/FilterField/FilterByParameterWrapper';

import { CardsQueryParametersField } from '@/card';

// TODO: rework functionality.
export const FilterByStats: React.FC<{
    submitSearch: (queryParameters: CardsQueryParametersField[]) => void;
    statistics?: CardsQueryParametersField[],
    setStatistics?: any,
}> = ({ submitSearch, setStatistics, statistics }) => {
    /** Indicates if FilterByStats component shown. */
    const [isFilterByStatsShown, setIsFilterByStatsShown] = useState(false);

    /** Shows and closes FilterByStats component. */
    const showFilterByStats = () => {
        setIsFilterByStatsShown(isFilterByStatsShown => !isFilterByStatsShown);
    };

    /** Submits query parameters by stats. */
    const handleSubmit = async () => {
        console.log('submit: ', statistics);
        statistics && await submitSearch([...statistics]);
        showFilterByStats();
    };

    const setField = (index: number) => (e: React.ChangeEvent<HTMLInputElement>) => {
        const newStatistics = statistics && [...statistics];
        if (newStatistics) {
            newStatistics[index] = { [e.target.name]: e.target.value };
        }
        console.log('new:', newStatistics)
        setStatistics(newStatistics)
        console.log('stat: ', statistics);
    };

    return (
        <FilterByParameterWrapper
            showComponent={showFilterByStats}
            isComponentShown={isFilterByStatsShown}
            title="Stats"
        >
            <div className="filter-item__dropdown-active__stats__wrapper">
                {statistics && statistics.map((stat, index: number) => {
                    for (let i in stat) {
                        return <div>
                            <input
                                name={i}
                                value={stat[i]}
                                onChange={setField(index)}
                                placeholder={`${index % 2 === 0 ? 'Min' : 'Max'}`}
                            />
                        </div>
                    }
                })}
                <div className="filter-item__dropdown-active__stats">
                    <input
                        value="APPLY"
                        type="submit"
                        className="filter-item__dropdown-active__stats__apply"
                        onClick={handleSubmit}
                    />
                </div>
                <div className="filter-item__dropdown-active__stats">
                    <input
                        type="submit"
                        className="filter-item__dropdown-active__stats__clear"
                        value="CLEAR ALL"
                    />
                </div>
            </div>
        </FilterByParameterWrapper>
    );
};
