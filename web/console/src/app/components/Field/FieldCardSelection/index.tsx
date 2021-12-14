// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useDispatch, useSelector } from 'react-redux';

import { Paginator } from '@components/common/Paginator';
import { PlayerCard } from '@components/common/PlayerCard';
import { FilterField } from '@/app/components/common/FilterField';
import { FilterByPrice } from '@components/common/FilterField/FilterByPrice';
import { FilterByStats } from '@components/common/FilterField/FilterByStats';
import { FilterByStatus } from '@components/common/FilterField/FilterByStatus';
import { FilterByVersion } from '@components/common/FilterField/FilterByVersion';

import { RootState } from '@/app/store';
import { listOfCards, createCardsQueryParameters } from '@/app/store/actions/cards';
import { addCard, cardSelectionVisibility } from '@/app/store/actions/clubs';
import { CardEditIdentificators } from '@/api/club';
import { Card, CardsPage, CardsQueryParametersField } from '@/card';
import { Squad, SquadCard } from '@/club';

import './index.scss';
import { useEffect } from 'react';
import { clearCardsQueryParameters } from '../../../store/actions/cards';

export const FieldCardSelection = () => {
    const dispatch = useDispatch();
    const squad: Squad = useSelector((state: RootState) => state.clubsReducer.activeClub.squad);
    const squadCards: SquadCard[] = useSelector((state: RootState) => state.clubsReducer.activeClub.squadCards);
    const isCardsVisible = useSelector((state: RootState) => state.clubsReducer.options.showCardSeletion);

    const { cards, page }: CardsPage = useSelector((state: RootState) => state.cardsReducer.cardsPage);
    const club = useSelector((state: RootState) => state.clubsReducer);

    const Y_SCROLL_POINT = 200;
    const X_SCROLL_POINT = 0;
    const DELAY = 10;

    /** Function filters card list each time when we add card on field */
    function getAvailableCards() {
        const squadCardsIds = squadCards.map(card => card.card.id);

        return cards.filter((card: Card) => !squadCardsIds.includes(card.id));
    };

    /** Add card to field, and hide card selection component */
    function addCardOnField(cardId: string) {
        dispatch(
            addCard(
                new CardEditIdentificators(squad.clubId, squad.id, cardId, club.options.chosedCard)
            ));
        dispatch(cardSelectionVisibility(false));
        setTimeout(() => {
            window.scroll(X_SCROLL_POINT, Y_SCROLL_POINT);
        }, DELAY);
    }

    const availableCards = getAvailableCards();
    /** Exposes default page number. */
    const DEFAULT_PAGE_INDEX: number = 1;

    /** Submits search by cards query parameters. */
    const submitSearch = async(cardsQueryParameters: CardsQueryParametersField[]) => {
        createCardsQueryParameters(cardsQueryParameters);
        await dispatch(listOfCards(DEFAULT_PAGE_INDEX));
    };

    useEffect(() => {
        (async() => {
            clearCardsQueryParameters();
            await dispatch(listOfCards(DEFAULT_PAGE_INDEX));
        })();
    }, [isCardsVisible]);

    return (
        <div id="cardList" className="card-selection">
            <FilterField >
                <FilterByVersion submitSearch={submitSearch} />
                <FilterByStats submitSearch={submitSearch} />
                <FilterByPrice />
                <FilterByStatus />
            </FilterField>
            <div className="card-selection__list">
                {cards &&
                    availableCards.map((card: Card, index: number) =>
                        <div
                            key={index}
                            className="card-selection__card"
                            onClick={() => addCardOnField(card.id)}
                        >
                            <PlayerCard
                                id={card.id}
                                className={'card-selection__card'}
                            />
                        </div>
                    )}
            </div>
            <Paginator
                getCardsOnPage={listOfCards}
                itemsCount={page.totalCount}
                selectedPage={page.currentPage}
            />
        </div>
    );
};
