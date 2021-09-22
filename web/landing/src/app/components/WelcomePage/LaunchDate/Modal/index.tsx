// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React, { SetStateAction, useState } from 'react';

import { Validator } from '@/user/validation';

import { UserDataArea } from '@components/common/UserDataArea';

import './index.scss';

export const Modal: React.FC<{ handleModal: () => void }> = ({
    handleModal
}) => {
    const [email, setEmail] = useState('');
    const [emailError, setEmailError] =
        useState<SetStateAction<null | string>>(null);
    /** checks if value does't valid then set an error message */
    const validateForm: () => boolean = () => {
        let isValidForm = true;
        if (!Validator.email(email)) {
            setEmailError('Email is not valid');
            isValidForm = false;
        };

        return isValidForm;
    };

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();

        if (!validateForm()) {
            return;
        };
        /** closes modal window */
        handleModal();
        /** TODO: send email for gets notifications to server */
    };

    const formValue = {
        value: email,
        placeHolder: 'Email',
        handleChange: setEmail,
        className: 'launch-date-modal__notification__send',
        type: 'email',
        error: emailError,
        clearError: setEmailError,
    };

    return (
        <div className="launch-date-modal">
            <div className="launch-date-modal__window">
                <div className="launch-date-modal__header"> 
                    <a onClick={handleModal}
                        className="launch-date-modal__close"
                    >
                        <p className="launch-date-modal__close__text">
                            &#215;
                        </p>
                    </a>
                    <h1 className="launch-date-modal__description">
                        Get notified on the launch
                    </h1>
                </div>
                <div>
                    <form
                        className="launch-date-modal__notification"
                        onSubmit={handleSubmit}
                    >
                        <label className="launch-date-modal__notification__label" htmlFor="email">Email</label>
                        <UserDataArea {...formValue} />
                        <input
                            value="SEND"
                            className="launch-date-modal__notification__confirm"
                            type="submit"
                        />
                    </form>
                </div>
                <div
                    className="launch-date-modal__wrapper"
                />
            </div>
        </div>
    )
};
