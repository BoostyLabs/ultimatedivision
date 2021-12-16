// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useAppSelector } from '@/app/store';

import { RootState } from '@/app/store';

import './index.scss';

export const MatchScore: React.FC = () => {
    const { firstTeam, secondTeam } = useAppSelector((state: RootState) => state.matches);

    return <div className="match__score">
        <div className="match__score__board">
            <div className="match__score__board__gradient"></div>
            <div className="match__score__board__timer">90:00</div>
            <div className="match__score__board__result">
                <div className="match__score__board-team-1">{firstTeam.quantityGoals}</div>
                <div className="match__score__board-dash">-</div>
                <div className="match__score__board-team-2">{secondTeam.quantityGoals}</div>
            </div>
        </div>
    </div>;
};
