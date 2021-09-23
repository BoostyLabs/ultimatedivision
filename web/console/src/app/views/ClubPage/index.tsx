// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { userCards } from '@/app/store/actions/cards';

import { filteredCards } from '@/app/store/actions/cards';

import { ClubCardsArea } from '@components/Club/ClubCardsArea';
import { FilterField } from '@components/common/FilterField';
import { Paginator } from '@components/common/Paginator';

import './index.scss';

const Club: React.FC = () => {

    return (
        <section className="club">
            <FilterField
                title="My cards"
                thunk={filteredCards}
            />
            <ClubCardsArea />
            <Paginator getCardsOnPage={userCards} />
        </section>
    );
};

export default Club;
