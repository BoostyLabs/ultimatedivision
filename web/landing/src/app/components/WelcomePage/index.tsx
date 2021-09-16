// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Roadmap } from '@components/WelcomePage/Roadmap';
import { Projects } from '@components/WelcomePage/Projects';
import { Authors } from '@components/WelcomePage/Authors';
import { Footer } from '@components/WelcomePage/Footer';
import { Description } from '@/app/views/Description';

const WelcomePage: React.FC = () => {
    return (
        <>
            <Description />
            <Roadmap />
            <Projects />
            <Authors />
            <Footer />
        </>
    );
};

export default WelcomePage;
