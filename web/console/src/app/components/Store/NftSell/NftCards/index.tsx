// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import './index.scss';
import { useEffect, useState } from 'react';

export const NftCards: React.FC<{ card: string }> = ({ card }) => {
    const AMOUNT_OF_CARDS = 7;

    const [cards, setCards] = useState(new Array(AMOUNT_OF_CARDS).fill({}));

    const fillActiveCards = () => {
        setCards(
            cards.map((_, index) => ({
                img: card,
            }))
        );
    };

    useEffect(() => {
        fillActiveCards();
    }, [card]);

    return (
        <div className="nft-card">
            {cards.map((card, index) =>
                <div key={index} className={`nft-card__item-${index}`}>
                    <img src={card.img} alt="card nft" />
                </div>
            )}
        </div>
    );
};
