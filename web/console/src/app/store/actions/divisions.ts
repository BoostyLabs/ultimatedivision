// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { DivisionsClient } from '@/api/divisions';
import { DivisionsService } from '@/divisions/service';
import { createAsyncThunk } from '@reduxjs/toolkit';

const client = new DivisionsClient();
const service = new DivisionsService(client);

/** thunk that handles gets current seasons divisions */
export const getCurrentDivisionSeasons = createAsyncThunk(
    'divisions/getCurrentDivisionSeasons',
    async function () {
        return await service.getCurrentDivisionSeasons();
    });

/** thunk that handles gets seasons statistics */
export const getDivisionSeasonsStatistics = createAsyncThunk(
    'divisions/getDivisionSeasonsStatistics',
    async function (id: string) {
        return await service.getDivisionSeasonsStatistics(
            id
        );
    });
