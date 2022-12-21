// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import configureStore from 'redux-mock-store';
import { afterEach, beforeEach, describe, expect, it } from '@jest/globals';
import { useDispatch, useSelector } from "react-redux";
import { cleanup } from "@testing-library/react";

import { UsersClient } from '@/api/users';
import { SET_USER } from '@/app/store/actions/users';
import { User } from '@/users';
import { UsersService } from '@/users/service';

const usersClient = new UsersClient();
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

/** Mock connected networks list. */
const MOCK_USER: User = 
    new User(
        "c00e750d541c86e5dc6892a8be43864c00000c854048a037c5d2befb3d0c6d70",
        "0202b2a13f20e71016aa8bbca5fbc1cca56af4092c000000e65dcfab2168ab051c92",
        "test@test.com",
        "00000000-0000-0000-0000-000000000000",
        "2021-12-17T00:31:52.437252Z",
        "test",
        "2022-12-17T00:31:51.874508Z",
        "0x0000000000000000000000000000000000000000",
        "casper-wallet",
    )
;

/** Mock initial networks state. */
const initialState = {
    usersReducer: {
        user: [],
        userService: new UsersService(usersClient)
    }
};

const reactRedux = { useDispatch, useSelector }
const useDispatchMock = jest.spyOn(reactRedux, "useDispatch");
const useSelectorMock = jest.spyOn(reactRedux, "useSelector");
let updatedStore: any = mockStore(initialState);
const mockDispatch = jest.fn();
useDispatchMock.mockReturnValue(mockDispatch);
updatedStore.dispatch = mockDispatch;

describe('Requests user.', () => {
    beforeEach(() => {
        successFetchMock(MOCK_USER);
    });

    afterEach(() => {
        globalThis.fetch = mockedGlobalFetch;
    });
    
    it('Success response', async () => {
        const connectedNetworks = await usersClient.getUser();
        expect(connectedNetworks).toEqual(MOCK_USER);
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
        
        it('Must be empty networks list', async () => {
            try {
                await usersClient.getUser();
            } catch (error) {
                mockDispatch(SET_USER, []);
                expect(updatedStore.getState().networksReducer.networks).toEqual([]);
            }
        });
    })
});

