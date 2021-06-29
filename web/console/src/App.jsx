/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
*/

import React from 'react';
import { Route, Switch } from 'react-router-dom';

import { MarketPlace }
    from './components/MarketPlacePage/MarketPlace/MarketPlace';

import './App.scss';

import { FootballerCard } from
    './components/FootballerCardPage/FootballerCard/FootballerCard';

export function App() {
    return (
        <>
            <Switch>
                <Route exact path="/ud/marketplace/">
                    <MarketPlace />
                </Route>
                <Route exact path="/ud/marketplace/card">
                    <FootballerCard />
                </Route>
            </Switch>
        </>
    );
}

