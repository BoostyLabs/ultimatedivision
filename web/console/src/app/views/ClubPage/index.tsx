// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useDispatch, useSelector } from 'react-redux';

import { ClubCardsArea } from '@components/Club/ClubCardsArea';
import { FilterField } from '@components/common/FilterField';
import { FilterByPrice } from '@components/common/FilterField/FilterByPrice';
import { FilterByStats } from '@components/common/FilterField/FilterByStats';
import { FilterByStatus } from '@components/common/FilterField/FilterByStatus';
import { FilterByVersion } from '@components/common/FilterField/FilterByVersion';
import { Paginator } from '@components/common/Paginator';

import { RootState } from '@/app/store';
import { listOfCards, createCardsQueryParameters } from '@/app/store/actions/cards';
import { CardsQueryParametersField } from '@/card';

import './index.scss';

const Club: React.FC = () => {
    /** Exposes default page number. */
    const DEFAULT_PAGE_INDEX: number = 1;

    const dispatch = useDispatch();

    const submitSearch = async(cardsQueryParameters: CardsQueryParametersField[]) => {
        createCardsQueryParameters(cardsQueryParameters);
        await dispatch(listOfCards(DEFAULT_PAGE_INDEX));
    };

    const { page } = useSelector((state: RootState) => state.cardsReducer.cardsPage);

    return (
        <section className="club">
            <h1 className="club__title">
                MY CARDS
            </h1>
            <FilterField >
                <FilterByVersion submitSearch={submitSearch}/>
                <FilterByStats submitSearch={submitSearch}/>
                <FilterByPrice/>
                <FilterByStatus />
            </FilterField>
            <ClubCardsArea />
            <Paginator
                getCardsOnPage={listOfCards}
                itemsCount={page.totalCount}
                selectedPage={page.currentPage}
            />
        </section>
    );
};

export default Club;
