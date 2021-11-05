// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { ToastContainer, ToastPosition } from 'react-toastify';

import './index.scss';

/** Custom component for notifications */
export const Notification: React.FC = () => {
    /** variables describes notification behaviour */
    const POSITION: ToastPosition = 'top-right';
    const AUTO_CLOSE_TIME: number = 5000;
    const IS_NEWEST_ON_TOP: boolean = false;
    const IS_CLOSED_ON_CLICK: boolean = false;
    const IS_RIGHT_TO_LEFT_LAYOUT: boolean = false;

    return <ToastContainer
        position={POSITION}
        autoClose={AUTO_CLOSE_TIME}
        hideProgressBar
        newestOnTop={IS_NEWEST_ON_TOP}
        closeOnClick={IS_CLOSED_ON_CLICK}
        rtl={IS_RIGHT_TO_LEFT_LAYOUT}
        pauseOnFocusLoss
        pauseOnHover
    />
};
