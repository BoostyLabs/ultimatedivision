// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import Player from "@/app/static/img/match/player.svg";

import "./index.scss";

export const PlayersScore: React.FC = () => {
    // TODO: Mock data players score.
    const playersScoreTeam1 = [
        { logo: Player, name: "Ronalculus", goalTime: "18:00", goals: 1 },
        { logo: Player, name: "Ronalculus", goalTime: "25:01", goals: 3 },
        { logo: Player, name: "Ronalculus", goalTime: "44:13", goals: 2 },
        { logo: Player, name: "Ronalculus", goalTime: "99:99", goals: 2 },
    ];

    // TODO: Mock data players score.
    const playersScoreTeam2 = [
        { logo: Player, name: "Ronalculus", goalTime: "01:00", goals: 2 },
        { logo: Player, name: "Ronalculus", goalTime: "34:58", goals: 2 },
        { logo: Player, name: "Ronalculus", goalTime: "82:44", goals: 10 },
    ];

    return (
        <div className="match__players-score">
            <div className="match__players-score__team-1">
                <span className="match__players-score__title">Team 1</span>
                {playersScoreTeam1 &&
                    playersScoreTeam1.map((player, index) => (
                        <div className="player" key={index}>
                            <img
                                src={player.logo}
                                alt={`${player.name} player`}
                            ></img>
                            <span className="player__name">{player.name}</span>
                            <span className="player__goal-time">
                                {player.goalTime}
                            </span>
                            <div className="player__goals">{player.goals}</div>
                        </div>
                    ))}
            </div>
            <div className="match__players-score__team-2">
                <span className="match__players-score__title">Team 2</span>
                {playersScoreTeam2 &&
                    playersScoreTeam2.map((player, index) => (
                        <div className="player" key={index}>
                            <img
                                src={player.logo}
                                alt={`${player.name} player`}
                            ></img>
                            <span className="player__name">{player.name}</span>
                            <span className="player__goal-time">
                                {player.goalTime}
                            </span>
                            <div className="player__goals">{player.goals}</div>
                        </div>
                    ))}
            </div>
        </div>
    );
};
