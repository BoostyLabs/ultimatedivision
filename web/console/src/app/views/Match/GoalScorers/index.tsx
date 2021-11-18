// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from "react";

import { GoalScorersTeam } from "@/app/views/Match/GoalScorers/GoalScorersTeam";

import Player from "@/app/static/img/match/player.svg";

import "./index.scss";

/** The maximum number of players scoring a goal, at which a block scroll is not needed. */
const SCORED_PLAYERS_COUNT: number = 3;

export const GoalScorers: React.FC = () => {
    const [isVisibleScroll, setIsVisibleScroll] = useState<{
        team_1: boolean;
        team_2: boolean;
    }>({ team_1: false, team_2: false });

    // TODO: Mock data players score.
    const firstTeamGoalScorers = [
        { logo: Player, name: "Ronalculus", goalTime: "18:00", goals: 1 },
        { logo: Player, name: "Ronalculus", goalTime: "25:01", goals: 2 },
        { logo: Player, name: "Ronalculus", goalTime: "44:13", goals: 3 },
        { logo: Player, name: "Ronalculus", goalTime: "55:99", goals: 4 },
        { logo: Player, name: "Ronalculus", goalTime: "78:99", goals: 5 },
        { logo: Player, name: "Ronalculus", goalTime: "88:99", goals: 6 },
        { logo: Player, name: "Ronalculus", goalTime: "98:99", goals: 7 },
        { logo: Player, name: "Ronalculus", goalTime: "99:99", goals: 8 },
    ];

    // TODO: Mock data players score.
    const secondTeamGoalScorers = [
        { logo: Player, name: "Ronalculus", goalTime: "01:00", goals: 1 },
        { logo: Player, name: "Ronalculus", goalTime: "34:58", goals: 2 },
        { logo: Player, name: "Ronalculus", goalTime: "82:44", goals: 3 },
    ];

    useEffect(() => {
        const copyIsVisibleScroll = { ...isVisibleScroll };

        /** If the length of the array of players who scored a goal is more than 3 - add a scroll for the block. */
        if (firstTeamGoalScorers.length > SCORED_PLAYERS_COUNT) {
            copyIsVisibleScroll.team_1 = true;
        }

        if (secondTeamGoalScorers.length > SCORED_PLAYERS_COUNT) {
            copyIsVisibleScroll.team_2 = true;
        }

        setIsVisibleScroll(copyIsVisibleScroll);
    }, []);

    return (
        <div className="match__goal-scorers">
            <div className="match__goal-scorers__team-1">
                <span className="match__goal-scorers__title">Team 1</span>
                <div
                    className={`scores${
                        isVisibleScroll.team_1 ? "-scroll" : ""
                    }`}
                >
                    <GoalScorersTeam team={firstTeamGoalScorers} />
                </div>
            </div>
            <div className="match__goal-scorers__team-2">
                <span className="match__goal-scorers__title">Team 2</span>
                <div
                    className={`scores${
                        isVisibleScroll.team_2 ? "-scroll" : ""
                    }`}
                >
                    <GoalScorersTeam team={secondTeamGoalScorers} />
                </div>
            </div>
        </div>
    );
};
