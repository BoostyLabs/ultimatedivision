// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Club, ClubState } from '@/club';

import {
    CARD_POSITION,
    SET_CLUBS,
    DRAG_START,
    SELECTION_VISIBILITY,
} from '@/app/store/actions/clubs';

/** TODO: replace by initial object */
const clubState = new ClubState();

export const clubsReducer = (state = clubState, action: any = {}) => {
    switch (action.type) {
        case SET_CLUBS:
            return {
                ...state,
                clubs: action.clubs,
                activeClub: action.clubs.find((club:Club) => club.status === 1)
            }
        case SELECTION_VISIBILITY:
            state.options.showCardSeletion = action.isVisible;
            break;
        case CARD_POSITION:
            state.options.chosedCard = action.index;
            break;
        case DRAG_START:
            state.options.dragStart = action.index;
            break;
        default:
            break;
    }

    return { ...state };
};
