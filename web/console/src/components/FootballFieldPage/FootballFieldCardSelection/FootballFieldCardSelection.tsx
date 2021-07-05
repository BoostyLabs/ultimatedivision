import React from 'react';
import './FootballFieldCardSelection.scss';

import { MarketPlaceCardsGroup }
    from '../../MarketPlacePage/MarketPlaceCardsGroup/MarketPlaceCardsGroup';
import { useSelector } from 'react-redux';
import { RootState } from '../../../store';

export const FootballFieldCardSelection = () => {
    const Cards = useSelector((state: RootState) => state.footballerCard)
    return (
        <div>
            <MarketPlaceCardsGroup cards={Cards} place={'FootballField'} />
        </div>
    )
}
