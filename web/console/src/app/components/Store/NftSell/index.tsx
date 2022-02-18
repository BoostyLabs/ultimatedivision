// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { NftCardDisplay } from "./NftCardDisplay";
import { MintingArea } from "./MintingArea";
import { CardsAmountDisplay } from "./CardsAmountDisplay";
import "./index.scss";

export const NftSell = () => {
  const MOCK_NFT_AMOUNT = 7;
  const MOCK_TIME = "22:12:03";

  const MAX_NFT_AMOUNT = 10;

  return (
    <div className="sell-nft">
      <div className="sell-nft__wrapper">
        <NftCardDisplay />
        <div className="sell-nft__info">
          <h1 className="sell-nft__title">DIAMOND CARD</h1>
          <div className="sell-nft__line"></div>
          <div className="sell-nft__remainder">
            <p className="sell-nft__remainder__text">Remaining</p>
            <span className="sell-nft__remainder__amount">
              {MOCK_NFT_AMOUNT} NFTS
            </span>
          </div>
          <CardsAmountDisplay
            maxAmount={MAX_NFT_AMOUNT}
            amountOfActive={MOCK_NFT_AMOUNT}
          />
          <div className="sell-nft__line"></div>
          <MintingArea isInactive={!MOCK_NFT_AMOUNT} time={MOCK_TIME} />
        </div>
      </div>
    </div>
  );
};
