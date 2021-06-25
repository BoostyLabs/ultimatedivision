import React from 'react';
import { Route, Switch } from 'react-router-dom';

import { MarketPlace } from './components/MarketPlacePage/MarketPlace/MarketPlace';

import './App.scss';

import { FootballerCard } from
    './components/FootballerCardPage/FootballerCard/FootballerCard';

export const App: React.FC = () => {
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
};

