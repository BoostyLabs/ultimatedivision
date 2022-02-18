// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { arrayFromNumber } from "@/app/internal/arrayFromNumber";
import emptyNft from "@static/img/StorePage/SellNft/empty-nft.svg";
import fillNft from "@static/img/StorePage/SellNft/fill-nft.svg";
import "./index.scss";

export const CardsAmountDisplay: React.FC<{
  maxAmount: number;
  amountOfActive: number;
}> = ({ maxAmount, amountOfActive }) => {
  const cards = arrayFromNumber(maxAmount);
  cards.forEach((card, index) => (card.active = index < amountOfActive));

  return (
    <div className="card">
      {cards.map((card) => (
        <div key={card.id} className="card__item">
          <img src={card.active ? fillNft : emptyNft} alt="card nft" />
        </div>
      ))}
    </div>
  );
};
