/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
*/

import React from 'react';
import './FootballerCardIllustrations.scss';
import { FootballerCardIllustrationsRadar } from '../FootballerCardIllustrationsRadar';
import { FootballerCardIllustrationsDiagramsArea } from '../FootballerCardIllustrationsDiagramsArea';

import icon from '../../../img/FootballerCardPage/diamond2.png';

export const FootballerCardIllustrations: React.FC = () =>
    <div className="footballer-card-illustrations">
        <img
            src={icon}
            alt="fotballer illustration"
            className="footballer-card-illustrations__logo"
        />
        <FootballerCardIllustrationsRadar />
        <FootballerCardIllustrationsDiagramsArea />
    </div>
;
