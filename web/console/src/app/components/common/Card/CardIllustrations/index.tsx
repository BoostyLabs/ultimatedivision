// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { FootballerCardIllustrationsDiagramsArea } from "@/app/components/common/Card/CardIllustrationsDiagramsArea";
import { FootballerCardIllustrationsRadar } from "@/app/components/common/Card/CardIllustrationsRadar";

import { Card } from "@/card";
import { PlayerCard } from "@components/common/PlayerCard";

import "./index.scss";

export const FootballerCardIllustrations: React.FC<{ card: Card }> = ({
  card,
}) => (
  <div className="footballer-card-illustrations">
    <div className="footballer-card-illustrations__card">
      <PlayerCard
        card={card}
        parentClassName="footballer-card-illustrations__card"
      />
    </div>
    <div className="footballer-card-illustrations__divider"></div>
    <FootballerCardIllustrationsRadar card={card} />
    <div className="footballer-card-illustrations__divider"></div>
    <FootballerCardIllustrationsDiagramsArea card={card} />
  </div>
);
