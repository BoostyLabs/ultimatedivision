//Copyright (C) 2021 Creditor Corp. Group.
//See LICENSE for copying information.

import "./index.scss";

export const GameScore: React.FC = () => {
    return (
        <div className="game-score">
            <div className="game-score__scoreboard">
                <div className="game-score__scoreboard__game-timer">90:00</div>
                <div className="game-score__scoreboard__match-result">
                    <div className="game-score__scoreboard-team-1">8</div>
                    <div className="game-score__scoreboard-dash">-</div>
                    <div className="game-score__scoreboard-time-2">5</div>
                </div>
            </div>
        </div>
    );
};
