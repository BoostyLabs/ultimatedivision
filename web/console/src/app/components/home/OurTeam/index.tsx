// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import ruslanAzarov from '@static/img/gameLanding/ourTeam/teammember.png';

import './index.scss';

/** Defines member of team  */
export class Member {
    constructor(
        public photo: string = '',
        public name: string = '',
        public position: string = '',
        public description: string = ''
    ) {}
}

export const OurTeam: React.FC = () => {
    /**  TODO: Mock data, will be replaced */
    const members: Member[] = [
        new Member(ruslanAzarov, 'Ruslan Azarov', 'CEO', '5 years in gamedev'),
        new Member(ruslanAzarov, 'Ruslan Azarov', 'CEO', '5 years in gamedev'),
        new Member(ruslanAzarov, 'Ruslan Azarov', 'CEO', '5 years in gamedev'),
        new Member(ruslanAzarov, 'Ruslan Azarov', 'CEO', '5 years in gamedev'),
        new Member(ruslanAzarov, 'Ruslan Azarov', 'CEO', '5 years in gamedev'),
        new Member(ruslanAzarov, 'Ruslan Azarov', 'CEO', '5 years in gamedev'),
        new Member(ruslanAzarov, 'Ruslan Azarov', 'CEO', '5 years in gamedev'),
        new Member(ruslanAzarov, 'Ruslan Azarov', 'CEO', '5 years in gamedev'),
    ];

    return (
        <section className="team">
            <div className="team__wrapper">
                <h2 className="team__title">Our Team</h2>
                <div className="team__members">
                    {members &&
                        members.map((member, index) =>
                            <div className="team__members__item" key={member.photo}>
                                <img src={member.photo} alt={member.name} className="team__members__item__photo" />
                                <p className="team__members__item__name">{member.name}</p>
                                <p className="team__members__item__position">{member.position}</p>
                                <p className="team__members__item__description">{member.description}</p>
                            </div>
                        )}
                </div>
            </div>
        </section>
    );
};
