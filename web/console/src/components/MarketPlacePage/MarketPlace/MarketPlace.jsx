/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React, { useState } from 'react';
import { useSelector } from 'react-redux';

import { MarketPlaceNavbar } from '../MarketPlaceNavbar/MarketPlaceNavbar';
import { MarketPlaceFilterField }
    from '../MarketPlaceFilterField/MarketPlaceFilterField';
import { MarketPlaceCardsGroup }
    from '../MarketPlaceCardsGroup/MarketPlaceCardsGroup';
import './MarketPlace.scss';
import { MarketPlacePaginator }
    from '../MarketPlacePaginator/MarketPlacePaginator';

export const MarketPlace = () => {
    const cards = useSelector(state => state.footballerCard);
    const [currentPage, setCurrentPage] = useState(1);

    const cardsOnPage = 21;
    /* find index of last card on page*/
    const indexOfLastCard = currentPage * cardsOnPage;
    /* find index of first card on page */
    const indexOfFirstCard = indexOfLastCard - cardsOnPage;
    /* get current list of cards on page */
    const currentCards = cards.slice(indexOfFirstCard, indexOfLastCard);
    /* change page on paginator component */
    const onPageChange = (type, pages, pageNumber = currentPage) => {
        switch (type) {
            case 'next page':
                return pageNumber < pages.length - 1 ? setCurrentPage(pageNumber + 1)
                    : null;
            case 'previous page':
                return pageNumber > 1 ? setCurrentPage(pageNumber - 1)
                    : null;
            case 'change page':
                setCurrentPage(pageNumber);
                return;
        }
    };

    return (
        <section className="marketplace">
            <MarketPlaceNavbar />
            <MarketPlaceFilterField />
            <MarketPlaceCardsGroup
                cards={currentCards} />
            <MarketPlacePaginator
                onPageChange={onPageChange}
                cardsOnPage={cardsOnPage}
                countCards={cards.length} />
        </section>
    );
};
