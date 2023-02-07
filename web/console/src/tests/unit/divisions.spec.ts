// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import configureStore from 'redux-mock-store';
import { afterEach, beforeEach, describe, expect, it } from '@jest/globals';
import { useDispatch, useSelector } from "react-redux";
import { cleanup } from "@testing-library/react";

import { DivisionsClient } from '@/api/divisions';
import { GET_CURRENT_DIVISION_SEASONS } from '@/app/store/actions/divisions';
import { CurrentDivisionSeasons, DivisionsState, DivisionSeasonsStatistics } from '@/divisions';

const divisionsClient = new DivisionsClient();

const mockStore = configureStore();

const successFetchMock = async (body: any) => {
    globalThis.fetch = () =>
        Promise.resolve({
            json: () => Promise.resolve(body),
            ok: true,
            status: 200,
        }) as Promise<Response>;
};

const failedFetchMock = async () => {
    globalThis.fetch = () => {
        throw new Error();
    };
};

const mockedGlobalFetch = globalThis.fetch;


/** Mock initial networks state. */
const initialState = {
    divisionsReducer: {
        currentDivisionSeasons: new CurrentDivisionSeasons('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', new Date('2023-02-07T01:13:52.114Z'), new Date('2023-02-07T01:13:52.114Z')),
        divisionSeasonsStatistics: new DivisionSeasonsStatistics(),
        activeDivision: new CurrentDivisionSeasons('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', new Date('2023-02-07T01:13:52.114Z'), new Date('2023-02-07T01:13:52.114Z'))
    }
};

/** Mock division state. */
const MOCK_DIVISIONS_STATE = {
    currentDivisionSeasons: [{
        id: '00000000-0000-0000-0000-000000000000',
        divisionId: '00000000-0000-0000-0000-000000000000',
        startedAt: new Date('2023-02-07T01:13:52.114Z'),
        endedAt: new Date('2023-02-07T01:13:52.114Z'),
    }],
    divisionSeasonsStatistics: new DivisionSeasonsStatistics(),
    activeDivision: {
        id: '00000000-0000-0000-0000-000000000000',
        divisionId: '00000000-0000-0000-0000-000000000000',
        startedAt: new Date('2023-02-07T01:13:52.114Z'),
        endedAt: new Date('2023-02-07T01:13:52.114Z'),
    },
}

const reactRedux = { useDispatch, useSelector }
const useDispatchMock = jest.spyOn(reactRedux, "useDispatch");
const useSelectorMock = jest.spyOn(reactRedux, "useSelector");
let updatedStore: any = mockStore(initialState);
const mockDispatch = jest.fn();
useDispatchMock.mockReturnValue(mockDispatch);
updatedStore.dispatch = mockDispatch;

describe('Requests list of cards.', () => {
    beforeEach(() => {
        successFetchMock(MOCK_DIVISIONS_STATE.currentDivisionSeasons);
    });

    afterEach(() => {
        globalThis.fetch = mockedGlobalFetch;
    });

    it('Requests divisions seasons.', async () => {
        const divisionSeasons = await divisionsClient.getCurrentDivisionSeasons();
        expect(divisionSeasons).toEqual(MOCK_DIVISIONS_STATE.currentDivisionSeasons);
    });

    describe('Failed response.', () => {
        beforeEach(() => {
            failedFetchMock();
            useSelectorMock.mockClear();
            useDispatchMock.mockClear();
        });

        afterEach(() => {
            globalThis.fetch = mockedGlobalFetch;
            cleanup();
        });

        it('Must be no cards', async () => {
            try {
                await divisionsClient.getCurrentDivisionSeasons();
            } catch (error) {
                mockDispatch(GET_CURRENT_DIVISION_SEASONS, {});
                expect(updatedStore.getState().divisionsReducer.currentDivisionSeasons).toEqual(MOCK_DIVISIONS_STATE.currentDivisionSeasons[0]);
            }
        });
    })
});

describe('Requests list of cards.', () => {
    beforeEach(() => {
        successFetchMock(MOCK_DIVISIONS_STATE.divisionSeasonsStatistics);
    });

    afterEach(() => {
        globalThis.fetch = mockedGlobalFetch;
    });

    it('Requests divisions seasons.', async () => {
        const divisionSeasons = await divisionsClient.getDivisionSeasonsStatistics('00000000-0000-0000-0000-000000000000');
        expect(divisionSeasons).toEqual(MOCK_DIVISIONS_STATE.divisionSeasonsStatistics);
    });

    describe('Failed response.', () => {
        beforeEach(() => {
            failedFetchMock();
            useSelectorMock.mockClear();
            useDispatchMock.mockClear();
        });

        afterEach(() => {
            globalThis.fetch = mockedGlobalFetch;
            cleanup();
        });

        it('Must be no cards', async () => {
            try {
                await divisionsClient.getCurrentDivisionSeasons();
            } catch (error) {
                mockDispatch(GET_CURRENT_DIVISION_SEASONS, {});
                expect(updatedStore.getState().divisionsReducer.divisionSeasonsStatistics).toEqual(MOCK_DIVISIONS_STATE.divisionSeasonsStatistics);
            }
        });
    })
});

