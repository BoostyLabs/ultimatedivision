// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useSelector } from 'react-redux';

import { RootState } from '@/app/store';

import { filteredLots, marketplaceLots } from '@/app/store/actions/cards';

import { MarketPlaceCardsGroup } from '@components/MarketPlace/MarketPlaceCardsGroup';
import { FilterField } from '@components/common/FilterField';
import { Paginator } from '@components/common/Paginator';

import './index.scss';

const MarketPlace: React.FC = () => {
    const currentPage = useSelector((state: RootState) => state.cardsReducer.marketplaceCurrentPage);
    const pagesCount = useSelector((state: RootState) => state.cardsReducer.marketplacePagesCount);
    return (
        <section className="marketplace">
            <FilterField
                title="MARKETPLACE"
                thunk={filteredLots}
            />
            <MarketPlaceCardsGroup />
            <Paginator
                getCardsOnPage={marketplaceLots}
                page={currentPage}
                pagesCount={pagesCount}
            />
        </section>
    );
};

export default MarketPlace;
