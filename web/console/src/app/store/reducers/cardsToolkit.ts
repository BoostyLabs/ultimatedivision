
import { createSlice, PayloadAction } from '@reduxjs/toolkit';

import { AppThunk, AppDispatch } from '@/app/store/indexToolkit';
import { Card, CardsPage, ICard } from '@/card';
import { CardClient } from '@/api/cards';
import { CardService } from '@/card/service';
import { Pagination } from '@/app/types/pagination';

const client = new CardClient();
const service = new CardService(client);

const DEFAULT_OFFSET_VALUE: number = 0;
const DEFAULT_LIMIT_VALUE: number = 24;
const FIRST_PAGE: number = 1;
const PAGES_COUNT: number = 1;
const CARDS_TOTAL_COUNT: number = 1;

class CardsState {
    /** class implementation */
    constructor(
        public cardsPage: CardsPage,
        public card: Card,
    ) { };
};

const page = {
    offset: DEFAULT_OFFSET_VALUE,
    limit: DEFAULT_LIMIT_VALUE,
    currentPage: FIRST_PAGE,
    pageCount: PAGES_COUNT,
    totalCount: CARDS_TOTAL_COUNT,
};

const cardsPage = new CardsPage([], page);
const openedCard = new Card();
const initialState: CardsState = new CardsState(cardsPage, openedCard);

export const cardsSlice = createSlice({
    name: 'cards',
    initialState,
    reducers: {
        getCards(state, action: PayloadAction<CardsPage>) {
            state.cardsPage = action.payload;
        },
        userCards(state, action: PayloadAction<Card>) {
            state.card = action.payload;
        }
    }
});

export const listOfCards = (
    { selectedPage, limit }: Pagination
): AppThunk => async (dispatch: AppDispatch) => {
    const response = await service.list({ selectedPage, limit });

    const cards = response.cards.
        map((card: Partial<ICard>) => new Card(card));
    const page = response.page;

    dispatch(cardsSlice.actions.getCards({ cards, page }))
}

export const openUserCard = (
    id: string
): AppThunk => async (dispatch: AppDispatch) => {
    const card = await service.getCardById(id);

    dispatch(cardsSlice.actions.userCards(new Card(card)))
}

export const filteredCards = (
    lowRange: string, topRange: string
): AppThunk => async (dispatch: AppDispatch) => {
    const filterParam = `${lowRange}&${topRange}`;
    const response = await service.filteredList(filterParam);
    ;
    const cards = response.cards.
        map((card: Partial<ICard>) => new Card(card));
    const page = response.page;

    dispatch(cardsSlice.actions.getCards({ cards, page }))
}

export const { getCards, userCards } = cardsSlice.actions;

export default cardsSlice.reducer;  