// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect } from "react";

import { useTimeCounter } from '@/app/hooks/useTimeCounter';

export const DelayTimer: React.FC = () => {
    const timeCounter = useTimeCounter();

    useEffect(() => {
    }, []);

    return <div className="match-finder__timer">
        <span className="match-finder__timer__text">
            {timeCounter.minutes} : {timeCounter.seconds}
        </span>
    </div>
};
