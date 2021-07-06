// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React from 'react';
import './FootballerCardIllustrationsDiagram.scss';
import { Doughnut } from 'react-chartjs-2';
import { Diagram } from '../../../types/fotballerCard';

export const FootballerCardIllustrationsDiagram: React.FC<{props: Diagram}> = ({ props }) => {
    const { name, min, max, value } = props
    const PERCENT = (Math.round((value - min) / max * 100))

    return (
        <div className="footballer-card-illustrations-diagram">
            <Doughnut
                type={Doughnut}
                data={{
                    datasets: [{
                        data: [PERCENT, (100 - PERCENT)],
                        backgroundColor: ['#3CCF5D', '#5E5EAA'],
                        borderColor: [
                            'transparent'
                        ],
                        cutout: '80%',
                        rotation: 270,
                        circumference: 180,
                        maintainAspectRatio: true,
                    }],
                }}
                options={{
                    plugins: {
                        title: {
                            display: true,
                            text: name.toUpperCase(),
                            color: 'white',
                        }
                    }
                }}
            />
            <div className="footballer-card-illustrations-diagram__values-area">
                <span className="footballer-card-illustrations-diagram__min">{min}</span>
                <span className="footballer-card-illustrations-diagram__value">{value}</span>
                <span className="footballer-card-illustrations-diagram__max">{max}</span>
            </div>
        </div>
    )
}
