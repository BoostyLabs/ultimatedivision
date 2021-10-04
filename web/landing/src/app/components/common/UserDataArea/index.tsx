// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React, { SetStateAction, useEffect, useState } from 'react';

import { useDebounce } from '@/app/hooks/useDebounce';

import check from '@static/images/registerPage/check.svg';

export const UserDataArea: React.FC<{
    value: string,
    placeHolder: string,
    onChange: any,
    className: string,
    type: string,
    error: SetStateAction<string | null>,
    clearError: any,
    validate: (value: string) => boolean,
}> = ({
    value,
    placeHolder,
    onChange,
    className,
    type,
    error,
    clearError,
    validate,
}) => {
        const DELAY: number = 500;
        /**
        * The value string from input returned by the useDebounce method after 500 milliseconds.
        */
        const debouncedValue: string = useDebounce(value, DELAY);

        /** inline styles for valid input field */
        const [trueCheckStyle, setTrueCheckStyle] =
            useState({});

        useEffect(() => {
            if (!validate(debouncedValue)) {
                setTrueCheckStyle({ backgroundImage: '' });
            } else {
                setTrueCheckStyle({ backgroundImage: `url(${check})` });
            };

        }, [debouncedValue]);

        const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
            onChange(e.target.value);
            clearError(null);
        };

        return <>
            <input
                className={error ? `${className}-error` : className}
                value={value}
                placeholder={placeHolder}
                onChange={handleChange}
                style={trueCheckStyle}
                type={type}
            />
            {error && <label className={`${className}__error`} htmlFor={value}>
                {error}
            </label>}
        </>;
    };
