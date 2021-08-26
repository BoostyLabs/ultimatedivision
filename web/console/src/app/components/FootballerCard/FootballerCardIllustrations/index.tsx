// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { FootballerCardIllustrationsDiagramsArea } from '@components/FootballerCard/FootballerCardIllustrationsDiagramsArea';
import { FootballerCardIllustrationsRadar } from '@components/FootballerCard/FootballerCardIllustrationsRadar';

import { Card } from '@/card';
import { PlayerCard } from '@components/common/PlayerCard';

import './index.scss';

export const FootballerCardIllustrations: React.FC<{ card: Card }> = ({ card }) =>
    <div className="footballer-card-illustrations">
        <div className="footballer-card-illustrations__card">
            <PlayerCard card={card} parentClassName="footballer-card-illustrations__card" />
        </div>
        <FootballerCardIllustrationsRadar card={card} />
        <FootballerCardIllustrationsDiagramsArea card={card} />
    </div>;
