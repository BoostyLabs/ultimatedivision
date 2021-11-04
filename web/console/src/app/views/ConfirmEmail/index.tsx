// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { SetStateAction, useEffect, useState } from 'react';

import { UserClient } from '@/api/user';
import { UserService } from '@/user/service';

import { useQueryToken } from '@/app/hooks/useQueryToken';
import { AuthRouteConfig } from '@/app/routes';

import './index.scss';

/** TODO: Rework this view after design solution */
const ConfirmEmail: React.FC = () => {
    const token = useQueryToken();

    const [confirmEmailMessage, setConfirmEmailMessage]
        = useState<SetStateAction<null | string>>(null);

    const userClient = new UserClient();
    const users = new UserService(userClient);

    const DELAY: number = 3000;
    /** catches error if token is not valid */
    async function checkEmailToken() {
        try {
            await users.checkEmailToken(token);
            setConfirmEmailMessage(`Your email has been successfully verified.
            You will be redirected to the sign-in page in 3 seconds.`);
            await setTimeout(() => {
                location.pathname = AuthRouteConfig.SignIn.path;
            }, DELAY);
        } catch (error: any) {
            setConfirmEmailMessage('Email verification failed');
        };
    };

    useEffect(() => {
        checkEmailToken();
    }, []);

    return <div className="confirm-email">
        <h1 className="confirm-email__title">
            {confirmEmailMessage}
        </h1>
    </div>;
};

export default ConfirmEmail;
