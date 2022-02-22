// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';

import './index.scss';


export const NftCards: React.FC<{ card: string }> = ({ card }) => {
    const AMOUNT_OF_CARDS = 7;

    /** Creates array with empty elements and then fill them */
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
