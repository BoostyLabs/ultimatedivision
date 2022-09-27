// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useLocation } from 'react-router-dom';

import { FootballGame } from '@components/home/FootballGame';
import { Roadmap } from '@components/home/Roadmap';
import { GameInfo } from '@/app/components/home/GameInfo';
import { Partnerships } from '@/app/components/home/Partnerships';
import Navbar from '@components/home/HomeNavbar';

import banner from '@static/img/gameLanding/banner.png';

import './index.scss';

const Home: React.FC = () => {
    /** Current path from hook */
    const location = useLocation();
    const currentPath = location.pathname;

    return (
        <>
            {currentPath === '/' && <Navbar />}
            <FootballGame />
            <GameInfo/>
            <Roadmap />
            <Partnerships />
            <div className="home__banner">
                <img src={banner} className="home__banner__image" alt="banner" />
            </div>
        </>
    );
};

export default Home;
