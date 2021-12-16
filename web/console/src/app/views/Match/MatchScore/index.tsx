// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useSelector } from 'react-redux';

import coin from '@static/img/match/money.svg';

import { RootState } from '@/app/store';

import './index.scss';

export const MatchScore: React.FC = () => {
    const { firstTeam, secondTeam } = useSelector(
        (state: RootState) => state.matchesReducer
    );

    return (
        <div className="match__score">
            <div className="match__score__board">
                <div className="match__score__board__gradient"></div>
                <div className="match__score__board__timer">90:00</div>
                <div className="match__score__board__result">
                    <div className="match__score__board-team-1">
                        {firstTeam.quantityGoals}
                    </div>
                    <div className="match__score__board-dash">-</div>
                    <div className="match__score__board-team-2">
                        {secondTeam.quantityGoals}
                    </div>
                </div>
                <div className="match__score__board__coins">
                    <img
                        className="match__score__board__coins-image"
                        src={coin}
                        alt="Coin"
                    />
                    <span className="match__score__board__coins-value">
                        1,200,000
                    </span>
                    <button className="match__score__board__coins__btn">
                        <span className="match__score__board__coins__btn-text">
                            GET
                        </span>
                    </button>
                </div>
            </div>
        </div>
    );
};
