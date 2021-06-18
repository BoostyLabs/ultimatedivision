import React from 'react'
import { useSelector } from 'react-redux';

import './FootballerCardIllustrationsRadar.scss';

import { Radar } from 'react-chartjs-2';

export const FootballerCardIllustrationsRadar = () => {

    const stats = useSelector(state => state.footballerCard[0]);

    const {
        tactics,
        physique,
        technique,
        offence,
        defence,
        goalkeeping
    } = stats;

    return (
        <div className="footballer-card-illustrations-radar">
                <Radar
                    data={{
                        labels: ['TAC', 'PHY', 'GK', 'DEF', 'OFF', 'TEC'],
                        datasets: [{
                            backgroundColor: '#66FF8866',
                            data: [
                                tactics.tactics,
                                physique.physique,
                                goalkeeping.goalkeeping,
                                defence.defence,
                                offence.offence,
                                technique.technique,
                            ],
                        }],
                    }}
                    options={{
                        animations: {
                            tension: {
                                duration: 1000,
                                easing: 'linear',
                                from: 0,
                                to: 0,
                            }
                        },
                        plugins: {
                            legend: {
                                display: false
                            },
                            interaction: {
                                display: false
                            }
                        },
                        scale: {
                            ticks: {
                                maxTicksLimit: 2,
                            },
                        },
                        scales: {
                            r: {
                                ticks: {
                                    display: false
                                },
                                pointLabels: {
                                    color: '#afafaf'
                                },
                                angleLines: {
                                    type: 'dashed',
                                    color: '#515180'
                                },
                                grid: {
                                    color: '#515180'
                                }
                            },

                        }
                    }
                    }
                />
            </div>
    )
}
