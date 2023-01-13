// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useMemo, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useParams } from 'react-router';
import { Link } from 'react-router-dom';

import { FootballerCardIllustrationsRadar } from '@/app/components/common/Card/CardIllustrationsRadar';
import { FootballerCardPrice } from '@/app/components/common/Card/CardPrice';
import { FootballerCardStatsArea } from '@/app/components/common/Card/CardStatsArea';
import { PlayerCard } from '@/app/components/common/PlayerCard';
import { ToastNotifications } from '@/notifications/service';

import { RootState } from '@/app/store';
import { openUserCard } from '@/app/store/actions/cards';
import { setCurrentUser } from '@/app/store/actions/users';

import WalletService from '@/wallet/service';

import CardPageBg from '@static/img/FootballerCardPage/background.png';

import './index.scss';

const Card: React.FC = () => {
    const dispatch = useDispatch();

    const [isMinted, setIsMinted] = useState<boolean>(false);
    const user = useSelector((state: RootState) => state.usersReducer.user);
    const { card } = useSelector((state: RootState) => state.cardsReducer);
    const { id }: { id: string } = useParams();

    /** implements opening new card */
    async function openCard() {
        try {
            await dispatch(openUserCard(id));
        } catch (error: any) {
            Notification
        }
    }
    /** implements opening new card */
    async function setUser() {
        try {
            await dispatch(setCurrentUser());
        } catch (error: any) {
            ToastNotifications.couldNotGetUser();
        }
    }

    const mint = async() => {
        const walletService = new WalletService(user);
        await walletService.mintNft(id);

        setIsMinted(true);
    };


    useEffect(() => {
        setUser();
        openCard();
    }, []);

    return (
        card &&
            <div className="card">
                <div className="card__border">
                    <div className="card__wrapper">

                        <div className="card__info">
                            <PlayerCard className="card__player" id={card.id} />
                            <div className="card__player__info">
                                <h2 className="card__name">{card.playerName}</h2>
                                <div className="card__mint-info">
                                    <div className="card__mint-info__nft">
                                        <span className="card__mint-info__nft-title">NFT:</span>
                                        <span className="card__mint-info__nft-value">
                                            {isMinted ? 'minted to Polygon' : 'not minted'}
                                        </span>
                                        {!isMinted &&
                                            <button className="card__mint" onClick={mint}>
                                                Mint now
                                            </button>
                                        }
                                    </div>
                                    <div className="card__mint-info__club">
                                        <span className="card__mint-info__club-title">Club:</span>
                                        <span className="card__mint-info__club-name">FC228</span>
                                    </div>
                                </div>
                            </div>
                            <div className="card__illustrator-radar">
                                <h2 className="card__illustrator-radar__title">Skills</h2>
                                <FootballerCardIllustrationsRadar card={card} />
                            </div>
                        </div>
                        <div className="card__stats-area">
                            <FootballerCardPrice card={card} isMinted={isMinted} />
                            <FootballerCardStatsArea card={card} />
                        </div>
                    </div>
                </div>
                <img src={CardPageBg} alt="background" className="card__bg" />
            </div>

    );
};

export default Card;
