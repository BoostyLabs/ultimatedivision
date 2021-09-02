// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useDispatch } from 'react-redux';

import { openLootbox } from '@/app/store/actions/lootboxes';

import { BoxCardQuality } from './BoxCardQuality';
import { BoxData } from '@/app/types/lootboxData';

import wood from '@static/img/StorePage/BoxCard/wood.svg';
import silver from '@static/img/StorePage/BoxCard/silver.svg';
import gold from '@static/img/StorePage/BoxCard/gold.svg';
import diamond from '@static/img/StorePage/BoxCard/diamond.svg';
import coin from '@static/img/MarketPlacePage/MyCard/goldPrice.svg';

import './index.scss';

export const BoxCard: React.FC<{ data: BoxData }> = ({ data }) => {
    const dispatch = useDispatch();
    const openLootbox1 = () => {
        dispatch(openLootbox({ uuid: data.id, name: data.title }))
    };

    const qualities = [
        {
            name: 'Wood',
            icon: wood,
        },
        {
            name: 'Silver',
            icon: silver,
        },
        {
            name: 'Gold',
            icon: gold,
        },
        {
            name: 'Diamond',
            icon: diamond,
        },
    ];

    return (
        <div className="box-card">
            <div className="box-card__wrapper">
                <div className="box-card__description">
                    <div
                        className="box-card__icon"
                        style={{ backgroundImage: `url(${data.icon})` }}
                    />
                    <h2 className="box-card__title">{data.title}</h2>
                    <div className="box-card__quantity">
                        <span className="box-card__quantity-label">Cards</span>
                        <span className="box-card__quantity-value">{data.quantity}</span>
                    </div>
                </div>
                <div className="box-card__qualities">
                    {data.dropChance.map((item, index) =>
                        <BoxCardQuality
                            label={qualities[index]}
                            chance={item}
                            key={index}
                        />
                    )}
                    <button
                        className="box-card__button"
                        onClick={openLootbox1}
                    >
                        <span className="box-card__button-text">OPEN</span>
                        <span className="box-card__button-value"><img src={coin} alt="" />{data.price}</span>
                    </button>
                </div>
            </div>
        </div>
    );
};
