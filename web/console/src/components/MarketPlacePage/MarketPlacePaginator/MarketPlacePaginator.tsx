import React, { useState, useEffect } from "react";

import { MarketPlacePaginatorBlockPages }
    from './MarketPlacePaginatorBlockPages/MarketPlacePaginatorBlockPages';

import './MarketPlacePaginator.scss';

import next
    from '../../../img/MarketPlacePage/MarketPlacePaginator/next.png';
import previous
    from '../../../img/MarketPlacePage/MarketPlacePaginator/previous.png';

export const MarketPlacePaginator: React.FC<{ countCards: number }> = ({ countCards }) => {
    let [currentPage, setCurrentPage] = useState<number>(1);
    /**
    * split the page into 3 blocks that can be needed
    * to separate page numbers
     */
    let [firstBlockPages, setFirstBlockPages] = useState<number[]>([]);
    let [middleBlockPages, setMiddleBlockPages] = useState<number[]>([]);
    let [lastBlockPages, setLastBlockPages] = useState<number[]>([]);

    useEffect(() => {
        populatePages();
    }, [currentPage]);

    const CARDS_ON_PAGE: number = 8;
    const MAX_PAGES_PER_BLOCK: number = 5;
    const MAX_PAGES_OFF_BLOCKS: number = 10;
    let pages: number[] = [];
    for (let i = 1; i <= Math.ceil(countCards / CARDS_ON_PAGE); i++) {
        pages.push(i);
    }
    /**
     * set block pages depends on current page
     */
    let setBlocksIfCurrentInFirstBlock = () => {
        setFirstBlockPages(pages.slice(0, 5));
        setMiddleBlockPages([]);
        setLastBlockPages(pages.slice(-1));
    };
    let setBlocksIfCurrentInMiddleBlock = () => {
        setFirstBlockPages(pages.slice(0, 1));
        setMiddleBlockPages(pages.slice(currentPage - 3, currentPage + 2));
        setLastBlockPages(pages.slice(-1));
    };
    let setBlocksIfCurrentInLastBlock = () => {
        setFirstBlockPages(pages.slice(0, 1));
        setMiddleBlockPages([]);
        setLastBlockPages(pages.slice(-5));
    };
    /**
    * Indicates visibility of dots after first pages block
     */
    let isFirstDotsShown: boolean =
        middleBlockPages.length <= MAX_PAGES_PER_BLOCK
        && pages.length > MAX_PAGES_OFF_BLOCKS;
    /*
    * Indicates visibility of dots after middle pages block
     */
    let isSecondDotsShown: boolean = !!middleBlockPages.length;
    /**
     * indicates in which block current page
     */
    let isCurrentInFirstBlock: boolean = currentPage < MAX_PAGES_PER_BLOCK;
    let isCurrentInLastBlock: boolean = pages.length - currentPage < MAX_PAGES_PER_BLOCK - 1;
    /**
     * change page blocks reorganization depends
     * on current page
     */
    let reorganizePagesBlock = () => {
        if (isOneBlockRequired) {
            return;
        }
        if (isCurrentInFirstBlock) {
            setBlocksIfCurrentInFirstBlock();
            return;
        }
        if (!isCurrentInFirstBlock && !isCurrentInLastBlock) {
            setBlocksIfCurrentInMiddleBlock();
            return;
        }
        if (isCurrentInLastBlock) {
            setBlocksIfCurrentInLastBlock();
        }
    };
    /*
    * indicates if dots delimiter is needed
    * to separate page numbers
    */
    let isOneBlockRequired: boolean = pages.length <= MAX_PAGES_OFF_BLOCKS;
    let populatePages = () => {
        if (!pages.length) return;
        if (isOneBlockRequired) {
            setFirstBlockPages(pages.slice());
            setMiddleBlockPages([]);
            setLastBlockPages([]);
            return;
        }
        reorganizePagesBlock();
    };
    /**
     * change current page and set pages block
     */
    const onPageChange = (type: string, pageNumber: number = currentPage): void => {
        switch (type) {
            case 'next page':
                if (pageNumber < pages.length)
                    setCurrentPage(pageNumber + 1);
                populatePages();
                return;
            case 'previous page':
                if (pageNumber > 1)
                    setCurrentPage(pageNumber - 1);
                populatePages();
                return;
            case 'change page':
                setCurrentPage(pageNumber);
                populatePages();
                return;
        }
    };

    return (
        <section className="marketplace-paginator">
            <div className="marketplace-paginator__wrapper">
                <a className="marketplace-paginator__previous"
                    onClick={() => onPageChange('previous page')}>
                    <img className="marketplace-paginator__previous__arrow"
                        src={previous}
                        alt="Previous page" />
                    <p className="marketplace-paginator__previous__title">
                        Previous page
                    </p>
                </a>
                <MarketPlacePaginatorBlockPages blockPages={firstBlockPages}
                    onPageChange={onPageChange} />
                {isFirstDotsShown
                    && <span className="marketplace-paginator__pages__dots">
                        ...</span>}
                <MarketPlacePaginatorBlockPages blockPages={middleBlockPages}
                    onPageChange={onPageChange} />
                {isSecondDotsShown
                    && <span className="marketplace-paginator__pages__dots">
                        ...</span>}
                <MarketPlacePaginatorBlockPages blockPages={lastBlockPages}
                    onPageChange={onPageChange} />
                <a className="marketplace-paginator__next"
                    onClick={() => onPageChange('next page')}>
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
