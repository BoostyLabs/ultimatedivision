// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';

import { FilterByParameterWrapper } from '@/app/components/common/FilterField/FilterByParameterWrapper';

import { CardsQueryParametersField, CardsQueryParameters } from '@/card';

import { useCardsQueryParameters } from '@/app/hooks/useCardsQueryParameters';

// TODO: rework functionality.
export const FilterByStats: React.FC<{
    submitSearch: (queryParameters: CardsQueryParametersField[]) => Promise<void>;
    clearsStatisticsField: (queryParameters: CardsQueryParametersField[]) => Promise<void>,
    cardsQueryParameters: CardsQueryParameters,
}> = ({ submitSearch, clearsStatisticsField, cardsQueryParameters }) => {

    /** Describes all statistics required fields. */
    const cardsQueryParametersFields = useCardsQueryParameters([
        'defence_gte',
        'defence_lt',
        'goalkeeping_gte',
        'goalkeeping_lt',
        'offense_gte',
        'offense_lt',
        'physique_gte',
        'physique_lt',
        'tactics_gte',
        'tactics_lt',
        'technique_gte',
        'technique_lt'
    ], cardsQueryParameters);

    /** Desrcribes stats values. */
    const [stats, setStats] = useState(cardsQueryParametersFields);

    /** Indicates if FilterByStats component shown. */
    const [isFilterByStatsShown, setIsFilterByStatsShown] = useState(false);

    /** Shows and closes FilterByStats component. */
    const showFilterByStats = () => {
        setIsFilterByStatsShown(isFilterByStatsShown => !isFilterByStatsShown);
    };

    /** Submits query parameters by stats. */
    const handleSubmit = async () => {
        await submitSearch([...stats]);
        showFilterByStats();
    };

    /** Clears all statistics values. */
    const clearAllValues = async () => {

        await clearsStatisticsField([...stats]);

        const clearedStats = [...stats];

        clearedStats.map((stat: CardsQueryParametersField, index: number) => {
            for (let property in stat) {
                clearedStats[index] = { [property]: '' }
            }
        });

        setStats(clearedStats);
    };

    /** Changes current stats field. */
    const changeCurrentStatsField = (index: number) => (e: React.ChangeEvent<HTMLInputElement>) => {
        const currentStats = [...stats];

        currentStats[index] = { [e.target.name]: e.target.value };

        setStats(currentStats);
    };

    useEffect(() => {
        setStats(cardsQueryParametersFields);
    }, [cardsQueryParametersFields]);

    return (
        <FilterByParameterWrapper
            showComponent={showFilterByStats}
            isComponentShown={isFilterByStatsShown}
            title="Stats"
        >
            <div className="filter-item__dropdown-active__stats__wrapper">
                {stats.map((stat: CardsQueryParametersField, index: number) => {
                    for (let property in stat) {
                        return <div>
                            <input
                                name={property}
                                value={stat[property]}
                                onChange={changeCurrentStatsField(index)}
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
                        onClick={clearAllValues}
                        type="submit"
                        className="filter-item__dropdown-active__stats__clear"
                        value="CLEAR ALL"
                    />
                </div>
            </div>
        </FilterByParameterWrapper>
    );
};
