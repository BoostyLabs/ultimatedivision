// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.


import { useEffect, useState } from 'react';

import { CardsQueryParameters, CardsQueryParametersField } from '@/card';

/** Hook useCardsQueryParameters returns array of current CardsQueryParametersField. */
export const useCardsQueryParameters = (fields: string[], cardsQueryParameters: CardsQueryParameters) => {
    const [cardsQueryParametersFields, setCardsQueryParametersFields] = useState<CardsQueryParametersField[]>([]);

    useEffect(() => {
        (() => {
            const currentCardsQueryParametersField: CardsQueryParametersField[] = []
            fields.forEach((field) => {
                currentCardsQueryParametersField.push({ [field]: cardsQueryParameters[field] })
            });
            setCardsQueryParametersFields(currentCardsQueryParametersField);
        })();

    }, [cardsQueryParameters]);

    return cardsQueryParametersFields;
};
