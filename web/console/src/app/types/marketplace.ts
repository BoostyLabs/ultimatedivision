// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/** base created marketplace Lot implementation */
export class CreatedLot {
    /** default created lot initital values */
    constructor(
        public itemId: string,
        public startPrice: number,
        public maxPrice: number,
        public period: number,
    ) { };
};
