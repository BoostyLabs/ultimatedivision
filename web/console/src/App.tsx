// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Suspense } from 'react';
import { BrowserRouter } from 'react-router-dom';
import { Routes } from '@/app/router';

import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.min.css';

import { Navbar } from '@components/common/Navbar';
import { AboutMenu } from '@components/common/AboutMenu';

/** initial App setup */
export function App() {
    return (
        <Suspense fallback={<div>Loading...</div>}>
            {/** TODO: LoadingPage */}
            <BrowserRouter basename="/">
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
                <AboutMenu />
                <Routes />
            </BrowserRouter>
        </Suspense>

    );
}

export default App;
