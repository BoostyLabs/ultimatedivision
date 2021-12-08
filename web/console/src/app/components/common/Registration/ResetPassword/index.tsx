// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { SetStateAction, useState } from 'react';
import { toast } from 'react-toastify';

import { UserDataArea } from '@components/common/UserDataArea';

import ultimate from '@static/img/registerPage/ultimate.svg';
import goBack from '@static/img/registerPage/goback.svg';

import { UsersClient } from '@/api/users';
import { UsersService } from '@/users/service';
import { Validator } from '@/users/validation';

export const ResetPassword: React.FC<{ showSignInComponent: () => void }> = ({
    showSignInComponent,
}) => {
    /** controlled values for form inputs */
    const [email, setEmail] = useState('');
    const [emailError, setEmailError]
        = useState<SetStateAction<null | string>>(null);
    /** checks if values does't valid then set an error messages */
    const validateForm: () => boolean = () => {
        let isFormValid = true;

        if (!Validator.isEmail(email)) {
            setEmailError('Email is not valid');
            isFormValid = false;
        };

        return isFormValid;
    };

    const usersClient = new UsersClient();
    const usersService = new UsersService(usersClient);
    /** submit changed password */
    const handleSubmit = async(e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        if (!validateForm()) {
            return;
        };

        try {
            await usersService.sendEmailForPasswordReset(email);
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
            validate: Validator.isEmail,
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
                <div
                    onClick={showSignInComponent}
                    className="register__reset__go-back">
                    <img
                        src={goBack}
                        alt="go back"
                    />
                    <span className="register__reset__go-back__title">
                        GO BACK
                    </span>
                </div>
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
                        value="CHANGE PASSWORD"
                        type="submit"
                    />
                </form>
            </div >
        </div>
    );
};
