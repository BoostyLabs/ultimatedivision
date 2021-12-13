// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';

import { FilterByParameterWrapper } from '@/app/components/common/FilterField/FilterByParameterWrapper';

import { CardsQueryParametersField } from '@/card';

// TODO: rework functionality.
export const FilterByStats: React.FC<{ submitSearch: (queryParameters: CardsQueryParametersField[]) => void }> = ({ submitSearch }) => {
    /** Indicates if FilterByStats component shown. */
    const [isFilterByStatsShown, setIsFilterByStatsShown] = useState(false);

    /** Shows and closes FilterByStats component. */
    const showFilterByStats = () => {
        setIsFilterByStatsShown(isFilterByStatsShown => !isFilterByStatsShown);
    };

    const [statistics, setStatisctis] = useState<CardsQueryParametersField[]>([
        { 'defence_gte': '' },
        { 'defence_lt': '' },
        { 'goalkeeping_gte': '' },
        { 'goalkeeping_lt': '' },
        { 'offense_gte': '' },
        { 'offense_lt': '' },
        { 'physique_gte': '' },
        { 'physique_lt': '' },
        { 'tactics_gte': '' },
        { 'tactics_lt': '' },
        { 'technique_gte': '' },
        { 'technique_lt': '' }
    ]);

    /** Submits query parameters by stats. */
    const handleSubmit = async () => {
        await submitSearch([...statistics]);
        showFilterByStats();
    };

    const setField = (index: number) => (e: React.ChangeEvent<HTMLInputElement>) => {
        const newStatistics = [...statistics];
        newStatistics[index] = { [e.target.name]: e.target.value };
        setStatisctis(newStatistics)
    };

    return (
        <FilterByParameterWrapper
            showComponent={showFilterByStats}
            isComponentShown={isFilterByStatsShown}
            title="Stats"
        >
            <div className="filter-item__dropdown-active__stats__wrapper">
                {statistics.map((stat, index: number) => {
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
