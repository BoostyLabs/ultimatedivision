import React from "react";
import { Switch, RouteProps } from "react-router-dom";
import { FootballerCard } from "../components/FootballerCardPage/FootballerCard/FootballerCard";
import { MarketPlace } from "../components/MarketPlacePage/MarketPlace/MarketPlace";

export class ComponentRoutes {
    constructor(
        public path: string,
        public component: React.FC,
        public exact: boolean,
    ) { }
};

export class RouteConfig {
    public static MarketPlace: ComponentRoutes = new ComponentRoutes(
        "/ud/marketplace",
        MarketPlace,
        true,
    );
    public static FootballerCard: ComponentRoutes = new ComponentRoutes(
        "/ud/marketplace/card",
        FootballerCard,
        true,
    );
    public static Default: ComponentRoutes = new ComponentRoutes(
        "/ud/",
        MarketPlace,
        true,
    );
    public static routes: ComponentRoutes[] = [
        RouteConfig.MarketPlace,
        RouteConfig.FootballerCard,
        RouteConfig.Default
    ]
};


type RoutesProps = { component: React.FC } & RouteProps;

const Route: React.FC<RoutesProps> = ({
    component: Component, ...children
}) => {
    return (
        <Component {...children} />
    )
};

export const Routes = () => {
    return (
        <Switch>
            {RouteConfig.routes.map((route, index) =>
                <Route
                    key={index}
                    path={route.path}
                    component={route.component}
                    exact={route.exact}
                />
            )}
        </Switch>
    );
};
