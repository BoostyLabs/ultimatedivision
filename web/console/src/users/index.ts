// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/** User describes user domain entity. */
export class User {
    /** User domain entity contains email, password, nickName, firstName, lastName. */
    public constructor(
        public email: string,
        public password: string,
        public nickName: string,
        public firstName: string,
        public lastName: string,
    ) { };
};

/** Defines fields for login */
export class LoginFields {
    constructor (
        public email: string = '',
        public password: string = ''
    ) {}
}

export class ChangePasswordFields {
    constructor (
        public password: string = '',
        public newPassword: string = '',
    ) {}
}
