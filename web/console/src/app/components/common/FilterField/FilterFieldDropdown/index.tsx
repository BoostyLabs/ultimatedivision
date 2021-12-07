// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { CardsClient } from '@/api/cards';
import { RootState } from '@/app/store';
import { CardService } from '@/card/service';
import { listOfCards, createCardsQueryParameters } from '@/app/store/actions/cards';


export const FilterFieldDropdown1: React.FC<{
    props: { label: string; image: string };
}> = ({ props }) => {
    const { label, image } = props;

    const [isDropdownShow, setIsDropdownShow] = useState(false);
    const dispatch = useDispatch();

    /** describes quality of each card. At list consist of wood, silver, gold, diamond. */
    const [quality, setQuality] = useState<string>('');

    /** describes technique skills of each card. */
    const [tacticsMin, setTacticsMin] = useState('');

    const [tacticsMax, setTacticsMax] = useState('');


    /** describes technic skills of each card. */
    const [technique, setTechnique] = useState<{ min: string | number; max: string | number }>({
        min: '',
        max: '',
    });
    const cardsClient = new CardsClient();
    const cardsService = new CardService(cardsClient);

    const DEFAULT_PAGE: number = 1;

    const handleSubmit = async() => {
        createCardsQueryParameters([{ quality }]);
        await dispatch(listOfCards(DEFAULT_PAGE));
    };

    return (
        <div
            className="filter-item"
            onClick={() => setIsDropdownShow(true)}
        >
            <span className="filter-item__title">{label}</span>
            <img
                className="filter-item__picture"
                src={image}
                alt={image && 'filter icon'}
            />
            <div
                className={`filter-item__dropdown${isDropdownShow ? '-active' : '-inactive'
                }`}
            >
                {/* <input
                    value={technique.min}
                    placeholder='min'
                    onChange={(e) => {
                        setTechnique({ ...technique, min: e.target.value });
                    }}
                />
                <input
                    value={technique.max}
                    placeholder='max'
                    onChange={(e) => {
                        setTechnique({ ...technique, max: e.target.value });
                    }}
                /> */}
                <input
                    value={quality}
                    placeholder="min"
                    onChange={(e) => {
                        setQuality(e.target.value);
                    }}
                />
                <input
                    value="apply"
                    onClick={handleSubmit}
                />
            </div>
        </div>
    );
};
