// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect } from "react";
import { useHistory, useLocation } from "react-router";

import { UserService } from "@/user/service";
import { UserClient } from "@/api/user";
import { RouteConfig } from "../router";

export const useConfirmEmail = () => {
    useEffect(() => {
        let token = query.get("token");
        if (token) {
            localStorage.setItem('confirmEmail', token);
            history.replace({
                pathname: RouteConfig.ConfirmEmail.path,
            });
        };

        confirmEmail();
    }, []);

    const history = useHistory();
    const useQuery = () => {
        return new URLSearchParams(useLocation().search);
    };
    const query = useQuery();

    const userClient = new UserClient();
    const users = new UserService(userClient);

    async function confirmEmail() {
        return await users.confirmEmail(localStorage.getItem('confirmEmail'));
    };
};
