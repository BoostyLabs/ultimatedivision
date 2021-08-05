//Copyright (C) 2021 Creditor Corp. Group.
//See LICENSE for copying information.

import { Suspense } from 'react';
import { HashRouter } from 'react-router-dom';

import { Navbar } from '@components/Navbar';
import { Routes } from '@/app/router';
import { AboutMenu } from '@components/AboutPage/AboutMenu';

/** initial App setup */
export function App() {
    return (
        <Suspense fallback={<div>Loading...</div>}>
            {/** TODO: LoadingPage */}
            <HashRouter basename="/">
                <Navbar />
                <AboutMenu />
                <Routes />
            </HashRouter>
        </Suspense>

    );
}

export default App;
