// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';
import { useSelector } from 'react-redux';

import { GoalScorersTeam } from '@/app/views/Match/GoalScorers/GoalScorersTeam';

import { RootState } from '@/app/store';

import './index.scss';

/** The maximum number of players scoring a goal, at which a block scroll is not needed. */
const MAX_SCORED_PLAYERS_COUNT: number = 3;

export const GoalScorers: React.FC = () => {
    const { matchResults } = useSelector((state: RootState) => state.matchesReducer.gameResult);

    /** FIRST_TEAM_INDEX is variable that describes first team index in teams array. */
    const FIRST_TEAM_INDEX: number = 0;
    /** SECOND_TEAM_INDEX is variable that describes second team index in teams array. */
    const SECOND_TEAM_INDEX: number = 1;

    const [visabilityTeamsAreaScroll, setVisabilityTeamsAreaScroll] = useState<{
        firstTeam: boolean;
        secondTeam: boolean;
    }>({ firstTeam: false, secondTeam: false });

    useEffect(() => {
        /** If the length of the array of players who scored a goal is more than 3 - add a scroll for the block. */
        setVisabilityTeamsAreaScroll({
            firstTeam: matchResults[FIRST_TEAM_INDEX].goals && matchResults[FIRST_TEAM_INDEX].goals.length > MAX_SCORED_PLAYERS_COUNT,
            secondTeam: matchResults[SECOND_TEAM_INDEX].goals && matchResults[SECOND_TEAM_INDEX].goals.length > MAX_SCORED_PLAYERS_COUNT,
        });
    }, []);

    return (
        <div className="match__goal-scorers">
            <div className="match__goal-scorers__team-1">
                <span className="match__goal-scorers__title">Team 1</span>
                <div
                    className={`scores${visabilityTeamsAreaScroll.firstTeam ? '-scroll' : ''
                    }`}
                >
                    {matchResults[FIRST_TEAM_INDEX].goals && <GoalScorersTeam goals={matchResults[FIRST_TEAM_INDEX].goals} />}
                </div>
            </div>
            <div className="match__goal-scorers__team-2">
                <span className="match__goal-scorers__title">Team 2</span>
                <div
                    className={`scores${visabilityTeamsAreaScroll.secondTeam ? '-scroll' : ''
                    }`}
                >
                    {matchResults[SECOND_TEAM_INDEX].goals && <GoalScorersTeam goals={matchResults[SECOND_TEAM_INDEX].goals} />}
                </div>
            </div>
        </div>
    );
};
