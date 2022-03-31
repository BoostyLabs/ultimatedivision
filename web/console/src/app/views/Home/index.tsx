// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useLocation } from 'react-router-dom';

import { FootballGame } from '@components/home/FootballGame';
import { Footer } from '@components/home/Footer';
import { Partnerships } from '@components/home/Partnerships';
import { Roadmap } from '@components/home/Roadmap';
import { OurTeam } from '@/app/components/home/OurTeam';
import Navbar from '@components/home/HomeNavbar';

import './index.scss';

const Home: React.FC = () => {
    /** Current path from hook */
    const location = useLocation();
    const currentPath = location.pathname;

    return (
        <>
            {currentPath === '/' && <Navbar />}
            <FootballGame />
            <Roadmap />
            <Partnerships />
            <OurTeam />
            <Footer />
        </>
    );
};

export default Home;
