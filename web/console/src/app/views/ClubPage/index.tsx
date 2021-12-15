// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState, useEffect, useMemo } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { ClubCardsArea } from '@components/Club/ClubCardsArea';
import { FilterField } from '@components/common/FilterField';
import { FilterByPrice } from '@components/common/FilterField/FilterByPrice';
import { FilterByStats } from '@components/common/FilterField/FilterByStats';
import { FilterByStatus } from '@components/common/FilterField/FilterByStatus';
import { FilterByVersion } from '@components/common/FilterField/FilterByVersion';
import { Paginator } from '@components/common/Paginator';
import { RegistrationPopup } from '@/app/components/common/Registration/Registration';

import { UnauthorizedError } from '@/api';
import { RootState } from '@/app/store';
import { listOfCards, clearConcretCardsQueryParameters, createCardsQueryParameters, getQueryParameters } from '@/app/store/actions/cards';
import { CardsQueryParametersField, CardsQueryParameters } from '@/card';

import './index.scss';

const Club: React.FC = () => {
    const { page } = useSelector((state: RootState) => state.cardsReducer.cardsPage);
    const isCardsVisible = useSelector((state: RootState) => state.clubsReducer.options.showCardSeletion);

    const dispatch = useDispatch();

    const cardsQueryParameters = getQueryParameters();

    /** Indicates if registration is required. */
    const [isRegistrationRequired, setIsRegistrationRequired] = useState(false);

    /** Exposes default page number. */
    const DEFAULT_PAGE_INDEX: number = 1;

    /** Submits search by cards query parameters. */
    const submitSearch = async(queryParameters: CardsQueryParametersField[]) => {
        createCardsQueryParameters(queryParameters);
        await dispatch(listOfCards(DEFAULT_PAGE_INDEX));
    };

    const clearsStatisticsField = async(queryParameters: CardsQueryParametersField[]) => {
        clearConcretCardsQueryParameters(queryParameters);
        await dispatch(listOfCards(DEFAULT_PAGE_INDEX));
    }

    /** Closes RegistrationPopup componnet. */
    const closeRegistrationPopup = () => {
        setIsRegistrationRequired(false);
    };

    useEffect(() => {
        (async () => {
            try {
                // clearCardsQueryParameters();
                await dispatch(listOfCards(DEFAULT_PAGE_INDEX));
            } catch (error: any) {
                if (error instanceof UnauthorizedError) {
                    setIsRegistrationRequired(true);

                    return;
                };
            };
        })();
    }, [isCardsVisible]);

    return (
        <section className="club">
            {isRegistrationRequired && <RegistrationPopup closeRegistrationPopup={closeRegistrationPopup} />}
            <h1 className="club__title">
                MY CARDS
            </h1>
            <FilterField >
                <FilterByVersion submitSearch={submitSearch} />
                <FilterByStats
                    cardsQueryParameters={cardsQueryParameters}
                    clearsStatisticsField={clearsStatisticsField}
                    submitSearch={submitSearch} />
                <FilterByPrice />
                <FilterByStatus />
            </FilterField>
            <ClubCardsArea />
            <Paginator
                getCardsOnPage={listOfCards}
                itemsCount={page.totalCount}
                selectedPage={page.currentPage}
            />
        </section>
    );
};

export default Club;
