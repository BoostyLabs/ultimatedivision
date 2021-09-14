// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect } from "react";
import { useHistory, useLocation } from "react-router"

import { UserClient } from "@/api/user";
import { UserService } from "@/user/service";
import { RouteConfig } from "../router";

export const useAuth = () => {
    useEffect(() => {
        let token = query.get('token');
        if (token) {
            localStorage.setItem('recoverPassword', token);
            history.replace({
                pathname: RouteConfig.RecoverPassword.path,
            });
        };

        checkToken();
    }, []);

    const history = useHistory();
    const useQuery = () => {
        return new URLSearchParams(useLocation().search);
    };
    const query = useQuery();

    const userClient = new UserClient();
    const users = new UserService(userClient);

    async function checkToken() {
        return await users.checkToken(localStorage.getItem('recoverPassword'));
    };
};
