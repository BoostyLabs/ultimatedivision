// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import boxBody from '@static/img/StorePage/BoxContent/boxBody.svg'
import boxCover from '@static/img/StorePage/BoxContent/boxCover.svg'

import './index.scss';


export const BoxAnimation = () => {
    return (
        <div className="box-animation">
            <img
                src={boxBody}
                alt="box body"
                className="box-animation__box-body"
            />
            <img
                src={boxCover}
                alt="box cover"
                className="box-animation__box-cover"
            />
        </div>
    )
}
