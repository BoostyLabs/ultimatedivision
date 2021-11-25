// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { toast } from 'react-toastify';

import realMadrid from '@static/img/divisions/realmadrid.png';
import rectangle from '@static/img/FilterField/rectangle.svg';

import { RootState } from '@/app/store';
import {
    listOfCurrentSeasonsDivisions,
    divisionSeasonsStatistics,
    setActiveDivision,
} from '@/app/store/actions/divisions';
import { DivisionClub } from '@/app/types/division';
import { CurrentSeasonsDivision } from '@/divisions';

import './index.scss';

const Division: React.FC = () => {
    const dispatch = useDispatch();

    const { currentSeasonsDivisions, seasonsStatistics, activeDivision } =
        useSelector((state: RootState) => state.divisionsReducer);

    /** Get current seasons divisions. */
    async function getCurrentSeasonsDivisions() {
        try {
            await dispatch(listOfCurrentSeasonsDivisions());
        } catch (error: any) {
            toast.error('Failed to get current seasons divisions', {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
        }
    }

    /** Get divisions seasons statistics. */
    async function getSeasonsStatistics() {
        try {
            await dispatch(divisionSeasonsStatistics());
        } catch (error: any) {
            toast.error('Failed to get seasons statistics', {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
        }
    }

    useEffect(() => {
        getCurrentSeasonsDivisions();
        getSeasonsStatistics();
    }, []);

    /** TODO: replase test datas. Just for test rigth now. */
    const divisionClub: DivisionClub = {
        position: '1',
        club: {
            name: 'Real Madrid',
            icon: realMadrid,
        },
        games: '30',
        wins: '20',
        draws: '6',
        defeats: '5',
        goalDifference: '+50',
        points: '50',
    };

    /** TODO: delete this. Rigth now for test */
    const CLUBS_COUNT: number = 50;
    const divisionClubs = new Array(CLUBS_COUNT).fill(divisionClub);

    /** Variables describes upper and lower clubs position
     * in general division table. */
    const COEFFICIENT: number = 0.1;

    const UPPER_BREAKPOINT: number = Math.round(
        divisionClubs.length * COEFFICIENT
    );
    const LOWER_BREAKPOINT: number = divisionClubs.length - UPPER_BREAKPOINT;

    const titles: string[] = [
        '#',
        'Club',
        'Games',
        'Wins',
        'Draws',
        'Defeats',
        'The difference of goals conceded',
        'Points',
    ];

    const changeGradationDivisionClassName = (position: string) => {
        let className: string = '';

        if (+position <= UPPER_BREAKPOINT) {
            className = '-upper';
        }

        if (+position >= LOWER_BREAKPOINT) {
            className = '-lower';
        }

        return className;
    };

    /** Check active division and change him className */
    const changeChoosedDivisionClass = (division: CurrentSeasonsDivision) => `division__item${
        division.id === activeDivision.id ? '-active' : '-inactive'
    }`;

    return (
        <section className="division">
            <div className="division__titles">
                <h1 className="division__titles__main">
                    Division {activeDivision.id}
                </h1>
                <span className="division__titles__count">
                    {CLUBS_COUNT}
                    <span className="division__titles__count__text">Teams</span>
                </span>
            </div>
            <div className="division__list">
                {currentSeasonsDivisions.map(
                    (division: CurrentSeasonsDivision, index: number) =>
                        <div
                            className={changeChoosedDivisionClass(division)}
                            key={index}
                            onClick={() => dispatch(setActiveDivision(division))}
                        >
                            Division {division.id}
                        </div>

                )}
            </div>
            <div className="division__filters">
                <span className="division__filters__title">Filters</span>
                <div className="division__filters__item">
                    <input
                        id="division-checkbox-1"
                        className="division__filters__item__checkbox"
                        type="checkbox"
                    />
                    <label
                        className="division__filters__item__text"
                        htmlFor="division-checkbox-1"
                    >
                        played all the matches
                    </label>
                </div>
                <div className="division__filters__item">
                    <input
                        id="division-checkbox-2"
                        className="division__filters__item__checkbox"
                        type="checkbox"
                    />
                    <label
                        className="division__filters__item__text"
                        htmlFor="division-checkbox-2"
                    >
                        did not play in the current draw
                    </label>
                </div>
            </div>
            <table className="division__clubs">
                <thead>
                    <tr className="division__clubs__titles">
                        {titles.map((title: string, index: number) =>
                            <th
                                key={index}
                                className="division__clubs__titles__item"
                            >
                                {title}
                            </th>
                        )}
                    </tr>
                </thead>
                <tbody>
                    {divisionClubs.map(
                        (divisionClub: DivisionClub, index: number) =>
                            <tr
                                className={`division__clubs__club${changeGradationDivisionClassName(
                                    divisionClub.position
                                )}`}
                                key={index}
                            >
                                <td className="division__clubs__club__item">
                                    <span className="division__clubs__club__item__position">
                                        {divisionClub.position}
                                    </span>
                                </td>
                                <td className="division__clubs__club__item">
                                    <img
                                        src={divisionClub.club.icon}
                                        className="division__clubs__club__item__icon"
                                    />
                                    {divisionClub.club.name}
                                </td>
                                <td className="division__clubs__club__item">
                                    {divisionClub.games}
                                </td>
                                <td className="division__clubs__club__item">
                                    {divisionClub.wins}
                                </td>
                                <td className="division__clubs__club__item">
                                    {divisionClub.draws}
                                </td>
                                <td className="division__clubs__club__item">
                                    {divisionClub.defeats}
                                </td>
                                <td className="division__clubs__club__item">
                                    {divisionClub.goalDifference}
                                </td>
                                <td className="division__clubs__club__item">
                                    {divisionClub.points}
                                </td>
                            </tr>

                    )}
                </tbody>
            </table>
        </section>
    );
};

export default Division;
