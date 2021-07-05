import React from "react";
import { Switch, RouteProps, BrowserRouter } from "react-router-dom";
import { FootballerCard } from "../components/FootballerCardPage/FootballerCard/FootballerCard";
import { MarketPlace } from "../components/MarketPlacePage/MarketPlace/MarketPlace";

export class ComponentRoutes {
    constructor(
        public path: string,
        public component: React.FC,
        public exact: boolean,
        // public subroutes: ComponentSubRouters
    ) { }
};

// class ComponentSubRouters {
//     constructor(
//         public path: string,
//         public component: React.FC,
//         public exact: boolean,
//     ) { }
// };

export class RouteConfig {
    public static MarketPlace: ComponentRoutes = new ComponentRoutes(
        "/ud/marketplace",
        MarketPlace,
        false,
        // new ComponentSubRouters(
        //     "/ud/marketplace",
        //     MarketPlace,
        //     true
        // )
    );
    public static Club: ComponentRoutes = new ComponentRoutes(
        "/ud/club",
        FootballerCard,
        false,
        // new ComponentSubRouters(
        //     "/ud/club",
        //     FootballerCard,
        //     true
        // )
    );
    public static Default: ComponentRoutes = new ComponentRoutes(
        "/",
        MarketPlace,
        true,
        // new ComponentSubRouters(
        //     "/ud/marketplace/card",
        //     FootballerCard,
        //     true
        // )
    );
    public static routes: ComponentRoutes[] = [
        RouteConfig.MarketPlace,
        RouteConfig.Club,
        RouteConfig.Default
    ]
};


type RoutesProps = { component: React.FC } & RouteProps;
type RoutesProps1 = { component: React.FC } & RouteProps;


const Route: React.FC<RoutesProps> = ({
    component: Component, ...children
}) => {
    return (

        <Component {...children}>
            {/* <SubRoute component={subroutes.component} exact={subroutes.exact} path={subroutes.path} /> */}
        </Component>
    )
};

const SubRoute: React.FC<RoutesProps1> = ({
    component: Component, ...children
}) => {
    console.log(children);
    return <Component {...children}></Component>
}

export const Routes = () => {
    return (
        <BrowserRouter>
                <Switch>
                    {RouteConfig.routes.map((route, index) =>
                        <Route
                            key={index}
                            path={route.path}
                            component={route.component}
                            exact={route.exact}
                        // subroutes={route.subroutes}
                        />
                    )}
                </Switch>
        </BrowserRouter>
    );
};
