// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { SetStateAction, useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';
import { useDispatch } from 'react-redux';
import { toast } from 'react-toastify';

import { UserClient } from '@/api/user';

import { useQueryToken } from '@/app/hooks/useQueryToken';
import { AuthRouteConfig } from '@/app/routes';
import { recoverUserPassword } from '@/app/store/actions/users';

import { UserService } from '@/user/service';
import { Validator } from '@/user/validation';

import { UserDataArea } from '@components/common/UserDataArea';

import ultimate from '@static/img/registerPage/ultimate.svg';

import './index.scss';

const ResetPassword: React.FC = () => {
    const dispatch = useDispatch();
    const history = useHistory();
    const token = useQueryToken();

    const [errorMessage, setErrorMessage]
        = useState<SetStateAction<null | string>>(null);

    const userClient = new UserClient();
    const users = new UserService(userClient);

    /** catches error if token is not valid */
    async function checkRecoverToken() {
        try {
            await users.checkRecoverToken(token);
        } catch (error: any) {
            /** TODO: handles error */
            setErrorMessage('Cannot get access');
        };
    };
    useEffect(() => {
        checkRecoverToken();
    }, []);

    /** controlled values for form inputs */
    const [password, setPassword] = useState('');
    const [passwordError, setPasswordError]
        = useState<SetStateAction<null | string>>(null);
    const [confirmedPassword, setConfirmedPassword] = useState('');
    const [confirmedPasswordError, setConfirmedPasswordError]
        = useState<SetStateAction<null | string>>(null);
    /** checks if values does't valid then set an error messages */
    const validateForm: () => boolean = () => {
        let isFormValid = true;

        if (!Validator.isPassword(password)) {
            setPasswordError('Password is not valid');
            isFormValid = false;
        };

        if (!Validator.isPassword(confirmedPassword)) {
            setConfirmedPasswordError('Confirmed password is not valid');
            isFormValid = false;
        };

        if (password !== confirmedPassword) {
            setConfirmedPasswordError('Passwords does not match, please try again');
            isFormValid = false;
        }

        return isFormValid;
    };

    const DELAY: number = 3000;
    /** sign in user data */
    const handleSubmit = async(e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        if (!validateForm()) {
            return;
        };

        try {
            await dispatch(recoverUserPassword(password));
            toast.success('Successfully! You will be redirected after 3 seconds', {
                position: toast.POSITION.TOP_RIGHT,
            });
            await setTimeout(() => {
                history.push(AuthRouteConfig.SignIn.path);
            }, DELAY);
        } catch (error: any) {
            toast.error('Please, make sure your password is correct.', {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
        }
    };
    /** user datas for recover password */
    const passwords = [
        {
            value: password,
            placeHolder: 'Enter a new password',
            onChange: setPassword,
            className: 'register__recover__sign-form__password',
            type: 'password',
            error: passwordError,
            clearError: setPasswordError,
            validate: Validator.isPassword,
        },
        {
            value: confirmedPassword,
            placeHolder: 'Enter a new password again',
            onChange: setConfirmedPassword,
            className: 'register__recover__sign-form__password',
            type: 'password',
            error: confirmedPasswordError,
            clearError: setConfirmedPasswordError,
            validate: Validator.isPassword,
        },
    ];

    if (errorMessage) {
        return <h1>{errorMessage}</h1>;
    };

    return (
        <div className="register">
            <div className="register__represent-reset">
                <img
                    src={ultimate}
                    alt="utlimate division logo"
                    className="register__represent-reset__ultimate"
                />
            </div>
            <div className="register__recover">
                <h1 className="register__recover__title">Recover password</h1>
                <form
                    className="register__recover__sign-form"
                    onSubmit={handleSubmit}
                >
                    {passwords.map((password, index) => <UserDataArea
                        key={index}
                        value={password.value}
                        placeHolder={password.placeHolder}
                        onChange={password.onChange}
                        className={password.className}
                        type={password.type}
                        error={password.error}
                        clearError={password.clearError}
                        validate={password.validate}
                    />)}
                    <input
                        className="register__recover__sign-form__confirm"
                        value="RECOVER YOUR PASSWORD"
                        type="submit"
                    />
                </form>
            </div >
        </div>
    );
};

export default ResetPassword;
