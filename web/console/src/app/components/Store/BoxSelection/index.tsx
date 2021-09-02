// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { BoxCard } from '../BoxCard';
import { BoxData } from '@/app/types/lootboxData';

import box from '@static/img/StorePage/BoxCard/box.svg';
import coolBox from '@static/img/StorePage/BoxCard/coolBox.svg';

import './index.scss';

export const BoxSelection = () => {
    const REGULAR_BOX_CARDS_QUANTITY = 5;
    const COOL_BOX_CARDS_QUANTITY = 10;
    /** TODO: remove test code */
    const boxesData = [
        new BoxData(
            '1',
            box,
            'Regular Box',
            REGULAR_BOX_CARDS_QUANTITY,
            // eslint-disable-next-line
            [80, 15, 4, 1],
            '200,000',
        ),
        new BoxData(
            '2',
            coolBox,
            'Cool Box',
            COOL_BOX_CARDS_QUANTITY,
            // eslint-disable-next-line
            [70, 20, 8, 2],
            '500,000',
        ),
    ];

    return (
        <div className="box-selection">
            {boxesData.map((item, index) =>
                <BoxCard data={item} key={index} />
            )}
        </div>
    );
};
