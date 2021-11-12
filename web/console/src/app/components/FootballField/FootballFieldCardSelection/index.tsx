// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useDispatch, useSelector } from 'react-redux';

import { Paginator } from '@components/common/Paginator';
import { PlayerCard } from '@components/common/PlayerCard';
import { FilterField } from '@components/FootballField/FootballFieldCardSelection/FilterField';

import { RootState } from '@/app/store';
import { listOfCards } from '@/app/store/actions/cards';
import { addCard, cardSelectionVisibility } from '@/app/store/actions/club';
import { Card } from '@/card';
import { Squad } from '@/club';

import './index.scss';

export const FootballFieldCardSelection = () => {
    const dispatch = useDispatch();
    const squad = useSelector((state: RootState) => state.clubReducer.squad);
    const { cards, page } = useSelector(
        (state: RootState) => state.cardsReducer.cardsPage
    );
    const fieldSetup = useSelector((state: RootState) => state.clubReducer);

    const Y_SCROLL_POINT = 200;
    const X_SCROLL_POINT = 0;
    const DELAY = 10;

    /** Add card to field, and hide card selection component */
    function handleClick(squad: Squad, cardId: string, position: number) {
        dispatch(addCard({ squad, cardId, position }));
        dispatch(cardSelectionVisibility(false));
        setTimeout(() => {
            window.scroll(X_SCROLL_POINT, Y_SCROLL_POINT);
        }, DELAY);
    }

    return (
        <div id="cardList" className="card-selection">
            <FilterField />
            <div className="card-selection__list">
                {cards.map((card: Card, index: number) =>
                    <div
                        key={index}
                        className="card-selection__card"
                        onClick={() =>
                            handleClick(
                                squad,
                                card.id,
                                fieldSetup.options.chosedCard,
                            )
                        }
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
