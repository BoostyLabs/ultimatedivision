// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Radar } from 'react-chartjs-2';

import { Card } from '@/card';

import './index.scss';

export const FootballerCardIllustrationsRadar: React.FC<{ card: Card }> = ({ card }) => {
    const stats = card.statsArea;
    const labels = ['TAC', 'PHY', 'TEC', 'OFF', 'DEF', 'GK'];

    return (
        <div className="footballer-card-illustrations-radar">
            <Radar
                type={Radar}
                data={{
                    labels: stats.map((item, index) => `${item.average} ${labels[index]}`),
                    datasets: [
                        {
                            backgroundColor: 'rgba(245, 255, 99, 0.34)',
                            data: stats.map((item) => item.average),
                        },
                    ],
                }}
                options={{
                    elements: {
                        line: {
                            borderWidth: 1,
                            borderColor: '#F5FF63',
                        },
                    },
                    animations: {
                        tension: {
                            duration: 1000,
                            easing: 'linear',
                            from: 0,
                            to: 0,
                        },
                    },
                    plugins: {
                        legend: {
                            display: false,
                        },
                        interaction: {
                            display: false,
                        },
                    },
                    scale: {
                        ticks: {
                            maxTicksLimit: 2,
                        },
                    },
                    scales: {
                        r: {
                            ticks: {
                                display: false,
                            },
                            pointLabels: {
                                color: '#F5FF63',
                            },
                            angleLines: {
                                /* eslint-disable */
                                borderDash: [7, 8],
                                color: "#515180",
                            },
                            grid: {
                                color: "#515180",
                            },
                        },
                    },
                }}
            />
        </div>
    );
};
