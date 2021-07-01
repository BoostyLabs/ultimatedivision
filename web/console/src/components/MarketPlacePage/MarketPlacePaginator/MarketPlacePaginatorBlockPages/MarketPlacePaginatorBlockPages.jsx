import React from 'react';

export const MarketPlacePaginatorBlockPages = ({ blockPages, onPageChange }) => {
    return (
        <ul className="marketplace-paginator__pages">
            {blockPages.map((page, index) =>
                <li
                    className="marketplace-paginator__pages__item"
                    key={index}
                    onClick={() => onPageChange('change page', page)}>
                    {page}
                </li>
            )}
        </ul>
    )
};
