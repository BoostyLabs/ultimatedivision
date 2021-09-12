// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useSelector } from 'react-redux';

import { MyCard } from '@/app/components/Club/ClubCardsArea/MyCard';
import { RootState } from '@/app/store';

import boxBody from '@static/img/StorePage/BoxContent/boxBody.svg';
import boxCover from '@static/img/StorePage/BoxContent/boxCover.svg';
import boxLight from '@static/img/StorePage/BoxContent/boxLight.svg';
import ribbons from '@static/img/StorePage/BoxContent/ribbons.svg';

import './index.scss';


export const LootboxOpening = () => {
    const FIRST_CARD = 0;
    const card = useSelector((state: RootState) => state.lootboxReducer.lootbox[FIRST_CARD]);

    return (
        <div className="box-animation">
            <div className="box-animation__box-wrapper">
                <img
                    src={boxBody}
                    alt="box body"
                    className="box-animation__box-body"
                />
                <img
                    src={boxCover}
                    alt="box cover"
                    className="box-animation__box-cover"
                />

                <img
                    src={boxLight}
                    alt="shadow image"
                    className="box-animation__light"
                />

                <img
                    src={ribbons}
                    alt="ribbons"
                    className="box-animation__ribbons"
                />
            </div>
            <div className="box-animation__card-wrapper">
                <div className="box-animation__card-wrapper-backlight">
                    <MyCard card={card} />
                </div>
            </div>
        </div>
    );
};
