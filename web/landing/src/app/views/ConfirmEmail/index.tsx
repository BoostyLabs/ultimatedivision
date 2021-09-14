// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useConfirmEmail } from '@/app/hooks/confirmEmail';

/** TODO: Rework this view after design solution */
const ConfirmEmail: React.FC = () => {
    useConfirmEmail();
    return <div>
        <h1>Your email has been successfully verified</h1>
    </div>;
};

export default ConfirmEmail;
