// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from "react";

import { Manager, Owner } from "@static/img/FieldPage/clubs";

import "./index.scss";

const Clubs: React.FC = () => {
    const [activeClub, setActiveClub] = useState<string>("CLUB 1");
    const [activeComposition, setActiveComposition] =
        useState<string>("Composition 1");
    const [isActiveDropdown, setIsActiveDropdown] = useState<boolean>(false);
    const [clubId, setClubId] = useState<number | null>(null);

    // TODO: Mock data
    const clubs: Array<{ logo: any; name: string; whose: string }> = [
        { logo: <Owner />, name: "CLUB 1", whose: "owner" },
        { logo: <Owner />, name: "CLUB 2", whose: "owner" },
        { logo: <Owner />, name: "CLUB 3", whose: "owner" },
        { logo: <Manager />, name: "CLUB 4", whose: "manager" },
    ];

    // TODO: Mock data
    const compositions: string[] = [
        "Composition 1",
        "Composition 2",
        "Composition 3",
        "Composition 4",
    ];

    const handleChooseComposition = (composition: string) => {
        setActiveComposition(composition);
        setIsActiveDropdown(false);
    };

    return (
        <div className="field-controls-area__clubs">
            <span className="field-controls-area__clubs-name">
                {activeClub}
            </span>
            <div className="clubs">
                {clubs.map((club, index) => (
                    <div
                        className={`club${
                            club.name === activeClub ? "-active" : ""
                        }`}
                        key={index}
                        onClick={() => setActiveClub(club.name)}
                        onMouseLeave={() => setClubId(null)}
                        onMouseEnter={() => setClubId(index)}
                    >
                        {club.logo}
                        <span className="club__name">{club.name}</span>
                        {clubId === index && (
                            <div className="club__info">
                                {club.whose === "owner"
                                    ? `are you the ${club.whose}`
                                    : "you are the manager"}
                            </div>
                        )}
                    </div>
                ))}
            </div>
            <div className="composition">
                <div
                    className={`composition__choosed-item ${
                        isActiveDropdown ? "active-dropdown" : ""
                    }`}
                    onClick={() => setIsActiveDropdown(!isActiveDropdown)}
                >
                    {activeComposition}
                </div>
                <div
                    className={`composition__list${
                        isActiveDropdown ? "-active" : ""
                    }`}
                >
                    {compositions.map((composition, index) => (
                        <div
                            className="composition__list-item"
                            key={index}
                            onClick={() => handleChooseComposition(composition)}
                        >
                            <span>{composition}</span>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    );
};

export default Clubs;
