// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from "react";
import Unity, { UnityContext } from "react-unity-webgl";

import "./index.scss";

const FootballGame: React.FC = () => {
    const unityContext = new UnityContext({
        companyName: "DefaultCompany",
        productName: "Football2019",
        loaderUrl: "/static/dist/webGl/UnityLoader.js",
        dataUrl: "/static/dist/webGl/footbal.data.unityweb",
        frameworkUrl: "/static/dist/webGl/footbal.wasm.framework.unityweb",
        codeUrl: "/static/dist/webGl/footbal.wasm",
    });

    return <Unity unityContext={unityContext} className="unity-container" />;
};

export default FootballGame;
