// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Roadmap } from '@components/WelcomePage/Roadmap';
import { Footer } from '@components/WelcomePage/Footer';
import { Description } from '@components/WelcomePage/Description';
import { LaunchRoadmap } from '@components/WelcomePage/LaunchRoadmap';
import { Navbar } from '@components/WelcomePage/NavBar';
import { Home } from '@components/WelcomePage/Home';
import { LaunchDate } from '@components/WelcomePage/LaunchDate';
import { Metaverse } from '@components/WelcomePage/Metaverse';
import { Authors } from '@components/WelcomePage/Authors';

const Main: React.FC = () => {
    return (
        <>
            <Navbar />
            <Home />
            <LaunchDate />
            <Metaverse />
            <Description />
            <LaunchRoadmap />
            <Roadmap />
            <Authors />
            <Footer />
        </>
    );
};

export default Main;