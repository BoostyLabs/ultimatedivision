// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { RootState } from '@/app/store';
import { divisionSeasonsStatistics } from '@/app/store/actions/divisions';
import { DivisionClub } from '@/app/types/division';
import { ToastNotifications } from '@/notifications/service';
import { DivisionsClient } from '@/api/divisions';
import { DivisionsService } from '@/divisions/service';
import { setCurrentUser } from '@/app/store/actions/users';
import WalletService from '@/wallet/service';

import realMadrid from '@static/img/divisions/realmadrid.png';

import './index.scss';

const Division: React.FC = () => {
    const DEFAULT_SEASONS_REWARD_TOKENS_STATUS = 0;
    const divisionClient = new DivisionsClient();
    const divisionService = new DivisionsService(divisionClient);

    const dispatch = useDispatch();

    const [seasonRewardStatus, setSeasonRewardStatus] = useState<number>(DEFAULT_SEASONS_REWARD_TOKENS_STATUS);

    const { currentDivisionsSeasons, seasonsStatistics, activeDivision } =
        useSelector((state: RootState) => state.divisionsReducer);

    const [activeDivisions, setActiveDivisions] = useState<string>('10');

    const user = useSelector((state: RootState) => state.usersReducer.user);

    /** sets user info */
    async function setUser() {
        try {
            await dispatch(setCurrentUser());
        } catch (error: any) {
            ToastNotifications.couldNotGetUser();
        }
    }

    /** sets user info */
    async function setSeasonStatus() {
        await divisionService.getSeasonStatus();
    }

    /** Get divisions seasons status. */
    async function seasonsRewardStatus() {
        try {
            const seasonsRewardStatus = await divisionService.seasonsRewardStatus();
            setSeasonRewardStatus(seasonsRewardStatus);
        }
        catch {
            setSeasonRewardStatus(DEFAULT_SEASONS_REWARD_TOKENS_STATUS);
        }
    }

    /** Get divisions seasons statistics. */
    async function getSeasonsStatistics() {
        try {
            await dispatch(divisionSeasonsStatistics(activeDivisions));
        } catch (error: any) {
            ToastNotifications.failedGettingSeasonStatistics();
        }
    }

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
    const changeChoosedDivisionClass = (div: string) =>
        `division__item${activeDivisions === div ? '-active' : '-inactive'}`;

    // TODO: Mock data (waiting backend).
    const divisions: string[] = [
        '1',
        '2',
        '3',
        '4',
        '5',
        '6',
        '7',
        '8',
        '9',
        '10',
    ];

    const signTokens = async() => {
        const transactionData = await divisionService.getDivisionSeasonsReward();
        const walletService = new WalletService(user);

        await walletService.mintSeasonToken(transactionData);
    };

    useEffect(() => {
        getSeasonsStatistics();
    }, [activeDivisions]);

    useEffect(() => {
        setUser();
        setSeasonStatus();
        seasonsRewardStatus();
    }, []);

    return (
        <section className="division">
            <div className="division__titles">
                <h1 className="division__titles__main">
                    Division {activeDivisions}
                </h1>
                <span className="division__titles__count">
                    {CLUBS_COUNT}
                    <span className="division__titles__count__text">Teams</span>
                </span>
                {
                    seasonRewardStatus > DEFAULT_SEASONS_REWARD_TOKENS_STATUS &&
                        <button onClick={() => signTokens()}>Get Reward</button>
                }
            </div>
            <div className="division__list">
                {divisions.map((division: string, index: number) =>
                    <div
                        className={changeChoosedDivisionClass(division)}
                        key={index}
                        onClick={() => setActiveDivisions(division)}
                    >
                        Division&nbsp;{division}
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

            {!seasonsStatistics.statistics ?
                <>
                    <h2 className="division__clubs__no-results">
                        You need to play at least 3 matches, but not more than 30
                    </h2>
                </>
                : <>

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

                            {/* TODO: Need change type of divisionClub */}
                            {seasonsStatistics.statistics.map(
                                (divisionClub: any, index: number) =>
                                    <tr
                                        className={`division__clubs__club${changeGradationDivisionClassName(
                                            divisionClub.club.name
                                        )}`}
                                        key={index}
                                    >
                                        <td className="division__clubs__club__item">
                                            <span className="division__clubs__club__item__position">
                                                {activeDivisions}
                                            </span>
                                        </td>
                                        <td className="division__clubs__club__item">
                                            <img
                                                src=""
                                                className="division__clubs__club__item__icon"
                                                alt=""
                                            />
                                            <span>{divisionClub.club.name}</span>
                                        </td>
                                        <td className="division__clubs__club__item">
                                            {divisionClub.matchPlayed}
                                        </td>
                                        <td className="division__clubs__club__item">
                                            {divisionClub.wins}
                                        </td>
                                        <td className="division__clubs__club__item">
                                            {divisionClub.draws}
                                        </td>
                                        <td className="division__clubs__club__item">
                                            {divisionClub.losses}
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
                </>
            }
        </section>
    );
};

export default Division;
