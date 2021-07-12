import React from 'react';
import { BrowserRouter } from 'react-router-dom';

import { UltimateDivisionNavbar }
    from './components/Navbar/Navbar';
import { Routes } from './routes/index'

import './App.scss';


export function App() {
    return (
        <BrowserRouter>
            <UltimateDivisionNavbar />
            <Routes />
        </BrowserRouter>
    );
}

export default App;
