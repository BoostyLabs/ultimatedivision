(self['webpackChunkcryptofotball'] = self['webpackChunkcryptofotball'] || []).push([[630],{

    /***/ 6630:
    /***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {
        
        __webpack_require__.r(__webpack_exports__);

        /* harmony export */ __webpack_require__.d(__webpack_exports__, {
            /* harmony export */ 'default': () => (__WEBPACK_DEFAULT_EXPORT__)
            /* harmony export */ });

        /* harmony import */ const react_jsx_runtime__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(5893);
        /* harmony import */ const react_redux__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(9226);
        /* harmony import */ const react_router__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(5977);
        /* harmony import */ const _app_store_actions_users__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(8096);

        // Copyright (C) 2021 Creditor Corp. Group.
        // See LICENSE for copying information.

        /** TODO: Rework this view after design solution */
        const ConfirmEmail = function () {
            const dispatch = (0,react_redux__WEBPACK_IMPORTED_MODULE_1__/* .useDispatch */ .I0)();
            const useQuery = function () {
                return new URLSearchParams((0,react_router__WEBPACK_IMPORTED_MODULE_3__/* .useLocation */ .TH)().search);
            };
            const query = useQuery();
            const confirmEmail = function () {
                return dispatch((0,_app_store_actions_users__WEBPACK_IMPORTED_MODULE_2__/* .confirmUserEmail */ .U4)(query.get('token')));
            };

            ;

            return (0,react_jsx_runtime__WEBPACK_IMPORTED_MODULE_0__.jsx)('div', { children: (0,react_jsx_runtime__WEBPACK_IMPORTED_MODULE_0__.jsx)('input', {
                value: 'Confirm Email', onClick: confirmEmail 
            }, void 0) }, void 0);
        };
        /* harmony default export */ const __WEBPACK_DEFAULT_EXPORT__ = (ConfirmEmail);

        /***/ })

}]);
//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly9jcnlwdG9mb3RiYWxsLy4vc3JjL2FwcC92aWV3cy9Db25maXJtRW1haWwvaW5kZXgudHN4Il0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiI7Ozs7Ozs7Ozs7Ozs7OztBQUFBLDJDQUEyQztBQUMzQyx1Q0FBdUM7QUFFRztBQUNDO0FBRWtCO0FBRTdELG1EQUFtRDtBQUNuRCxJQUFNLFlBQVksR0FBYTtJQUMzQixJQUFNLFFBQVEsR0FBRyxrRUFBVyxFQUFFLENBQUM7SUFFL0IsSUFBTSxRQUFRLEdBQUc7UUFDYixPQUFPLElBQUksZUFBZSxDQUFDLG1FQUFXLEVBQUUsQ0FBQyxNQUFNLENBQUMsQ0FBQztJQUNyRCxDQUFDLENBQUM7SUFFRixJQUFNLEtBQUssR0FBRyxRQUFRLEVBQUUsQ0FBQztJQUV6QixJQUFNLFlBQVksR0FBRztRQUNqQixlQUFRLENBQUMsb0ZBQWdCLENBQUMsS0FBSyxDQUFDLEdBQUcsQ0FBQyxPQUFPLENBQUMsQ0FBQyxDQUFDO0lBQTlDLENBQThDLENBQUM7SUFFbkQsQ0FBQztJQUVELE9BQU8sMEVBQ0gsa0VBQ0ksS0FBSyxFQUFDLGVBQWUsRUFDckIsT0FBTyxFQUFFLFlBQVksV0FDdkIsV0FDQSxDQUFDO0FBQ1gsQ0FBQyxDQUFDO0FBRUYsaUVBQWUsWUFBWSxFQUFDIiwiZmlsZSI6IjYzMC43MGViYWQwZDk1ZmYyODc0ODVlNC5qcyIsInNvdXJjZXNDb250ZW50IjpbIi8vIENvcHlyaWdodCAoQykgMjAyMSBDcmVkaXRvciBDb3JwLiBHcm91cC5cbi8vIFNlZSBMSUNFTlNFIGZvciBjb3B5aW5nIGluZm9ybWF0aW9uLlxuXG5pbXBvcnQgeyB1c2VEaXNwYXRjaCB9IGZyb20gJ3JlYWN0LXJlZHV4JztcbmltcG9ydCB7IHVzZUxvY2F0aW9uIH0gZnJvbSAncmVhY3Qtcm91dGVyJztcblxuaW1wb3J0IHsgY29uZmlybVVzZXJFbWFpbCB9IGZyb20gJ0AvYXBwL3N0b3JlL2FjdGlvbnMvdXNlcnMnO1xuXG4vKiogVE9ETzogUmV3b3JrIHRoaXMgdmlldyBhZnRlciBkZXNpZ24gc29sdXRpb24gKi9cbmNvbnN0IENvbmZpcm1FbWFpbDogUmVhY3QuRkMgPSAoKSA9PiB7XG4gICAgY29uc3QgZGlzcGF0Y2ggPSB1c2VEaXNwYXRjaCgpO1xuXG4gICAgY29uc3QgdXNlUXVlcnkgPSAoKSA9PiB7XG4gICAgICAgIHJldHVybiBuZXcgVVJMU2VhcmNoUGFyYW1zKHVzZUxvY2F0aW9uKCkuc2VhcmNoKTtcbiAgICB9O1xuXG4gICAgY29uc3QgcXVlcnkgPSB1c2VRdWVyeSgpO1xuXG4gICAgY29uc3QgY29uZmlybUVtYWlsID0gKCkgPT5cbiAgICAgICAgZGlzcGF0Y2goY29uZmlybVVzZXJFbWFpbChxdWVyeS5nZXQoJ3Rva2VuJykpKTtcblxuICAgIDtcblxuICAgIHJldHVybiA8ZGl2PlxuICAgICAgICA8aW5wdXRcbiAgICAgICAgICAgIHZhbHVlPVwiQ29uZmlybSBFbWFpbFwiXG4gICAgICAgICAgICBvbkNsaWNrPXtjb25maXJtRW1haWx9XG4gICAgICAgIC8+XG4gICAgPC9kaXY+O1xufTtcblxuZXhwb3J0IGRlZmF1bHQgQ29uZmlybUVtYWlsO1xuIl0sInNvdXJjZVJvb3QiOiIifQ==