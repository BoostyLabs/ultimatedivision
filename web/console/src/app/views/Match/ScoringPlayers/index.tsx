//Copyright (C) 2021 Creditor Corp. Group.
//See LICENSE for copying information.

import "./index.scss";

export const ScoringPlayers: React.FC = () => {
    // TODO: Mock data scoring players.
    const scoringPlayetsTeam1 = [
        { logo: "logo", name: "Ronalculus", goals: 1 },
        { logo: "logo", name: "Ronalculus", goals: 3 },
        { logo: "logo", name: "Ronalculus", goals: 2 },
        { logo: "logo", name: "Ronalculus", goals: 2 },
    ];

    // TODO: Mock data scoring players.
    const scoringPlayetsTeam2 = [
        { logo: "logo", name: "Ronalculus", goals: 2 },
        { logo: "logo", name: "Ronalculus", goals: 2 },
        { logo: "logo", name: "Ronalculus", goals: 1 },
    ];

    return (
        <div className="scoring-players">
            <div className="scoring-players__team-1">
                <span className="scoring-players__title">Team 1</span>
                {scoringPlayetsTeam1 &&
                    scoringPlayetsTeam1.map((player, index) => {
                        return (
                            <div className="player" key={index}>
                                <img src="" alt={`${player.name} player`}></img>
                                <span className="player__name">
                                    {player.name}
                                </span>
                                <div className="player__goals">
                                    {player.goals}
                                </div>
                            </div>
                        );
                    })}
            </div>
            <div className="scoring-players__team-2">
                <span className="scoring-players__title">Team 2</span>
                {scoringPlayetsTeam1 &&
                    scoringPlayetsTeam2.map((player, index) => {
                        return (
                            <div className="player" key={index}>
                                <img src="" alt={`${player.name} player`}></img>
                                <span className="player__name">
                                    {player.name}
                                </span>
                                <div className="player__goals">
                                    {player.goals}
                                </div>
                            </div>
                        );
                    })}
            </div>
        </div>
    );
};
