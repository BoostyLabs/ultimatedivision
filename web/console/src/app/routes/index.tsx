//Copyright (C) 2021 Creditor Corp. Group.
//See LICENSE for copying information.

import { lazy } from 'react';
import { Switch } from 'react-router-dom';

const FootballerCard = lazy(() => import('@components/FootballerCardPage/FootballerCard'));
const FootballField = lazy(() => import('@components/FootballFieldPage/FootballField'));
const MarketPlace = lazy(() => import('@components/MarketPlacePage/MarketPlace'));
const About = lazy(() => import('@components/AboutPage/About'));
import Summary from '@/app/components/AboutPage/WhitePaperPage/Summary';
import GameMechanics from '@/app/components/AboutPage/WhitePaperPage/GameMechanics';
import PayToEarnEconomy from '@components/AboutPage/WhitePaperPage/PayToEarnEconomy';
import Technology from '@components/AboutPage/WhitePaperPage/Technology';
import Fund from '@components/AboutPage/TokenomicsPage/Fund';
import PayToEarn from '@components/AboutPage/TokenomicsPage/PayToEarn';
import Spending from '@components/AboutPage/TokenomicsPage/Spending';
import Staking from '@components/AboutPage/TokenomicsPage/Staking';
/** Route base config implementation */
/**interfafe fot AboutPage subroutes */
export interface AboutPageSubroutes {
    whitePaper: ComponentRoutes[];
    tokenomics: ComponentRoutes[];
}
export class ComponentRoutes {
    /** data route config*/
    constructor(
        public path: string,
        public component: React.FC<any>,
        public exact: boolean,
        public subRoutes?: AboutPageSubroutes
    ) { }
};
/** Route config implementation */
export class RouteConfig {
    public static MarketPlace: ComponentRoutes = new ComponentRoutes(
        '/test/marketplace',
        MarketPlace,
        true,
    );
    public static FootballerCard: ComponentRoutes = new ComponentRoutes(
        '/test/marketplace/card',
        FootballerCard,
        true,
    );
    public static FootballField: ComponentRoutes = new ComponentRoutes(
        '/test/field',
        FootballField,
        true,
    );
    public static MyCards: ComponentRoutes = new ComponentRoutes(
        '/test/marketplace/club',
        MarketPlace,
        true,
    );
    public static About: ComponentRoutes = new ComponentRoutes(
        '/test/about/',
        About,
        false,
        {
            whitePaper: [
                new ComponentRoutes(
                    '/test/about/whitepaper',
                    Summary,
                    true
                ),
                new ComponentRoutes(
                    '/test/about/whitepaper/game-mechanicks',
                    GameMechanics,
                    true
                ),
                new ComponentRoutes(
                    '/test/about/whitepaper/pay-to-earn-and-economy',
                    PayToEarnEconomy,
                    true
                ),
                new ComponentRoutes(
                    '/test/about/whitepaper/technology',
                    Technology,
                    true
                ),
            ],
            tokenomics: [
                new ComponentRoutes(
                    '/test/about/tokenomics/udt-spending',
                    Spending,
                    true
                ),
                new ComponentRoutes(
                    '/test/about/tokenomics/pay-to-earn',
                    PayToEarn,
                    true
                ),
                new ComponentRoutes(
                    '/test/about/tokenomics/staking',
                    Staking,
                    true
                ),
                new ComponentRoutes(
                    '/test/about/tokenomics/ud-dao-fund',
                    Fund,
                    true
                ),
            ],
        }
    );
    public static Default: ComponentRoutes = new ComponentRoutes(
        '/test/',
        MarketPlace,
        true,
    );
    public static routes: ComponentRoutes[] = [
        RouteConfig.MarketPlace,
        RouteConfig.FootballerCard,
        RouteConfig.FootballField,
        RouteConfig.MyCards,
        RouteConfig.About,
        RouteConfig.Default,
    ];
};
export const Route: React.FC<ComponentRoutes> = ({
    component: Component,
    ...children
}) => <Component {...children} />;
export const Routes = () =>
    <Switch>
        {RouteConfig.routes.map((route, index) =>
            <Route
                key={index}
                path={route.path}
                component={route.component}
                exact={route.exact}
                children={route.subRoutes}
            />,
        )}
    </Switch>;