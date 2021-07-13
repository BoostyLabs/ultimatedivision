/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
*/

import React from 'react';
import { useSelector } from 'react-redux';

import { FootballerCardIllustrationsRadar }
    from '../FootballerCardIllustrationsRadar/FootballerCardIllustrationsRadar';
import { FootballerCardIllustrationsDiagramsArea }
    from '../FootballerCardIllustrationsDiagramsArea/FootballerCardIllustrationsDiagramsArea';
import { PlayerCard } from '../../PlayerCard/PlayerCard';

import { RootState } from '../../../store';

import './FootballerCardIllustrations.scss';

export const FootballerCardIllustrations: React.FC = () => {
    //TODO: Route with card ID
    const card = useSelector((state: RootState) => state.cardReducer[0]);
    return (
        <div className="footballer-card-illustrations">
            <PlayerCard
                card={card}
                parentClassName={"footballer-card-illustrations"}
            />
            <FootballerCardIllustrationsRadar />
            <FootballerCardIllustrationsDiagramsArea />
        </div>
    );
};
