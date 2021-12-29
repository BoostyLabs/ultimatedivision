// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';
import Unity, { UnityContext } from 'react-unity-webgl';

import './index.scss';

const FootballGame: React.FC = () => {
    const unityContext = new UnityContext({
        loaderUrl: '/static/dist/Build/footbal.loader.js',
        dataUrl: '/static/dist/Build/footbal.data',
        frameworkUrl: '/static/dist/Build/footbal.framework.js',
        codeUrl: '/static/dist/Build/footbal.wasm',
    });

    unityContext.on('progress', (progression) => {
        setProgress(progression);
    });

    return (
        <div id="unity-container" className="unity-desktop">
            <Unity unityContext={unityContext}>
                <canvas id="unity-canvas"></canvas>
                <div id="unity-loading-bar">
                    <div id="unity-logo"></div>
                    <div id="unity-progress-bar-empty">
                        <div id="unity-progress-bar-full"></div>
                    </div>
                </div>
                <div id="unity-mobile-warning">
                    WebGL builds are not supported on mobile devices.
                </div>
                <div id="unity-footer">
                    <div id="unity-webgl-logo"></div>
                    <div id="unity-fullscreen-button"></div>
                    <div id="unity-build-title">Footbal</div>
                </div>
            </Unity>
        </div>
    );
};

export default FootballGame;
