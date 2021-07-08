/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';

import './FilterField.scss';

import rectangle from '../../../../img/FootballFieldPage/FilterField/rectangle.png'
import search from '../../../../img/FootballFieldPage/FilterField/search.png'

export const FilterField: React.FC = () => {
    const filterFieldTitles: { title: string, src: string }[] = [
        {
            title: 'Player`s name',
            src: search
        },
        {
            title: 'Card quality',
            src: rectangle
        },
        {
            title: 'Overal rating',
            src: rectangle
        },
        {
            title: 'Player`s position',
            src: rectangle
        },

    ];

    return (
        <section className="football-field-filter">
            <div className="football-field-filter__wrapper">
                <ul className="football-field-filter__list">
                    {filterFieldTitles.map((item, index) => {
                        return (
                            <li key={index}
                                className="football-field-filter__list__item">
                                {item.title}
                                <img
                                    src={item.src}
                                    alt="Filter icon"
                                    className="football-field-filter__list__item__picture"
                                />
                            </li>
                        );
                    })}
                </ul>
            </div>
        </section>
    );
};
