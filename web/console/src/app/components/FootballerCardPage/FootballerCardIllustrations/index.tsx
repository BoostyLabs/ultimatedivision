/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
*/

import { FootballerCardIllustrationsDiagramsArea } from '@FootballerCard/FootballerCardIllustrationsDiagramsArea';
import { FootballerCardIllustrationsRadar } from '@FootballerCard/FootballerCardIllustrationsRadar';

import { Card } from '@Store/reducers/footballerCard';
import { PlayerCard } from '@PlayerCard';

import './index.scss';

export const FootballerCardIllustrations: React.FC<{ card: Card }> = ({ card }) =>
    <div className="footballer-card-illustrations">
        <div className="footballer-card-illustrations__card">
            <PlayerCard card={card} parentClassName="footballer-card-illustrations__card" />
        </div>
        <FootballerCardIllustrationsRadar card={card} />
        <FootballerCardIllustrationsDiagramsArea card={card} />
    </div>;
