// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect } from "react";
import { useLocation } from "react-router";

import { UserService } from "@/user/service";
import { UserClient } from "@/api/user";

export const useConfirmEmail = () => {
    useEffect(() => {
        confirmEmail();
    }, []);

    const useQuery = () => {
        return new URLSearchParams(useLocation().search);
    };
    const query = useQuery();

    const userClient = new UserClient();
    const users = new UserService(userClient);

    async function confirmEmail() {
        return await users.confirmEmail(query.get('token'));
    };
};
