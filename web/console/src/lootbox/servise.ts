// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { LootboxClient } from "@/api/lootboxes";
import { Lootbox } from ".";

/** exposes all lootbox related logic */
export class Lootboxes {
    private readonly lootboxes: LootboxClient;
    public constructor(lootboxes: LootboxClient) {
        this.lootboxes = lootboxes;
    };
    /** handles lootbox buying */
    public async buy(lootbox: Lootbox): Promise<Response> {
        return await this.lootboxes.buy(lootbox);
    };
    /** handles lootbox opening */
    public async open(lootbox: Lootbox): Promise<Response> {
        return await this.lootboxes.open(lootbox);
    };
};
