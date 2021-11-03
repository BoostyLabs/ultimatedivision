// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { HttpClient } from '@/private/http/client';

/**
 * ErrorUnauthorized is a custom error type for performing unauthorized operations.
 */
export class UnauthorizedError extends Error {
    /** Error message while unautorized */
    public constructor(message = 'authorization required') {
        super(message);
    }
}

/**
 * NotFoundError is a custom error type for performing bad request.
 */
export class NotFoundError extends Error {
    /** Error message while bad request */
    public constructor(message = 'bad request') {
        super(message);
    }
}

/**
 * InternalError is a custom error type for internal server error.
 */
export class InternalError extends Error {
    /** Error message for internal server error */
    public constructor(message = 'internal server error') {
        super(message);
    }
}
const UNAUTORISED_ERROR = 401;
const NOT_FOUND_ERROR = 404;
const INTERNAL_ERROR = 500;

/**
 * APIClient is base client that holds http client and error handler.
 */
export class APIClient {
    protected readonly http: HttpClient = new HttpClient();

    /**
     * handles error due to response code.
     * @param response - response from server.
     *
     * @throws {@link NotFoundError}
     * This exception is thrown if the input is not a valid ISBN number.
     *
     * @throws {@link UnauthorizedError}
     * Thrown if the ISBN number is valid, but no such book exists in the catalog.
     *
     * @throws {@link InternalError}
     * Thrown if the ISBN number is valid, but no such book exists in the catalog.
     *
     * @private
     */
    /* eslint-disable */
    protected async handleError(response: Response): Promise<void> {

        switch (response.status) {
        case UNAUTORISED_ERROR: throw new UnauthorizedError();
        case NOT_FOUND_ERROR: throw new NotFoundError();
        case INTERNAL_ERROR:
        default:
            throw new InternalError();
        }
    }
}
