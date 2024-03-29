// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import diamondShadow from '@static/img/MarketPlacePage/marketPlaceCardsGroup/diamondShadow.svg';
import goldShadow from '@static/img/MarketPlacePage/marketPlaceCardsGroup/goldShadow.svg';
import silverShadow from '@static/img/MarketPlacePage/marketPlaceCardsGroup/silverShadow.svg';
import woodShadow from '@static/img/MarketPlacePage/marketPlaceCardsGroup/woodShadow.svg';

const DEFAULT_VALUE = 0;
const IS_MINTED_VALUE = 1;

/** class for our getters to get label and value while mapping */
export class CardField {
    /** label and value for mapping */
    constructor(public label: string, public value: string | number) { }
}

/** CardsQueryParametersField is an interface for cards and lots query parameters field, that consist of key and value. */
export interface CardsQueryParametersField {
    [key: string]: string[] | string | undefined;
};

/** CardsQueryParameters is class that uses for filtering cards and lots by queries. */
export class CardsQueryParameters {
    [key: string]: string | string[] | undefined;
    public 'defence_gte': string | undefined = undefined;
    public 'defence_lte': string | undefined = undefined;
    public 'goalkeeping_gte': string | undefined = undefined;
    public 'goalkeeping_lte': string | undefined = undefined;
    public 'offense_gte': string | undefined = undefined;
    public 'offense_lte': string | undefined = undefined;
    public 'physique_gte': string | undefined = undefined;
    public 'physique_lte': string | undefined = undefined;
    public 'tactics_gte': string | undefined = undefined;
    public 'tactics_lte': string | undefined = undefined;
    public 'quality': string[] = [];
    public 'status': string | undefined = undefined;
    public 'technique_gte': string | undefined = undefined;
    public 'technique_lte': string | undefined = undefined;
};


/* eslint-disable */
/** player stats implementation */
export class CardStats {
    /** main stat with substats */
    constructor(
        public title: string = '',
        public fields: CardField[] = []
    ) { }
    /** Returns average value of fields */
    get average() {
        return Math.round(
            this.fields
                .map((item) => +item.value)
                .reduce((prev, current) => prev + current) / this.fields.length
        );
    }

    /** stat giagram color depend on avarage stat value */
    get color(): string {
        const STATISTIC_UPPER_BOUND = 90;
        const STATISTIC_LOWER_BOUND = 50;

        const STATISTIC_UPPER_BOUND_COLOR = '#3CCF5D';
        const STATISTIC_MEDIUM_BOUND_COLOR = '#E8EC16';
        const STATISTIC_LOWER_BOUND_COLOR = '#FF4200';

        switch (true) {
            case this.average >= STATISTIC_UPPER_BOUND:
                return STATISTIC_UPPER_BOUND_COLOR;
            case this.average >= STATISTIC_LOWER_BOUND:
                return STATISTIC_MEDIUM_BOUND_COLOR;
            default:
                return STATISTIC_LOWER_BOUND_COLOR;
        }
    }
}

