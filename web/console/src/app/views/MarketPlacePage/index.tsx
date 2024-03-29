// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useDispatch, useSelector } from 'react-redux';
import { useEffect, useState } from 'react';

import { FilterField } from '@components/common/FilterField';
import { FilterByPrice } from '@components/common/FilterField/FilterByPrice';
import { FilterByStats } from '@components/common/FilterField/FilterByStats';
import { FilterByStatus } from '@components/common/FilterField/FilterByStatus';
import { FilterByVersion } from '@components/common/FilterField/FilterByVersion';
import { Paginator } from '@components/common/Paginator';
import { MarketPlaceCardsGroup } from '@components/MarketPlace/MarketPlaceCardsGroup';
import { MarketPlaceModal } from '@/app/components/MarketPlace/MarketplaceModal';

import { RootState } from '@/app/store';
import {
    createLotsQueryParameters,
    getCurrentLotsQueryParameters,
    listOfLots,
} from '@/app/store/actions/marketplace';
import { CardsQueryParametersField } from '@/card';
import { Lot } from '@/marketplace';
import { setCurrentUser } from '@/app/store/actions/users';
import { ToastNotifications } from '@/notifications/service';

import './index.scss';

const MarketPlace: React.FC = () => {
    const dispatch = useDispatch();

    const { lots, page } = useSelector(
        (state: RootState) => state.marketplaceReducer.marketplacePage
    );
    const [currentLot, setCurrentLot] = useState<Lot>();
    const [showModal, setShowModal] = useState<boolean>(false);

    const handleShowModal = (lot: Lot) => {
        setShowModal(true);
        setCurrentLot(lot);
    };

    /** Exposes default page number. */
    const DEFAULT_PAGE_INDEX: number = 1;

    const lotsQueryParameters = getCurrentLotsQueryParameters();

    /** Submits search by lots query parameters. */
    const submitSearch = async(
        queryParameters: CardsQueryParametersField[]
    ) => {
        createLotsQueryParameters(queryParameters);
        await dispatch(listOfLots(DEFAULT_PAGE_INDEX));
    };

    /** sets user info */
    async function setUser() {
        try {
            await dispatch(setCurrentUser());
        } catch (error: any) {
            ToastNotifications.couldNotGetUser();
        }
    }

    useEffect(() => {
        setUser();
    }, []);

    return (
        <>
            {showModal && currentLot && <MarketPlaceModal lot={currentLot} setShowModal={setShowModal} />}
            <section className="marketplace" >
                <FilterField>
                    <FilterByVersion
                        cardsQueryParameters={lotsQueryParameters}
                        submitSearch={submitSearch}
                    />
                    <FilterByStats
                        cardsQueryParameters={lotsQueryParameters}
                        submitSearch={submitSearch}
                    />
                    <FilterByPrice />
                    <FilterByStatus />
                </FilterField>
                <MarketPlaceCardsGroup lots={lots} handleShowModal={handleShowModal} />
                <Paginator
                    getCardsOnPage={listOfLots}
                    itemsCount={page.totalCount}
                    selectedPage={page.currentPage}
                />
            </section >
        </>
    );
};

export default MarketPlace;
