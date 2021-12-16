// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { LootboxClient } from '@/api/lootboxes';
import { Lootbox } from '@/lootbox';
import { LootboxService } from '@/lootbox/service';
import { createAsyncThunk } from '@reduxjs/toolkit';

const client = new LootboxClient();
const service = new LootboxService(client);
/** thunk that handles opening new lootbox */
export const openLootbox = createAsyncThunk(
    'lootboxes/openLootbox',
    async function (lootbox: Lootbox) {
        return await service.buy(lootbox);
    });
