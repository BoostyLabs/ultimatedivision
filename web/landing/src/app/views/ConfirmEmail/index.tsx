// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useDispatch } from "react-redux";
import { useLocation } from "react-router";

import { confirmUserEmail } from "@/app/store/actions/users";
/** TODO: Rework this view after design solution */
const ConfirmEmail: React.FC = () => {
    const dispatch = useDispatch();
    const currentLocation = useLocation();
    /** TODO: IT SHOULD TO BE REWORKED. But it doesnt work with useParams */
    let pathname = currentLocation.pathname;
    const TOKEN_INDEX = 3;

    const confirm = () =>
        dispatch(confirmUserEmail(pathname.split('/')[TOKEN_INDEX]));
    ;

    return <div>
        <input
            value="Confirm Email"
            onClick={confirm}
        />
    </div>
};

export default ConfirmEmail;
