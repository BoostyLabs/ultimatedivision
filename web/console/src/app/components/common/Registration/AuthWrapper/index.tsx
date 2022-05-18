import { useEffect } from 'react';
import { useHistory } from 'react-router-dom';

import { vaclient } from '@/app/velas/service';
import { toast } from 'react-toastify';

import { AuthRouteConfig, RouteConfig } from '@/app/routes';
import { InternalError } from '@/api';

import ulimatedivisionLogo from '@static/img/registerPage/ultimate.svg';

import './index.scss';
import { UsersClient } from '@/api/users';
import { UsersService } from '@/users/service';

const AuthWrapper = () => {
    const history = useHistory();
    const usersClient = new UsersClient();
    const usersService = new UsersService(usersClient);

    const sendAuthData = async(authResult: any) => {
        try {
            await usersService.velasRegister(authResult);
            const nonce = await usersService.velasNonce(authResult.access_token_payload.sub);
            await usersService.velasLogin(nonce, authResult);
            history.push(RouteConfig.MarketPlace.path);
            window.location.reload();
        } catch (error) {
            if (error instanceof InternalError) {
                history.push(AuthRouteConfig.SignIn.path);
            }

            toast.error(`${error}`, {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
        }
    };

    const processAuthResult = (e: any, authResult: any) => {
        if (authResult && authResult.access_token_payload) {
            window.history.replaceState({}, document.title, window.location.pathname);

            sendAuthData(authResult);
        } else if (e) {
            window.history.replaceState({}, document.title, window.location.pathname);

            toast.error(`${e.description}`, {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
        }
    };

    const authHandleRedirect = () => {
        vaclient.handleRedirectCallback(processAuthResult);
    };
    useEffect(authHandleRedirect, []);

    return (
        <div className="auth-wrapper">
            <img src={ulimatedivisionLogo} alt="ultimatedivision-logo" />
        </div>
    );
};
export default AuthWrapper;
