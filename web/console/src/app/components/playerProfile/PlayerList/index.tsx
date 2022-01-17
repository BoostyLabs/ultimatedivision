// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import './index.scss';

const PlayerList: React.FC<{title: string, logo: string}> = ({title, logo}) =>
    <div className="profile__player-list">
        <div>
            <img src={logo} alt={`${logo} logo`} />
            <span>Clubs {title} by Player</span>
        </div>
        <div></div>
        <button><span>show more</span></button>
    </div>;

export default PlayerList;