export interface ICard {
    id: string;
    playerName: string;
    quality: string;
    pictureType: number;
    height: number;
    weight: number;
    skinColor: number;
    hairStyle: number;
    hairColor: number;
    accessories: number[];
    dominantFoot: string;
    isTattoos: boolean;
    status: number;
    type: string;
    userId: string;
    tactics: number;
    positioning: number;
    composure: number;
    aggression: number;
    vision: number;
    awareness: number;
    crosses: number;
    physique: number;
    acceleration: number;
    runningSpeed: number;
    reactionSpeed: number;
    agility: number;
    stamina: number;
    strength: number;
    jumping: number;
    balance: number;
    technique: number;
    dribbling: number;
    ballControl: number;
    weakFoot: number;
    skillMoves: number;
    finesse: number;
    curve: number;
    volleys: number;
    shortPassing: number;
    longPassing: number;
    forwardPass: number;
    offense: number;
    finishingAbility: number;
    shotPower: number;
    accuracy: number;
    distance: number;
    penalty: number;
    freeKicks: number
    corners: number;
    headingAccuracy: number;
    defence: number;
    offsideTrap: number;
    sliding: number;
    tackles: number;
    ballFocus: number;
    interceptions: number;
    vigilance: number;
    goalkeeping: number;
    reflexes: number;
    diving: number;
    handling: number;
    sweeping: number;
    throwing: number;

}
/** Class defines with needed getters */
export class Card {
    id: string = '00000000-0000-0000-0000-000000000000';
    playerName: string = 'Taras';
    quality: string = '';
    pictureType: number = DEFAULT_VALUE;
    height: number = DEFAULT_VALUE;
    weight: number = DEFAULT_VALUE;
    skinColor: number = DEFAULT_VALUE;
    hairStyle: number = DEFAULT_VALUE;
    hairColor: number = DEFAULT_VALUE;
    accessories: number[] = [];
    dominantFoot: string = '';
    isTattoos: boolean = false;
    status: number = DEFAULT_VALUE;
    type: string = '';
    userId: string = '';
    tactics: number = DEFAULT_VALUE;
    positioning: number = DEFAULT_VALUE;
    composure: number = DEFAULT_VALUE;
    aggression: number = DEFAULT_VALUE;
    vision: number = DEFAULT_VALUE;
    awareness: number = DEFAULT_VALUE;
    crosses: number = DEFAULT_VALUE;
    physique: number = DEFAULT_VALUE;
    acceleration: number = DEFAULT_VALUE;
    runningSpeed: number = DEFAULT_VALUE;
    reactionSpeed: number = DEFAULT_VALUE;
    agility: number = DEFAULT_VALUE;
    stamina: number = DEFAULT_VALUE;
    strength: number = DEFAULT_VALUE;
    jumping: number = DEFAULT_VALUE;
    balance: number = DEFAULT_VALUE;
    technique: number = DEFAULT_VALUE;
    dribbling: number = DEFAULT_VALUE;
    ballControl: number = DEFAULT_VALUE;
    weakFoot: number = DEFAULT_VALUE;
    skillMoves: number = DEFAULT_VALUE;
    finesse: number = DEFAULT_VALUE;
    curve: number = DEFAULT_VALUE;
    volleys: number = DEFAULT_VALUE;
    shortPassing: number = DEFAULT_VALUE;
    longPassing: number = DEFAULT_VALUE;
    forwardPass: number = DEFAULT_VALUE;
    offense: number = DEFAULT_VALUE;
    finishingAbility: number = DEFAULT_VALUE;
    shotPower: number = DEFAULT_VALUE;
    accuracy: number = DEFAULT_VALUE;
    distance: number = DEFAULT_VALUE;
    penalty: number = DEFAULT_VALUE;
    freeKicks: number = DEFAULT_VALUE;
    corners: number = DEFAULT_VALUE;
    headingAccuracy: number = DEFAULT_VALUE;
    defence: number = DEFAULT_VALUE;
    offsideTrap: number = DEFAULT_VALUE;
    sliding: number = DEFAULT_VALUE;
    tackles: number = DEFAULT_VALUE;
    ballFocus: number = DEFAULT_VALUE;
    interceptions: number = DEFAULT_VALUE;
    vigilance: number = DEFAULT_VALUE;
    goalkeeping: number = DEFAULT_VALUE;
    reflexes: number = DEFAULT_VALUE;
    diving: number = DEFAULT_VALUE;
    handling: number = DEFAULT_VALUE;
    sweeping: number = DEFAULT_VALUE;
    throwing: number = DEFAULT_VALUE;
    isMinted: number = DEFAULT_VALUE | IS_MINTED_VALUE;

    /** Card fields */
    constructor(card?: ICard) {
        Object.assign(this, card);
    }

    /** Returns background type and shadow type according to quality */
    get shadow() {
        switch (this.quality) {
            case 'wood':
                return woodShadow;
            case 'silver':
                return silverShadow;
            case 'gold':
                return goldShadow;
            case 'diamond':
                return diamondShadow;
            default:
                return woodShadow;
        }
    }

    /** TODO: for testing, will be replaced */
    /* eslint-disable */
    get cardPrice() {
        const prp = 75;
        const pr = '1,142,000 - 15,000,000';
        const updated = 16;
        const price = '10,868,000';
        /** get stat giagram color depend on price value  */
        const PRICE_UPPER_BOUND = 80;
        const PRICE_MEDIUM_BOUND = 70;
        const PRICE_LOWER_BOUND = 50;

        const PRICE_UPPER_BOUND_COLOR = '#1898D7';
        const PRICE_MEDIUM_BOUND_COLOR = '#3CCF5D';
        const PRICE_LOWER_BOUND_COLOR = '#E86C27';
        const PRICE_DEFAULT_BOUND_COLOR = '#FF4200';
        let color: string;

        switch (true) {
            case prp >= PRICE_UPPER_BOUND:
                color = PRICE_UPPER_BOUND_COLOR;
                break;
            case prp >= PRICE_MEDIUM_BOUND:
                color = PRICE_MEDIUM_BOUND_COLOR;
                break;
            case prp >= PRICE_LOWER_BOUND:
                color = PRICE_LOWER_BOUND_COLOR;
                break;
            default:
                color = PRICE_DEFAULT_BOUND_COLOR;
        }

        return {
            prp,
            color,
            pr,
            updated,
            price,
        };
    }

