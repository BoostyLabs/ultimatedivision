// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Suspense } from 'react';
import { BrowserRouter } from 'react-router-dom';

import { AboutMenu } from '@components/common/AboutMenu';
import { Notification } from '@components/common/Notification';

import { Routes } from '@/app/routes';

import 'react-toastify/dist/ReactToastify.min.css';

/** initial App setup */
export function App() {
    return (
        <Suspense fallback={<div>Loading...</div>}>
            {/** TODO: LoadingPage */}
            <BrowserRouter basename="/">
                <Notification />
                <AboutMenu />
                <Routes />
            </BrowserRouter>
        </Suspense>
    );
};

export default App;
