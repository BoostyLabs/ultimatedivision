// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';

import rectangle
    from '@static/img/MarketPlacePage/marketPlaceFilterField/rectangle.svg';
import search
    from '@static/img/MarketPlacePage/marketPlaceFilterField/search.svg';
import star
    from '@static/img/MarketPlacePage/marketPlaceFilterField/star.svg';
import fut
    from '@static/img/MarketPlacePage/marketPlaceFilterField/fut.svg';
import eye
    from '@static/img/MarketPlacePage/marketPlaceFilterField/eye.svg';
import stars
    from '@static/img/MarketPlacePage/marketPlaceFilterField/stars.svg';
import parametres
    from '@static/img/MarketPlacePage/marketPlaceFilterField/parametres.svg';

import { useDispatch } from 'react-redux';
import { filterList } from '@/app/store/reducers/footballerCard';

import './index.scss';

export const MarketPlaceFilterField: React.FC<{ title: string }> = ({ title }) => {
    const [searchData, setSearchData] = useState('');

    const handleSerchChange = (event: any) => {
        setSearchData(event.target.value);
    };
    const dispatch = useDispatch();

    const filterFieldTitles: Array<{ title: string; src: string }> = [
        {
            title: 'Version',
            src: rectangle,
        },
        {
            title: 'Positions',
            src: rectangle,
        },
        {
            title: 'Nations',
            src: rectangle,
        },
        {
            title: 'Leagues',
            src: rectangle,
        },
        {
            title: 'WRF',
            src: rectangle,
        },
        {
            title: 'Stats',
            src: rectangle,
        },
        {
            title: '',
            src: star,
        },
        {
            title: 'PS',
            src: fut,
        },
        {
            title: 'T&S',
            src: rectangle,
        },
        {
            title: '',
            src: eye,
        },
        {
            title: '',
            src: stars,
        },
        {
            title: 'RPP',
            src: rectangle,
        },
        {
            title: '',
            src: parametres,
        },
        {
            title: 'Misc',
            src: rectangle,
        },
    ];

    return (
        <section className="marketplace-filter">
            <h1 className="marketplace-filter__title">
                {title}
            </h1>
            <div className="marketplace-filter__wrapper">
                <ul className="marketplace-filter__list">
                    <li className="marketplace-filter__list__item">
                        <img
                            src={search}
                            alt="Filter Icon"
                            className="marketplace-filter__list__item__picture"
                            onClick={() => dispatch(filterList(searchData))}
                        />
                        <input
                            value={searchData}
                            placeholder="Search"
                            className="marketplace-filter__list__item__search"
                            onChange={handleSerchChange}
                        />
                    </li>
                    {filterFieldTitles.map((item, index) =>
                        <li key={index}
                            className="marketplace-filter__list__item">
                            <span className="marketplace-filter__list__item__title" >
                                {item.title}
                            </span>
                            <img
                                src={item.src}
                                alt="Filter icon"
                                className="marketplace-filter__list__item__picture"
                            />
                        </li>,
                    )}
                </ul>
            </div>
        </section >
    );
};
