import { useEffect } from 'react';
import { useHistory } from 'react-router-dom';

import { vaclient } from '@/app/velas/service';
import { toast } from 'react-toastify';

import { AuthRouteConfig } from '@/app/routes';
import { InternalError } from '@/api';

import ulimatedivisionLogo from '@static/img/registerPage/ultimate.svg';

import './index.scss';

const AuthWrapper = () => {
    const history = useHistory();

    const sendAuthData = (authResult: any) => {
        try {
            // TODO: add function sending data to backend.
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
