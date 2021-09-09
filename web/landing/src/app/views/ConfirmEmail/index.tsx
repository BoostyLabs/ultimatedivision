// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useDispatch } from 'react-redux';
import { useLocation } from 'react-router';

import { confirmUserEmail } from '@/app/store/actions/users';

/** TODO: Rework this view after design solution */
const ConfirmEmail: React.FC = () => {
    const dispatch = useDispatch();
    const currentLocation = useLocation();
    /** TODO: IT SHOULD TO BE REWORKED. */
    const pathname = currentLocation.pathname;
    const TOKEN_INDEX = 3;

    const confirmEmail = () =>
        dispatch(confirmUserEmail(pathname.split('/')[TOKEN_INDEX]));

    ;

    return <div>
        <input
            value="Confirm Email"
            onClick={confirmEmail}
        />
    </div>;
};

export default ConfirmEmail;
