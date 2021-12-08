// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/** FilterFieldStatsAreaProps is an interface that uses for incoming props FilterFieldStatsArea component. */
export interface FilterFieldStatsAreaProps {
    label: string;
    minValue: string;
    maxValue: string;
    changeMinValue: (e: React.ChangeEvent<HTMLInputElement>) => void;
    changeMaxValue: (e: React.ChangeEvent<HTMLInputElement>) => void;
};

export const FilterFieldStatsArea: React.FC<FilterFieldStatsAreaProps> = ({
    label,
    minValue,
    maxValue,
    changeMinValue,
    changeMaxValue,
}) => <div className="filter-item__dropdown-active__stats">
    <span className="filter-item__dropdown-active__stats__label">
        {label}
    </span>
    <input
        value={minValue}
        placeholder="Min Rating"
        className="filter-item__dropdown-active__stats__area"
        type="text"
        onChange={changeMinValue}
    />
    <input
        value={maxValue}
        placeholder="Max Rating"
        className="filter-item__dropdown-active__stats__area"
        type="text"
        onChange={changeMaxValue}
    />
</div>;