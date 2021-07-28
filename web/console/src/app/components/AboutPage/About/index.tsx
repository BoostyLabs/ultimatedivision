import React from 'react';

import { Switch } from 'react-router-dom';

import { AboutMenu } from '../AboutMenu';
import { ComponentRoutes, Route } from '@/app/routes';

import './index.scss';

const About:React.FC<{routes: ComponentRoutes[]}> = ({routes}) => {
    return (
        <div className="about">
            <AboutMenu />
            <div className="about__wrapper">
                <Switch>
                    {routes.map((item: ComponentRoutes, index: number) =>
                        <Route
                            key={index}
                            path={item.path}
                            component={item.component}
                            exact={item.exact}
                        />
                    )}
                </Switch>
            </div>
        </div>
    )
}

export default About
