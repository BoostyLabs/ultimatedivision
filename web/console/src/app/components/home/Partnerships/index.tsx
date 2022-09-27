// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import devdao from '@static/img/gameLanding/partnerships/devdao.svg';
import boosty from '@static/img/gameLanding/partnerships/boosty.svg';
import casper from '@static/img/gameLanding/partnerships/casper.svg';
import storj from '@static/img/gameLanding/partnerships/storj.svg';
import polygon from '@static/img/gameLanding/partnerships/polygon.svg';
import chickenfish from '@static/img/gameLanding/partnerships/chickenfish.svg';
import velas from '@static/img/gameLanding/partnerships/velas.svg';

import './index.scss';

/** Domain entity Partnership implementation */
class Partnership {
    /** default partnership implementation */
    constructor(public name: string = '', public logo: string = '') {}
}

export const Partnerships: React.FC = () => {
    /** Defines logos of partner companies */
    const logos = [
        new Partnership('polygon', polygon),
        new Partnership('velas', velas),
        new Partnership('casper', casper),
        new Partnership('devdao', devdao),
        new Partnership('storj', storj),
        new Partnership('boosty', boosty),
        new Partnership('chickenfish', chickenfish),
    ];

    return (
        <section className="partnerships">
            <div className="partnerships__wrapper">
                <h2 className="partnerships__title">Our <span className="partnerships__title__second-part">Partners</span></h2>
                <div className="partnerships__area">
                    {logos.map((logo, index: number) =>
                        <div key={index} className="partnerships__area__item">
                            <div className="partnerships__area__item__wrapper">
                                <img className={`partnerships__area__item__logo partnerships__area__item__logo__${logo.name}`} key={index} src={logo.logo} alt="logo" />
                            </div>
                        </div>
                    )}
                </div>
            </div>
        </section>
    );
};
