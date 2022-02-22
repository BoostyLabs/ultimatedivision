// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';

import emptyNft from '@static/img/StorePage/SellNft/empty-nft.svg';
import fillNft from '@static/img/StorePage/SellNft/fill-nft.svg';

import './index.scss';

export const CardMintingProgress: React.FC<{
    max: number;
    activeCardsCount: number;
}> = ({ max, activeCardsCount }) => {
    /** Creates array with empty elements and then fill them */
    const [cards, setCards] = useState(new Array(max).fill({}));

    const fillActiveCards = () => {
        setCards(
            cards.map((_, index) => ({
                active: index < activeCardsCount,
            }))
        );
    };

    useEffect(() => {
        fillActiveCards();
    }, [activeCardsCount]);

    return (
        <div className="card">
            {cards.map((card, index) =>
                <div key={index} className="card__item">
                    <img src={card.active ? fillNft : emptyNft} alt="card nft" />
                </div>
            )}
        </div>
    );
};
