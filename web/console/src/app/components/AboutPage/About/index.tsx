// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.
import { Switch } from 'react-router-dom';
import { AboutPageSubroutes, ComponentRoutes, Route } from '@/app/routes';
import { AboutMenu } from '../AboutMenu';
import './index.scss';
const About: React.FC<{ children: AboutPageSubroutes }> = ({ children }) => {
    const { whitePaper, tokenomics } = children;
    const routes = whitePaper.concat(tokenomics);
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
    );
};
export default About;