    /** Using in footballerCard in info block */
    get infoBlock() {
        return [
            new CardField('name', this.playerName),
            // TODO: at this momenty nation does not exist.
            new CardField('nation', 'this.nation'),
            new CardField('skills', '5'),
            new CardField('weak foot', this.weakFoot),
            new CardField('intl. rep', '5'),
            new CardField('foot', this.dominantFoot),
            new CardField('height', this.height),
            new CardField('nation', this.weight),
            // TODO: at this momenty revision does not exist or it is designer mistake or it is quality.
            new CardField('revision', 'rare'),
            // TODO: create method to convert attack and defence values into this scale.
            new CardField('def. wr', 'low'),
            new CardField('arr. wr', 'high'),
            // next fields does not exist in card at this moment.
            new CardField('added on', '2020-09-10'),
            new CardField('origin', 'na'),
            new CardField('r. Face', 'low'),
            new CardField('b. type', ''),
            new CardField('age', '36 years old'),
        ];
    }

    /** Using in diagramm area in footballerCard */
    get diagramArea() {
        // TODO: need to get real min and max values to convert into diagram value.
        // TODO: this fields does not exist.
        return [
            new CardField('physical', DEFAULT_VALUE),
            new CardField('mental', DEFAULT_VALUE),
            new CardField('skill', DEFAULT_VALUE),
            new CardField('cham. style', DEFAULT_VALUE),
            new CardField('base stats', DEFAULT_VALUE),
            new CardField('in game stats', DEFAULT_VALUE),
        ];
    }

    /** Returns fields for card stats area in footballerCard */
    get statsArea() {
        return [
            new CardStats('tactics', [
                new CardField('positioning', this.positioning),
                new CardField('composure', this.composure),
                new CardField('aggression', this.aggression),
                new CardField('vision', this.vision),
                new CardField('awareness', this.awareness),
                new CardField('crosses', this.crosses),
            ]),
            new CardStats('physique', [
                new CardField('acceleration', this.acceleration),
                new CardField('running speed', this.runningSpeed),
                new CardField('reaction speed', this.reactionSpeed),
                new CardField('agility', this.agility),
                new CardField('stamina', this.stamina),
                new CardField('strength', this.strength),
                new CardField('jumping', this.jumping),
                new CardField('balance', this.jumping),
            ]),
            new CardStats('technique', [
                new CardField('dribbing', this.dribbling),
                new CardField('ball fontrol', this.ballControl),
                new CardField('weak foot', this.weakFoot),
                new CardField('skill moves', this.skillMoves),
                new CardField('finesse', this.finesse),
                new CardField('curve', this.curve),
                new CardField('volleys', this.volleys),
                new CardField('short passing', this.shortPassing),
                new CardField('long passing', this.longPassing),
                new CardField('forward pass', this.forwardPass),
            ]),
            new CardStats('offence', [
                new CardField('finishing ability', this.finishingAbility),
                new CardField('shot power', this.shotPower),
                new CardField('accuracy', this.accuracy),
                new CardField('distance', this.distance),
                new CardField('penalty', this.penalty),
                new CardField('free Kicks', this.freeKicks),
                new CardField('corners', this.corners),
                new CardField('heading accuracy', this.headingAccuracy),
            ]),
            new CardStats('defence', [
                new CardField('offside trap', this.offsideTrap),
                new CardField('tackles', this.tackles),
                new CardField('ball focus', this.ballFocus),
                new CardField('interceptions', this.interceptions),
                new CardField('vigilance', this.vigilance),
            ]),
            new CardStats('goalkeeping', [
                new CardField('diving', this.diving),
                new CardField('handling', this.handling),
                new CardField('sweeping', this.sweeping),
                new CardField('throwing', this.throwing),
            ]),
        ];
    }
}
/** Cards domain entity */
export class CardsPage {
    /** default Cards initial values */
    constructor(
        public cards: Card[],
        public page: {
            offset: number;
            limit: number;
            currentPage: number;
            pageCount: number;
            totalCount: number;
        }
    ) { }
}
