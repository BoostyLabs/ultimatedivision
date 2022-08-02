// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';
import { useSelector } from 'react-redux';

import { PlayerCard } from '@components/common/PlayerCard';
import { RootState } from '@/app/store';
import { Card } from '@/card';

import boxOpening from '@static/img/StorePage/BoxContent/boxOpening.gif';
import box from '@static/img/StorePage/BoxContent/box.png';

import './index.scss';

export const LootboxOpening = () => {
    const [isOpenedLootBox, handleOpenedLootbox] = useState(false);

    const FIRST_CARD = 0;
    const ANIMATION_LOOTBOX_OPENING_DELAY = 10000;

    const cards: Card[] = useSelector((state: RootState) => state.lootboxReducer.lootbox);

    useEffect(() => {
        setTimeout(() => handleOpenedLootbox(true), ANIMATION_LOOTBOX_OPENING_DELAY);
    });

    return (
        <div className="box-animation">
            <div
                className={`box-animation__box-container ${
                    isOpenedLootBox ? 'box-animation__box-container--opened' : ''
                }`}
            >
                {isOpenedLootBox ?
                    <img src={box} alt="box body" className={'box-animation__box--opened'} />
                    :
                    <img src={boxOpening} alt="box body" className={'box-animation__box'} />
                }
            </div>
            <div className="box-animation__card__container">
                <div className="box-animation__card__container__backlight">
                    <PlayerCard className={'box-animation__card'} id={cards[FIRST_CARD].id} />
                </div>
            </div>
        </div>
    );
};
