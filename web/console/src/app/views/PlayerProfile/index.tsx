// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import PlayerList from '@/app/components/playerProfile/PlayerList';
import Statistics from '@/app/components/playerProfile/Statistics';

import profileLogo from '@static/img/profile/player-logo.png';
import owner from '@static/img/profile/owner.svg';
import manager from '@static/img/profile/manager.svg';

import './index.scss';

const PlayerProfile: React.FC = () =>
    <section className="profile">
        <div className="profile__wrapper">
            <div className="profile__info">
            <div className="profile__info__gradient"></div>
                <img className="logo" src={profileLogo} alt="Player logo" />
                <span className="player-name">player one</span>
            </div>
            <PlayerList title='Owned' logo={owner} />
            <PlayerList title='Managed' logo={manager} />
            <Statistics />
        </div>
    </section>;

export default PlayerProfile;
