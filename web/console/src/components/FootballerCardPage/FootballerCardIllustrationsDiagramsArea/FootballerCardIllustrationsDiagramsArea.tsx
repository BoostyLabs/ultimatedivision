// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React from 'react'
import { useSelector } from 'react-redux';
import { RootState } from '../../../store';

import './FootballerCardIllustrationsDiagramsArea.scss';

import { FootballerCardIllustrationsDiagram }
    from '../FootballerCardIllustrationsDiagram/FootballerCardIllustrationsDiagram'

export const FootballerCardIllustrationsDiagramsArea: React.FC = () => {
    const DIAGRAM_DATA = useSelector((state: RootState) => state.cardReducer[0].diagram)

    return (
        <div className="footballer-card-illustrations-diagram-area">
            {DIAGRAM_DATA.map(item => (
                <FootballerCardIllustrationsDiagram
                    key={item.id}
                    props={item}
                />
            ))}
        </div>
    )
}
