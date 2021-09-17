/******/ (() => { // webpackBootstrap
/******/ 	

    /******/ 	const __webpack_modules__ = ({

        /***/ 4800:
        /***/ ((module, __webpack_exports__, __webpack_require__) => {

            /* harmony export */ __webpack_require__.d(__webpack_exports__, {
                /* harmony export */ 'Z': () => (__WEBPACK_DEFAULT_EXPORT__)
                /* harmony export */ });

            /* harmony import */ const _node_modules_css_loader_dist_runtime_cssWithMappingToString_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(4015);
            /* harmony import */ const _node_modules_css_loader_dist_runtime_cssWithMappingToString_js__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(_node_modules_css_loader_dist_runtime_cssWithMappingToString_js__WEBPACK_IMPORTED_MODULE_0__);
            /* harmony import */ const _node_modules_css_loader_dist_runtime_api_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(3645);
            /* harmony import */ const _node_modules_css_loader_dist_runtime_api_js__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(_node_modules_css_loader_dist_runtime_api_js__WEBPACK_IMPORTED_MODULE_1__);
            /* harmony import */ const _node_modules_css_loader_dist_runtime_getUrl_js__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(1667);
            /* harmony import */ const _node_modules_css_loader_dist_runtime_getUrl_js__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(_node_modules_css_loader_dist_runtime_getUrl_js__WEBPACK_IMPORTED_MODULE_2__);
            // Imports

            const ___CSS_LOADER_URL_IMPORT_0___ = new URL(/* asset import */ __webpack_require__(8395), __webpack_require__.b);
            const ___CSS_LOADER_EXPORT___ = _node_modules_css_loader_dist_runtime_api_js__WEBPACK_IMPORTED_MODULE_1___default()((_node_modules_css_loader_dist_runtime_cssWithMappingToString_js__WEBPACK_IMPORTED_MODULE_0___default()));
            const ___CSS_LOADER_URL_REPLACEMENT_0___ = _node_modules_css_loader_dist_runtime_getUrl_js__WEBPACK_IMPORTED_MODULE_2___default()(___CSS_LOADER_URL_IMPORT_0___);

            // Module
            ___CSS_LOADER_EXPORT___.push([module.id, '*{box-sizing:border-box}body,section,div,span,h1,h2,h3,h4,h5,h6,p,ul,li{margin:0;padding:0}*{box-sizing:border-box}body,section,div,span,h1,h2,h3,h4,h5,h6,p,ul,li{margin:0;padding:0}@font-face{font-family:"BebasNeue";src:url(' + ___CSS_LOADER_URL_REPLACEMENT_0___ + ');font-style:normal;font-weight:lighter}html{font-family:"BebasNeue",sans-serif}', '',{
                'version':3,'sources':['webpack://./src/app/static/styles/_reset.scss','webpack://./src/index.scss'],'names':[],'mappings':'AAAA,EACI,qBAAA,CAGJ,gDAaI,QAAA,CACA,SAAA,CAlBJ,EACI,qBAAA,CAGJ,gDAaI,QAAA,CACA,SAAA,CCbJ,WACI,uBAAA,CACA,2CAAA,CACA,iBAAA,CACA,mBAAA,CAGJ,KACI,kCAAA','sourcesContent':['* {\n    box-sizing: border-box;\n}\n\nbody,\nsection,\ndiv,\nspan,\nh1,\nh2,\nh3,\nh4,\nh5,\nh6,\np,\nul,\nli {\n    margin: 0;\n    padding: 0;\n}\n','@import "@static/styles/reset";\n@import "@static/styles/variables";\n@import "@static/styles/extends";\n@import "@static/styles/mixins";\n\n@font-face {\n    font-family: "BebasNeue";\n    src: url("@static/fonts/BebasNeue-Regular.ttf");\n    font-style: normal;\n    font-weight: lighter;\n}\n\nhtml {\n    font-family: "BebasNeue", sans-serif;\n}\n'],'sourceRoot':''
            }]);

            // Exports
            /* harmony default export */ const __WEBPACK_DEFAULT_EXPORT__ = (___CSS_LOADER_EXPORT___);

            /***/ }),

        /***/ 9443:
        /***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

            // EXPORTS
            __webpack_require__.d(__webpack_exports__, {
                'K': () => (/* binding */ UserClient)
            });

            // EXTERNAL MODULE: ./node_modules/tslib/tslib.es6.js
            const tslib_es6 = __webpack_require__(655);

            ;// CONCATENATED MODULE: ./src/private/http/client.ts
            // Copyright (C) 2021 Creditor Corp. Group.
            // See LICENSE for copying information.

            const HttpClient = /** @class */ (function () {
                function HttpClient() {
                }

                /* do sends an HTTP request and returns
    * an HTTP response as configured on the client. */
                HttpClient.prototype.do = function (method, path, body) {
                    return (0,tslib_es6/* __awaiter */.mG)(this, void 0, void 0, function () {
                        let request;

                        return (0,tslib_es6/* __generator */.Jh)(this, function (_a) {
                            switch (_a.label) {
                            case 0:
                                request = {
                                    method: method,
                                    body: body,
                                };

                                request.headers = {
                                    'Accept': 'application/json',
                                    'Content-type': 'application/json',
                                };

                                return [4 /*yield*/, fetch(path, request)];
                            case 1: return [2 /*return*/, _a.sent()];
                            }
                        });
                    });
                };
                ;

                HttpClient.prototype.post = function (path, body) {
                    return (0,tslib_es6/* __awaiter */.mG)(this, void 0, void 0, function () {
                        return (0,tslib_es6/* __generator */.Jh)(this, function (_a) {
                            switch (_a.label) {
                            case 0: return [4 /*yield*/, this.do('POST', path, body)];
                            case 1: return [2 /*return*/, _a.sent()];
                            }
                        });
                    });
                };
                ;

                HttpClient.prototype.get = function (path) {
                    return (0,tslib_es6/* __awaiter */.mG)(this, void 0, void 0, function () {
                        return (0,tslib_es6/* __generator */.Jh)(this, function (_a) {
                            switch (_a.label) {
                            case 0: return [4 /*yield*/, this.do('GET', path, null)];
                            case 1: return [2 /*return*/, _a.sent()];
                            }
                        });
                    });
                };
                ;

                HttpClient.prototype.put = function (path, body) {
                    return (0,tslib_es6/* __awaiter */.mG)(this, void 0, void 0, function () {
                        return (0,tslib_es6/* __generator */.Jh)(this, function (_a) {
                            switch (_a.label) {
                            case 0: return [4 /*yield*/, this.do('PUT', path, '')];
                            case 1: return [2 /*return*/, _a.sent()];
                            }
                        });
                    });
                };
                ;

                HttpClient.prototype.delete = function (path) {
                    return (0,tslib_es6/* __awaiter */.mG)(this, void 0, void 0, function () {
                        return (0,tslib_es6/* __generator */.Jh)(this, function (_a) {
                            switch (_a.label) {
                            case 0: return [4 /*yield*/, this.do('DELETE', path, '')];
                            case 1: return [2 /*return*/, _a.sent()];
                            }
                        });
                    });
                };
                ;

                HttpClient.prototype.patch = function (path, body) {
                    return (0,tslib_es6/* __awaiter */.mG)(this, void 0, void 0, function () {
                        return (0,tslib_es6/* __generator */.Jh)(this, function (_a) {
                            switch (_a.label) {
                            case 0: return [4 /*yield*/, this.do('PATCH', path, body)];
                            case 1: return [2 /*return*/, _a.sent()];
                            }
                        });
                    });
                };

                return HttpClient;
            }());

            ;

            ;// CONCATENATED MODULE: ./src/api/index.ts
            // Copyright (C) 2021 Storj Labs, Inc.
            // See LICENSE for copying information.

            /**
 * ErrorUnauthorized is a custom error type
 * for performing unauthorized operations.
 */
            const UnauthorizedError = /** @class */ (function (_super) {
                (0,tslib_es6/* __extends */.ZT)(UnauthorizedError, _super);

                function UnauthorizedError(message) {
                    if (message === void 0) {
                        message = 'authorization required'; 
                    }

                    return _super.call(this, message) || this;
                }

                return UnauthorizedError;
            }(Error));

            ;

            /**
 * BadRequestError is a custom error type for performing bad request.
 */
            const BadRequestError = /** @class */ (function (_super) {
                (0,tslib_es6/* __extends */.ZT)(BadRequestError, _super);

                function BadRequestError(message) {
                    if (message === void 0) {
                        message = 'bad request'; 
                    }

                    return _super.call(this, message) || this;
                }

                return BadRequestError;
            }(Error));

            ;

            /**
 * InternalError is a custom error type for internal server error.
 */
            const InternalError = /** @class */ (function (_super) {
                (0,tslib_es6/* __extends */.ZT)(InternalError, _super);

                function InternalError(message) {
                    if (message === void 0) {
                        message = 'internal server error'; 
                    }

                    return _super.call(this, message) || this;
                }

                return InternalError;
            }(Error));

            ;

            /**
 * APIClient is base client that holds http client and error handler.
 */
            const APIClient = /** @class */ (function () {
                function APIClient() {
                    this.http = new HttpClient();
                }

                APIClient.prototype.handleError = function (response) {
                    return (0,tslib_es6/* __awaiter */.mG)(this, void 0, void 0, function () {
                        return (0,tslib_es6/* __generator */.Jh)(this, function (_a) {
                            switch (response.status) {
                            case 401: throw new UnauthorizedError();
                            case 400: throw new BadRequestError();
                            case 500: throw new InternalError();
                            default:
                                break;
                            }

                            return [2 /*return*/];
                        });
                    });
                };

                return APIClient;
            }());

            ;

            ;// CONCATENATED MODULE: ./src/api/user.ts
            // Copyright (C) 2021 Creditor Corp. Group.
            // See LICENSE for copying information.

            /** Client for user controller of api */
            var UserClient = /** @class */ (function (_super) {
                (0,tslib_es6/* __extends */.ZT)(UserClient, _super);

                function UserClient() {
                    const _this = _super !== null && _super.apply(this, arguments) || this;

                    _this.ROOT_PATH = '/api/v0/auth';

                    return _this;
                }

                /** Register new user  */
                UserClient.prototype.register = function (user) {
                    return (0,tslib_es6/* __awaiter */.mG)(this, void 0, void 0, function () {
                        let path, response;

                        return (0,tslib_es6/* __generator */.Jh)(this, function (_a) {
                            switch (_a.label) {
                            case 0:
                                path = this.ROOT_PATH + '/register';

                                return [4 /*yield*/, this.http.post(path, JSON.stringify(user))];
                            case 1:
                                response = _a.sent();

                                if (!!response.ok) {
                                    return [3 /*break*/, 3];
                                }

                                return [4 /*yield*/, this.handleError(response)];
                            case 2:
                                _a.sent();
                                _a.label = 3;
                            case 3:
                                ;

                                return [2 /*return*/];
                            }
                        });
                    });
                };
                ;

                /** user login */
                UserClient.prototype.login = function (email, password) {
                    return (0,tslib_es6/* __awaiter */.mG)(this, void 0, void 0, function () {
                        let path, response;

                        return (0,tslib_es6/* __generator */.Jh)(this, function (_a) {
                            switch (_a.label) {
                            case 0:
                                path = this.ROOT_PATH + '/login';

                                return [4 /*yield*/, this.http.post(path, JSON.stringify({
                                    email: email,
                                    password: password
                                }))];
                            case 1:
                                response = _a.sent();

                                if (!!response.ok) {
                                    return [3 /*break*/, 3];
                                }

                                return [4 /*yield*/, this.handleError(response)];
                            case 2:
                                _a.sent();
                                _a.label = 3;
                            case 3:
                                ;

                                return [2 /*return*/];
                            }
                        });
                    });
                };
                ;

                /** change user password implementation */
                UserClient.prototype.changePassword = function (password, newPassword) {
                    return (0,tslib_es6/* __awaiter */.mG)(this, void 0, void 0, function () {
                        let path, response;

                        return (0,tslib_es6/* __generator */.Jh)(this, function (_a) {
                            switch (_a.label) {
                            case 0:
                                path = this.ROOT_PATH + '/change-password';

                                return [4 /*yield*/, this.http.post(path, JSON.stringify({
                                    password: password,
                                    newPassword: newPassword
                                }))];
                            case 1:
                                response = _a.sent();

                                if (!!response.ok) {
                                    return [3 /*break*/, 3];
                                }

                                return [4 /*yield*/, this.handleError(response)];
                            case 2:
                                _a.sent();
                                _a.label = 3;
                            case 3:
                                ;

                                return [2 /*return*/];
                            }
                        });
                    });
                };
                ;

                /** confirm user email implementation */
                UserClient.prototype.confirmEmail = function (token) {
                    return (0,tslib_es6/* __awaiter */.mG)(this, void 0, void 0, function () {
                        let path, response;

                        return (0,tslib_es6/* __generator */.Jh)(this, function (_a) {
                            switch (_a.label) {
                            case 0:
                                path = this.ROOT_PATH + '/email/confirm/' + token;

                                return [4 /*yield*/, this.http.get(path)];
                            case 1:
                                response = _a.sent();

                                if (!!response.ok) {
                                    return [3 /*break*/, 3];
                                }

                                return [4 /*yield*/, this.handleError(response)];
                            case 2:
                                _a.sent();
                                _a.label = 3;
                            case 3:
                                ;

                                return [2 /*return*/, response];
                            }
                        });
                    });
                };
                ;

                /** check user auth token */
                UserClient.prototype.checkToken = function (token) {
                    return (0,tslib_es6/* __awaiter */.mG)(this, void 0, void 0, function () {
                        let path, response;

                        return (0,tslib_es6/* __generator */.Jh)(this, function (_a) {
                            switch (_a.label) {
                            case 0:
                                path = this.ROOT_PATH + '/reset-password/' + token;

                                return [4 /*yield*/, this.http.get(path)];
                            case 1:
                                response = _a.sent();

                                if (!!response.ok) {
                                    return [3 /*break*/, 3];
                                }

                                return [4 /*yield*/, this.handleError(response)];
                            case 2:
                                _a.sent();
                                _a.label = 3;
                            case 3:
                                ;

                                return [2 /*return*/, response];
                            }
                        });
                    });
                };
                ;

                /** recover user password */
                UserClient.prototype.recoverPassword = function (newPassword) {
                    return (0,tslib_es6/* __awaiter */.mG)(this, void 0, void 0, function () {
                        let path, response;

                        return (0,tslib_es6/* __generator */.Jh)(this, function (_a) {
                            switch (_a.label) {
                            case 0:
                                path = this.ROOT_PATH + '/reset-password';

                                return [4 /*yield*/, this.http.patch(path, JSON.stringify({ newPassword: newPassword }))];
                            case 1:
                                response = _a.sent();

                                if (!!response.ok) {
                                    return [3 /*break*/, 3];
                                }

                                return [4 /*yield*/, this.handleError(response)];
                            case 2:
                                _a.sent();
                                _a.label = 3;
                            case 3:
                                ;

                                return [2 /*return*/, response];
                            }
                        });
                    });
                };
                ;

                return UserClient;
            }(APIClient));

            ;

            /***/ }),

        /***/ 2631:
        /***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

            /* harmony export */ __webpack_require__.d(__webpack_exports__, {
                /* harmony export */ 'FM': () => (/* binding */ RouteConfig),
                /* harmony export */ 'Z5': () => (/* binding */ Routes)
                /* harmony export */ });

            /* unused harmony exports ComponentRoutes, Route */
            /* harmony import */ const tslib__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(655);
            /* harmony import */ const react_jsx_runtime__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(5893);
            /* harmony import */ const react__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(7294);
            /* harmony import */ const react_router_dom__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(5977);

            // Copyright (C) 2021 Creditor Corp. Group.
            // See LICENSE for copying information.

            const WelcomePage = react__WEBPACK_IMPORTED_MODULE_1__.lazy(function () {
                return Promise.all(/* import() */[__webpack_require__.e(639), __webpack_require__.e(387)]).then(__webpack_require__.bind(__webpack_require__, 4387)); 
            });
            const SignIn = react__WEBPACK_IMPORTED_MODULE_1__.lazy(function () {
                return __webpack_require__.e(/* import() */ 246).then(__webpack_require__.bind(__webpack_require__, 5246)); 
            });
            const SignUp = react__WEBPACK_IMPORTED_MODULE_1__.lazy(function () {
                return __webpack_require__.e(/* import() */ 525).then(__webpack_require__.bind(__webpack_require__, 8525)); 
            });
            const ChangePassword = react__WEBPACK_IMPORTED_MODULE_1__.lazy(function () {
                return __webpack_require__.e(/* import() */ 904).then(__webpack_require__.bind(__webpack_require__, 1904)); 
            });
            const ConfirmEmail = react__WEBPACK_IMPORTED_MODULE_1__.lazy(function () {
                return __webpack_require__.e(/* import() */ 630).then(__webpack_require__.bind(__webpack_require__, 6630)); 
            });
            const RecoverPassword = react__WEBPACK_IMPORTED_MODULE_1__.lazy(function () {
                return __webpack_require__.e(/* import() */ 103).then(__webpack_require__.bind(__webpack_require__, 7103)); 
            });
            /** Route base config implementation */
            const ComponentRoutes = /** @class */ (function () {
                /** data route config*/
                function ComponentRoutes(path, component, exact, children) {
                    this.path = path;
                    this.component = component;
                    this.exact = exact;
                    this.children = children;
                }
                ;

                /* change path for children routes */
                ComponentRoutes.prototype.with = function (child, parrent) {
                    child.path = parrent.path + '/' + child.path;

                    return this;
                };
                ;

                /* adds children routes to route */
                ComponentRoutes.prototype.addChildren = function (children) {
                    const _this = this;

                    this.children = children.map(function (child) {
                        return child.with(child, _this); 
                    });

                    return this;
                };
                ;

                return ComponentRoutes;
            }());

            ;

            /** Route config implementation */
            var RouteConfig = /** @class */ (function () {
                function RouteConfig() {
                }
                RouteConfig.WelcomePage = new ComponentRoutes('/', WelcomePage, true);
                RouteConfig.SignIn = new ComponentRoutes('/sign-in', SignIn, true);
                RouteConfig.SignUp = new ComponentRoutes('/sign-up', SignUp, true);
                RouteConfig.ResetPassword = new ComponentRoutes('/change-password', ChangePassword, true);
                RouteConfig.ConfirmEmail = new ComponentRoutes('/email/confirm/', ConfirmEmail, true);
                RouteConfig.RecoverPassword = new ComponentRoutes('/recover-password', RecoverPassword, true);

                RouteConfig.routes = [
                    RouteConfig.WelcomePage,
                    RouteConfig.SignIn,
                    RouteConfig.SignUp,
                    RouteConfig.ResetPassword,
                    RouteConfig.ConfirmEmail,
                    RouteConfig.RecoverPassword,
                ];

                return RouteConfig;
            }());

            ;

            const Route = function (_a) {
                const Component = _a.component, children = (0,tslib__WEBPACK_IMPORTED_MODULE_2__/* .__rest */ ._T)(_a, ['component']);

                return (0,react_jsx_runtime__WEBPACK_IMPORTED_MODULE_0__.jsx)(Component, (0,tslib__WEBPACK_IMPORTED_MODULE_2__/* .__assign */ .pi)({}, children), void 0);
            };
            var Routes = function () {
                return ((0,react_jsx_runtime__WEBPACK_IMPORTED_MODULE_0__.jsx)(react_router_dom__WEBPACK_IMPORTED_MODULE_3__/* .Switch */ .rs, { children: RouteConfig.routes.map(function (route, index) {
                    return ((0,react_jsx_runtime__WEBPACK_IMPORTED_MODULE_0__.jsx)(Route, {
                        path: route.path, component: route.component, exact: route.exact, children: route.children 
                    }, index)); 
                }) }, void 0));
            };

            /***/ }),

        /***/ 8096:
        /***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

            /* harmony export */ __webpack_require__.d(__webpack_exports__, {
                /* harmony export */ 'ym': () => (/* binding */ LOGIN),
                /* harmony export */ 'oj': () => (/* binding */ CHANGE_PASSWORD),
                /* harmony export */ 'k4': () => (/* binding */ CONFIRM_EMAIL),
                /* harmony export */ 'Em': () => (/* binding */ RECOVER_PASSWORD),
                /* harmony export */ 'a$': () => (/* binding */ registerUser),
                /* harmony export */ 'pH': () => (/* binding */ loginUser),
                /* harmony export */ 'oi': () => (/* binding */ changeUserPassword),
                /* harmony export */ 'U4': () => (/* binding */ confirmUserEmail),
                /* harmony export */ 'oz': () => (/* binding */ recoverUserPassword)
                /* harmony export */ });

            /* unused harmony exports REGISTER, register, login, changePassword, confirmEmail, recoverPassword */
            /* harmony import */ const tslib__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(655);
            /* harmony import */ const _app_router__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(2631);
            /* harmony import */ const _api_user__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(9443);
            /* harmony import */ const _user_service__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(4526);
            // Copyright (C) 2021 Creditor Corp. Group.
            // See LICENSE for copying information.

            /** action types implementation */
            const REGISTER = 'REGISTER';
            var LOGIN = 'LOGIN';
            var CHANGE_PASSWORD = 'CHANGE_PASSWORD';
            var CONFIRM_EMAIL = 'CONFIRM_EMAIL';
            var RECOVER_PASSWORD = 'RECOVER_PASSWORD';
            /** implement registration of new user */
            const register = function (user) {
                return ({
                    type: REGISTER,
                    user: user,
                }); 
            };
            /** get registred user by id */
            const login = function (email, password) {
                return ({
                    type: LOGIN,
                    user: {
                        email: email,
                        password: password,
                    }
                }); 
            };
            /** changing user password */
            const changePassword = function (password, newPassword) {
                return ({
                    type: CHANGE_PASSWORD,
                    passwords: {
                        password: password,
                        newPassword: newPassword,
                    }
                }); 
            };
            /** user email confirm */
            const confirmEmail = function (token) {
                return ({
                    type: CONFIRM_EMAIL,
                    token: token,
                }); 
            };
            /** recover user password */
            const recoverPassword = function (password) {
                return ({
                    type: RECOVER_PASSWORD,
                    password: password
                }); 
            };
            const client = new _api_user__WEBPACK_IMPORTED_MODULE_1__/* .UserClient */ .K();
            const users = new _user_service__WEBPACK_IMPORTED_MODULE_2__/* .UserService */ .K(client);
            /** thunk that implements user registration */
            var registerUser = function (user) {
                return function (dispatch) {
                    return (0,tslib__WEBPACK_IMPORTED_MODULE_3__/* .__awaiter */ .mG)(this, void 0, void 0, function () {
                        let error_1;

                        return (0,tslib__WEBPACK_IMPORTED_MODULE_3__/* .__generator */ .Jh)(this, function (_a) {
                            switch (_a.label) {
                            case 0:
                                _a.trys.push([0, 2, , 3]);

                                return [4 /*yield*/, users.register(user)];
                            case 1:
                                _a.sent();
                                dispatch(register(user));
                                location.pathname = _app_router__WEBPACK_IMPORTED_MODULE_0__/* .RouteConfig.SignIn.path */ .FM.SignIn.path;

                                return [3 /*break*/, 3];
                            case 2:
                                error_1 = _a.sent();
                                // TODO: rework catching errors
                                /* eslint-disable */
                        console.log(error_1.message);
                        return [3 /*break*/, 3];
                    case 3:
                        ;
                        return [2 /*return*/];
                }
            });
        });
    };
};
/** thunk that implements user login */
var loginUser = function (email, password) {
    return function (dispatch) {
        return (0,tslib__WEBPACK_IMPORTED_MODULE_3__/* .__awaiter */ .mG)(this, void 0, void 0, function () {
            var whitepaperPath, error_2;
            return (0,tslib__WEBPACK_IMPORTED_MODULE_3__/* .__generator */ .Jh)(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        whitepaperPath = '/whitepaper';
                        _a.label = 1;
                    case 1:
                        _a.trys.push([1, 3, , 4]);
                        return [4 /*yield*/, users.login(email, password)];
                    case 2:
                        _a.sent();
                        dispatch(login(email, password));
                        location.pathname = whitepaperPath;
                        return [3 /*break*/, 4];
                    case 3:
                        error_2 = _a.sent();
                        // TODO: rework catching errors
                        /* eslint-disable */
                        console.log(error_2.message);
                        return [3 /*break*/, 4];
                    case 4:
                        ;
                        return [2 /*return*/];
                }
            });
        });
    };
};
/** thunk that implements user changing password */
var changeUserPassword = function (password, newPassword) {
    return function (dispatch) {
        return (0,tslib__WEBPACK_IMPORTED_MODULE_3__/* .__awaiter */ .mG)(this, void 0, void 0, function () {
            var marketplacePath, error_3;
            return (0,tslib__WEBPACK_IMPORTED_MODULE_3__/* .__generator */ .Jh)(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        marketplacePath = '/marketplace';
                        _a.label = 1;
                    case 1:
                        _a.trys.push([1, 3, , 4]);
                        return [4 /*yield*/, users.changePassword(password, newPassword)];
                    case 2:
                        _a.sent();
                        dispatch(changePassword(password, newPassword));
                        location.pathname = marketplacePath;
                        return [3 /*break*/, 4];
                    case 3:
                        error_3 = _a.sent();
                        // TODO: rework catching errors
                        /* eslint-disable */
                        console.log(error_3.message);
                        return [3 /*break*/, 4];
                    case 4:
                        ;
                        return [2 /*return*/];
                }
            });
        });
    };
};
/** thunk that implements user email confirm */
var confirmUserEmail = function (token) {
    return function (dispatch) {
        return (0,tslib__WEBPACK_IMPORTED_MODULE_3__/* .__awaiter */ .mG)(this, void 0, void 0, function () {
            var error_4;
            return (0,tslib__WEBPACK_IMPORTED_MODULE_3__/* .__generator */ .Jh)(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        _a.trys.push([0, 2, , 3]);
                        return [4 /*yield*/, users.confirmEmail(token)];
                    case 1:
                        _a.sent();
                        dispatch(confirmEmail(token));
                        return [3 /*break*/, 3];
                    case 2:
                        error_4 = _a.sent();
                        /** TODO: rework catching errros */
                        /* eslint-disable */
                        console.log(error_4.message);
                        return [3 /*break*/, 3];
                    case 3: return [2 /*return*/];
                }
            });
        });
    };
};
/** thunk that implements user reset password */
var recoverUserPassword = function (password) {
    return function (dispatch) {
        return (0,tslib__WEBPACK_IMPORTED_MODULE_3__/* .__awaiter */ .mG)(this, void 0, void 0, function () {
            var error_5;
            return (0,tslib__WEBPACK_IMPORTED_MODULE_3__/* .__generator */ .Jh)(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        _a.trys.push([0, 2, , 3]);
                        return [4 /*yield*/, users.recoverPassword(password)];
                    case 1:
                        _a.sent();
                        dispatch(recoverPassword(password));
                        return [3 /*break*/, 3];
                    case 2:
                        error_5 = _a.sent();
                        /** TODO: rework catching errros */
                        /* eslint-disable */
                        console.log(error_5.message);
                        return [3 /*break*/, 3];
                    case 3: return [2 /*return*/];
                }
            });
        });
    };
};


