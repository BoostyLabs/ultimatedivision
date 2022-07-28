// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';

import { LootboxContent } from '@/app/components/Store/LootboxContent';
import { LootboxSelection } from '@/app/components/Store/LootboxSelection';

import './index.scss';

const Store = () => {
    const [isOpening, handleOpening] = useState(false);

    const header = document.querySelector('.ultimatedivision-navbar');
    useEffect(
        () =>
            isOpening
                ? header?.classList.add('ultimatedivision-navbar__store-opening')
                : header?.classList.remove('ultimatedivision-navbar__store-opening'),
        [isOpening]
    );

    return (
        <section className="store">
            {!isOpening ?
                <div className="store__content">
                    <h1>Box</h1>
                    <LootboxSelection handleOpening={handleOpening} />
                </div>
                :
                <LootboxContent handleOpening={handleOpening} />
            }
        </section>
    );
};

export default Store;
