// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useDispatch, useSelector } from 'react-redux';

import { MarketPlaceCardsGroup } from '@components/MarketPlacePage/MarketPlaceCardsGroup';
import { FilterField } from '@components/common/FilterField';
import { Paginator } from '@components/common/Paginator';

import { RootState } from '@/app/store';
import { createCardList } from '@/app/store/reducers/marketplace';

import './index.scss';

const MarketPlace = () => {
    const dispatch = useDispatch();
    dispatch(createCardList());
    const cards = useSelector((state: RootState) => state.marketplaceReducer);

    return (
        <section className="marketplace">
            <FilterField
                title="MARKETPLACE"
            />
            <MarketPlaceCardsGroup
                cards={cards}
            />
            <Paginator
                itemCount={cards.length} />
        </section>
    );
};

export default MarketPlace;