/***/ }),

/***/ 4555:
/***/ ((__unused_webpack_module, __unused_webpack___webpack_exports__, __webpack_require__) => {


// EXTERNAL MODULE: ./node_modules/tslib/tslib.es6.js
var tslib_es6 = __webpack_require__(655);
// EXTERNAL MODULE: ./node_modules/react/jsx-runtime.js
var jsx_runtime = __webpack_require__(5893);
// EXTERNAL MODULE: ./node_modules/react-dom/index.js
var react_dom = __webpack_require__(3935);
// EXTERNAL MODULE: ./node_modules/react-redux/es/index.js + 17 modules
var es = __webpack_require__(9226);
// EXTERNAL MODULE: ./node_modules/react/index.js
var react = __webpack_require__(7294);
// EXTERNAL MODULE: ./node_modules/react-router-dom/esm/react-router-dom.js
var react_router_dom = __webpack_require__(3727);
// EXTERNAL MODULE: ./src/app/router/index.tsx
var router = __webpack_require__(2631);
;// CONCATENATED MODULE: ./src/App.tsx


// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.



var App = function () { return ((0,jsx_runtime.jsx)(react_router_dom/* BrowserRouter */.VK, (0,tslib_es6/* __assign */.pi)({ basename: "/" }, { children: (0,jsx_runtime.jsx)(react.Suspense, (0,tslib_es6/* __assign */.pi)({ fallback: (0,jsx_runtime.jsx)("div", { children: "Loading..." }, void 0) }, { children: (0,jsx_runtime.jsx)(router/* Routes */.Z5, {}, void 0) }), void 0) }), void 0)); };

// EXTERNAL MODULE: ./node_modules/redux/es/redux.js + 2 modules
var redux = __webpack_require__(8676);
// EXTERNAL MODULE: ./node_modules/redux-thunk/es/index.js
var redux_thunk_es = __webpack_require__(3894);
// EXTERNAL MODULE: ./src/api/user.ts + 2 modules
var user = __webpack_require__(9443);
// EXTERNAL MODULE: ./src/user/service.ts
var service = __webpack_require__(4526);
// EXTERNAL MODULE: ./src/app/store/actions/users.ts
var users = __webpack_require__(8096);
;// CONCATENATED MODULE: ./src/app/store/reducers/users.ts
// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.




/** implementation of user state */
var UsersState = /** @class */ (function () {
    function UsersState(userService) {
        this.user = {
            email: '',
            password: '',
            status: null,
        };
        this.userService = userService;
    }
    ;
    return UsersState;
}());

;
var client = new user/* UserClient */.K();
var users_service = new service/* UserService */.K(client);
var usersReducer = function (state, action) {
    if (state === void 0) { state = new UsersState(users_service); }
    if (action === void 0) { action = {}; }
    var user = state.user;
    switch (action.type) {
        case users/* LOGIN */.ym:
            user.email = action.user.email;
            user.password = action.user.password;
            break;
        case users/* CHANGE_PASSWORD */.oj:
            user.password = action.passwords.newPassword;
            break;
        case users/* CONFIRM_EMAIL */.k4:
            user.status = action.token;
            break;
        case users/* RECOVER_PASSWORD */.Em:
            user.password = action.password;
            break;
        default:
            break;
    }
    ;
    return (0,tslib_es6/* __assign */.pi)({}, state);
};

;// CONCATENATED MODULE: ./src/app/store/index.ts
// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.



var reducer = (0,redux/* combineReducers */.UY)({
    users: usersReducer,
});
var store = (0,redux/* createStore */.MT)(reducer, (0,redux/* applyMiddleware */.md)(redux_thunk_es/* default */.Z));

// EXTERNAL MODULE: ./node_modules/style-loader/dist/runtime/injectStylesIntoStyleTag.js
var injectStylesIntoStyleTag = __webpack_require__(3379);
var injectStylesIntoStyleTag_default = /*#__PURE__*/__webpack_require__.n(injectStylesIntoStyleTag);
// EXTERNAL MODULE: ./node_modules/style-loader/dist/runtime/styleDomAPI.js
var styleDomAPI = __webpack_require__(7795);
var styleDomAPI_default = /*#__PURE__*/__webpack_require__.n(styleDomAPI);
// EXTERNAL MODULE: ./node_modules/style-loader/dist/runtime/insertBySelector.js
var insertBySelector = __webpack_require__(569);
var insertBySelector_default = /*#__PURE__*/__webpack_require__.n(insertBySelector);
// EXTERNAL MODULE: ./node_modules/style-loader/dist/runtime/setAttributesWithoutAttributes.js
var setAttributesWithoutAttributes = __webpack_require__(3565);
var setAttributesWithoutAttributes_default = /*#__PURE__*/__webpack_require__.n(setAttributesWithoutAttributes);
// EXTERNAL MODULE: ./node_modules/style-loader/dist/runtime/insertStyleElement.js
var insertStyleElement = __webpack_require__(9216);
var insertStyleElement_default = /*#__PURE__*/__webpack_require__.n(insertStyleElement);
// EXTERNAL MODULE: ./node_modules/style-loader/dist/runtime/styleTagTransform.js
var styleTagTransform = __webpack_require__(4589);
var styleTagTransform_default = /*#__PURE__*/__webpack_require__.n(styleTagTransform);
// EXTERNAL MODULE: ./node_modules/css-loader/dist/cjs.js!./node_modules/sass-loader/dist/cjs.js!./src/index.scss
var cjs_js_src = __webpack_require__(4800);
;// CONCATENATED MODULE: ./src/index.scss

      
      
      
      
      
      
      
      
      

var options = {};

options.styleTagTransform = (styleTagTransform_default());
options.setAttributes = (setAttributesWithoutAttributes_default());

      options.insert = insertBySelector_default().bind(null, "head");
    
options.domAPI = (styleDomAPI_default());
options.insertStyleElement = (insertStyleElement_default());

var update = injectStylesIntoStyleTag_default()(cjs_js_src/* default */.Z, options);




       /* harmony default export */ const src = (cjs_js_src/* default */.Z && cjs_js_src/* default.locals */.Z.locals ? cjs_js_src/* default.locals */.Z.locals : undefined);

;// CONCATENATED MODULE: ./src/index.tsx


// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.





react_dom.render((0,jsx_runtime.jsx)(es/* Provider */.zt, (0,tslib_es6/* __assign */.pi)({ store: store }, { children: (0,jsx_runtime.jsx)(App, {}, void 0) }), void 0), document.getElementById('root'));


/***/ }),

/***/ 4526:
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "K": () => (/* binding */ UserService)
/* harmony export */ });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(655);
// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/** exposes all user related logic  */
var UserService = /** @class */ (function () {
    function UserService(users) {
        this.users = users;
    }
    ;
    /** handles user registration */
    UserService.prototype.register = function (user) {
        return (0,tslib__WEBPACK_IMPORTED_MODULE_0__/* .__awaiter */ .mG)(this, void 0, void 0, function () {
            return (0,tslib__WEBPACK_IMPORTED_MODULE_0__/* .__generator */ .Jh)(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, this.users.register(user)];
                    case 1: return [2 /*return*/, _a.sent()];
                }
            });
        });
    };
    ;
    /** return registred user */
    UserService.prototype.login = function (email, password) {
        return (0,tslib__WEBPACK_IMPORTED_MODULE_0__/* .__awaiter */ .mG)(this, void 0, void 0, function () {
            return (0,tslib__WEBPACK_IMPORTED_MODULE_0__/* .__generator */ .Jh)(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, this.users.login(email, password)];
                    case 1: return [2 /*return*/, _a.sent()];
                }
            });
        });
    };
    ;
    /** handles user changing password */
    UserService.prototype.changePassword = function (password, newPassword) {
        return (0,tslib__WEBPACK_IMPORTED_MODULE_0__/* .__awaiter */ .mG)(this, void 0, void 0, function () {
            return (0,tslib__WEBPACK_IMPORTED_MODULE_0__/* .__generator */ .Jh)(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, this.users.changePassword(password, newPassword)];
                    case 1: return [2 /*return*/, _a.sent()];
                }
            });
        });
    };
    ;
    /** handles user email confirmation */
    UserService.prototype.confirmEmail = function (token) {
        return (0,tslib__WEBPACK_IMPORTED_MODULE_0__/* .__awaiter */ .mG)(this, void 0, void 0, function () {
            return (0,tslib__WEBPACK_IMPORTED_MODULE_0__/* .__generator */ .Jh)(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, this.users.confirmEmail(token)];
                    case 1: return [2 /*return*/, _a.sent()];
                }
            });
        });
    };
    ;
    /** handles user check token */
    UserService.prototype.checkToken = function (token) {
        return (0,tslib__WEBPACK_IMPORTED_MODULE_0__/* .__awaiter */ .mG)(this, void 0, void 0, function () {
            return (0,tslib__WEBPACK_IMPORTED_MODULE_0__/* .__generator */ .Jh)(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, this.users.checkToken(token)];
                    case 1: return [2 /*return*/, _a.sent()];
                }
            });
        });
    };
    ;
    /** handles user recover password */
    UserService.prototype.recoverPassword = function (password) {
        return (0,tslib__WEBPACK_IMPORTED_MODULE_0__/* .__awaiter */ .mG)(this, void 0, void 0, function () {
            return (0,tslib__WEBPACK_IMPORTED_MODULE_0__/* .__generator */ .Jh)(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, this.users.recoverPassword(password)];
                    case 1: return [2 /*return*/, _a.sent()];
                }
            });
        });
    };
    ;
    return UserService;
}());

