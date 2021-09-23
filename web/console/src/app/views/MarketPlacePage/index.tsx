// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { filteredLots, marketplaceLots } from '@/app/store/actions/cards';

import { MarketPlaceCardsGroup } from '@components/MarketPlace/MarketPlaceCardsGroup';
import { FilterField } from '@components/common/FilterField';
import { Paginator } from '@components/common/Paginator';

import './index.scss';

const MarketPlace: React.FC = () => {
    return (
        <section className="marketplace">
            <FilterField
                title="MARKETPLACE"
                thunk={filteredLots}
            />
            <MarketPlaceCardsGroup />
            <Paginator getCardsOnPage={marketplaceLots} />
        </section>
    );
};

export default MarketPlace;
