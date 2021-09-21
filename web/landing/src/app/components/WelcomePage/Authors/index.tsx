// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { BoostyLogo, ChickenfishLogo, FirstRowProjects, SecondRowProjects } from '@static/images/authorsPage/authors';
import './index.scss';

export const Authors: React.FC = () => (
    <section className="authors">
        <div className="authors__wrapper">
            <span className="authors__wrapper-title">The game was created by a team involved in the development of well-know crypto projects</span>
            <div className="authors__wrapper-projects">
                <FirstRowProjects />
                <SecondRowProjects />
            </div>
            <span className="authors__wrapper-title">Created by</span>
            <div className="authors__wrapper-created-by">
                <ChickenfishLogo />
                <BoostyLogo />
            </div>
            <div className="authors__wrapper-created-by__name">
                <span className="chickenfish">CHICKENFISH GAMES</span>
                <span className="boosty-labs">BOOSTY LABS</span>
            </div>
        </div>
    </section>
);
