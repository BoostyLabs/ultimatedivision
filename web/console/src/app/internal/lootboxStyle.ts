// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import coolBoxCover from "@static/img/StorePage/BoxContent/coolBoxCover.svg";
import regularBoxCover from "@static/img/StorePage/BoxContent/regularBoxCover.svg";
import boxBg from "@static/img/StorePage/BoxContent/boxLight.svg";

/** function for getting right box for animation */
export function boxStyle(length: number) {
    const REGULAR_BOX_LENGTH = 5;
    const box = {
        body: boxBg,
        cover: regularBoxCover,
    };
    if (length > REGULAR_BOX_LENGTH) {
        box.body = boxBg;
        box.cover = coolBoxCover;
    }

    return box;
}
