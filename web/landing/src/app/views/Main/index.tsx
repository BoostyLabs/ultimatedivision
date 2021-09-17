
// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Roadmap } from '@components/WelcomePage/Roadmap';
import { Projects } from '@components/WelcomePage/Projects';
import { Authors } from '@components/WelcomePage/Authors';
import { Footer } from '@components/WelcomePage/Footer';
import { LaunchRoadmap } from '@components/WelcomePage/LaunchRoadmap';
import { Navbar } from '@components/WelcomePage/NavBar';
import { Home } from '@components/WelcomePage/Home';

const Main: React.FC = () => {
    return (
        <>
            <Navbar />
            <Home />
            <LaunchRoadmap />
            <Roadmap />
            <Projects />
            <Authors />
            <Footer />
        </>
    );
};

export default Main;