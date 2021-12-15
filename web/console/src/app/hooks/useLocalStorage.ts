// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/** Hook for extracting a string from local storage and converting it to boolean. */
export const useLocalStorage = () => {
    /* Boolean value from localStorage, which indicates whether the user is logged in or not. */
    const isLoggined = window.localStorage.getItem('IS_LOGGINED');

    return isLoggined === 'true';
};
