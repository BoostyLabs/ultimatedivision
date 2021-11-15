//Copyright (C) 2021 Creditor Corp. Group.
//See LICENSE for copying information.

import { GameScore } from "./GameScore";
import { ScoringPlayers } from "./ScoringPlayers";
import { PlayingField } from "./PlayingField";

import "./index.scss";

const Match: React.FC = () => {
    return (
        <div className="match">
            <GameScore />
            <ScoringPlayers />
            <PlayingField />
        </div>
    );
};

export default Match;
