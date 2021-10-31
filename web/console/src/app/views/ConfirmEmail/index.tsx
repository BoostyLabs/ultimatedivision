// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { SetStateAction, useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';

import { UserClient } from '@/api/user';
import { UserService } from '@/user/service';

import { useQueryToken } from '@/app/hooks/useQueryToken';
import { AuthRouteConfig } from '@/app/routes';

/** TODO: Rework this view after design solution */
const ConfirmEmail: React.FC = () => {
    const token = useQueryToken();
    const history = useHistory();

    const [errorMessage, setErrorMessage]
        = useState<SetStateAction<null | string>>(null);

    const userClient = new UserClient();
    const users = new UserService(userClient);

    const DELAY: number = 3000;
    /** catches error if token is not valid */
    async function checkEmailToken() {
        try {
            await users.checkEmailToken(token);

            await setTimeout(() => {
                history.push(AuthRouteConfig.SignIn.path);
            }, DELAY);
        } catch (error: any) {
            setErrorMessage('Email verification failed');
        };
    };

    useEffect(() => {
        checkEmailToken();
    }, []);

    if (errorMessage) {
        return <h1>{errorMessage}</h1>;
    };

    return <div>
        <h1>
            Your email has been successfully verified.
            You will be redirected to the sign-in page in 3 seconds.
        </h1>
    </div>;
};

export default ConfirmEmail;
