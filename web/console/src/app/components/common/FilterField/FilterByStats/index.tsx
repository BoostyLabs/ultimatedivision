// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState, useEffect, useContext } from 'react';

import { FilterByParameterWrapper } from '@/app/components/common/FilterField/FilterByParameterWrapper';

import { CardsQueryParameters, CardsQueryParametersField } from '@/card';
import { useCardsQueryParameters } from '@/app/hooks/useCardsQueryParameters';

import { FilterContext } from '../index';

export const FilterByStats: React.FC<{
    submitSearch: (queryParameters: CardsQueryParametersField[]) => Promise<void>;
    clearsStatisticsField: (queryParameters: CardsQueryParametersField[]) => Promise<void>;
    cardsQueryParameters: CardsQueryParameters;
}> = ({ submitSearch, clearsStatisticsField, cardsQueryParameters }) => {
    const { activeFilterIndex, setActiveFilterIndex }: {
        activeFilterIndex: number;
        setActiveFilterIndex: React.Dispatch<React.SetStateAction<number>>;
    } = useContext(FilterContext);
    /** Exposes default index which does not exist in array. */
    const DEFAULT_FILTER_ITEM_INDEX = -1;
    const FILTER_BY_STATS_INDEX = 2;
    /** Indicates if FilterByStats component shown. */
    const [isFilterByStatsShown, setIsFilterByStatsShown] = useState(false);

    const isVisible = FILTER_BY_STATS_INDEX === activeFilterIndex && isFilterByStatsShown;

    /** Shows and closes FilterByStats component. */
    const showFilterByStats = () => {
        setActiveFilterIndex(FILTER_BY_STATS_INDEX);
        setIsFilterByStatsShown(isFilterByStatsShown => !isFilterByStatsShown);
    };

    /** Describes all statistics required fields. */
    const cardsQueryParametersFields = useCardsQueryParameters([
        'tactics_gte',
        'tactics_lt',
        'offense_gte',
        'offense_lt',
        'technique_gte',
        'technique_lt',
        'physique_gte',
        'physique_lt',
        'defence_gte',
        'defence_lt',
        'goalkeeping_gte',
        'goalkeeping_lt',
    ], cardsQueryParameters);

    /** Desrcribes stats values. */
    const [stats, setStats] = useState(cardsQueryParametersFields);

    /** Submits query parameters by stats. */
    const handleSubmit = async () => {
        await submitSearch([...stats]);
        setIsFilterByStatsShown(false);
        setActiveFilterIndex(DEFAULT_FILTER_ITEM_INDEX);
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
        FILTER_BY_STATS_INDEX !== activeFilterIndex && setIsFilterByStatsShown(false);
        setStats(cardsQueryParametersFields);
    }, [activeFilterIndex, cardsQueryParametersFields]);

    return (
        <FilterByParameterWrapper
            showComponent={showFilterByStats}
            isVisible={isVisible}
            title="Stats"
        >
            <div className="filter-item__dropdown-active__stats__wrapper">
                <div className="filter-item__dropdown-active__stats">
                    <span className="filter-item__dropdown-active__stats__label">
                        TAC
                    </span>
                    {stats.slice(0, 2).map((stat: CardsQueryParametersField, index: number) => {
                        for (let property in stat) {
                            return < input
                                className="filter-item__dropdown-active__stats__area"
                                name={property}
                                type='text'
                                value={stat[property]}
                                onChange={changeCurrentStatsField(index)}
                                placeholder={`${index === 0 ? 'Min' : 'Max'} Rating`
                                }
                            />
                        }
                    })}
                </div>
                <div className="filter-item__dropdown-active__stats">
                    <span className="filter-item__dropdown-active__stats__label">
                        OFF
                    </span>
                    {stats.slice(2, 4).map((stat: CardsQueryParametersField, index: number) => {
                        for (let property in stat) {
                            return < input
                                className="filter-item__dropdown-active__stats__area"
                                name={property}
                                type='text'
                                value={stat[property]}
                                onChange={changeCurrentStatsField(index + 2)}
                                placeholder={`${index === 2 ? 'Min' : 'Max'} Rating`
                                }
                            />
                        }
                    })}
                </div>
                <div className="filter-item__dropdown-active__stats">
                    <span className="filter-item__dropdown-active__stats__label">
                        TEC
                    </span>
                    {stats.slice(4, 6).map((stat: CardsQueryParametersField, index: number) => {
                        for (let property in stat) {
                            return < input
                                className="filter-item__dropdown-active__stats__area"
                                name={property}
                                type='text'
                                value={stat[property]}
                                onChange={changeCurrentStatsField(index + 4)}
                                placeholder={`${index === 4 ? 'Min' : 'Max'} Rating`
                                }
                            />
                        }
                    })}
                </div>
                <div className="filter-item__dropdown-active__stats">
                    <span className="filter-item__dropdown-active__stats__label">
                        PHY
                    </span>
                    {stats.slice(6, 8).map((stat: CardsQueryParametersField, index: number) => {
                        for (let property in stat) {
                            return < input
                                className="filter-item__dropdown-active__stats__area"
                                name={property}
                                type='text'
                                value={stat[property]}
                                onChange={changeCurrentStatsField(index + 6)}
                                placeholder={`${index === 6 ? 'Min' : 'Max'} Rating`
                                }
                            />
                        }
                    })}
                </div>
                <div className="filter-item__dropdown-active__stats">
                    <span className="filter-item__dropdown-active__stats__label">
                        DEF
                    </span>
                    {stats.slice(8, 10).map((stat: CardsQueryParametersField, index: number) => {
                        for (let property in stat) {
                            return < input
                                className="filter-item__dropdown-active__stats__area"
                                name={property}
                                type='text'
                                value={stat[property]}
                                onChange={changeCurrentStatsField(index + 8)}
                                placeholder={`${index === 8 ? 'Min' : 'Max'} Rating`
                                }
                            />
                        }
                    })}
                </div>
                <div className="filter-item__dropdown-active__stats">
                    <span className="filter-item__dropdown-active__stats__label">
                        GK
                    </span>
                    {stats.slice(10, 12).map((stat: CardsQueryParametersField, index: number) => {
                        for (let property in stat) {
                            return < input
                                className="filter-item__dropdown-active__stats__area"
                                name={property}
                                type='text'
                                value={stat[property]}
                                onChange={changeCurrentStatsField(index + 10)}
                                placeholder={`${index === 10 ? 'Min' : 'Max'} Rating`
                                }
                            />
                        }
                    })}
                </div>
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
