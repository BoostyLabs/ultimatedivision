// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { SetStateAction, useState } from 'react';
import { Link } from 'react-router-dom';
import { toast } from 'react-toastify';

import { AuthRouteConfig } from '@/app/routes';

import { UserClient } from '@/api/user';
import { UserService } from '@/user/service';
import { Validator } from '@/user/validation';

import { UserDataArea } from '@components/common/UserDataArea';

import ultimate from '@static/img/registerPage/ultimate.svg';
import goBack from '@static/img/registerPage/goback.svg';

import './index.scss';

const ChangePassword: React.FC = () => {
    /** controlled values for form inputs */
    const [email, setEmail] = useState('');
    const [emailError, setEmailError]
        = useState<SetStateAction<null | string>>(null);
    /** checks if values does't valid then set an error messages */
    const validateForm: () => boolean = () => {
        let isValidForm = true;

        if (!Validator.email(email)) {
            setEmailError('Email is not valid');
            isValidForm = false;
        };

        return isValidForm;
    };

    const userClient = new UserClient();
    const users = new UserService(userClient);
    /** sign in user data */
    const handleSubmit = async(e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        if (!validateForm()) {
            return;
        };

        try {
            await users.sendEmailForResetPassword(email);
            toast.success('Successfully! Please, check your mail box.', {
                position: toast.POSITION.TOP_RIGHT,
            });
        } catch (error: any) {
            toast.error('Please, make sure your email is correct.', {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
        };
    };
    /** user datas for registration */
    const resetPasswordDatas = [
        {
            value: email,
            placeHolder: 'Enter your email',
            onChange: setEmail,
            className: 'register__reset__sign-form__email',
            type: 'text',
            error: emailError,
            clearError: setEmailError,
            validate: Validator.email,
        },
    ];

    return (
        <div className="register">
            <div className="register__represent-reset">
                <img
                    src={ultimate}
                    alt="utlimate division logo"
                    className="register__represent-reset__ultimate"
                />
            </div>
            <div className="register__reset">
                <Link
                    className="register__reset__go-back"
                    to={AuthRouteConfig.SignIn.path}>
                    <img
                        alt="go back"
                        src={goBack}
                    />
                    <span className="register__reset__go-back__title">
                        GO BACK
                    </span>
                </Link>
                <h1 className="register__reset__title">Change your password</h1>
                <form
                    className="register__reset__sign-form"
                    onSubmit={handleSubmit}
                >
                    {resetPasswordDatas.map((data, index) => <UserDataArea
                        key={index}
                        value={data.value}
                        placeHolder={data.placeHolder}
                        onChange={data.onChange}
                        className={data.className}
                        type={data.type}
                        error={data.error}
                        clearError={data.clearError}
                        validate={data.validate}
                    />)}
                    <input
                        className="register__reset__sign-form__confirm"
                        value="CONFIRM"
                        type="submit"
                    />
                </form>
            </div >
        </div>
    );
};

export default ChangePassword;
