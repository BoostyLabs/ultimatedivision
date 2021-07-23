/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
*/

import { BrowserRouter } from 'react-router-dom';

import { Navbar } from '@Navbar';
import { Routes } from '@Routes/index';

/** initial App setup */
export function App() {
    return (
        <BrowserRouter>
            <Navbar />
            <Routes />
        </BrowserRouter>
    );
}

export default App;
