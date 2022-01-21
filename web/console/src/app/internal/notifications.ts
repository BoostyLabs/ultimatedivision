// Copyright (C) 2021-2022 Amuzed GmbH finn@amuzed.io.
// This file is part of the project AMUZED.
// AMUZED can not be copied and/or distributed without the express.
// permission of Amuzed GmbH.

import { toast } from 'react-toastify';
import { InternalError, TooManyRequestError } from '@/api';

/** Code which indicates that 'eth_requestAccounts' already processing */
const RPC_ERROR_CODE = -32002;
const ALREADY_MINTED_ERROR_CODE = -32603;
const DENIED_TRANSACTION_CODE = 4001;

export function metamaskNotifications(error: any) {
    if (error.code === RPC_ERROR_CODE) {
        toast.error('Please open metamask manually!', {
            position: toast.POSITION.TOP_RIGHT,
            theme: 'colored',
        });
    } else if (error.code === DENIED_TRANSACTION_CODE) {
        toast.error('You denied transaction', {
            position: toast.POSITION.TOP_RIGHT,
            theme: 'colored',
        });
    } else if (error instanceof TooManyRequestError) {
        toast.error(error.message, {
            position: toast.POSITION.TOP_RIGHT,
            theme: 'colored',
        });
    } else if (error instanceof InternalError) {
        toast.error('Something went wrong', {
            position: toast.POSITION.TOP_RIGHT,
            theme: 'colored',
        });
    } else if (error.error.code === ALREADY_MINTED_ERROR_CODE) {
        toast.error('Token already minted', {
            position: toast.POSITION.TOP_RIGHT,
            theme: 'colored',
        });
    }
}
