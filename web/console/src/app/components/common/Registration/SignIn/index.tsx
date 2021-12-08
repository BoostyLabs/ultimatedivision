// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { SetStateAction, useState } from 'react';
import { useDispatch } from 'react-redux';
import { useHistory } from 'react-router-dom';
import { toast } from 'react-toastify';

import { UserDataArea } from '@components/common/UserDataArea';

import facebook from '@static/img/registerPage/facebook_logo.svg';
import google from '@static/img/registerPage/google_logo.svg';
import ultimate from '@static/img/registerPage/ultimate.svg';

import { RouteConfig } from '@/app/routes';
import { loginUser } from '@/app/store/actions/users';
import { Validator } from '@/users/validation';

export const SignIn: React.FC<{
    showResetPasswordComponent: () => void;
    showSignUpComponent: () => void;
}> = ({ showResetPasswordComponent, showSignUpComponent }) => {
    const dispatch = useDispatch();
    const history = useHistory();

    /** controlled values for form inputs */
    const [email, setEmail] = useState('');
    const [emailError, setEmailError] = useState<SetStateAction<null | string>>(null);
    const [password, setPassword] = useState('');
    const [passwordError, setPasswordError] = useState<SetStateAction<null | string>>(null);

    const [isRemember, setIsRemember] = useState(false);

    /** TODO: rework remember me implementation  */
    const handleIsRemember = () => setIsRemember(prev => !prev);
    /** checks if values does't valid then set an error messages */
    const validateForm: () => boolean = () => {
        let isFormValid = true;

        if (!Validator.isEmail(email)) {
            setEmailError('Email is not valid');
            isFormValid = false;
        };

        if (!Validator.isPassword(password)) {
            setPasswordError('Password is not valid');
            isFormValid = false;
        };

        return isFormValid;
    };

    /** user data that will send to server */
    const handleSubmit = async(e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        if (!validateForm()) {
            return;
        };

        try {
            await dispatch(loginUser(email, password));
            history.push(RouteConfig.MarketPlace.path);
        } catch (error: any) {
            toast.error('Incorrect email or password', {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
        };
    };

    /** user datas for registration */
    const signInDatas = [
        {
            value: email,
            placeHolder: 'E-mail',
            onChange: setEmail,
            className: 'register__sign-in__sign-form__email',
            type: 'email',
            error: emailError,
            clearError: setEmailError,
            validate: Validator.isEmail,
        },
        {
            value: password,
            placeHolder: 'Password',
            onChange: setPassword,
            className: 'register__sign-in__sign-form__password',
            type: 'password',
            error: passwordError,
            clearError: setPasswordError,
            validate: Validator.isPassword,
        },
    ];

    return <div className="register">
        <div className="register__represent">
            <img
                alt="utlimate division logo"
                src={ultimate}
                className="register__represent__ultimate"
            />
        </div>
        <div className="register__sign-in">
            <h1 className="register__sign-in__title">SIGN IN</h1>
            <form
                className="register__sign-in__sign-form"
                onSubmit={handleSubmit}
            >
                {signInDatas.map((data, index) =>
                    <UserDataArea
                        key={index}
                        value={data.value}
                        placeHolder={data.placeHolder}
                        onChange={data.onChange}
                        className={data.className}
                        type={data.type}
                        error={data.error}
                        clearError={data.clearError}
                        validate={data.validate}
                    />
                )}
                <span
                    onClick={showResetPasswordComponent}
                    className="register__sign-in__sign-form__forgot-password"
                >
                    Forgot Password?
                </span>
                <div className="register__sign-in__sign-form__auth-internal">
                    <input
                        className="register__sign-in__sign-form__confirm"
                        value="SIGN IN"
                        type="submit"
                    />
                    or
                    <div className="register__sign-in__sign-form__logos">
                        <img
                            src={google}
                            alt="Google logo"
                            className="register__sign-in__sign-form__logos__google"
                        />
                        <img
                            src={facebook}
                            alt="Facebook logo"
                            className="register__sign-in__sign-form__logos__facebook"
                        />
                    </div>
                </div>
            </form>
            <div className="register__sign-in__description">
                <p className="register__sign-up__description__information">
                    Don't have an account?
                    <span
                        className="register__sign-in__description__information__sign"
                        onClick={showSignUpComponent}
                    >
                        Sign up
                    </span>
                </p>
            </div>
        </div >
    </div>;
};
