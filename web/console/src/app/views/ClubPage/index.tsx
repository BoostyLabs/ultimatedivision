// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { filteredCards } from '@/app/store/actions/cards';

import { useClub } from '@/app/hooks/club';

import { ClubCardsArea } from '@components/Club/ClubCardsArea';
import { FilterField } from '@components/common/FilterField';
import { Paginator } from '@components/common/Paginator';

import './index.scss';

const Club: React.FC = () => {
    /** TODO: decide use custom hook or directly dispatch thunk into useEffect*/
    const cards = useClub();

    return (
        <section className="club">
            <FilterField
                title="My cards"
                thunk={filteredCards}
            />
            <ClubCardsArea
                cards={cards}
            />
            <Paginator
                itemCount={cards.length} />
        </section>
    );
};

export default Club;
