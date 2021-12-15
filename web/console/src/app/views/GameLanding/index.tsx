// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Authors } from "@components/gameLanding/Authors";
import { FootballGame } from "@components/gameLanding/FootballGame";
import { Footer } from "@components/gameLanding/Footer";
import { Navbar } from "@components/gameLanding/Navbar";
import { Projects } from "@components/gameLanding/Projects";
import { Roadmap } from "@components/gameLanding/Roadmap";

import "./index.scss";

const GameLanding: React.FC = () => {
    return (
        <>
            <Navbar />
            <FootballGame />
            <Roadmap />
            <Projects />
            <Authors />
            <Footer />
        </>
    );
};

export default GameLanding;
