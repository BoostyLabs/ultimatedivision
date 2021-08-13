// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.
import { useDispatch, useSelector } from 'react-redux';
import { RootState } from '@/app/store';
import { addCard, cardSelectionVisibility } from '@/app/store/reducers/footballField';

import { FilterField } from
    '@components/FootballFieldPage/FootballFieldCardSelection/FilterField';
import { PlayerCard } from '@/app/components/common/PlayerCard';

import './index.scss';
import { Paginator } from '@/app/components/common/Paginator';
import { Card } from '@/Card';

export const FootballFieldCardSelection = () => {
    const cardList = useSelector((state: RootState) => state.cardReducer);
    const dispatch = useDispatch();
    const fieldSetup = useSelector((state: RootState) => state.fieldReducer);

    const Y_SCROLL_POINT = 200;
    const X_SCROLL_POINT = 0;
    const DELAY = 10;

    /** Add card to field, and hide card selection component */
    function handleClick(card: Card, index: number) {
        dispatch(addCard(card, index));
        dispatch(cardSelectionVisibility(false));
        setTimeout(() => {
            window.scroll(X_SCROLL_POINT, Y_SCROLL_POINT);
        }, DELAY);
    }

    return (
        <div id="cardList" className="card-selection">
            <FilterField />
            <div className="card-selection__list">
                {cardList.map((card, index) =>
                    <div
                        key={index}
                        className="card-selection__card"
                        onClick={() => handleClick(card, fieldSetup.options.chosedCard)}
                    >
                        <PlayerCard
                            card={card}
                            parentClassName={'card-selection__card'}
                        />
                    </div>,
                )}
            </div>
            <Paginator itemCount={cardList.length} />
        </div>
    );
};
