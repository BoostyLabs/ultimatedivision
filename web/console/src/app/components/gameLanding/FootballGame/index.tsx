// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { JoinButton } from '@components/common/JoinButton';

import footballField from '@static/img/gameLanding/main/football-field.svg';

import './index.scss';

export const FootballGame: React.FC = () =>
    <section className="football-game">
        <picture>
            <img src="" alt="Player cards" />
        </picture>
        <span>ULTIMATE DIVISION</span>
        <span>Football P2E Game</span>
        <JoinButton />
        <picture>
            <img src={footballField} alt="Football field" />
        </picture>
    </section>;

