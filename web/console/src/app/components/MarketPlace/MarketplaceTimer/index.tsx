// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';

import { Lot } from '@/marketplace';

const DEFAULT_TIME_VALUE = 0;
const HOURS = 24;
const MINUTES_AND_SECONDS = 60;
const INTERVAL = 1000;
const CONVERT_MILLISECONDS = 1000;

export const MarketplaceTimer: React.FC<{ lot: Lot; className?: string }> = ({ lot, className }) => {
    const [hours, setHours] = useState(DEFAULT_TIME_VALUE);
    const [minutes, setMinutes] = useState(DEFAULT_TIME_VALUE);
    const [seconds, setSeconds] = useState(DEFAULT_TIME_VALUE);

    const getTime = (deadline: string) => {
        const time = Date.parse(deadline) - Date.now();

        setHours(Math.floor(time / (CONVERT_MILLISECONDS * MINUTES_AND_SECONDS * MINUTES_AND_SECONDS) % HOURS));
        setMinutes(Math.floor(time / CONVERT_MILLISECONDS / MINUTES_AND_SECONDS % MINUTES_AND_SECONDS));
        setSeconds(Math.floor(time / CONVERT_MILLISECONDS % MINUTES_AND_SECONDS));
    };

    useEffect(() => {
        const interval = setInterval(() => getTime(lot.endTime), INTERVAL);

        return () => clearInterval(interval);
    }, [lot]);

    return (
        <div className={`${className ? className : ''}`}>
            {hours} : {minutes} : {seconds}
        </div>
    );
};
