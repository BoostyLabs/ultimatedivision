import React from 'react';

import { UltimateDivisionNavbar }
    from './components/UltimateDivisionNavbar/UltimateDivisionNavbar';
import { Routes } from './routes/index'

import './App.scss';
import { BrowserRouter } from 'react-router-dom';

export function App() {
    return (
        <BrowserRouter>
            <UltimateDivisionNavbar />
            <Routes />
        </BrowserRouter>
    );
}

export default App;
