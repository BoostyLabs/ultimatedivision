// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { ToastContainer } from 'react-toastify';

import './index.scss';

export const Notification: React.FC = () => {
    /** variables describes notification behaviour */
    const POSITION: string = 'top-right';
    const AUTO_CLOSE_TIME: number = 5000;
    const IS_NEWEST_ON_TOP: boolean = false;
    const IS_CLOSED_ON_CLICK: boolean = false;
    const IS_RIGHT_TO_LEFT_LAYOUT: boolean = false;

    return <ToastContainer
        position="top-right"
        autoClose={5000}
        hideProgressBar
        newestOnTop={false}
        closeOnClick={false}
        rtl={false}
        pauseOnFocusLoss
        pauseOnHover
    />
};
