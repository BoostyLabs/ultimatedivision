import React from 'react';
import './FootballerCardIllustrations.scss';
import { FootballerCardIllustrationsRadar }
    from '../FootballerCardIllustrationsRadar/FootballerCardIllustrationsRadar';
import  { FootballerCardIllustrationsDiagramsArea }
    from '../FootballerCardIllustrationsDiagramsArea/FootballerCardIllustrationsDiagramsArea';

import icon from '../../../img/FootballerCardPage/diamond2.png';

export const FootballerCardIllustrations = () => {

    return (
        <div className="footballer-card-illustrations">
            <img
                src={icon}
                alt=""
                className="footballer-card-illustrations__logo"
            />
            <FootballerCardIllustrationsRadar />
            <FootballerCardIllustrationsDiagramsArea />
        </div>
    );
};
