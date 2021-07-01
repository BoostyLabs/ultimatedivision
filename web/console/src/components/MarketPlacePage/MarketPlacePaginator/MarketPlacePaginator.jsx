import React from "react";
import { PropTypes } from 'prop-types';

import './MarketPlacePaginator.scss';

import next
    from '../../../img/MarketPlacePage/MarketPlacePaginator/next.png';
import previous
    from '../../../img/MarketPlacePage/MarketPlacePaginator/previous.png';

export const MarketPlacePaginator = ({ onPageChange, cardsOnPage, countCards }) => {
    const pages = [];
    for (let i = 1; i <= Math.ceil(countCards / cardsOnPage); i++) {
        pages.push(i);
    }
    return (
        <section className="marketplace-paginator">
            <div className="marketplace-paginator__wrapper">
                <a className={"marketplace-paginator__previous"}
                    onClick={() => onPageChange('previous page', pages)}>
                    <img className="marketplace-paginator__previous__arrow"
                        src={previous}
                        alt="Previous page" />
                    <p className="marketplace-paginator__previous__title">
                        Previous page
                    </p>
                </a>
                <ul className="marketplace-paginator__pages">
                    {pages.map((page, index) =>
                        <li
                            className="marketplace-paginator__pages__item"
                            key={index}
                            onClick={() => onPageChange('change page', pages, page)}>
                            {page}
                        </li>
                    )}
                </ul>
                <a className="marketplace-paginator__next"
                    onClick={() => onPageChange('next page', pages)}>
                    <p className="marketplace-paginator__next__title">
                        Next page
                    </p>
                    <img className="marketplace-paginator__next__arrow"
                        src={next}
                        alt="Next page" />
                </a>
            </div>
        </section>
    );
};

MarketPlacePaginator.propTypes = {
    onPageChange: PropTypes.func.isRequired,
    cardsOnPage: PropTypes.number.isRequired,
    countCards: PropTypes.number.isRequired,
};
