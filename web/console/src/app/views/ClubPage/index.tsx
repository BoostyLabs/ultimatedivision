// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useSelector } from 'react-redux';

import { RootState } from '@/app/store';
import { filteredCards, userCards } from '@/app/store/actions/cards';

import { ClubCardsArea } from '@components/Club/ClubCardsArea';
import { FilterField } from '@components/common/FilterField';
import { Paginator } from '@components/common/Paginator';

import './index.scss';

const Club: React.FC = () => {
    const currentPage = useSelector((state: RootState) => state.cardsReducer.clubCurrentPage);
    const pagesCount = useSelector((state: RootState) => state.cardsReducer.clubPagesCount);

    return (
        <section className="club">
            <FilterField
                title="My cards"
                thunk={filteredCards}
            />
            <ClubCardsArea />
            <Paginator
                getCardsOnPage={userCards}
                pagesCount={pagesCount}
                page={currentPage}
            />
        </section>
    );
};

export default Club;
