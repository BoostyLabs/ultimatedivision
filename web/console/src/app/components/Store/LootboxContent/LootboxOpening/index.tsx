// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useSelector } from 'react-redux';

import { PlayerCard } from '@components/common/PlayerCard';

import boxBg from '@static/img/StorePage/BoxContent/boxLight.png';

import { RootState } from '@/app/store';
import { boxStyle } from '@/app/internal/lootboxStyle';
import { Card } from '@/card';

import './index.scss';

export const LootboxOpening = () => {
    const FIRST_CARD = 0;
    const REGULAR_CARDS_AMOUNT: number = 5;
    const cards: Card[] = useSelector((state: RootState) => state.lootboxReducer.lootbox);

    const box = boxStyle(cards.length);

    return (
        <div className="box-animation">
            <div className="box-animation__box-container">
                <img
                    src={boxBg}
                    alt="box body"
                    className={`box-animation__box-body ${
                        cards.length > REGULAR_CARDS_AMOUNT && 'box-animation__box-body__cool'
                    }`}
                />
                <img src={box.cover} alt="box cover" className="box-animation__box-cover" />
            </div>
            <div className="box-animation__card-container">
                <div className="box-animation__card-container-backlight">
                    <PlayerCard className={'box-animation__card'} id={cards[FIRST_CARD].id} />
                </div>
            </div>
        </div>
    );
};
