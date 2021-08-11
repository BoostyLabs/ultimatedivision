// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Card } from '@/app/types/fotballerCard';

/** create list of player cards (implementation for test)*/
function cardList(count: number): Card[] {
    const list: Card[] = [];
    while (count > 0) {
        list.push(new Card(0), new Card(1), new Card(2), new Card(3));
        count--;
    }

    return list;
}

const SEARCH_FILTER = 'SearchFilter';


interface FilterAction {
    type: string;
    action: string;
}

export const filterList: (filterValue: string) => FilterAction = (filterValue: string) => ({
    type: SEARCH_FILTER,
    action: filterValue,
});

// FilterField functionality base implementation

export const cardReducer = (cardState = cardList(20), action: any = {}) => {
    let marketPlaceCardList = [...cardState];
    switch (action.type) {
    case SEARCH_FILTER:
        marketPlaceCardList = cardState.filter(card =>
            card.overalInfo[0].value.includes(action.action));
        break;
    default:
        break;
    }

    return marketPlaceCardList;
};