;


/***/ }),

/***/ 8395:
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

module.exports = __webpack_require__.p + "fonts/BebasNeue-Regular8c5b2.ttf";

/***/ })

/******/ 	});
/************************************************************************/
/******/ 	// The module cache
/******/ 	var __webpack_module_cache__ = {};
/******/ 	
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/ 		// Check if module is in cache
/******/ 		var cachedModule = __webpack_module_cache__[moduleId];
/******/ 		if (cachedModule !== undefined) {
/******/ 			return cachedModule.exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = __webpack_module_cache__[moduleId] = {
/******/ 			id: moduleId,
/******/ 			// no module.loaded needed
/******/ 			exports: {}
/******/ 		};
/******/ 	
/******/ 		// Execute the module function
/******/ 		__webpack_modules__[moduleId].call(module.exports, module, module.exports, __webpack_require__);
/******/ 	
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/ 	
/******/ 	// expose the modules object (__webpack_modules__)
/******/ 	__webpack_require__.m = __webpack_modules__;
/******/ 	
/************************************************************************/
/******/ 	/* webpack/runtime/chunk loaded */
/******/ 	(() => {
/******/ 		var deferred = [];
/******/ 		__webpack_require__.O = (result, chunkIds, fn, priority) => {
/******/ 			if(chunkIds) {
/******/ 				priority = priority || 0;
/******/ 				for(var i = deferred.length; i > 0 && deferred[i - 1][2] > priority; i--) deferred[i] = deferred[i - 1];
/******/ 				deferred[i] = [chunkIds, fn, priority];
/******/ 				return;
/******/ 			}
/******/ 			var notFulfilled = Infinity;
/******/ 			for (var i = 0; i < deferred.length; i++) {
/******/ 				var [chunkIds, fn, priority] = deferred[i];
/******/ 				var fulfilled = true;
/******/ 				for (var j = 0; j < chunkIds.length; j++) {
/******/ 					if ((priority & 1 === 0 || notFulfilled >= priority) && Object.keys(__webpack_require__.O).every((key) => (__webpack_require__.O[key](chunkIds[j])))) {
/******/ 						chunkIds.splice(j--, 1);
/******/ 					} else {
/******/ 						fulfilled = false;
/******/ 						if(priority < notFulfilled) notFulfilled = priority;
/******/ 					}
/******/ 				}
/******/ 				if(fulfilled) {
/******/ 					deferred.splice(i--, 1)
/******/ 					var r = fn();
/******/ 					if (r !== undefined) result = r;
/******/ 				}
/******/ 			}
/******/ 			return result;
/******/ 		};
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/compat get default export */
/******/ 	(() => {
/******/ 		// getDefaultExport function for compatibility with non-harmony modules
/******/ 		__webpack_require__.n = (module) => {
/******/ 			var getter = module && module.__esModule ?
/******/ 				() => (module['default']) :
/******/ 				() => (module);
/******/ 			__webpack_require__.d(getter, { a: getter });
/******/ 			return getter;
/******/ 		};
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/define property getters */
/******/ 	(() => {
/******/ 		// define getter functions for harmony exports
/******/ 		__webpack_require__.d = (exports, definition) => {
/******/ 			for(var key in definition) {
/******/ 				if(__webpack_require__.o(definition, key) && !__webpack_require__.o(exports, key)) {
/******/ 					Object.defineProperty(exports, key, { enumerable: true, get: definition[key] });
/******/ 				}
/******/ 			}
/******/ 		};
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/ensure chunk */
/******/ 	(() => {
/******/ 		__webpack_require__.f = {};
/******/ 		// This file contains only the entry chunk.
/******/ 		// The chunk loading function for additional chunks
/******/ 		__webpack_require__.e = (chunkId) => {
/******/ 			return Promise.all(Object.keys(__webpack_require__.f).reduce((promises, key) => {
/******/ 				__webpack_require__.f[key](chunkId, promises);
/******/ 				return promises;
/******/ 			}, []));
/******/ 		};
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/get javascript chunk filename */
/******/ 	(() => {
/******/ 		// This function allow to reference async chunks
/******/ 		__webpack_require__.u = (chunkId) => {
/******/ 			// return url for filenames based on template
/******/ 			return "" + chunkId + "." + __webpack_require__.h() + ".js";
/******/ 		};
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/get mini-css chunk filename */
/******/ 	(() => {
/******/ 		// This function allow to reference all chunks
/******/ 		__webpack_require__.miniCssF = (chunkId) => {
/******/ 			// return url for filenames not based on template
/******/ 			if (chunkId === 369) return "369.css";
/******/ 			if (chunkId === 179) return "main.css";
/******/ 			// return url for filenames based on template
/******/ 			return "" + chunkId + ".css";
/******/ 		};
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/getFullHash */
/******/ 	(() => {
/******/ 		__webpack_require__.h = () => ("70ebad0d95ff287485e4")
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/global */
/******/ 	(() => {
/******/ 		__webpack_require__.g = (function() {
/******/ 			if (typeof globalThis === 'object') return globalThis;
/******/ 			try {
/******/ 				return this || new Function('return this')();
/******/ 			} catch (e) {
/******/ 				if (typeof window === 'object') return window;
/******/ 			}
/******/ 		})();
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/hasOwnProperty shorthand */
/******/ 	(() => {
/******/ 		__webpack_require__.o = (obj, prop) => (Object.prototype.hasOwnProperty.call(obj, prop))
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/load script */
/******/ 	(() => {
/******/ 		var inProgress = {};
/******/ 		var dataWebpackPrefix = "cryptofotball:";
/******/ 		// loadScript function to load a script via script tag
/******/ 		__webpack_require__.l = (url, done, key, chunkId) => {
/******/ 			if(inProgress[url]) { inProgress[url].push(done); return; }
/******/ 			var script, needAttach;
/******/ 			if(key !== undefined) {
/******/ 				var scripts = document.getElementsByTagName("script");
/******/ 				for(var i = 0; i < scripts.length; i++) {
/******/ 					var s = scripts[i];
/******/ 					if(s.getAttribute("src") == url || s.getAttribute("data-webpack") == dataWebpackPrefix + key) { script = s; break; }
/******/ 				}
/******/ 			}
/******/ 			if(!script) {
/******/ 				needAttach = true;
/******/ 				script = document.createElement('script');
/******/ 		
/******/ 				script.charset = 'utf-8';
/******/ 				script.timeout = 120;
/******/ 				if (__webpack_require__.nc) {
/******/ 					script.setAttribute("nonce", __webpack_require__.nc);
/******/ 				}
/******/ 				script.setAttribute("data-webpack", dataWebpackPrefix + key);
/******/ 				script.src = url;
/******/ 			}
/******/ 			inProgress[url] = [done];
/******/ 			var onScriptComplete = (prev, event) => {
/******/ 				// avoid mem leaks in IE.
/******/ 				script.onerror = script.onload = null;
/******/ 				clearTimeout(timeout);
/******/ 				var doneFns = inProgress[url];
/******/ 				delete inProgress[url];
/******/ 				script.parentNode && script.parentNode.removeChild(script);
/******/ 				doneFns && doneFns.forEach((fn) => (fn(event)));
/******/ 				if(prev) return prev(event);
/******/ 			}
/******/ 			;
/******/ 			var timeout = setTimeout(onScriptComplete.bind(null, undefined, { type: 'timeout', target: script }), 120000);
/******/ 			script.onerror = onScriptComplete.bind(null, script.onerror);
/******/ 			script.onload = onScriptComplete.bind(null, script.onload);
/******/ 			needAttach && document.head.appendChild(script);
/******/ 		};
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/make namespace object */
/******/ 	(() => {
/******/ 		// define __esModule on exports
/******/ 		__webpack_require__.r = (exports) => {
/******/ 			if(typeof Symbol !== 'undefined' && Symbol.toStringTag) {
/******/ 				Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });
/******/ 			}
/******/ 			Object.defineProperty(exports, '__esModule', { value: true });
/******/ 		};
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/publicPath */
/******/ 	(() => {
/******/ 		__webpack_require__.p = "/";
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/jsonp chunk loading */
/******/ 	(() => {
/******/ 		__webpack_require__.b = document.baseURI || self.location.href;
/******/ 		
/******/ 		// object to store loaded and loading chunks
/******/ 		// undefined = chunk not loaded, null = chunk preloaded/prefetched
/******/ 		// [resolve, reject, Promise] = chunk loading, 0 = chunk loaded
/******/ 		var installedChunks = {
/******/ 			179: 0
/******/ 		};
/******/ 		
/******/ 		__webpack_require__.f.j = (chunkId, promises) => {
/******/ 				// JSONP chunk loading for javascript
/******/ 				var installedChunkData = __webpack_require__.o(installedChunks, chunkId) ? installedChunks[chunkId] : undefined;
/******/ 				if(installedChunkData !== 0) { // 0 means "already installed".
/******/ 		
/******/ 					// a Promise means "currently loading".
/******/ 					if(installedChunkData) {
/******/ 						promises.push(installedChunkData[2]);
/******/ 					} else {
/******/ 						if(true) { // all chunks have JS
/******/ 							// setup Promise in chunk cache
/******/ 							var promise = new Promise((resolve, reject) => (installedChunkData = installedChunks[chunkId] = [resolve, reject]));
/******/ 							promises.push(installedChunkData[2] = promise);
/******/ 		
/******/ 							// start chunk loading
/******/ 							var url = __webpack_require__.p + __webpack_require__.u(chunkId);
/******/ 							// create error before stack unwound to get useful stacktrace later
/******/ 							var error = new Error();
/******/ 							var loadingEnded = (event) => {
/******/ 								if(__webpack_require__.o(installedChunks, chunkId)) {
/******/ 									installedChunkData = installedChunks[chunkId];
/******/ 									if(installedChunkData !== 0) installedChunks[chunkId] = undefined;
/******/ 									if(installedChunkData) {
/******/ 										var errorType = event && (event.type === 'load' ? 'missing' : event.type);
/******/ 										var realSrc = event && event.target && event.target.src;
/******/ 										error.message = 'Loading chunk ' + chunkId + ' failed.\n(' + errorType + ': ' + realSrc + ')';
/******/ 										error.name = 'ChunkLoadError';
/******/ 										error.type = errorType;
/******/ 										error.request = realSrc;
/******/ 										installedChunkData[1](error);
/******/ 									}
/******/ 								}
/******/ 							};
/******/ 							__webpack_require__.l(url, loadingEnded, "chunk-" + chunkId, chunkId);
/******/ 						} else installedChunks[chunkId] = 0;
/******/ 					}
/******/ 				}
/******/ 		};
/******/ 		
/******/ 		// no prefetching
/******/ 		
/******/ 		// no preloaded
/******/ 		
/******/ 		// no HMR
/******/ 		
/******/ 		// no HMR manifest
/******/ 		
/******/ 		__webpack_require__.O.j = (chunkId) => (installedChunks[chunkId] === 0);
/******/ 		
/******/ 		// install a JSONP callback for chunk loading
/******/ 		var webpackJsonpCallback = (parentChunkLoadingFunction, data) => {
/******/ 			var [chunkIds, moreModules, runtime] = data;
/******/ 			// add "moreModules" to the modules object,
/******/ 			// then flag all "chunkIds" as loaded and fire callback
/******/ 			var moduleId, chunkId, i = 0;
/******/ 			for(moduleId in moreModules) {
/******/ 				if(__webpack_require__.o(moreModules, moduleId)) {
/******/ 					__webpack_require__.m[moduleId] = moreModules[moduleId];
/******/ 				}
/******/ 			}
/******/ 			if(runtime) var result = runtime(__webpack_require__);
/******/ 			if(parentChunkLoadingFunction) parentChunkLoadingFunction(data);
/******/ 			for(;i < chunkIds.length; i++) {
/******/ 				chunkId = chunkIds[i];
/******/ 				if(__webpack_require__.o(installedChunks, chunkId) && installedChunks[chunkId]) {
/******/ 					installedChunks[chunkId][0]();
/******/ 				}
/******/ 				installedChunks[chunkIds[i]] = 0;
/******/ 			}
/******/ 			return __webpack_require__.O(result);
/******/ 		}
/******/ 		
/******/ 		var chunkLoadingGlobal = self["webpackChunkcryptofotball"] = self["webpackChunkcryptofotball"] || [];
/******/ 		chunkLoadingGlobal.forEach(webpackJsonpCallback.bind(null, 0));
/******/ 		chunkLoadingGlobal.push = webpackJsonpCallback.bind(null, chunkLoadingGlobal.push.bind(chunkLoadingGlobal));
/******/ 	})();
/******/ 	
/************************************************************************/
/******/ 	
/******/ 	// startup
/******/ 	// Load entry module and return exports
/******/ 	// This entry module depends on other loaded chunks and execution need to be delayed
/******/ 	var __webpack_exports__ = __webpack_require__.O(undefined, [369], () => (__webpack_require__(4555)))
/******/ 	__webpack_exports__ = __webpack_require__.O(__webpack_exports__);
/******/ 	
/******/ })()
;
//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly9jcnlwdG9mb3RiYWxsLy4vc3JjL2luZGV4LnNjc3MiLCJ3ZWJwYWNrOi8vY3J5cHRvZm90YmFsbC8uL3NyYy9wcml2YXRlL2h0dHAvY2xpZW50LnRzIiwid2VicGFjazovL2NyeXB0b2ZvdGJhbGwvLi9zcmMvYXBpL2luZGV4LnRzIiwid2VicGFjazovL2NyeXB0b2ZvdGJhbGwvLi9zcmMvYXBpL3VzZXIudHMiLCJ3ZWJwYWNrOi8vY3J5cHRvZm90YmFsbC8uL3NyYy9hcHAvcm91dGVyL2luZGV4LnRzeCIsIndlYnBhY2s6Ly9jcnlwdG9mb3RiYWxsLy4vc3JjL2FwcC9zdG9yZS9hY3Rpb25zL3VzZXJzLnRzIiwid2VicGFjazovL2NyeXB0b2ZvdGJhbGwvLi9zcmMvQXBwLnRzeCIsIndlYnBhY2s6Ly9jcnlwdG9mb3RiYWxsLy4vc3JjL2FwcC9zdG9yZS9yZWR1Y2Vycy91c2Vycy50cyIsIndlYnBhY2s6Ly9jcnlwdG9mb3RiYWxsLy4vc3JjL2FwcC9zdG9yZS9pbmRleC50cyIsIndlYnBhY2s6Ly9jcnlwdG9mb3RiYWxsLy4vc3JjL2luZGV4LnNjc3M/NzIyMyIsIndlYnBhY2s6Ly9jcnlwdG9mb3RiYWxsLy4vc3JjL2luZGV4LnRzeCIsIndlYnBhY2s6Ly9jcnlwdG9mb3RiYWxsLy4vc3JjL3VzZXIvc2VydmljZS50cyIsIndlYnBhY2s6Ly9jcnlwdG9mb3RiYWxsL3dlYnBhY2svYm9vdHN0cmFwIiwid2VicGFjazovL2NyeXB0b2ZvdGJhbGwvd2VicGFjay9ydW50aW1lL2NodW5rIGxvYWRlZCIsIndlYnBhY2s6Ly9jcnlwdG9mb3RiYWxsL3dlYnBhY2svcnVudGltZS9jb21wYXQgZ2V0IGRlZmF1bHQgZXhwb3J0Iiwid2VicGFjazovL2NyeXB0b2ZvdGJhbGwvd2VicGFjay9ydW50aW1lL2RlZmluZSBwcm9wZXJ0eSBnZXR0ZXJzIiwid2VicGFjazovL2NyeXB0b2ZvdGJhbGwvd2VicGFjay9ydW50aW1lL2Vuc3VyZSBjaHVuayIsIndlYnBhY2s6Ly9jcnlwdG9mb3RiYWxsL3dlYnBhY2svcnVudGltZS9nZXQgamF2YXNjcmlwdCBjaHVuayBmaWxlbmFtZSIsIndlYnBhY2s6Ly9jcnlwdG9mb3RiYWxsL3dlYnBhY2svcnVudGltZS9nZXQgbWluaS1jc3MgY2h1bmsgZmlsZW5hbWUiLCJ3ZWJwYWNrOi8vY3J5cHRvZm90YmFsbC93ZWJwYWNrL3J1bnRpbWUvZ2V0RnVsbEhhc2giLCJ3ZWJwYWNrOi8vY3J5cHRvZm90YmFsbC93ZWJwYWNrL3J1bnRpbWUvZ2xvYmFsIiwid2VicGFjazovL2NyeXB0b2ZvdGJhbGwvd2VicGFjay9ydW50aW1lL2hhc093blByb3BlcnR5IHNob3J0aGFuZCIsIndlYnBhY2s6Ly9jcnlwdG9mb3RiYWxsL3dlYnBhY2svcnVudGltZS9sb2FkIHNjcmlwdCIsIndlYnBhY2s6Ly9jcnlwdG9mb3RiYWxsL3dlYnBhY2svcnVudGltZS9tYWtlIG5hbWVzcGFjZSBvYmplY3QiLCJ3ZWJwYWNrOi8vY3J5cHRvZm90YmFsbC93ZWJwYWNrL3J1bnRpbWUvcHVibGljUGF0aCIsIndlYnBhY2s6Ly9jcnlwdG9mb3RiYWxsL3dlYnBhY2svcnVudGltZS9qc29ucCBjaHVuayBsb2FkaW5nIiwid2VicGFjazovL2NyeXB0b2ZvdGJhbGwvd2VicGFjay9zdGFydHVwIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiI7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQTtBQUNzSDtBQUM3QjtBQUNPO0FBQ2hHLDRDQUE0QyxtRUFBc0Q7QUFDbEcsOEJBQThCLG1GQUEyQixDQUFDLHdHQUFxQztBQUMvRix5Q0FBeUMsc0ZBQStCO0FBQ3hFO0FBQ0EsNENBQTRDLHNCQUFzQixnREFBZ0QsU0FBUyxVQUFVLEVBQUUsc0JBQXNCLGdEQUFnRCxTQUFTLFVBQVUsV0FBVywwQkFBMEIsb0RBQW9ELGtCQUFrQixvQkFBb0IsS0FBSyxxQ0FBcUMsT0FBTyx5U0FBeVMsNkJBQTZCLEdBQUcsNkVBQTZFLGdCQUFnQixpQkFBaUIsR0FBRyxzQ0FBc0MsdUNBQXVD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QUNBO1VBQ0EiLCJmaWxlIjoibWFpbi43MGViYWQwZDk1ZmYyODc0ODVlNC5qcyIsInNvdXJjZXNDb250ZW50IjpbIi8vIEltcG9ydHNcbmltcG9ydCBfX19DU1NfTE9BREVSX0FQSV9TT1VSQ0VNQVBfSU1QT1JUX19fIGZyb20gXCIuLi9ub2RlX21vZHVsZXMvY3NzLWxvYWRlci9kaXN0L3J1bnRpbWUvY3NzV2l0aE1hcHBpbmdUb1N0cmluZy5qc1wiO1xuaW1wb3J0IF9fX0NTU19MT0FERVJfQVBJX0lNUE9SVF9fXyBmcm9tIFwiLi4vbm9kZV9tb2R1bGVzL2Nzcy1sb2FkZXIvZGlzdC9ydW50aW1lL2FwaS5qc1wiO1xuaW1wb3J0IF9fX0NTU19MT0FERVJfR0VUX1VSTF9JTVBPUlRfX18gZnJvbSBcIi4uL25vZGVfbW9kdWxlcy9jc3MtbG9hZGVyL2Rpc3QvcnVudGltZS9nZXRVcmwuanNcIjtcbnZhciBfX19DU1NfTE9BREVSX1VSTF9JTVBPUlRfMF9fXyA9IG5ldyBVUkwoXCJAc3RhdGljL2ZvbnRzL0JlYmFzTmV1ZS1SZWd1bGFyLnR0ZlwiLCBpbXBvcnQubWV0YS51cmwpO1xudmFyIF9fX0NTU19MT0FERVJfRVhQT1JUX19fID0gX19fQ1NTX0xPQURFUl9BUElfSU1QT1JUX19fKF9fX0NTU19MT0FERVJfQVBJX1NPVVJDRU1BUF9JTVBPUlRfX18pO1xudmFyIF9fX0NTU19MT0FERVJfVVJMX1JFUExBQ0VNRU5UXzBfX18gPSBfX19DU1NfTE9BREVSX0dFVF9VUkxfSU1QT1JUX19fKF9fX0NTU19MT0FERVJfVVJMX0lNUE9SVF8wX19fKTtcbi8vIE1vZHVsZVxuX19fQ1NTX0xPQURFUl9FWFBPUlRfX18ucHVzaChbbW9kdWxlLmlkLCBcIip7Ym94LXNpemluZzpib3JkZXItYm94fWJvZHksc2VjdGlvbixkaXYsc3BhbixoMSxoMixoMyxoNCxoNSxoNixwLHVsLGxpe21hcmdpbjowO3BhZGRpbmc6MH0qe2JveC1zaXppbmc6Ym9yZGVyLWJveH1ib2R5LHNlY3Rpb24sZGl2LHNwYW4saDEsaDIsaDMsaDQsaDUsaDYscCx1bCxsaXttYXJnaW46MDtwYWRkaW5nOjB9QGZvbnQtZmFjZXtmb250LWZhbWlseTpcXFwiQmViYXNOZXVlXFxcIjtzcmM6dXJsKFwiICsgX19fQ1NTX0xPQURFUl9VUkxfUkVQTEFDRU1FTlRfMF9fXyArIFwiKTtmb250LXN0eWxlOm5vcm1hbDtmb250LXdlaWdodDpsaWdodGVyfWh0bWx7Zm9udC1mYW1pbHk6XFxcIkJlYmFzTmV1ZVxcXCIsc2Fucy1zZXJpZn1cIiwgXCJcIix7XCJ2ZXJzaW9uXCI6MyxcInNvdXJjZXNcIjpbXCJ3ZWJwYWNrOi8vLi9zcmMvYXBwL3N0YXRpYy9zdHlsZXMvX3Jlc2V0LnNjc3NcIixcIndlYnBhY2s6Ly8uL3NyYy9pbmRleC5zY3NzXCJdLFwibmFtZXNcIjpbXSxcIm1hcHBpbmdzXCI6XCJBQUFBLEVBQ0kscUJBQUEsQ0FHSixnREFhSSxRQUFBLENBQ0EsU0FBQSxDQWxCSixFQUNJLHFCQUFBLENBR0osZ0RBYUksUUFBQSxDQUNBLFNBQUEsQ0NiSixXQUNJLHVCQUFBLENBQ0EsMkNBQUEsQ0FDQSxpQkFBQSxDQUNBLG1CQUFBLENBR0osS0FDSSxrQ0FBQVwiLFwic291cmNlc0NvbnRlbnRcIjpbXCIqIHtcXG4gICAgYm94LXNpemluZzogYm9yZGVyLWJveDtcXG59XFxuXFxuYm9keSxcXG5zZWN0aW9uLFxcbmRpdixcXG5zcGFuLFxcbmgxLFxcbmgyLFxcbmgzLFxcbmg0LFxcbmg1LFxcbmg2LFxcbnAsXFxudWwsXFxubGkge1xcbiAgICBtYXJnaW46IDA7XFxuICAgIHBhZGRpbmc6IDA7XFxufVxcblwiLFwiQGltcG9ydCBcXFwiQHN0YXRpYy9zdHlsZXMvcmVzZXRcXFwiO1xcbkBpbXBvcnQgXFxcIkBzdGF0aWMvc3R5bGVzL3ZhcmlhYmxlc1xcXCI7XFxuQGltcG9ydCBcXFwiQHN0YXRpYy9zdHlsZXMvZXh0ZW5kc1xcXCI7XFxuQGltcG9ydCBcXFwiQHN0YXRpYy9zdHlsZXMvbWl4aW5zXFxcIjtcXG5cXG5AZm9udC1mYWNlIHtcXG4gICAgZm9udC1mYW1pbHk6IFxcXCJCZWJhc05ldWVcXFwiO1xcbiAgICBzcmM6IHVybChcXFwiQHN0YXRpYy9mb250cy9CZWJhc05ldWUtUmVndWxhci50dGZcXFwiKTtcXG4gICAgZm9udC1zdHlsZTogbm9ybWFsO1xcbiAgICBmb250LXdlaWdodDogbGlnaHRlcjtcXG59XFxuXFxuaHRtbCB7XFxuICAgIGZvbnQtZmFtaWx5OiBcXFwiQmViYXNOZXVlXFxcIiwgc2Fucy1zZXJpZjtcXG59XFxuXCJdLFwic291cmNlUm9vdFwiOlwiXCJ9XSk7XG4vLyBFeHBvcnRzXG5leHBvcnQgZGVmYXVsdCBfX19DU1NfTE9BREVSX0VYUE9SVF9fXztcbiIsIi8vIENvcHlyaWdodCAoQykgMjAyMSBDcmVkaXRvciBDb3JwLiBHcm91cC5cbi8vIFNlZSBMSUNFTlNFIGZvciBjb3B5aW5nIGluZm9ybWF0aW9uLlxuXG5leHBvcnQgY2xhc3MgSHR0cENsaWVudCB7XG4gICAgLyogZG8gc2VuZHMgYW4gSFRUUCByZXF1ZXN0IGFuZCByZXR1cm5zXG4gICAgKiBhbiBIVFRQIHJlc3BvbnNlIGFzIGNvbmZpZ3VyZWQgb24gdGhlIGNsaWVudC4gKi9cbiAgICBwcml2YXRlIGFzeW5jIGRvKFxuICAgICAgICBtZXRob2Q6IHN0cmluZyxcbiAgICAgICAgcGF0aDogc3RyaW5nLFxuICAgICAgICBib2R5OiBzdHJpbmcgfCBudWxsXG4gICAgKTogUHJvbWlzZTxSZXNwb25zZT4ge1xuICAgICAgICBjb25zdCByZXF1ZXN0OiBSZXF1ZXN0SW5pdCA9IHtcbiAgICAgICAgICAgIG1ldGhvZDogbWV0aG9kLFxuICAgICAgICAgICAgYm9keTogYm9keSxcbiAgICAgICAgfTtcblxuICAgICAgICByZXF1ZXN0LmhlYWRlcnMgPSB7XG4gICAgICAgICAgICAnQWNjZXB0JzogJ2FwcGxpY2F0aW9uL2pzb24nLFxuICAgICAgICAgICAgJ0NvbnRlbnQtdHlwZSc6ICdhcHBsaWNhdGlvbi9qc29uJyxcbiAgICAgICAgfTtcblxuICAgICAgICByZXR1cm4gYXdhaXQgZmV0Y2gocGF0aCwgcmVxdWVzdCk7XG4gICAgfTtcbiAgICBwdWJsaWMgYXN5bmMgcG9zdChwYXRoOiBzdHJpbmcsIGJvZHk6IHN0cmluZyB8IG51bGwpOiBQcm9taXNlPFJlc3BvbnNlPiB7XG4gICAgICAgIHJldHVybiBhd2FpdCB0aGlzLmRvKCdQT1NUJywgcGF0aCwgYm9keSk7XG4gICAgfTtcbiAgICBwdWJsaWMgYXN5bmMgZ2V0KHBhdGg6IHN0cmluZyk6IFByb21pc2U8UmVzcG9uc2U+IHtcbiAgICAgICAgcmV0dXJuIGF3YWl0IHRoaXMuZG8oJ0dFVCcsIHBhdGgsIG51bGwpO1xuICAgIH07XG4gICAgcHVibGljIGFzeW5jIHB1dChwYXRoOiBzdHJpbmcsIGJvZHk6IHN0cmluZyB8IG51bGwpIHtcbiAgICAgICAgcmV0dXJuIGF3YWl0IHRoaXMuZG8oJ1BVVCcsIHBhdGgsICcnKTtcbiAgICB9O1xuICAgIHB1YmxpYyBhc3luYyBkZWxldGUocGF0aDogc3RyaW5nKTogUHJvbWlzZTxSZXNwb25zZT4ge1xuICAgICAgICByZXR1cm4gYXdhaXQgdGhpcy5kbygnREVMRVRFJywgcGF0aCwgJycpO1xuICAgIH07XG4gICAgcHVibGljIGFzeW5jIHBhdGNoKHBhdGg6IHN0cmluZywgYm9keTogc3RyaW5nKTogUHJvbWlzZTxSZXNwb25zZT4ge1xuICAgICAgICByZXR1cm4gYXdhaXQgdGhpcy5kbygnUEFUQ0gnLCBwYXRoLCBib2R5KTtcbiAgICB9XG59O1xuIiwiLy8gQ29weXJpZ2h0IChDKSAyMDIxIFN0b3JqIExhYnMsIEluYy5cbi8vIFNlZSBMSUNFTlNFIGZvciBjb3B5aW5nIGluZm9ybWF0aW9uLlxuXG5pbXBvcnQgeyBIdHRwQ2xpZW50IH0gZnJvbSAnQC9wcml2YXRlL2h0dHAvY2xpZW50JztcblxuLyoqXG4gKiBFcnJvclVuYXV0aG9yaXplZCBpcyBhIGN1c3RvbSBlcnJvciB0eXBlXG4gKiBmb3IgcGVyZm9ybWluZyB1bmF1dGhvcml6ZWQgb3BlcmF0aW9ucy5cbiAqL1xuZXhwb3J0IGNsYXNzIFVuYXV0aG9yaXplZEVycm9yIGV4dGVuZHMgRXJyb3Ige1xuICAgIHB1YmxpYyBjb25zdHJ1Y3RvcihtZXNzYWdlID0gJ2F1dGhvcml6YXRpb24gcmVxdWlyZWQnKSB7XG4gICAgICAgIHN1cGVyKG1lc3NhZ2UpO1xuICAgIH1cbn07XG5cbi8qKlxuICogQmFkUmVxdWVzdEVycm9yIGlzIGEgY3VzdG9tIGVycm9yIHR5cGUgZm9yIHBlcmZvcm1pbmcgYmFkIHJlcXVlc3QuXG4gKi9cbmV4cG9ydCBjbGFzcyBCYWRSZXF1ZXN0RXJyb3IgZXh0ZW5kcyBFcnJvciB7XG4gICAgcHVibGljIGNvbnN0cnVjdG9yKG1lc3NhZ2UgPSAnYmFkIHJlcXVlc3QnKSB7XG4gICAgICAgIHN1cGVyKG1lc3NhZ2UpO1xuICAgIH1cbn07XG5cbi8qKlxuICogSW50ZXJuYWxFcnJvciBpcyBhIGN1c3RvbSBlcnJvciB0eXBlIGZvciBpbnRlcm5hbCBzZXJ2ZXIgZXJyb3IuXG4gKi9cbmV4cG9ydCBjbGFzcyBJbnRlcm5hbEVycm9yIGV4dGVuZHMgRXJyb3Ige1xuICAgIHB1YmxpYyBjb25zdHJ1Y3RvcihtZXNzYWdlID0gJ2ludGVybmFsIHNlcnZlciBlcnJvcicpIHtcbiAgICAgICAgc3VwZXIobWVzc2FnZSk7XG4gICAgfVxufTtcblxuLyoqXG4gKiBBUElDbGllbnQgaXMgYmFzZSBjbGllbnQgdGhhdCBob2xkcyBodHRwIGNsaWVudCBhbmQgZXJyb3IgaGFuZGxlci5cbiAqL1xuZXhwb3J0IGNsYXNzIEFQSUNsaWVudCB7XG4gICAgcHJvdGVjdGVkIHJlYWRvbmx5IGh0dHA6IEh0dHBDbGllbnQgPSBuZXcgSHR0cENsaWVudCgpO1xuICAgIHByb3RlY3RlZCBhc3luYyBoYW5kbGVFcnJvcihyZXNwb25zZTogUmVzcG9uc2UpOiBQcm9taXNlPHZvaWQ+IHtcbiAgICAgICAgc3dpdGNoIChyZXNwb25zZS5zdGF0dXMpIHtcbiAgICAgICAgY2FzZSA0MDE6IHRocm93IG5ldyBVbmF1dGhvcml6ZWRFcnJvcigpO1xuICAgICAgICBjYXNlIDQwMDogdGhyb3cgbmV3IEJhZFJlcXVlc3RFcnJvcigpO1xuICAgICAgICBjYXNlIDUwMDogdGhyb3cgbmV3IEludGVybmFsRXJyb3IoKTtcbiAgICAgICAgZGVmYXVsdDpcbiAgICAgICAgICAgIGJyZWFrO1xuICAgICAgICB9XG4gICAgfVxufTtcbiIsIi8vIENvcHlyaWdodCAoQykgMjAyMSBDcmVkaXRvciBDb3JwLiBHcm91cC5cbi8vIFNlZSBMSUNFTlNFIGZvciBjb3B5aW5nIGluZm9ybWF0aW9uLlxuXG5pbXBvcnQgeyBBUElDbGllbnQgfSBmcm9tICcuJztcbmltcG9ydCB7IFVzZXIgfSBmcm9tICdAL3VzZXInO1xuXG4vKiogQ2xpZW50IGZvciB1c2VyIGNvbnRyb2xsZXIgb2YgYXBpICovXG5leHBvcnQgY2xhc3MgVXNlckNsaWVudCBleHRlbmRzIEFQSUNsaWVudCB7XG4gICAgcHJpdmF0ZSByZWFkb25seSBST09UX1BBVEg6IHN0cmluZyA9ICcvYXBpL3YwL2F1dGgnO1xuICAgIC8qKiBSZWdpc3RlciBuZXcgdXNlciAgKi9cbiAgICBwdWJsaWMgYXN5bmMgcmVnaXN0ZXIodXNlcjogVXNlcik6IFByb21pc2U8dm9pZD4ge1xuICAgICAgICBjb25zdCBwYXRoID0gYCR7dGhpcy5ST09UX1BBVEh9L3JlZ2lzdGVyYDtcbiAgICAgICAgY29uc3QgcmVzcG9uc2UgPSBhd2FpdCB0aGlzLmh0dHAucG9zdChwYXRoLCBKU09OLnN0cmluZ2lmeSh1c2VyKSk7XG5cbiAgICAgICAgaWYgKCFyZXNwb25zZS5vaykge1xuICAgICAgICAgICAgYXdhaXQgdGhpcy5oYW5kbGVFcnJvcihyZXNwb25zZSk7XG4gICAgICAgIH07XG4gICAgfTtcbiAgICAvKiogdXNlciBsb2dpbiAqL1xuICAgIHB1YmxpYyBhc3luYyBsb2dpbihlbWFpbDogc3RyaW5nLCBwYXNzd29yZDogc3RyaW5nKTogUHJvbWlzZTx2b2lkPiB7XG4gICAgICAgIGNvbnN0IHBhdGggPSBgJHt0aGlzLlJPT1RfUEFUSH0vbG9naW5gO1xuICAgICAgICBjb25zdCByZXNwb25zZSA9IGF3YWl0IHRoaXMuaHR0cC5wb3N0KHBhdGgsIEpTT04uc3RyaW5naWZ5KHtcbiAgICAgICAgICAgIGVtYWlsLCBwYXNzd29yZFxuICAgICAgICB9KSk7XG5cbiAgICAgICAgaWYgKCFyZXNwb25zZS5vaykge1xuICAgICAgICAgICAgYXdhaXQgdGhpcy5oYW5kbGVFcnJvcihyZXNwb25zZSk7XG4gICAgICAgIH07XG4gICAgfTtcbiAgICAvKiogY2hhbmdlIHVzZXIgcGFzc3dvcmQgaW1wbGVtZW50YXRpb24gKi9cbiAgICBwdWJsaWMgYXN5bmMgY2hhbmdlUGFzc3dvcmQocGFzc3dvcmQ6IHN0cmluZywgbmV3UGFzc3dvcmQ6IHN0cmluZyk6IFByb21pc2U8dm9pZD4ge1xuICAgICAgICBjb25zdCBwYXRoID0gYCR7dGhpcy5ST09UX1BBVEh9L2NoYW5nZS1wYXNzd29yZGA7XG4gICAgICAgIGNvbnN0IHJlc3BvbnNlID0gYXdhaXQgdGhpcy5odHRwLnBvc3QocGF0aCwgSlNPTi5zdHJpbmdpZnkoe1xuICAgICAgICAgICAgcGFzc3dvcmQsIG5ld1Bhc3N3b3JkXG4gICAgICAgIH0pKTtcblxuICAgICAgICBpZiAoIXJlc3BvbnNlLm9rKSB7XG4gICAgICAgICAgICBhd2FpdCB0aGlzLmhhbmRsZUVycm9yKHJlc3BvbnNlKTtcbiAgICAgICAgfTtcbiAgICB9O1xuICAgIC8qKiBjb25maXJtIHVzZXIgZW1haWwgaW1wbGVtZW50YXRpb24gKi9cbiAgICBwdWJsaWMgYXN5bmMgY29uZmlybUVtYWlsKHRva2VuOiBzdHJpbmcgfCBudWxsKTogUHJvbWlzZTxSZXNwb25zZT4ge1xuICAgICAgICBjb25zdCBwYXRoID0gYCR7dGhpcy5ST09UX1BBVEh9L2VtYWlsL2NvbmZpcm0vJHt0b2tlbn1gO1xuICAgICAgICBjb25zdCByZXNwb25zZSA9IGF3YWl0IHRoaXMuaHR0cC5nZXQocGF0aCk7XG5cbiAgICAgICAgaWYgKCFyZXNwb25zZS5vaykge1xuICAgICAgICAgICAgYXdhaXQgdGhpcy5oYW5kbGVFcnJvcihyZXNwb25zZSk7XG4gICAgICAgIH07XG5cbiAgICAgICAgcmV0dXJuIHJlc3BvbnNlO1xuICAgIH07XG4gICAgLyoqIGNoZWNrIHVzZXIgYXV0aCB0b2tlbiAqL1xuICAgIHB1YmxpYyBhc3luYyBjaGVja1Rva2VuKHRva2VuOiBzdHJpbmcgfCBudWxsKTogUHJvbWlzZTxSZXNwb25zZT4ge1xuICAgICAgICBjb25zdCBwYXRoID0gYCR7dGhpcy5ST09UX1BBVEh9L3Jlc2V0LXBhc3N3b3JkLyR7dG9rZW59YDtcbiAgICAgICAgY29uc3QgcmVzcG9uc2UgPSBhd2FpdCB0aGlzLmh0dHAuZ2V0KHBhdGgpO1xuXG4gICAgICAgIGlmICghcmVzcG9uc2Uub2spIHtcbiAgICAgICAgICAgIGF3YWl0IHRoaXMuaGFuZGxlRXJyb3IocmVzcG9uc2UpO1xuICAgICAgICB9O1xuXG4gICAgICAgIHJldHVybiByZXNwb25zZTtcbiAgICB9O1xuICAgIC8qKiByZWNvdmVyIHVzZXIgcGFzc3dvcmQgKi9cbiAgICBwdWJsaWMgYXN5bmMgcmVjb3ZlclBhc3N3b3JkKG5ld1Bhc3N3b3JkOiBzdHJpbmcpOiBQcm9taXNlPFJlc3BvbnNlPiB7XG4gICAgICAgIGNvbnN0IHBhdGggPSBgJHt0aGlzLlJPT1RfUEFUSH0vcmVzZXQtcGFzc3dvcmRgO1xuICAgICAgICBjb25zdCByZXNwb25zZSA9IGF3YWl0IHRoaXMuaHR0cC5wYXRjaChwYXRoLCBKU09OLnN0cmluZ2lmeSh7IG5ld1Bhc3N3b3JkIH0pKTtcblxuICAgICAgICBpZiAoIXJlc3BvbnNlLm9rKSB7XG4gICAgICAgICAgICBhd2FpdCB0aGlzLmhhbmRsZUVycm9yKHJlc3BvbnNlKTtcbiAgICAgICAgfTtcblxuICAgICAgICByZXR1cm4gcmVzcG9uc2U7XG4gICAgfTtcbn07XG4iLCIvLyBDb3B5cmlnaHQgKEMpIDIwMjEgQ3JlZGl0b3IgQ29ycC4gR3JvdXAuXG4vLyBTZWUgTElDRU5TRSBmb3IgY29weWluZyBpbmZvcm1hdGlvbi5cblxuaW1wb3J0IFJlYWN0LCB7IExhenlFeG90aWNDb21wb25lbnQgfSBmcm9tICdyZWFjdCc7XG5pbXBvcnQgeyBTd2l0Y2ggfSBmcm9tICdyZWFjdC1yb3V0ZXItZG9tJztcblxuY29uc3QgV2VsY29tZVBhZ2UgPSBSZWFjdC5sYXp5KCgpID0+IGltcG9ydCgnQGNvbXBvbmVudHMvV2VsY29tZVBhZ2UnKSk7XG5jb25zdCBTaWduSW4gPSBSZWFjdC5sYXp5KCgpID0+IGltcG9ydCgnQC9hcHAvdmlld3MvU2lnbkluJykpO1xuY29uc3QgU2lnblVwID0gUmVhY3QubGF6eSgoKSA9PiBpbXBvcnQoJ0AvYXBwL3ZpZXdzL1NpZ25VcCcpKTtcbmNvbnN0IENoYW5nZVBhc3N3b3JkID0gUmVhY3QubGF6eSgoKSA9PiBpbXBvcnQoJ0AvYXBwL3ZpZXdzL0NoYW5nZVBhc3N3b3JkJykpO1xuY29uc3QgQ29uZmlybUVtYWlsID0gUmVhY3QubGF6eSgoKSA9PiBpbXBvcnQoJ0AvYXBwL3ZpZXdzL0NvbmZpcm1FbWFpbCcpKTtcbmNvbnN0IFJlY292ZXJQYXNzd29yZCA9IFJlYWN0LmxhenkoKCkgPT4gaW1wb3J0KCdAL2FwcC92aWV3cy9SZWNvdmVyUGFzc3dvcmQnKSk7XG5cbmV4cG9ydCBpbnRlcmZhY2UgUm91dGVJdGVtIHtcbiAgICBwYXRoOiBzdHJpbmcsXG4gICAgY29tcG9uZW50OiBSZWFjdC5GQzxhbnk+LFxuICAgIGV4YWN0OiBib29sZWFuLFxuICAgIGNoaWxkcmVuPzogQ29tcG9uZW50Um91dGVzW10sXG4gICAgd2l0aD86IChcbiAgICAgICAgY2hpbGQ6IENvbXBvbmVudFJvdXRlcyxcbiAgICAgICAgcGFycmVudDogQ29tcG9uZW50Um91dGVzXG4gICAgKSA9PiBDb21wb25lbnRSb3V0ZXMsXG4gICAgYWRkQ2hpbGRyZW4/OiAoY2hpbGRyZW46IENvbXBvbmVudFJvdXRlc1tdKSA9PiBDb21wb25lbnRSb3V0ZXNcbn1cblxuLyoqIFJvdXRlIGJhc2UgY29uZmlnIGltcGxlbWVudGF0aW9uICovXG5leHBvcnQgY2xhc3MgQ29tcG9uZW50Um91dGVzIGltcGxlbWVudHMgUm91dGVJdGVtIHtcbiAgICAvKiogZGF0YSByb3V0ZSBjb25maWcqL1xuICAgIGNvbnN0cnVjdG9yKFxuICAgICAgICBwdWJsaWMgcGF0aDogc3RyaW5nLFxuICAgICAgICBwdWJsaWMgY29tcG9uZW50OiBSZWFjdC5GQyB8XG4gICAgICAgICAgICBMYXp5RXhvdGljQ29tcG9uZW50PFJlYWN0LkZDPHsgY2hpbGRyZW46IENvbXBvbmVudFJvdXRlc1tdIH0+PixcbiAgICAgICAgcHVibGljIGV4YWN0OiBib29sZWFuLFxuICAgICAgICBwdWJsaWMgY2hpbGRyZW4/OiBDb21wb25lbnRSb3V0ZXNbXSxcbiAgICApIHsgfTtcbiAgICAvKiBjaGFuZ2UgcGF0aCBmb3IgY2hpbGRyZW4gcm91dGVzICovXG4gICAgcHVibGljIHdpdGgoXG4gICAgICAgIGNoaWxkOiBDb21wb25lbnRSb3V0ZXMsXG4gICAgICAgIHBhcnJlbnQ6IENvbXBvbmVudFJvdXRlc1xuICAgICk6IENvbXBvbmVudFJvdXRlcyB7XG4gICAgICAgIGNoaWxkLnBhdGggPSBgJHtwYXJyZW50LnBhdGh9LyR7Y2hpbGQucGF0aH1gO1xuXG4gICAgICAgIHJldHVybiB0aGlzO1xuICAgIH07XG4gICAgLyogYWRkcyBjaGlsZHJlbiByb3V0ZXMgdG8gcm91dGUgKi9cbiAgICBwdWJsaWMgYWRkQ2hpbGRyZW4oY2hpbGRyZW46IENvbXBvbmVudFJvdXRlc1tdKTogQ29tcG9uZW50Um91dGVzIHtcbiAgICAgICAgdGhpcy5jaGlsZHJlbiA9IGNoaWxkcmVuLm1hcCgoY2hpbGQpID0+IGNoaWxkLndpdGgoY2hpbGQsIHRoaXMpKTtcblxuICAgICAgICByZXR1cm4gdGhpcztcbiAgICB9O1xufTtcblxuLyoqIFJvdXRlIGNvbmZpZyBpbXBsZW1lbnRhdGlvbiAqL1xuZXhwb3J0IGNsYXNzIFJvdXRlQ29uZmlnIHtcbiAgICBwdWJsaWMgc3RhdGljIFdlbGNvbWVQYWdlOiBDb21wb25lbnRSb3V0ZXMgPSBuZXcgQ29tcG9uZW50Um91dGVzKFxuICAgICAgICAnLycsXG4gICAgICAgIFdlbGNvbWVQYWdlLFxuICAgICAgICB0cnVlLFxuICAgICk7XG4gICAgcHVibGljIHN0YXRpYyBTaWduSW46IENvbXBvbmVudFJvdXRlcyA9IG5ldyBDb21wb25lbnRSb3V0ZXMoXG4gICAgICAgICcvc2lnbi1pbicsXG4gICAgICAgIFNpZ25JbixcbiAgICAgICAgdHJ1ZVxuICAgICk7XG4gICAgcHVibGljIHN0YXRpYyBTaWduVXA6IENvbXBvbmVudFJvdXRlcyA9IG5ldyBDb21wb25lbnRSb3V0ZXMoXG4gICAgICAgICcvc2lnbi11cCcsXG4gICAgICAgIFNpZ25VcCxcbiAgICAgICAgdHJ1ZVxuICAgICk7XG4gICAgcHVibGljIHN0YXRpYyBSZXNldFBhc3N3b3JkOiBDb21wb25lbnRSb3V0ZXMgPSBuZXcgQ29tcG9uZW50Um91dGVzKFxuICAgICAgICAnL2NoYW5nZS1wYXNzd29yZCcsXG4gICAgICAgIENoYW5nZVBhc3N3b3JkLFxuICAgICAgICB0cnVlXG4gICAgKTtcbiAgICBwdWJsaWMgc3RhdGljIENvbmZpcm1FbWFpbDogQ29tcG9uZW50Um91dGVzID0gbmV3IENvbXBvbmVudFJvdXRlcyhcbiAgICAgICAgJy9lbWFpbC9jb25maXJtLycsXG4gICAgICAgIENvbmZpcm1FbWFpbCxcbiAgICAgICAgdHJ1ZSxcbiAgICApO1xuICAgIHB1YmxpYyBzdGF0aWMgUmVjb3ZlclBhc3N3b3JkOiBDb21wb25lbnRSb3V0ZXMgPSBuZXcgQ29tcG9uZW50Um91dGVzKFxuICAgICAgICAnL3JlY292ZXItcGFzc3dvcmQnLFxuICAgICAgICBSZWNvdmVyUGFzc3dvcmQsXG4gICAgICAgIHRydWUsXG4gICAgKTtcbiAgICBwdWJsaWMgc3RhdGljIHJvdXRlczogQ29tcG9uZW50Um91dGVzW10gPSBbXG4gICAgICAgIFJvdXRlQ29uZmlnLldlbGNvbWVQYWdlLFxuICAgICAgICBSb3V0ZUNvbmZpZy5TaWduSW4sXG4gICAgICAgIFJvdXRlQ29uZmlnLlNpZ25VcCxcbiAgICAgICAgUm91dGVDb25maWcuUmVzZXRQYXNzd29yZCxcbiAgICAgICAgUm91dGVDb25maWcuQ29uZmlybUVtYWlsLFxuICAgICAgICBSb3V0ZUNvbmZpZy5SZWNvdmVyUGFzc3dvcmQsXG4gICAgXTtcbn07XG5cbmV4cG9ydCBjb25zdCBSb3V0ZTogUmVhY3QuRkM8Um91dGVJdGVtPiA9ICh7XG4gICAgY29tcG9uZW50OiBDb21wb25lbnQsIC4uLmNoaWxkcmVuXG59KSA9PlxuICAgIDxDb21wb25lbnQgey4uLmNoaWxkcmVufSAvPjtcblxuZXhwb3J0IGNvbnN0IFJvdXRlcyA9ICgpID0+IHtcbiAgICByZXR1cm4gKFxuICAgICAgICA8U3dpdGNoPlxuICAgICAgICAgICAge1xuICAgICAgICAgICAgICAgIFJvdXRlQ29uZmlnLnJvdXRlcy5tYXAoKHJvdXRlLCBpbmRleCkgPT4gKFxuICAgICAgICAgICAgICAgICAgICA8Um91dGVcbiAgICAgICAgICAgICAgICAgICAgICAgIGtleT17aW5kZXh9XG4gICAgICAgICAgICAgICAgICAgICAgICBwYXRoPXtyb3V0ZS5wYXRofVxuICAgICAgICAgICAgICAgICAgICAgICAgY29tcG9uZW50PXtyb3V0ZS5jb21wb25lbnR9XG4gICAgICAgICAgICAgICAgICAgICAgICBleGFjdD17cm91dGUuZXhhY3R9XG4gICAgICAgICAgICAgICAgICAgICAgICBjaGlsZHJlbj17cm91dGUuY2hpbGRyZW59XG4gICAgICAgICAgICAgICAgICAgIC8+XG4gICAgICAgICAgICAgICAgKVxuICAgICAgICAgICAgICAgIClcbiAgICAgICAgICAgIH1cbiAgICAgICAgPC9Td2l0Y2g+XG4gICAgKTtcbn07XG4iLCIvLyBDb3B5cmlnaHQgKEMpIDIwMjEgQ3JlZGl0b3IgQ29ycC4gR3JvdXAuXG4vLyBTZWUgTElDRU5TRSBmb3IgY29weWluZyBpbmZvcm1hdGlvbi5cblxuaW1wb3J0IHsgRGlzcGF0Y2ggfSBmcm9tICdyZWR1eCc7XG5cbmltcG9ydCB7IFJvdXRlQ29uZmlnIH0gZnJvbSAnQC9hcHAvcm91dGVyJztcblxuaW1wb3J0IHsgVXNlciB9IGZyb20gJ0AvdXNlcic7XG5pbXBvcnQgeyBVc2VyQ2xpZW50IH0gZnJvbSAnQC9hcGkvdXNlcic7XG5pbXBvcnQgeyBVc2VyU2VydmljZSB9IGZyb20gJ0AvdXNlci9zZXJ2aWNlJztcblxuLyoqIGFjdGlvbiB0eXBlcyBpbXBsZW1lbnRhdGlvbiAqL1xuZXhwb3J0IGNvbnN0IFJFR0lTVEVSID0gJ1JFR0lTVEVSJztcbmV4cG9ydCBjb25zdCBMT0dJTiA9ICdMT0dJTic7XG5leHBvcnQgY29uc3QgQ0hBTkdFX1BBU1NXT1JEID0gJ0NIQU5HRV9QQVNTV09SRCc7XG5leHBvcnQgY29uc3QgQ09ORklSTV9FTUFJTCA9ICdDT05GSVJNX0VNQUlMJztcbmV4cG9ydCBjb25zdCBSRUNPVkVSX1BBU1NXT1JEID0gJ1JFQ09WRVJfUEFTU1dPUkQnO1xuLyoqIGltcGxlbWVudCByZWdpc3RyYXRpb24gb2YgbmV3IHVzZXIgKi9cbmV4cG9ydCBjb25zdCByZWdpc3RlciA9ICh1c2VyOiBVc2VyKSA9PiAoe1xuICAgIHR5cGU6IFJFR0lTVEVSLFxuICAgIHVzZXIsXG59KTtcbi8qKiBnZXQgcmVnaXN0cmVkIHVzZXIgYnkgaWQgKi9cbmV4cG9ydCBjb25zdCBsb2dpbiA9IChlbWFpbDogc3RyaW5nLCBwYXNzd29yZDogc3RyaW5nKSA9PiAoe1xuICAgIHR5cGU6IExPR0lOLFxuICAgIHVzZXI6IHtcbiAgICAgICAgZW1haWwsXG4gICAgICAgIHBhc3N3b3JkLFxuICAgIH1cbn0pO1xuLyoqIGNoYW5naW5nIHVzZXIgcGFzc3dvcmQgKi9cbmV4cG9ydCBjb25zdCBjaGFuZ2VQYXNzd29yZCA9IChwYXNzd29yZDogc3RyaW5nLCBuZXdQYXNzd29yZDogc3RyaW5nKSA9PiAoe1xuICAgIHR5cGU6IENIQU5HRV9QQVNTV09SRCxcbiAgICBwYXNzd29yZHM6IHtcbiAgICAgICAgcGFzc3dvcmQsXG4gICAgICAgIG5ld1Bhc3N3b3JkLFxuICAgIH1cbn0pO1xuLyoqIHVzZXIgZW1haWwgY29uZmlybSAqL1xuZXhwb3J0IGNvbnN0IGNvbmZpcm1FbWFpbCA9ICh0b2tlbjogc3RyaW5nIHwgbnVsbCkgPT4gKHtcbiAgICB0eXBlOiBDT05GSVJNX0VNQUlMLFxuICAgIHRva2VuLFxufSk7XG4vKiogcmVjb3ZlciB1c2VyIHBhc3N3b3JkICovXG5leHBvcnQgY29uc3QgcmVjb3ZlclBhc3N3b3JkID0gKHBhc3N3b3JkOiBzdHJpbmcpID0+ICh7XG4gICAgdHlwZTogUkVDT1ZFUl9QQVNTV09SRCxcbiAgICBwYXNzd29yZFxufSk7XG5cbmNvbnN0IGNsaWVudCA9IG5ldyBVc2VyQ2xpZW50KCk7XG5jb25zdCB1c2VycyA9IG5ldyBVc2VyU2VydmljZShjbGllbnQpO1xuXG4vKiogdGh1bmsgdGhhdCBpbXBsZW1lbnRzIHVzZXIgcmVnaXN0cmF0aW9uICovXG5leHBvcnQgY29uc3QgcmVnaXN0ZXJVc2VyID0gKHVzZXI6IFVzZXIpID0+XG4gICAgYXN5bmMgZnVuY3Rpb24gKGRpc3BhdGNoOiBEaXNwYXRjaCkge1xuICAgICAgICB0cnkge1xuICAgICAgICAgICAgYXdhaXQgdXNlcnMucmVnaXN0ZXIodXNlcik7XG4gICAgICAgICAgICBkaXNwYXRjaChyZWdpc3Rlcih1c2VyKSk7XG4gICAgICAgICAgICBsb2NhdGlvbi5wYXRobmFtZSA9IFJvdXRlQ29uZmlnLlNpZ25Jbi5wYXRoO1xuICAgICAgICB9IGNhdGNoIChlcnJvcjogYW55KSB7XG4gICAgICAgICAgICAvLyBUT0RPOiByZXdvcmsgY2F0Y2hpbmcgZXJyb3JzXG4gICAgICAgICAgICAvKiBlc2xpbnQtZGlzYWJsZSAqL1xuICAgICAgICAgICAgY29uc29sZS5sb2coZXJyb3IubWVzc2FnZSk7XG4gICAgICAgIH07XG4gICAgfTtcblxuLyoqIHRodW5rIHRoYXQgaW1wbGVtZW50cyB1c2VyIGxvZ2luICovXG5leHBvcnQgY29uc3QgbG9naW5Vc2VyID0gKGVtYWlsOiBzdHJpbmcsIHBhc3N3b3JkOiBzdHJpbmcpID0+XG4gICAgYXN5bmMgZnVuY3Rpb24gKGRpc3BhdGNoOiBEaXNwYXRjaCkge1xuICAgICAgICBjb25zdCB3aGl0ZXBhcGVyUGF0aCA9ICcvd2hpdGVwYXBlcic7XG4gICAgICAgIHRyeSB7XG4gICAgICAgICAgICBhd2FpdCB1c2Vycy5sb2dpbihlbWFpbCwgcGFzc3dvcmQpO1xuICAgICAgICAgICAgZGlzcGF0Y2gobG9naW4oZW1haWwsIHBhc3N3b3JkKSk7XG4gICAgICAgICAgICBsb2NhdGlvbi5wYXRobmFtZSA9IHdoaXRlcGFwZXJQYXRoO1xuICAgICAgICB9IGNhdGNoIChlcnJvcjogYW55KSB7XG4gICAgICAgICAgICAvLyBUT0RPOiByZXdvcmsgY2F0Y2hpbmcgZXJyb3JzXG4gICAgICAgICAgICAvKiBlc2xpbnQtZGlzYWJsZSAqL1xuICAgICAgICAgICAgY29uc29sZS5sb2coZXJyb3IubWVzc2FnZSk7XG4gICAgICAgIH07XG4gICAgfTtcblxuLyoqIHRodW5rIHRoYXQgaW1wbGVtZW50cyB1c2VyIGNoYW5naW5nIHBhc3N3b3JkICovXG5leHBvcnQgY29uc3QgY2hhbmdlVXNlclBhc3N3b3JkID0gKHBhc3N3b3JkOiBzdHJpbmcsIG5ld1Bhc3N3b3JkOiBzdHJpbmcpID0+XG4gICAgYXN5bmMgZnVuY3Rpb24gKGRpc3BhdGNoOiBEaXNwYXRjaCkge1xuICAgICAgICBjb25zdCBtYXJrZXRwbGFjZVBhdGggPSAnL21hcmtldHBsYWNlJztcbiAgICAgICAgdHJ5IHtcbiAgICAgICAgICAgIGF3YWl0IHVzZXJzLmNoYW5nZVBhc3N3b3JkKHBhc3N3b3JkLCBuZXdQYXNzd29yZCk7XG4gICAgICAgICAgICBkaXNwYXRjaChjaGFuZ2VQYXNzd29yZChwYXNzd29yZCwgbmV3UGFzc3dvcmQpKTtcbiAgICAgICAgICAgIGxvY2F0aW9uLnBhdGhuYW1lID0gbWFya2V0cGxhY2VQYXRoO1xuICAgICAgICB9IGNhdGNoIChlcnJvcjogYW55KSB7XG4gICAgICAgICAgICAvLyBUT0RPOiByZXdvcmsgY2F0Y2hpbmcgZXJyb3JzXG4gICAgICAgICAgICAvKiBlc2xpbnQtZGlzYWJsZSAqL1xuICAgICAgICAgICAgY29uc29sZS5sb2coZXJyb3IubWVzc2FnZSk7XG4gICAgICAgIH07XG4gICAgfTtcblxuLyoqIHRodW5rIHRoYXQgaW1wbGVtZW50cyB1c2VyIGVtYWlsIGNvbmZpcm0gKi9cbmV4cG9ydCBjb25zdCBjb25maXJtVXNlckVtYWlsID0gKHRva2VuOiBzdHJpbmcgfCBudWxsKSA9PlxuICAgIGFzeW5jIGZ1bmN0aW9uIChkaXNwYXRjaDogRGlzcGF0Y2gpIHtcbiAgICAgICAgdHJ5IHtcbiAgICAgICAgICAgIGF3YWl0IHVzZXJzLmNvbmZpcm1FbWFpbCh0b2tlbik7XG4gICAgICAgICAgICBkaXNwYXRjaChjb25maXJtRW1haWwodG9rZW4pKTtcbiAgICAgICAgfSBjYXRjaCAoZXJyb3I6IGFueSkge1xuICAgICAgICAgICAgLyoqIFRPRE86IHJld29yayBjYXRjaGluZyBlcnJyb3MgKi9cbiAgICAgICAgICAgIC8qIGVzbGludC1kaXNhYmxlICovXG4gICAgICAgICAgICBjb25zb2xlLmxvZyhlcnJvci5tZXNzYWdlKTtcbiAgICAgICAgfVxuICAgIH07XG5cbi8qKiB0aHVuayB0aGF0IGltcGxlbWVudHMgdXNlciByZXNldCBwYXNzd29yZCAqL1xuZXhwb3J0IGNvbnN0IHJlY292ZXJVc2VyUGFzc3dvcmQgPSAocGFzc3dvcmQ6IHN0cmluZykgPT5cbiAgICBhc3luYyBmdW5jdGlvbiAoZGlzcGF0Y2g6IERpc3BhdGNoKSB7XG4gICAgICAgIHRyeSB7XG4gICAgICAgICAgICBhd2FpdCB1c2Vycy5yZWNvdmVyUGFzc3dvcmQocGFzc3dvcmQpO1xuICAgICAgICAgICAgZGlzcGF0Y2gocmVjb3ZlclBhc3N3b3JkKHBhc3N3b3JkKSk7XG4gICAgICAgIH0gY2F0Y2ggKGVycm9yOiBhbnkpIHtcbiAgICAgICAgICAgIC8qKiBUT0RPOiByZXdvcmsgY2F0Y2hpbmcgZXJycm9zICovXG4gICAgICAgICAgICAvKiBlc2xpbnQtZGlzYWJsZSAqL1xuICAgICAgICAgICAgY29uc29sZS5sb2coZXJyb3IubWVzc2FnZSk7XG4gICAgICAgIH1cbiAgICB9O1xuIiwiLy8gQ29weXJpZ2h0IChDKSAyMDIxIENyZWRpdG9yIENvcnAuIEdyb3VwLlxuLy8gU2VlIExJQ0VOU0UgZm9yIGNvcHlpbmcgaW5mb3JtYXRpb24uXG5cbmltcG9ydCB7IFN1c3BlbnNlIH0gZnJvbSAncmVhY3QnO1xuaW1wb3J0IHsgQnJvd3NlclJvdXRlciB9IGZyb20gJ3JlYWN0LXJvdXRlci1kb20nO1xuXG5pbXBvcnQgeyBSb3V0ZXMgfSBmcm9tICdAL2FwcC9yb3V0ZXInO1xuXG5leHBvcnQgY29uc3QgQXBwID0gKCkgPT4gKFxuICAgIDxCcm93c2VyUm91dGVyIGJhc2VuYW1lPVwiL1wiPlxuICAgICAgICB7LyoqIFRPRE86IExvYWRpbmdQYWdlIG9yIGluZGljYXRvciovfVxuICAgICAgICA8U3VzcGVuc2UgZmFsbGJhY2s9ezxkaXY+TG9hZGluZy4uLjwvZGl2Pn0gPlxuICAgICAgICAgICAgPFJvdXRlcyAvPlxuICAgICAgICA8L1N1c3BlbnNlPlxuICAgIDwvQnJvd3NlclJvdXRlcj5cbik7XG4iLCIvLyBDb3B5cmlnaHQgKEMpIDIwMjEgQ3JlZGl0b3IgQ29ycC4gR3JvdXAuXG4vLyBTZWUgTElDRU5TRSBmb3IgY29weWluZyBpbmZvcm1hdGlvbi5cblxuaW1wb3J0IHsgVXNlckNsaWVudCB9IGZyb20gJ0AvYXBpL3VzZXInO1xuaW1wb3J0IHsgVXNlclNlcnZpY2UgfSBmcm9tICdAL3VzZXIvc2VydmljZSc7XG5cbmltcG9ydCB7XG4gICAgQ0hBTkdFX1BBU1NXT1JELFxuICAgIENPTkZJUk1fRU1BSUwsXG4gICAgTE9HSU4sXG4gICAgUkVDT1ZFUl9QQVNTV09SRFxufSBmcm9tICcuLi9hY3Rpb25zL3VzZXJzJztcblxuLyoqIGltcGxlbWVudGF0aW9uIG9mIHVzZXIgc3RhdGUgKi9cbmV4cG9ydCBjbGFzcyBVc2Vyc1N0YXRlIHtcbiAgICBwdWJsaWMgcmVhZG9ubHkgdXNlclNlcnZpY2U6IFVzZXJTZXJ2aWNlO1xuICAgIHB1YmxpYyB1c2VyID0ge1xuICAgICAgICBlbWFpbDogJycsXG4gICAgICAgIHBhc3N3b3JkOiAnJyxcbiAgICAgICAgc3RhdHVzOiBudWxsLFxuICAgIH07XG4gICAgcHVibGljIGNvbnN0cnVjdG9yKHVzZXJTZXJ2aWNlOiBVc2VyU2VydmljZSkge1xuICAgICAgICB0aGlzLnVzZXJTZXJ2aWNlID0gdXNlclNlcnZpY2U7XG4gICAgfTtcbn07XG5cbmNvbnN0IGNsaWVudCA9IG5ldyBVc2VyQ2xpZW50KCk7XG5jb25zdCBzZXJ2aWNlID0gbmV3IFVzZXJTZXJ2aWNlKGNsaWVudCk7XG5cbmV4cG9ydCBjb25zdCB1c2Vyc1JlZHVjZXIgPSAoXG4gICAgc3RhdGUgPSBuZXcgVXNlcnNTdGF0ZShzZXJ2aWNlKSxcbiAgICBhY3Rpb246IGFueSA9IHt9XG4pID0+IHtcbiAgICBjb25zdCB1c2VyID0gc3RhdGUudXNlcjtcblxuICAgIHN3aXRjaCAoYWN0aW9uLnR5cGUpIHtcbiAgICBjYXNlIExPR0lOOlxuICAgICAgICB1c2VyLmVtYWlsID0gYWN0aW9uLnVzZXIuZW1haWw7XG4gICAgICAgIHVzZXIucGFzc3dvcmQgPSBhY3Rpb24udXNlci5wYXNzd29yZDtcbiAgICAgICAgYnJlYWs7XG4gICAgY2FzZSBDSEFOR0VfUEFTU1dPUkQ6XG4gICAgICAgIHVzZXIucGFzc3dvcmQgPSBhY3Rpb24ucGFzc3dvcmRzLm5ld1Bhc3N3b3JkO1xuICAgICAgICBicmVhaztcbiAgICBjYXNlIENPTkZJUk1fRU1BSUw6XG4gICAgICAgIHVzZXIuc3RhdHVzID0gYWN0aW9uLnRva2VuO1xuICAgICAgICBicmVhaztcbiAgICBjYXNlIFJFQ09WRVJfUEFTU1dPUkQ6XG4gICAgICAgIHVzZXIucGFzc3dvcmQgPSBhY3Rpb24ucGFzc3dvcmQ7XG4gICAgICAgIGJyZWFrO1xuICAgIGRlZmF1bHQ6XG4gICAgICAgIGJyZWFrO1xuICAgIH07XG5cbiAgICByZXR1cm4geyAuLi5zdGF0ZSB9O1xufTtcbiIsIi8vIENvcHlyaWdodCAoQykgMjAyMSBDcmVkaXRvciBDb3JwLiBHcm91cC5cbi8vIFNlZSBMSUNFTlNFIGZvciBjb3B5aW5nIGluZm9ybWF0aW9uLlxuXG5pbXBvcnQgeyBhcHBseU1pZGRsZXdhcmUsIGNyZWF0ZVN0b3JlLCBjb21iaW5lUmVkdWNlcnMgfSBmcm9tICdyZWR1eCc7XG5pbXBvcnQgdGh1bmtNaWRkbGV3YXJlIGZyb20gJ3JlZHV4LXRodW5rJztcblxuaW1wb3J0IHsgdXNlcnNSZWR1Y2VyIH0gZnJvbSAnLi9yZWR1Y2Vycy91c2Vycyc7XG5cbmNvbnN0IHJlZHVjZXIgPSBjb21iaW5lUmVkdWNlcnMoe1xuICAgIHVzZXJzOiB1c2Vyc1JlZHVjZXIsXG59KTtcblxuZXhwb3J0IGNvbnN0IHN0b3JlID0gY3JlYXRlU3RvcmUocmVkdWNlciwgYXBwbHlNaWRkbGV3YXJlKHRodW5rTWlkZGxld2FyZSkpO1xuXG5leHBvcnQgdHlwZSBSb290U3RhdGUgPSBSZXR1cm5UeXBlPHR5cGVvZiBzdG9yZS5nZXRTdGF0ZT47IiwiXG4gICAgICBpbXBvcnQgQVBJIGZyb20gXCIhLi4vbm9kZV9tb2R1bGVzL3N0eWxlLWxvYWRlci9kaXN0L3J1bnRpbWUvaW5qZWN0U3R5bGVzSW50b1N0eWxlVGFnLmpzXCI7XG4gICAgICBpbXBvcnQgZG9tQVBJIGZyb20gXCIhLi4vbm9kZV9tb2R1bGVzL3N0eWxlLWxvYWRlci9kaXN0L3J1bnRpbWUvc3R5bGVEb21BUEkuanNcIjtcbiAgICAgIGltcG9ydCBpbnNlcnRGbiBmcm9tIFwiIS4uL25vZGVfbW9kdWxlcy9zdHlsZS1sb2FkZXIvZGlzdC9ydW50aW1lL2luc2VydEJ5U2VsZWN0b3IuanNcIjtcbiAgICAgIGltcG9ydCBzZXRBdHRyaWJ1dGVzIGZyb20gXCIhLi4vbm9kZV9tb2R1bGVzL3N0eWxlLWxvYWRlci9kaXN0L3J1bnRpbWUvc2V0QXR0cmlidXRlc1dpdGhvdXRBdHRyaWJ1dGVzLmpzXCI7XG4gICAgICBpbXBvcnQgaW5zZXJ0U3R5bGVFbGVtZW50IGZyb20gXCIhLi4vbm9kZV9tb2R1bGVzL3N0eWxlLWxvYWRlci9kaXN0L3J1bnRpbWUvaW5zZXJ0U3R5bGVFbGVtZW50LmpzXCI7XG4gICAgICBpbXBvcnQgc3R5bGVUYWdUcmFuc2Zvcm1GbiBmcm9tIFwiIS4uL25vZGVfbW9kdWxlcy9zdHlsZS1sb2FkZXIvZGlzdC9ydW50aW1lL3N0eWxlVGFnVHJhbnNmb3JtLmpzXCI7XG4gICAgICBpbXBvcnQgY29udGVudCwgKiBhcyBuYW1lZEV4cG9ydCBmcm9tIFwiISEuLi9ub2RlX21vZHVsZXMvY3NzLWxvYWRlci9kaXN0L2Nqcy5qcyEuLi9ub2RlX21vZHVsZXMvc2Fzcy1sb2FkZXIvZGlzdC9janMuanMhLi9pbmRleC5zY3NzXCI7XG4gICAgICBcbiAgICAgIFxuXG52YXIgb3B0aW9ucyA9IHt9O1xuXG5vcHRpb25zLnN0eWxlVGFnVHJhbnNmb3JtID0gc3R5bGVUYWdUcmFuc2Zvcm1Gbjtcbm9wdGlvbnMuc2V0QXR0cmlidXRlcyA9IHNldEF0dHJpYnV0ZXM7XG5cbiAgICAgIG9wdGlvbnMuaW5zZXJ0ID0gaW5zZXJ0Rm4uYmluZChudWxsLCBcImhlYWRcIik7XG4gICAgXG5vcHRpb25zLmRvbUFQSSA9IGRvbUFQSTtcbm9wdGlvbnMuaW5zZXJ0U3R5bGVFbGVtZW50ID0gaW5zZXJ0U3R5bGVFbGVtZW50O1xuXG52YXIgdXBkYXRlID0gQVBJKGNvbnRlbnQsIG9wdGlvbnMpO1xuXG5cblxuZXhwb3J0ICogZnJvbSBcIiEhLi4vbm9kZV9tb2R1bGVzL2Nzcy1sb2FkZXIvZGlzdC9janMuanMhLi4vbm9kZV9tb2R1bGVzL3Nhc3MtbG9hZGVyL2Rpc3QvY2pzLmpzIS4vaW5kZXguc2Nzc1wiO1xuICAgICAgIGV4cG9ydCBkZWZhdWx0IGNvbnRlbnQgJiYgY29udGVudC5sb2NhbHMgPyBjb250ZW50LmxvY2FscyA6IHVuZGVmaW5lZDtcbiIsIi8vIENvcHlyaWdodCAoQykgMjAyMSBDcmVkaXRvciBDb3JwLiBHcm91cC5cbi8vIFNlZSBMSUNFTlNFIGZvciBjb3B5aW5nIGluZm9ybWF0aW9uLlxuXG5pbXBvcnQgUmVhY3RET00gZnJvbSAncmVhY3QtZG9tJztcbmltcG9ydCB7IFByb3ZpZGVyIH0gZnJvbSAncmVhY3QtcmVkdXgnO1xuXG5pbXBvcnQgeyBBcHAgfSBmcm9tICdAL0FwcCc7XG5cbmltcG9ydCB7IHN0b3JlIH0gZnJvbSAnQC9hcHAvc3RvcmUnO1xuXG5pbXBvcnQgJy4vaW5kZXguc2Nzcyc7XG5cblJlYWN0RE9NLnJlbmRlcihcbiAgICA8UHJvdmlkZXIgc3RvcmU9e3N0b3JlfT5cbiAgICAgICAgPEFwcCAvPlxuICAgIDwvUHJvdmlkZXI+LFxuICAgIGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKCdyb290JyksXG4pO1xuIiwiLy8gQ29weXJpZ2h0IChDKSAyMDIxIENyZWRpdG9yIENvcnAuIEdyb3VwLlxuLy8gU2VlIExJQ0VOU0UgZm9yIGNvcHlpbmcgaW5mb3JtYXRpb24uXG5cbmltcG9ydCB7IFVzZXIgfSBmcm9tICcuJztcbmltcG9ydCB7IFVzZXJDbGllbnQgfSBmcm9tICdAL2FwaS91c2VyJztcblxuLyoqIGV4cG9zZXMgYWxsIHVzZXIgcmVsYXRlZCBsb2dpYyAgKi9cbmV4cG9ydCBjbGFzcyBVc2VyU2VydmljZSB7XG4gICAgcHJpdmF0ZSByZWFkb25seSB1c2VyczogVXNlckNsaWVudDtcbiAgICBwdWJsaWMgY29uc3RydWN0b3IodXNlcnM6IFVzZXJDbGllbnQpIHtcbiAgICAgICAgdGhpcy51c2VycyA9IHVzZXJzO1xuICAgIH07XG4gICAgLyoqIGhhbmRsZXMgdXNlciByZWdpc3RyYXRpb24gKi9cbiAgICBwdWJsaWMgYXN5bmMgcmVnaXN0ZXIodXNlcjogVXNlcik6IFByb21pc2U8dm9pZD4ge1xuICAgICAgICByZXR1cm4gYXdhaXQgdGhpcy51c2Vycy5yZWdpc3Rlcih1c2VyKTtcbiAgICB9O1xuICAgIC8qKiByZXR1cm4gcmVnaXN0cmVkIHVzZXIgKi9cbiAgICBwdWJsaWMgYXN5bmMgbG9naW4oZW1haWw6IHN0cmluZywgcGFzc3dvcmQ6IHN0cmluZyk6IFByb21pc2U8dm9pZD4ge1xuICAgICAgICByZXR1cm4gYXdhaXQgdGhpcy51c2Vycy5sb2dpbihlbWFpbCwgcGFzc3dvcmQpO1xuICAgIH07XG4gICAgLyoqIGhhbmRsZXMgdXNlciBjaGFuZ2luZyBwYXNzd29yZCAqL1xuICAgIHB1YmxpYyBhc3luYyBjaGFuZ2VQYXNzd29yZChwYXNzd29yZDogc3RyaW5nLCBuZXdQYXNzd29yZDogc3RyaW5nKTogUHJvbWlzZTx2b2lkPiB7XG4gICAgICAgIHJldHVybiBhd2FpdCB0aGlzLnVzZXJzLmNoYW5nZVBhc3N3b3JkKHBhc3N3b3JkLCBuZXdQYXNzd29yZCk7XG4gICAgfTtcbiAgICAvKiogaGFuZGxlcyB1c2VyIGVtYWlsIGNvbmZpcm1hdGlvbiAqL1xuICAgIHB1YmxpYyBhc3luYyBjb25maXJtRW1haWwodG9rZW46IHN0cmluZyB8IG51bGwpOiBQcm9taXNlPFJlc3BvbnNlPiB7XG4gICAgICAgIHJldHVybiBhd2FpdCB0aGlzLnVzZXJzLmNvbmZpcm1FbWFpbCh0b2tlbik7XG4gICAgfTtcbiAgICAvKiogaGFuZGxlcyB1c2VyIGNoZWNrIHRva2VuICovXG4gICAgcHVibGljIGFzeW5jIGNoZWNrVG9rZW4odG9rZW46IHN0cmluZyB8IG51bGwpOiBQcm9taXNlPFJlc3BvbnNlPiB7XG4gICAgICAgIHJldHVybiBhd2FpdCB0aGlzLnVzZXJzLmNoZWNrVG9rZW4odG9rZW4pO1xuICAgIH07XG4gICAgLyoqIGhhbmRsZXMgdXNlciByZWNvdmVyIHBhc3N3b3JkICovXG4gICAgcHVibGljIGFzeW5jIHJlY292ZXJQYXNzd29yZChwYXNzd29yZDogc3RyaW5nKTogUHJvbWlzZTxSZXNwb25zZT4ge1xuICAgICAgICByZXR1cm4gYXdhaXQgdGhpcy51c2Vycy5yZWNvdmVyUGFzc3dvcmQocGFzc3dvcmQpO1xuICAgIH07XG59O1xuIiwiLy8gVGhlIG1vZHVsZSBjYWNoZVxudmFyIF9fd2VicGFja19tb2R1bGVfY2FjaGVfXyA9IHt9O1xuXG4vLyBUaGUgcmVxdWlyZSBmdW5jdGlvblxuZnVuY3Rpb24gX193ZWJwYWNrX3JlcXVpcmVfXyhtb2R1bGVJZCkge1xuXHQvLyBDaGVjayBpZiBtb2R1bGUgaXMgaW4gY2FjaGVcblx0dmFyIGNhY2hlZE1vZHVsZSA9IF9fd2VicGFja19tb2R1bGVfY2FjaGVfX1ttb2R1bGVJZF07XG5cdGlmIChjYWNoZWRNb2R1bGUgIT09IHVuZGVmaW5lZCkge1xuXHRcdHJldHVybiBjYWNoZWRNb2R1bGUuZXhwb3J0cztcblx0fVxuXHQvLyBDcmVhdGUgYSBuZXcgbW9kdWxlIChhbmQgcHV0IGl0IGludG8gdGhlIGNhY2hlKVxuXHR2YXIgbW9kdWxlID0gX193ZWJwYWNrX21vZHVsZV9jYWNoZV9fW21vZHVsZUlkXSA9IHtcblx0XHRpZDogbW9kdWxlSWQsXG5cdFx0Ly8gbm8gbW9kdWxlLmxvYWRlZCBuZWVkZWRcblx0XHRleHBvcnRzOiB7fVxuXHR9O1xuXG5cdC8vIEV4ZWN1dGUgdGhlIG1vZHVsZSBmdW5jdGlvblxuXHRfX3dlYnBhY2tfbW9kdWxlc19fW21vZHVsZUlkXS5jYWxsKG1vZHVsZS5leHBvcnRzLCBtb2R1bGUsIG1vZHVsZS5leHBvcnRzLCBfX3dlYnBhY2tfcmVxdWlyZV9fKTtcblxuXHQvLyBSZXR1cm4gdGhlIGV4cG9ydHMgb2YgdGhlIG1vZHVsZVxuXHRyZXR1cm4gbW9kdWxlLmV4cG9ydHM7XG59XG5cbi8vIGV4cG9zZSB0aGUgbW9kdWxlcyBvYmplY3QgKF9fd2VicGFja19tb2R1bGVzX18pXG5fX3dlYnBhY2tfcmVxdWlyZV9fLm0gPSBfX3dlYnBhY2tfbW9kdWxlc19fO1xuXG4iLCJ2YXIgZGVmZXJyZWQgPSBbXTtcbl9fd2VicGFja19yZXF1aXJlX18uTyA9IChyZXN1bHQsIGNodW5rSWRzLCBmbiwgcHJpb3JpdHkpID0+IHtcblx0aWYoY2h1bmtJZHMpIHtcblx0XHRwcmlvcml0eSA9IHByaW9yaXR5IHx8IDA7XG5cdFx0Zm9yKHZhciBpID0gZGVmZXJyZWQubGVuZ3RoOyBpID4gMCAmJiBkZWZlcnJlZFtpIC0gMV1bMl0gPiBwcmlvcml0eTsgaS0tKSBkZWZlcnJlZFtpXSA9IGRlZmVycmVkW2kgLSAxXTtcblx0XHRkZWZlcnJlZFtpXSA9IFtjaHVua0lkcywgZm4sIHByaW9yaXR5XTtcblx0XHRyZXR1cm47XG5cdH1cblx0dmFyIG5vdEZ1bGZpbGxlZCA9IEluZmluaXR5O1xuXHRmb3IgKHZhciBpID0gMDsgaSA8IGRlZmVycmVkLmxlbmd0aDsgaSsrKSB7XG5cdFx0dmFyIFtjaHVua0lkcywgZm4sIHByaW9yaXR5XSA9IGRlZmVycmVkW2ldO1xuXHRcdHZhciBmdWxmaWxsZWQgPSB0cnVlO1xuXHRcdGZvciAodmFyIGogPSAwOyBqIDwgY2h1bmtJZHMubGVuZ3RoOyBqKyspIHtcblx0XHRcdGlmICgocHJpb3JpdHkgJiAxID09PSAwIHx8IG5vdEZ1bGZpbGxlZCA+PSBwcmlvcml0eSkgJiYgT2JqZWN0LmtleXMoX193ZWJwYWNrX3JlcXVpcmVfXy5PKS5ldmVyeSgoa2V5KSA9PiAoX193ZWJwYWNrX3JlcXVpcmVfXy5PW2tleV0oY2h1bmtJZHNbal0pKSkpIHtcblx0XHRcdFx0Y2h1bmtJZHMuc3BsaWNlKGotLSwgMSk7XG5cdFx0XHR9IGVsc2Uge1xuXHRcdFx0XHRmdWxmaWxsZWQgPSBmYWxzZTtcblx0XHRcdFx0aWYocHJpb3JpdHkgPCBub3RGdWxmaWxsZWQpIG5vdEZ1bGZpbGxlZCA9IHByaW9yaXR5O1xuXHRcdFx0fVxuXHRcdH1cblx0XHRpZihmdWxmaWxsZWQpIHtcblx0XHRcdGRlZmVycmVkLnNwbGljZShpLS0sIDEpXG5cdFx0XHR2YXIgciA9IGZuKCk7XG5cdFx0XHRpZiAociAhPT0gdW5kZWZpbmVkKSByZXN1bHQgPSByO1xuXHRcdH1cblx0fVxuXHRyZXR1cm4gcmVzdWx0O1xufTsiLCIvLyBnZXREZWZhdWx0RXhwb3J0IGZ1bmN0aW9uIGZvciBjb21wYXRpYmlsaXR5IHdpdGggbm9uLWhhcm1vbnkgbW9kdWxlc1xuX193ZWJwYWNrX3JlcXVpcmVfXy5uID0gKG1vZHVsZSkgPT4ge1xuXHR2YXIgZ2V0dGVyID0gbW9kdWxlICYmIG1vZHVsZS5fX2VzTW9kdWxlID9cblx0XHQoKSA9PiAobW9kdWxlWydkZWZhdWx0J10pIDpcblx0XHQoKSA9PiAobW9kdWxlKTtcblx0X193ZWJwYWNrX3JlcXVpcmVfXy5kKGdldHRlciwgeyBhOiBnZXR0ZXIgfSk7XG5cdHJldHVybiBnZXR0ZXI7XG59OyIsIi8vIGRlZmluZSBnZXR0ZXIgZnVuY3Rpb25zIGZvciBoYXJtb255IGV4cG9ydHNcbl9fd2VicGFja19yZXF1aXJlX18uZCA9IChleHBvcnRzLCBkZWZpbml0aW9uKSA9PiB7XG5cdGZvcih2YXIga2V5IGluIGRlZmluaXRpb24pIHtcblx0XHRpZihfX3dlYnBhY2tfcmVxdWlyZV9fLm8oZGVmaW5pdGlvbiwga2V5KSAmJiAhX193ZWJwYWNrX3JlcXVpcmVfXy5vKGV4cG9ydHMsIGtleSkpIHtcblx0XHRcdE9iamVjdC5kZWZpbmVQcm9wZXJ0eShleHBvcnRzLCBrZXksIHsgZW51bWVyYWJsZTogdHJ1ZSwgZ2V0OiBkZWZpbml0aW9uW2tleV0gfSk7XG5cdFx0fVxuXHR9XG59OyIsIl9fd2VicGFja19yZXF1aXJlX18uZiA9IHt9O1xuLy8gVGhpcyBmaWxlIGNvbnRhaW5zIG9ubHkgdGhlIGVudHJ5IGNodW5rLlxuLy8gVGhlIGNodW5rIGxvYWRpbmcgZnVuY3Rpb24gZm9yIGFkZGl0aW9uYWwgY2h1bmtzXG5fX3dlYnBhY2tfcmVxdWlyZV9fLmUgPSAoY2h1bmtJZCkgPT4ge1xuXHRyZXR1cm4gUHJvbWlzZS5hbGwoT2JqZWN0LmtleXMoX193ZWJwYWNrX3JlcXVpcmVfXy5mKS5yZWR1Y2UoKHByb21pc2VzLCBrZXkpID0+IHtcblx0XHRfX3dlYnBhY2tfcmVxdWlyZV9fLmZba2V5XShjaHVua0lkLCBwcm9taXNlcyk7XG5cdFx0cmV0dXJuIHByb21pc2VzO1xuXHR9LCBbXSkpO1xufTsiLCIvLyBUaGlzIGZ1bmN0aW9uIGFsbG93IHRvIHJlZmVyZW5jZSBhc3luYyBjaHVua3Ncbl9fd2VicGFja19yZXF1aXJlX18udSA9IChjaHVua0lkKSA9PiB7XG5cdC8vIHJldHVybiB1cmwgZm9yIGZpbGVuYW1lcyBiYXNlZCBvbiB0ZW1wbGF0ZVxuXHRyZXR1cm4gXCJcIiArIGNodW5rSWQgKyBcIi5cIiArIF9fd2VicGFja19yZXF1aXJlX18uaCgpICsgXCIuanNcIjtcbn07IiwiLy8gVGhpcyBmdW5jdGlvbiBhbGxvdyB0byByZWZlcmVuY2UgYWxsIGNodW5rc1xuX193ZWJwYWNrX3JlcXVpcmVfXy5taW5pQ3NzRiA9IChjaHVua0lkKSA9PiB7XG5cdC8vIHJldHVybiB1cmwgZm9yIGZpbGVuYW1lcyBub3QgYmFzZWQgb24gdGVtcGxhdGVcblx0aWYgKGNodW5rSWQgPT09IDM2OSkgcmV0dXJuIFwiMzY5LmNzc1wiO1xuXHRpZiAoY2h1bmtJZCA9PT0gMTc5KSByZXR1cm4gXCJtYWluLmNzc1wiO1xuXHQvLyByZXR1cm4gdXJsIGZvciBmaWxlbmFtZXMgYmFzZWQgb24gdGVtcGxhdGVcblx0cmV0dXJuIFwiXCIgKyBjaHVua0lkICsgXCIuY3NzXCI7XG59OyIsIl9fd2VicGFja19yZXF1aXJlX18uaCA9ICgpID0+IChcIjcwZWJhZDBkOTVmZjI4NzQ4NWU0XCIpIiwiX193ZWJwYWNrX3JlcXVpcmVfXy5nID0gKGZ1bmN0aW9uKCkge1xuXHRpZiAodHlwZW9mIGdsb2JhbFRoaXMgPT09ICdvYmplY3QnKSByZXR1cm4gZ2xvYmFsVGhpcztcblx0dHJ5IHtcblx0XHRyZXR1cm4gdGhpcyB8fCBuZXcgRnVuY3Rpb24oJ3JldHVybiB0aGlzJykoKTtcblx0fSBjYXRjaCAoZSkge1xuXHRcdGlmICh0eXBlb2Ygd2luZG93ID09PSAnb2JqZWN0JykgcmV0dXJuIHdpbmRvdztcblx0fVxufSkoKTsiLCJfX3dlYnBhY2tfcmVxdWlyZV9fLm8gPSAob2JqLCBwcm9wKSA9PiAoT2JqZWN0LnByb3RvdHlwZS5oYXNPd25Qcm9wZXJ0eS5jYWxsKG9iaiwgcHJvcCkpIiwidmFyIGluUHJvZ3Jlc3MgPSB7fTtcbnZhciBkYXRhV2VicGFja1ByZWZpeCA9IFwiY3J5cHRvZm90YmFsbDpcIjtcbi8vIGxvYWRTY3JpcHQgZnVuY3Rpb24gdG8gbG9hZCBhIHNjcmlwdCB2aWEgc2NyaXB0IHRhZ1xuX193ZWJwYWNrX3JlcXVpcmVfXy5sID0gKHVybCwgZG9uZSwga2V5LCBjaHVua0lkKSA9PiB7XG5cdGlmKGluUHJvZ3Jlc3NbdXJsXSkgeyBpblByb2dyZXNzW3VybF0ucHVzaChkb25lKTsgcmV0dXJuOyB9XG5cdHZhciBzY3JpcHQsIG5lZWRBdHRhY2g7XG5cdGlmKGtleSAhPT0gdW5kZWZpbmVkKSB7XG5cdFx0dmFyIHNjcmlwdHMgPSBkb2N1bWVudC5nZXRFbGVtZW50c0J5VGFnTmFtZShcInNjcmlwdFwiKTtcblx0XHRmb3IodmFyIGkgPSAwOyBpIDwgc2NyaXB0cy5sZW5ndGg7IGkrKykge1xuXHRcdFx0dmFyIHMgPSBzY3JpcHRzW2ldO1xuXHRcdFx0aWYocy5nZXRBdHRyaWJ1dGUoXCJzcmNcIikgPT0gdXJsIHx8IHMuZ2V0QXR0cmlidXRlKFwiZGF0YS13ZWJwYWNrXCIpID09IGRhdGFXZWJwYWNrUHJlZml4ICsga2V5KSB7IHNjcmlwdCA9IHM7IGJyZWFrOyB9XG5cdFx0fVxuXHR9XG5cdGlmKCFzY3JpcHQpIHtcblx0XHRuZWVkQXR0YWNoID0gdHJ1ZTtcblx0XHRzY3JpcHQgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdzY3JpcHQnKTtcblxuXHRcdHNjcmlwdC5jaGFyc2V0ID0gJ3V0Zi04Jztcblx0XHRzY3JpcHQudGltZW91dCA9IDEyMDtcblx0XHRpZiAoX193ZWJwYWNrX3JlcXVpcmVfXy5uYykge1xuXHRcdFx0c2NyaXB0LnNldEF0dHJpYnV0ZShcIm5vbmNlXCIsIF9fd2VicGFja19yZXF1aXJlX18ubmMpO1xuXHRcdH1cblx0XHRzY3JpcHQuc2V0QXR0cmlidXRlKFwiZGF0YS13ZWJwYWNrXCIsIGRhdGFXZWJwYWNrUHJlZml4ICsga2V5KTtcblx0XHRzY3JpcHQuc3JjID0gdXJsO1xuXHR9XG5cdGluUHJvZ3Jlc3NbdXJsXSA9IFtkb25lXTtcblx0dmFyIG9uU2NyaXB0Q29tcGxldGUgPSAocHJldiwgZXZlbnQpID0+IHtcblx0XHQvLyBhdm9pZCBtZW0gbGVha3MgaW4gSUUuXG5cdFx0c2NyaXB0Lm9uZXJyb3IgPSBzY3JpcHQub25sb2FkID0gbnVsbDtcblx0XHRjbGVhclRpbWVvdXQodGltZW91dCk7XG5cdFx0dmFyIGRvbmVGbnMgPSBpblByb2dyZXNzW3VybF07XG5cdFx0ZGVsZXRlIGluUHJvZ3Jlc3NbdXJsXTtcblx0XHRzY3JpcHQucGFyZW50Tm9kZSAmJiBzY3JpcHQucGFyZW50Tm9kZS5yZW1vdmVDaGlsZChzY3JpcHQpO1xuXHRcdGRvbmVGbnMgJiYgZG9uZUZucy5mb3JFYWNoKChmbikgPT4gKGZuKGV2ZW50KSkpO1xuXHRcdGlmKHByZXYpIHJldHVybiBwcmV2KGV2ZW50KTtcblx0fVxuXHQ7XG5cdHZhciB0aW1lb3V0ID0gc2V0VGltZW91dChvblNjcmlwdENvbXBsZXRlLmJpbmQobnVsbCwgdW5kZWZpbmVkLCB7IHR5cGU6ICd0aW1lb3V0JywgdGFyZ2V0OiBzY3JpcHQgfSksIDEyMDAwMCk7XG5cdHNjcmlwdC5vbmVycm9yID0gb25TY3JpcHRDb21wbGV0ZS5iaW5kKG51bGwsIHNjcmlwdC5vbmVycm9yKTtcblx0c2NyaXB0Lm9ubG9hZCA9IG9uU2NyaXB0Q29tcGxldGUuYmluZChudWxsLCBzY3JpcHQub25sb2FkKTtcblx0bmVlZEF0dGFjaCAmJiBkb2N1bWVudC5oZWFkLmFwcGVuZENoaWxkKHNjcmlwdCk7XG59OyIsIi8vIGRlZmluZSBfX2VzTW9kdWxlIG9uIGV4cG9ydHNcbl9fd2VicGFja19yZXF1aXJlX18uciA9IChleHBvcnRzKSA9PiB7XG5cdGlmKHR5cGVvZiBTeW1ib2wgIT09ICd1bmRlZmluZWQnICYmIFN5bWJvbC50b1N0cmluZ1RhZykge1xuXHRcdE9iamVjdC5kZWZpbmVQcm9wZXJ0eShleHBvcnRzLCBTeW1ib2wudG9TdHJpbmdUYWcsIHsgdmFsdWU6ICdNb2R1bGUnIH0pO1xuXHR9XG5cdE9iamVjdC5kZWZpbmVQcm9wZXJ0eShleHBvcnRzLCAnX19lc01vZHVsZScsIHsgdmFsdWU6IHRydWUgfSk7XG59OyIsIl9fd2VicGFja19yZXF1aXJlX18ucCA9IFwiL1wiOyIsIl9fd2VicGFja19yZXF1aXJlX18uYiA9IGRvY3VtZW50LmJhc2VVUkkgfHwgc2VsZi5sb2NhdGlvbi5ocmVmO1xuXG4vLyBvYmplY3QgdG8gc3RvcmUgbG9hZGVkIGFuZCBsb2FkaW5nIGNodW5rc1xuLy8gdW5kZWZpbmVkID0gY2h1bmsgbm90IGxvYWRlZCwgbnVsbCA9IGNodW5rIHByZWxvYWRlZC9wcmVmZXRjaGVkXG4vLyBbcmVzb2x2ZSwgcmVqZWN0LCBQcm9taXNlXSA9IGNodW5rIGxvYWRpbmcsIDAgPSBjaHVuayBsb2FkZWRcbnZhciBpbnN0YWxsZWRDaHVua3MgPSB7XG5cdDE3OTogMFxufTtcblxuX193ZWJwYWNrX3JlcXVpcmVfXy5mLmogPSAoY2h1bmtJZCwgcHJvbWlzZXMpID0+IHtcblx0XHQvLyBKU09OUCBjaHVuayBsb2FkaW5nIGZvciBqYXZhc2NyaXB0XG5cdFx0dmFyIGluc3RhbGxlZENodW5rRGF0YSA9IF9fd2VicGFja19yZXF1aXJlX18ubyhpbnN0YWxsZWRDaHVua3MsIGNodW5rSWQpID8gaW5zdGFsbGVkQ2h1bmtzW2NodW5rSWRdIDogdW5kZWZpbmVkO1xuXHRcdGlmKGluc3RhbGxlZENodW5rRGF0YSAhPT0gMCkgeyAvLyAwIG1lYW5zIFwiYWxyZWFkeSBpbnN0YWxsZWRcIi5cblxuXHRcdFx0Ly8gYSBQcm9taXNlIG1lYW5zIFwiY3VycmVudGx5IGxvYWRpbmdcIi5cblx0XHRcdGlmKGluc3RhbGxlZENodW5rRGF0YSkge1xuXHRcdFx0XHRwcm9taXNlcy5wdXNoKGluc3RhbGxlZENodW5rRGF0YVsyXSk7XG5cdFx0XHR9IGVsc2Uge1xuXHRcdFx0XHRpZih0cnVlKSB7IC8vIGFsbCBjaHVua3MgaGF2ZSBKU1xuXHRcdFx0XHRcdC8vIHNldHVwIFByb21pc2UgaW4gY2h1bmsgY2FjaGVcblx0XHRcdFx0XHR2YXIgcHJvbWlzZSA9IG5ldyBQcm9taXNlKChyZXNvbHZlLCByZWplY3QpID0+IChpbnN0YWxsZWRDaHVua0RhdGEgPSBpbnN0YWxsZWRDaHVua3NbY2h1bmtJZF0gPSBbcmVzb2x2ZSwgcmVqZWN0XSkpO1xuXHRcdFx0XHRcdHByb21pc2VzLnB1c2goaW5zdGFsbGVkQ2h1bmtEYXRhWzJdID0gcHJvbWlzZSk7XG5cblx0XHRcdFx0XHQvLyBzdGFydCBjaHVuayBsb2FkaW5nXG5cdFx0XHRcdFx0dmFyIHVybCA9IF9fd2VicGFja19yZXF1aXJlX18ucCArIF9fd2VicGFja19yZXF1aXJlX18udShjaHVua0lkKTtcblx0XHRcdFx0XHQvLyBjcmVhdGUgZXJyb3IgYmVmb3JlIHN0YWNrIHVud291bmQgdG8gZ2V0IHVzZWZ1bCBzdGFja3RyYWNlIGxhdGVyXG5cdFx0XHRcdFx0dmFyIGVycm9yID0gbmV3IEVycm9yKCk7XG5cdFx0XHRcdFx0dmFyIGxvYWRpbmdFbmRlZCA9IChldmVudCkgPT4ge1xuXHRcdFx0XHRcdFx0aWYoX193ZWJwYWNrX3JlcXVpcmVfXy5vKGluc3RhbGxlZENodW5rcywgY2h1bmtJZCkpIHtcblx0XHRcdFx0XHRcdFx0aW5zdGFsbGVkQ2h1bmtEYXRhID0gaW5zdGFsbGVkQ2h1bmtzW2NodW5rSWRdO1xuXHRcdFx0XHRcdFx0XHRpZihpbnN0YWxsZWRDaHVua0RhdGEgIT09IDApIGluc3RhbGxlZENodW5rc1tjaHVua0lkXSA9IHVuZGVmaW5lZDtcblx0XHRcdFx0XHRcdFx0aWYoaW5zdGFsbGVkQ2h1bmtEYXRhKSB7XG5cdFx0XHRcdFx0XHRcdFx0dmFyIGVycm9yVHlwZSA9IGV2ZW50ICYmIChldmVudC50eXBlID09PSAnbG9hZCcgPyAnbWlzc2luZycgOiBldmVudC50eXBlKTtcblx0XHRcdFx0XHRcdFx0XHR2YXIgcmVhbFNyYyA9IGV2ZW50ICYmIGV2ZW50LnRhcmdldCAmJiBldmVudC50YXJnZXQuc3JjO1xuXHRcdFx0XHRcdFx0XHRcdGVycm9yLm1lc3NhZ2UgPSAnTG9hZGluZyBjaHVuayAnICsgY2h1bmtJZCArICcgZmFpbGVkLlxcbignICsgZXJyb3JUeXBlICsgJzogJyArIHJlYWxTcmMgKyAnKSc7XG5cdFx0XHRcdFx0XHRcdFx0ZXJyb3IubmFtZSA9ICdDaHVua0xvYWRFcnJvcic7XG5cdFx0XHRcdFx0XHRcdFx0ZXJyb3IudHlwZSA9IGVycm9yVHlwZTtcblx0XHRcdFx0XHRcdFx0XHRlcnJvci5yZXF1ZXN0ID0gcmVhbFNyYztcblx0XHRcdFx0XHRcdFx0XHRpbnN0YWxsZWRDaHVua0RhdGFbMV0oZXJyb3IpO1xuXHRcdFx0XHRcdFx0XHR9XG5cdFx0XHRcdFx0XHR9XG5cdFx0XHRcdFx0fTtcblx0XHRcdFx0XHRfX3dlYnBhY2tfcmVxdWlyZV9fLmwodXJsLCBsb2FkaW5nRW5kZWQsIFwiY2h1bmstXCIgKyBjaHVua0lkLCBjaHVua0lkKTtcblx0XHRcdFx0fSBlbHNlIGluc3RhbGxlZENodW5rc1tjaHVua0lkXSA9IDA7XG5cdFx0XHR9XG5cdFx0fVxufTtcblxuLy8gbm8gcHJlZmV0Y2hpbmdcblxuLy8gbm8gcHJlbG9hZGVkXG5cbi8vIG5vIEhNUlxuXG4vLyBubyBITVIgbWFuaWZlc3RcblxuX193ZWJwYWNrX3JlcXVpcmVfXy5PLmogPSAoY2h1bmtJZCkgPT4gKGluc3RhbGxlZENodW5rc1tjaHVua0lkXSA9PT0gMCk7XG5cbi8vIGluc3RhbGwgYSBKU09OUCBjYWxsYmFjayBmb3IgY2h1bmsgbG9hZGluZ1xudmFyIHdlYnBhY2tKc29ucENhbGxiYWNrID0gKHBhcmVudENodW5rTG9hZGluZ0Z1bmN0aW9uLCBkYXRhKSA9PiB7XG5cdHZhciBbY2h1bmtJZHMsIG1vcmVNb2R1bGVzLCBydW50aW1lXSA9IGRhdGE7XG5cdC8vIGFkZCBcIm1vcmVNb2R1bGVzXCIgdG8gdGhlIG1vZHVsZXMgb2JqZWN0LFxuXHQvLyB0aGVuIGZsYWcgYWxsIFwiY2h1bmtJZHNcIiBhcyBsb2FkZWQgYW5kIGZpcmUgY2FsbGJhY2tcblx0dmFyIG1vZHVsZUlkLCBjaHVua0lkLCBpID0gMDtcblx0Zm9yKG1vZHVsZUlkIGluIG1vcmVNb2R1bGVzKSB7XG5cdFx0aWYoX193ZWJwYWNrX3JlcXVpcmVfXy5vKG1vcmVNb2R1bGVzLCBtb2R1bGVJZCkpIHtcblx0XHRcdF9fd2VicGFja19yZXF1aXJlX18ubVttb2R1bGVJZF0gPSBtb3JlTW9kdWxlc1ttb2R1bGVJZF07XG5cdFx0fVxuXHR9XG5cdGlmKHJ1bnRpbWUpIHZhciByZXN1bHQgPSBydW50aW1lKF9fd2VicGFja19yZXF1aXJlX18pO1xuXHRpZihwYXJlbnRDaHVua0xvYWRpbmdGdW5jdGlvbikgcGFyZW50Q2h1bmtMb2FkaW5nRnVuY3Rpb24oZGF0YSk7XG5cdGZvcig7aSA8IGNodW5rSWRzLmxlbmd0aDsgaSsrKSB7XG5cdFx0Y2h1bmtJZCA9IGNodW5rSWRzW2ldO1xuXHRcdGlmKF9fd2VicGFja19yZXF1aXJlX18ubyhpbnN0YWxsZWRDaHVua3MsIGNodW5rSWQpICYmIGluc3RhbGxlZENodW5rc1tjaHVua0lkXSkge1xuXHRcdFx0aW5zdGFsbGVkQ2h1bmtzW2NodW5rSWRdWzBdKCk7XG5cdFx0fVxuXHRcdGluc3RhbGxlZENodW5rc1tjaHVua0lkc1tpXV0gPSAwO1xuXHR9XG5cdHJldHVybiBfX3dlYnBhY2tfcmVxdWlyZV9fLk8ocmVzdWx0KTtcbn1cblxudmFyIGNodW5rTG9hZGluZ0dsb2JhbCA9IHNlbGZbXCJ3ZWJwYWNrQ2h1bmtjcnlwdG9mb3RiYWxsXCJdID0gc2VsZltcIndlYnBhY2tDaHVua2NyeXB0b2ZvdGJhbGxcIl0gfHwgW107XG5jaHVua0xvYWRpbmdHbG9iYWwuZm9yRWFjaCh3ZWJwYWNrSnNvbnBDYWxsYmFjay5iaW5kKG51bGwsIDApKTtcbmNodW5rTG9hZGluZ0dsb2JhbC5wdXNoID0gd2VicGFja0pzb25wQ2FsbGJhY2suYmluZChudWxsLCBjaHVua0xvYWRpbmdHbG9iYWwucHVzaC5iaW5kKGNodW5rTG9hZGluZ0dsb2JhbCkpOyIsIi8vIHN0YXJ0dXBcbi8vIExvYWQgZW50cnkgbW9kdWxlIGFuZCByZXR1cm4gZXhwb3J0c1xuLy8gVGhpcyBlbnRyeSBtb2R1bGUgZGVwZW5kcyBvbiBvdGhlciBsb2FkZWQgY2h1bmtzIGFuZCBleGVjdXRpb24gbmVlZCB0byBiZSBkZWxheWVkXG52YXIgX193ZWJwYWNrX2V4cG9ydHNfXyA9IF9fd2VicGFja19yZXF1aXJlX18uTyh1bmRlZmluZWQsIFszNjldLCAoKSA9PiAoX193ZWJwYWNrX3JlcXVpcmVfXyg0NTU1KSkpXG5fX3dlYnBhY2tfZXhwb3J0c19fID0gX193ZWJwYWNrX3JlcXVpcmVfXy5PKF9fd2VicGFja19leHBvcnRzX18pO1xuIl0sInNvdXJjZVJvb3QiOiIifQ==