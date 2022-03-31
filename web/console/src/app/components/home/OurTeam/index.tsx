// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import RuslanAzarov from '@static/img/gameLanding/ourTeam/teammember.png';

import './index.scss';

export class Team {
    constructor(
        public photo: string = '',
        public name: string = '',
        public position: string = '',
        public description: string = ''
    ) {
        this.photo = photo;
        this.name = name;
        this.position = position;
        this.description = description;
    }
}

export const OurTeam: React.FC = () => {
    const TEAM: Team[] = [
        new Team(RuslanAzarov, 'Ruslan Azarov', 'CEO', '5 years in gamedev'),
        new Team(RuslanAzarov, 'Ruslan Azarov', 'CEO', '5 years in gamedev'),
        new Team(RuslanAzarov, 'Ruslan Azarov', 'CEO', '5 years in gamedev'),
        new Team(RuslanAzarov, 'Ruslan Azarov', 'CEO', '5 years in gamedev'),
        new Team(RuslanAzarov, 'Ruslan Azarov', 'CEO', '5 years in gamedev'),
        new Team(RuslanAzarov, 'Ruslan Azarov', 'CEO', '5 years in gamedev'),
        new Team(RuslanAzarov, 'Ruslan Azarov', 'CEO', '5 years in gamedev'),
        new Team(RuslanAzarov, 'Ruslan Azarov', 'CEO', '5 years in gamedev'),
    ];

    return (
        <section className="team">
            <div className="team__wrapper">
                <h2 className="team__title">Our Team</h2>
                <div className="team__members">
                    {TEAM.map((teammember, index) =>
                        <div className="team__members__item" key={teammember.photo}>
                            <img src={teammember.photo} alt={teammember.name} className="team__members__item__photo" />
                            <p className="team__members__item__name">{teammember.name}</p>
                            <p className="team__members__item__position">{teammember.position}</p>
                            <p className="team__members__item__description">{teammember.description}</p>
                        </div>
                    )}
                </div>
            </div>
        </section>
    );
};
