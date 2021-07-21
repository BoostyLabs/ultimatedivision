/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
*/

import { BrowserRouter } from 'react-router-dom';

import { Navbar }
    from './app/components/Navbar';
import { Routes } from './app/routes/index';

import './app/static/styles/App.scss';
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
