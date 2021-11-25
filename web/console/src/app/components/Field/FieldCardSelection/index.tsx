// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useDispatch, useSelector } from 'react-redux';

import { Paginator } from '@components/common/Paginator';
import { PlayerCard } from '@components/common/PlayerCard';
import { FilterField } from '@/app/components/Field/FieldCardSelection/FilterField';

import { RootState } from '@/app/store';
import { listOfCards } from '@/app/store/actions/cards';
import { CardWithStats } from '@/card';
import { addCard, cardSelectionVisibility } from '@/app/store/actions/clubs';
import { CardEditIdentificators } from '@/api/club';
import { Squad, SquadCard } from '@/club';

import './index.scss';

export const FieldCardSelection = () => {
    const dispatch = useDispatch();
    const squad: Squad = useSelector((state: RootState) => state.clubsReducer.activeClub.squad);
    const squadCards: SquadCard[] = useSelector((state: RootState) => state.clubsReducer.activeClub.squadCards);
    const { cards, page } = useSelector((state: RootState) => state.cardsReducer.cardsPage);
    const club = useSelector((state: RootState) => state.clubsReducer);

    const Y_SCROLL_POINT = 200;
    const X_SCROLL_POINT = 0;
    const DELAY = 10;

    /** Function filters card list each time when we add card on filed */
    function availableCards() {
        let userCards = [...cards];
        for (const squadCard of squadCards) {
            userCards = userCards.filter((card: CardWithStats) => card.id !== squadCard.cardId);
        };

        return userCards;
    }

    /** Add card to field, and hide card selection component */
    function setCard(cardId: string) {
        dispatch(
            addCard(
                new CardEditIdentificators(squad.clubId, squad.id, cardId, club.options.chosedCard)
            ));
        dispatch(cardSelectionVisibility(false));
        setTimeout(() => {
            window.scroll(X_SCROLL_POINT, Y_SCROLL_POINT);
        }, DELAY);
    }

    return (
        <div id="cardList" className="card-selection">
            <FilterField />
            <div className="card-selection__list">
                {cards &&
                    availableCards().map((card: CardWithStats, index: number) =>
                        <div
                            key={index}
                            className="card-selection__card"
                            onClick={() => setCard(card.id)}
                        >
                            <PlayerCard
                                card={card}
                                parentClassName={'card-selection__card'}
                            />
                        </div>
                    )}
            </div>
            <Paginator
                getCardsOnPage={listOfCards}
                pagesCount={page.pageCount}
                selectedPage={page.currentPage}
            />
        </div>
    );
};
