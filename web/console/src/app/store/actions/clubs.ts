// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch } from 'redux';

import { CardEditIdentificators, ClubsClient } from '@/api/club';
import { Club, Squad } from '@/club';
import { Tactic, TacticsType, Formations, FormationsType, } from '@/app/types/club'
import { ClubService } from '@/club/service';
import { createAsyncThunk } from '@reduxjs/toolkit';
import { getList } from './cards';

type dragParamType = number | null;
const DEFAULT_CARD_INDEX = null;

const client = new ClubsClient();
const service = new ClubService(client);

// Thunks

export const createClub = createAsyncThunk(
    'clubs/createClub',
    async function () {
        const clubId = await service.createClub();
        const squadId = await service.createSquad(clubId);
        return await service.getClubs();
    })

export const getClubs = createAsyncThunk(
    'clubs/getClubs',
    async function () {
        return await service.getClubs();
    })

export const setFormation = createAsyncThunk(
    'clubs/setFormation',
    async function (formation: FormationsType) {
        const club = await service.getActiveClub()
        await service.updateFormation(club.squad, Formations[formation]);
        return await service.getClubs();
    })

export const setCaptain = createAsyncThunk(
    'clubs/setCaptain',
    async function (captainId: string) {
        const club = await service.getActiveClub()
        await service.updateCaptain(club.squad, captainId);
        return await service.getClubs();
    });

export const setTactic = createAsyncThunk(
    'clubs/setTactic',
    async function (tactic: TacticsType) {
        const club = await service.getActiveClub()
        await service.updateTactic(club.squad, Tactic[tactic]);
        return await service.getClubs();
    });

export const addCard = createAsyncThunk(
    'clubs/addCard',
    async function (path: CardEditIdentificators) {
        await service.addCard(path);
        return await service.getClubs();
    });

export const deleteCard = createAsyncThunk(
    'clubs/addCard',
    async function (path: CardEditIdentificators) {
        await service.deleteCard(path);
        return await service.getClubs();
    });

export const changePosition = createAsyncThunk(
    'clubs/changePosition',
    async function (path: CardEditIdentificators) {
        await service.changeCardPosition(path);
        return await service.getClubs();
    });
export const swapCards = createAsyncThunk(
    'clubs/swapCards',
    async function (path: { currentCard: CardEditIdentificators, existCard: CardEditIdentificators }) {
        await service.changeCardPosition(path.currentCard);
        await service.changeCardPosition(path.existCard);
        return await service.getClubs();
    });

    export const changeActiveClub = createAsyncThunk(
        'clubs/changeActiveClub',
    async function (id: string) {
        await service.changeActiveClub(id);
        return await service.getClubs();
    })ZZz;
