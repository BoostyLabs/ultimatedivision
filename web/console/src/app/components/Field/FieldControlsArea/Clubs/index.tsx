// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";

import { Manager, Owner, Cross } from "@static/img/FieldPage/clubs";

import { RootState } from "@/app/store";
import { Club } from "@/club";

import "./index.scss";

const Clubs: React.FC = () => {
    const [activeClub, setActiveClub] = useState<string>("CLUB 1");
    const [clubId, setClubId] = useState<number | null>(null);

    const dispatch = useDispatch();

    const currentClub = useSelector(
        (state: RootState) => state.clubsReducer.activeClub
    );

    // const clubs = useSelector(
    //     (state: RootState) => state.clubsReducer.clubs
    // );

    // console.log("clubs", clubs);

    // TODO: Mock data
    const clubs: Array<{ logo: any; name: string; whose: string }> = [
        { logo: <Owner />, name: "CLUB 1", whose: "owner" },
        { logo: <Owner />, name: "CLUB 2", whose: "owner" },
        { logo: <Owner />, name: "CLUB 3", whose: "owner" },
        { logo: <Manager />, name: "CLUB 4", whose: "manager" },
        { logo: <Manager />, name: "CLUB 5", whose: "manager" },
        { logo: <Manager />, name: "CLUB 6", whose: "manager" },
    ];

    /** Show or hide helper for clubs. */
    const visabilityClubsHelper = (index: number, club: any) =>
        clubId === index && (
            <div className="club__info">
                {club.whose === "owner"
                    ? `are you the ${club.whose}`
                    : "you are the manager"}
            </div>
        );

    return (
        <div className="field-controls-area__clubs">
            <span className="field-controls-area__clubs-name">
                {/* {currentClub.name} */}
                {activeClub}
            </span>
            <div className="field-controls-area__clubs__add">
                <Cross />
            </div>
            <div className="clubs">
                {clubs &&
                    clubs.map((club: any, index: number) => (
                        <>
                            <div
                                className={`club ${
                                    // index === currentClub.index ? "active" : ""
                                    club.name === activeClub ? "active" : ""
                                }`}
                                key={index}
                                onClick={() => setActiveClub(club.name)}
                                onMouseLeave={() => setClubId(null)}
                                onMouseEnter={() => setClubId(index)}
                            >
                                {club.logo}
                                <span className="club__name" key={index}>
                                    {club.name}
                                </span>
                                {visabilityClubsHelper(index, club)}
                            </div>
                        </>
                    ))}
            </div>
        </div>
    );
};

export default Clubs;
