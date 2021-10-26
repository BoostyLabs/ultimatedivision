import { ToastContainer, toast } from 'react-toastify';
// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Roadmap } from '@components/Roadmap';
import { Projects } from '@components/Projects';
import { Footer } from '@components/Footer';
import { LaunchRoadmap } from '@components/LaunchRoadmap';
import { Navbar } from '@components/NavBar';
import { Home } from '@components/Home';
import { LaunchDate } from '@components/LaunchDate';
import { Description } from '@components/Description';
import { Metaverse } from '@components/Metaverse';
import { Authors } from '@components/Authors';

export const App = () => (
    <main className="main">
        <ToastContainer
            position="top-right"
            autoClose={5000}
            hideProgressBar
            newestOnTop={false}
            closeOnClick={false}
            rtl={false}
            pauseOnFocusLoss
            pauseOnHover
        />
        <Navbar />
        <Home />
        <LaunchDate />
        <Metaverse />
        <Description />
        <LaunchRoadmap />
        <Roadmap />
        <Projects />
        <Authors />
        <Footer />
    </main>
);
