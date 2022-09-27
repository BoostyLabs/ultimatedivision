// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

const content = document.querySelector('.page');

/** Sets or unsets scrolling page. */
export const setScrollAble = () => {
    content?.classList.add('scroll-unset');
};

/** Unsets or unsets scrolling page. */
export const unsetScrollAble = () => {
    content?.classList.remove('scroll-unset');
};


