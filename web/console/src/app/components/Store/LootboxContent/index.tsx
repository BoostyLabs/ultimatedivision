// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch, SetStateAction, useEffect, useState } from 'react';

import { LootboxKeeping } from './LootboxKeeping';
import { LootboxOpening } from './LootboxOpening';

import './index.scss';

export const LootboxContent: React.FC<{
    handleOpening: Dispatch<SetStateAction<boolean>>;
}> = ({ handleOpening }) => {
    const [isAnimated, handleAnimation] = useState(true);

    useEffect(() => {
        const ANIMATION_DELAY = 13000;
        setTimeout(() => handleAnimation(false), ANIMATION_DELAY);
    });

    return (
        <div className="box-content">
            {isAnimated ? <LootboxOpening /> : <LootboxKeeping handleOpening={handleOpening} />}
        </div>
    );
};
