import { useEffect } from "react";
import { useHistory } from "react-router-dom";
import { toast } from "react-toastify";

import { AuthRouteConfig, RouteConfig } from "@/app/routes";
import { InternalError } from "@/api";
import { UsersClient } from "@/api/users";
import { UsersService } from "@/users/service";

import ulimatedivisionLogo from "@static/img/registerPage/ultimate.svg";

// @ts-ignore
import { VAClient } from "@velas/account-client";
// @ts-ignore
import StorageHandler from "../../../../velas/storageHandler";
// @ts-ignore
import KeyStorageHandler from "../../../../velas/keyStorageHandler";

import "./index.scss";

const AuthWrapper = () => {
    const history = useHistory();

    const usersClient = new UsersClient();
    const usersService = new UsersService(usersClient);

    const sendAuthData = async (authResult: any) => {
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
                theme: "colored",
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
                theme: "colored",
            });
        }
    };

    /** generates vaclient with the help of creds  */
    const vaclientService = async () => {
        try {
            const vaclientCreds = await usersService.velasVaclientCreds();

            const vaclient = new VAClient({
                mode: "redirect",
                clientID: vaclientCreds.clientId,
                redirectUri: vaclientCreds.redirectUri,
                StorageHandler,
                KeyStorageHandler,
                accountProviderHost: vaclientCreds.accountProviderHost,
                networkApiHost: vaclientCreds.networkApiHost,
                transactionsSponsorApiHost: vaclientCreds.transactionsSponsorApiHost,
                transactionsSponsorPubKey: vaclientCreds.transactionsSponsorPubKey,
            });

            return vaclient;
        } catch (e) {
            toast.error(`${e}`, {
                position: toast.POSITION.TOP_RIGHT,
                theme: "colored",
            });
        }

        return null;
    };

    const authorization = async () => {
        try {
            const vaclient = await vaclientService();
            vaclient.handleRedirectCallback(processAuthResult);
        } catch (e) {
            toast.error(`${e}`, {
                position: toast.POSITION.TOP_RIGHT,
                theme: "colored",
            });
        }
    };

    useEffect(() => {
        authorization();
    }, []);

    return (
        <div className="auth-wrapper">
            <img src={ulimatedivisionLogo} alt="ultimatedivision-logo" />
        </div>
    );
};
export default AuthWrapper;
