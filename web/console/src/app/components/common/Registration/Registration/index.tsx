// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';

import { ResetPassword }
    from '@/app/components/common/Registration/ResetPassword';
import { SignUp }
    from '@/app/components/common/Registration/SignUp';
import { SignIn }
    from '@/app/components/common/Registration/SignIn';

import closeButton from '@static/img/registerPage/close.svg';

import './index.scss';

export const RegistrationPopup: React.FC<{ closeRegistrationPopup: () => void }> = ({
    closeRegistrationPopup,
}) => {
    /** Checks if concret popup component is visible */
    const [isShowSignIn, setIsShowSignIn] = useState(true);
    const [isShowSignUp, setIsShowSignUp] = useState(false);
    const [isShowResetPassword, setIsShowResetPassword] = useState(false);

    /** Shows SignIn popup and closes others. */
    const showSignInComponent = () => {
        setIsShowResetPassword(!isShowResetPassword);
        setIsShowSignIn(!isShowSignIn);
    };

    /** Shows SignUp popup and closes others. */
    const showSignUpComponent = () => {
        setIsShowSignIn(!isShowSignIn);
        setIsShowSignUp(!isShowSignUp);
    };

    /** Shows ResetPasswordPopUp component and closes others. */
    const showResetPasswordComponent = () => {
        setIsShowSignIn(!isShowSignIn);
        setIsShowResetPassword(!isShowResetPassword);
    };

    return <div className="pop-up-registration">
        <div className="pop-up-registration__wrapper">
            <div
                className="pop-up-registration__wrapper__close"
                onClick={closeRegistrationPopup}
            >
                <img
                    src={closeButton}
                    alt="close button"
                />
            </div>
            {isShowSignIn && <SignIn
                showResetPasswordComponent={showResetPasswordComponent}
                showSignUpComponent={showSignUpComponent}
            />}
            {isShowSignUp && <SignUp showSignUpComponent={showSignUpComponent} />}
            {isShowResetPassword && <ResetPassword
                showSignInComponent={showSignInComponent}
            />}
        </div>
    </div>;
};
