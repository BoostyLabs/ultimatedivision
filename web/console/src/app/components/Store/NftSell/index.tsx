// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { NftCards } from './NftCards';
import { MintingArea } from './MintingArea';
import { CardMintingProgress } from './CardMintingProgress';

import mockCard from '@static/img/StorePage/SellNft/mock-card.svg';

import './index.scss';

export const NftSell = () => {
    // TODO: change on real data
    const MOCK_NFT_AMOUNT = 7;
    const MOCK_TIME = '22:12:03';
    const MOCK_TITLE = 'DIAMOND CARD';
    const MOCK_CARD = mockCard;

    const MAX_NFT_AMOUNT = 10;

    return (
        <div className="sell-nft">
            <div className="sell-nft__wrapper">
                <NftCards card={mockCard} />
                <div className="sell-nft__info">
                    <h1 className="sell-nft__title">{MOCK_TITLE}</h1>
                    <div className="sell-nft__line"></div>
                    <div className="sell-nft__remainder">
                        <p className="sell-nft__remainder__text">Remaining</p>
                        <span className="sell-nft__remainder__amount">
                            {MOCK_NFT_AMOUNT} NFTS
                        </span>
                    </div>
                    <CardMintingProgress
                        max={MAX_NFT_AMOUNT}
                        activeCardsCount={MOCK_NFT_AMOUNT}
                    />
                    <div className="sell-nft__line"></div>
                    <MintingArea isInactive={!MOCK_NFT_AMOUNT} time={MOCK_TIME} />
                </div>
            </div>
        </div>
    );
};
