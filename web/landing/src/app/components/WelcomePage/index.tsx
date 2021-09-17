// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Heading } from '@components/WelcomePage/Heading';
import { Capabilities } from '@components/WelcomePage/Capabilities';
import { Roadmap } from '@components/WelcomePage/Roadmap';
import { Projects } from '@components/WelcomePage/Projects';
import { Authors } from '@components/WelcomePage/Authors';
import { Footer } from '@components/WelcomePage/Footer';
import { LaunchDate } from '@components/WelcomePage/LaunchDate';

const WelcomePage: React.FC = () => {
    return (
        <>
            <Heading />
            <LaunchDate />
            <Capabilities />
            <Roadmap />
            <Projects />
            <Authors />
            <Footer />
        </>
    );
};

export default WelcomePage;
