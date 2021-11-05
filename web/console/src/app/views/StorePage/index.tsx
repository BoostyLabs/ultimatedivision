// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';

import { LootboxContent } from '@/app/components/Store/LootboxContent';
import { LootboxSelection } from '@/app/components/Store/LootboxSelection';

import './index.scss';

const Store = () => {
    const [isOpening, handleOpening] = useState(false);

    return (
        <section className="store">
            {!isOpening ?
                <>
                    <h1 className="store__title">Box</h1>
                    <LootboxSelection handleOpening={handleOpening} />
                </>
                :
                <LootboxContent handleOpening={handleOpening} />
            }
        </section>
    );
};

export default Store;
