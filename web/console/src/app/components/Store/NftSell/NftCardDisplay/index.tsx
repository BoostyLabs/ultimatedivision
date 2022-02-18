// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import mockCard from "@static/img/StorePage/SellNft/mock-card.svg";
import { arrayFromNumber } from "@/app/internal/arrayFromNumber";
import "./index.scss";

export const NftCardDisplay = () => {
  const AMOUNT_OF_CARDS = 7;

  const TOP_INDENT = 15;
  const SIDE_INDENT = 32;
  const HIGHT_DECREASE = 60;
  const SIDE_INDENT_DECREASE = 1.5;
  const LINE_SEPARATOR_AMOUNT = 2;
  const DETERMINAL_OF_EVEN = 2;

  const CARDS_POSITION = "50%";
  const INTIAL_CARD_HEIGHT = "452px";

  const cards = arrayFromNumber(AMOUNT_OF_CARDS);
  cards.forEach((card) => (card.img = mockCard));

  return (
    <div className="nft-card">
      {cards.map((card, index) => {
        const CARD_ROW = Math.ceil(index / LINE_SEPARATOR_AMOUNT);
        const SIMILAR_STYLES = {
          top: `${CARD_ROW * TOP_INDENT}px`,
          zIndex: AMOUNT_OF_CARDS - index,
        };

        return index ? (
          <div
            key={card.id}
            className="nft-card__item"
            style={
              index % DETERMINAL_OF_EVEN === 0
                ? {
                    right: `calc(${CARDS_POSITION} - ${
                      CARD_ROW * (SIDE_INDENT - CARD_ROW * SIDE_INDENT_DECREASE)
                    }px`,

                    transform: `translate(calc(${CARDS_POSITION} + ${
                      CARD_ROW * (SIDE_INDENT - CARD_ROW * SIDE_INDENT_DECREASE)
                    }px ), ${CARD_ROW * TOP_INDENT}px)`,
                    ...SIMILAR_STYLES,
                  }
                : {
                    left: `calc(${CARDS_POSITION} - ${
                      CARD_ROW * (SIDE_INDENT - CARD_ROW * SIDE_INDENT_DECREASE)
                    }px`,
                    transform: `translate(calc(-${CARDS_POSITION} - ${
                      CARD_ROW * (SIDE_INDENT - CARD_ROW * SIDE_INDENT_DECREASE)
                    }px ), ${CARD_ROW * TOP_INDENT}px)`,
                    ...SIMILAR_STYLES,
                  }
            }
          >
            <img
              src={card.img}
              alt="card nft"
              style={{
                height: `calc(${INTIAL_CARD_HEIGHT} - ${
                  CARD_ROW * HIGHT_DECREASE
                }px`,
              }}
            />
          </div>
        ) : (
          <div className="nft-card__main-item" key={card.id}>
            <img src={mockCard} alt="card nft" />
          </div>
        );
      })}
    </div>
  );
};
