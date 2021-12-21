// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';

import { Manager, Owner, Cross } from '@static/img/FieldPage/clubs';

import './index.scss';

const Squads: React.FC = () => {
    const [activeComposition, setActiveComposition] =
        useState<string>('Composition 1');
    const [isActiveDropdown, setIsActiveDropdown] = useState<boolean>(false);

    // TODO: Mock data
    const compositions: string[] = [
        'Composition 1',
        'Composition 2',
        'Composition 3',
        'Composition 4',
    ];

    /** Method for set choosed composition to state and close dropdown block. */
    const handleChooseComposition = (composition: string) => {
        setActiveComposition(composition);
        setIsActiveDropdown(false);
    };

    return (
        <div className="field-controls-area__squads">
            <div className="field-controls-area__squads-composition">
                <div className="composition">
                    <div
                        className={`composition__choosed-item ${
                            isActiveDropdown ? 'active-dropdown' : ''
                        }`}
                        onClick={() => setIsActiveDropdown(!isActiveDropdown)}
                    >
                        {activeComposition}
                    </div>
                    <div
                        className={`composition__list${
                            isActiveDropdown ? '-active' : ''
                        }`}
                    >
                        {compositions.map((composition, index) =>
                            <div
                                className="composition__list-item"
                                key={index}
                                onClick={() =>
                                    handleChooseComposition(composition)
                                }
                            >
                                <span>{composition}</span>
                            </div>
                        )}
                    </div>
                </div>
            </div>
            <div className="field-controls-area__squads-add">
                <Cross />
            </div>
        </div>
    );
};

export default Squads;
